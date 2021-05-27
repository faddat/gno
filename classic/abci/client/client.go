package abcicli

import (
	"fmt"
	"sync"

	"github.com/tendermint/classic/abci/types"
	cmn "github.com/tendermint/classic/libs/common"
)

const (
	dialRetryIntervalSeconds = 3
	echoRetryIntervalSeconds = 1
)

// Client defines an interface for an ABCI client.
// All `Async` methods return a `ReqRes` object.
// All `Sync` methods return the appropriate protobuf ResponseXxx struct and an error.
// Note these are client errors, eg. ABCI socket connectivity issues.
// Application-related errors are reflected in response via ABCI error codes and logs.
type Client interface {
	cmn.Service

	SetResponseCallback(Callback)
	Error() error

	FlushAsync() *ReqRes
	EchoAsync(msg string) *ReqRes
	InfoAsync(abci.RequestInfo) *ReqRes
	SetOptionAsync(abci.RequestSetOption) *ReqRes
	DeliverTxAsync(abci.RequestDeliverTx) *ReqRes
	CheckTxAsync(abci.RequestCheckTx) *ReqRes
	QueryAsync(abci.RequestQuery) *ReqRes
	CommitAsync() *ReqRes
	InitChainAsync(abci.RequestInitChain) *ReqRes
	BeginBlockAsync(abci.RequestBeginBlock) *ReqRes
	EndBlockAsync(abci.RequestEndBlock) *ReqRes

	FlushSync() error
	EchoSync(msg string) (abci.ResponseEcho, error)
	InfoSync(abci.RequestInfo) (abci.ResponseInfo, error)
	SetOptionSync(abci.RequestSetOption) (abci.ResponseSetOption, error)
	DeliverTxSync(abci.RequestDeliverTx) (abci.ResponseDeliverTx, error)
	CheckTxSync(abci.RequestCheckTx) (abci.ResponseCheckTx, error)
	QuerySync(abci.RequestQuery) (abci.ResponseQuery, error)
	CommitSync() (abci.ResponseCommit, error)
	InitChainSync(abci.RequestInitChain) (abci.ResponseInitChain, error)
	BeginBlockSync(abci.RequestBeginBlock) (abci.ResponseBeginBlock, error)
	EndBlockSync(abci.RequestEndBlock) (abci.ResponseEndBlock, error)
}

//----------------------------------------

// NewClient returns a new ABCI client of the specified transport type.
// It returns an error if the transport is not "socket".
func NewClient(addr, transport string, mustConnect bool) (client Client, err error) {
	switch transport {
	case "socket":
		client = NewSocketClient(addr, mustConnect)
	default:
		err = fmt.Errorf("Unknown abci transport %s", transport)
	}
	return
}

//----------------------------------------

type Callback func(abci.Request, abci.Response)

//----------------------------------------

type ReqRes struct {
	abci.Request
	wg            *sync.WaitGroup
	abci.Response // Not set atomically, so be sure to use WaitGroup.

	mtx  sync.Mutex
	done bool                // Gets set to true once *after* WaitGroup.Done().
	cb   func(abci.Response) // A single callback that may be set.
}

func NewReqRes(req abci.Request) *ReqRes {
	return &ReqRes{
		Request:  req,
		wg:       waitGroup1(),
		Response: nil,

		done: false,
		cb:   nil,
	}
}

// Sets the callback for this ReqRes atomically.
// If reqRes is already done, calls cb immediately.
// NOTE: reqRes.cb should not change if reqRes.done.
// NOTE: only one callback is supported.
func (reqRes *ReqRes) SetCallback(cb func(res abci.Response)) {
	reqRes.mtx.Lock()

	if reqRes.done {
		reqRes.mtx.Unlock()
		cb(reqRes.Response)
		return
	}

	reqRes.cb = cb
	reqRes.mtx.Unlock()
}

func (reqRes *ReqRes) GetCallback() func(abci.Response) {
	reqRes.mtx.Lock()
	defer reqRes.mtx.Unlock()
	return reqRes.cb
}

func (reqRes *ReqRes) Wait() {
	reqRes.wg.Wait()
}

// NOTE: it should be safe to read reqRes.cb without locks after this.
func (reqRes *ReqRes) Done() {
	reqRes.mtx.Lock()
	reqRes.done = true
	reqRes.mtx.Unlock()

	// Finally, release the hounds.
	reqRes.wg.Done()
}

func waitGroup1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}

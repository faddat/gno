package utils

import (
	"encoding/hex"
	"errors"
	"strings"
	"time"

	ctypes "github.com/tendermint/classic/rpc/core/types"
	"github.com/tendermint/go-amino-x"

	"github.com/tendermint/classic/sdk/client/context"
	sdk "github.com/tendermint/classic/sdk/types"
	"github.com/tendermint/classic/sdk/x/auth/types"
)

// QueryTxsByEvents performs a search for transactions for a given set of events
// via the Tendermint RPC. An event takes the form of:
// "{eventAttribute}.{attributeKey} = '{attributeValue}'". Each event is
// concatenated with an 'AND' operand. It returns a slice of Info object
// containing txs and metadata. An error is returned if the query fails.
func QueryTxsByEvents(cliCtx context.CLIContext, events []string, page, limit int) (*sdk.SearchTxsResult, error) {
	if len(events) == 0 {
		return nil, errors.New("must declare at least one event to search")
	}

	if page <= 0 {
		return nil, errors.New("page must greater than 0")
	}

	if limit <= 0 {
		return nil, errors.New("limit must greater than 0")
	}

	// XXX: implement ANY
	query := strings.Join(events, " AND ")

	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	prove := !cliCtx.TrustNode

	resTxs, err := node.TxSearch(query, prove, page, limit)
	if err != nil {
		return nil, err
	}

	if prove {
		for _, tx := range resTxs.Txs {
			err := ValidateTxResult(cliCtx, tx)
			if err != nil {
				return nil, err
			}
		}
	}

	resBlocks, err := getBlocksForTxResults(cliCtx, resTxs.Txs)
	if err != nil {
		return nil, err
	}

	txs, err := formatTxResults(resTxs.Txs, resBlocks)
	if err != nil {
		return nil, err
	}

	result := sdk.NewSearchTxsResult(resTxs.TotalCount, len(txs), page, limit, txs)

	return &result, nil
}

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTx(cliCtx context.CLIContext, hashHexStr string) (sdk.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	resTx, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	if !cliCtx.TrustNode {
		if err = ValidateTxResult(cliCtx, resTx); err != nil {
			return sdk.TxResponse{}, err
		}
	}

	resBlocks, err := getBlocksForTxResults(cliCtx, []*ctypes.ResultTx{resTx})
	if err != nil {
		return sdk.TxResponse{}, err
	}

	out, err := formatTxResult(resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}

// formatTxResults parses the indexed txs into a slice of TxResponse objects.
func formatTxResults(resTxs []*ctypes.ResultTx, resBlocks map[int64]*ctypes.ResultBlock) ([]sdk.TxResponse, error) {
	var err error
	out := make([]sdk.TxResponse, len(resTxs))
	for i := range resTxs {
		out[i], err = formatTxResult(resTxs[i], resBlocks[resTxs[i].Height])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

// ValidateTxResult performs transaction verification.
func ValidateTxResult(cliCtx context.CLIContext, resTx *ctypes.ResultTx) error {
	if !cliCtx.TrustNode {
		check, err := cliCtx.Verify(resTx.Height)
		if err != nil {
			return err
		}
		err = resTx.Proof.Validate(check.Header.DataHash)
		if err != nil {
			return err
		}
	}
	return nil
}

func getBlocksForTxResults(cliCtx context.CLIContext, resTxs []*ctypes.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlocks := make(map[int64]*ctypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(&resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func formatTxResult(resTx *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (sdk.TxResponse, error) {
	tx, err := parseTx(resTx.Tx)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return sdk.NewResponseResultTx(resTx, tx, resBlock.Block.Time.Format(time.RFC3339)), nil
}

func parseTx(txBytes []byte) (sdk.Tx, error) {
	var tx types.StdTx

	err := amino.UnmarshalSized(txBytes, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

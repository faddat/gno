// PKGPATH: gno.land/r/test
package test

var root Node

type Node interface{}
type Key interface{}

type InnerNode struct {
	Key   Key
	Left  Node `gno:owned`
	Right Node `gno:owned`
}

func main() {
	key := "somekey"
	root = InnerNode{
		Key:   key,
		Left:  nil,
		Right: nil,
	}
}

// Realm:
// c[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:2]=VI[object:OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:0#F4D429297D647D907C26D462A4252F09C71F3549@0&1]:
// - EI[:473287F8298DBA7163A897908958F7C0EAE733E2(#4B5916C2E9A29495310FAA63903E3C9C7FB8CBAB)] // VI[text:"somekey"]
// - EI[nil]
// - EI[nil]
// u[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:0]=VI[object:OIDNONE:0#4CB6B9B545DD4C65409E601CF4D1F801FF591150@1&0]:
// - EI[:1AF40977153D0FABAB9803BF33EDEBA8EB420CC5(#96EED8EB3C33E427143CC83BC3BD112A962B5A40)] // VI[type:0FA10A8AA773837EA2180964EDE25A2617D1558F]
// - EI[OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:2:#A766D206B6B98AD2208EA1C4C60BBF9F15826AC0] // TypeID:D4EB8490A382914D14B504ABFA2AF0F5144253FA VI[object:OIDA8ADA09DEE16D791FD406D629FE29BB0ED084A30:0#F4D429297D647D907C26D462A4252F09C71F3549@0&1]
// - EI[:1AF40977153D0FABAB9803BF33EDEBA8EB420CC5(#C59AF0F8D72349E62D988F6C70B18B6EDD4AB045)] // VI[type:47EDC0DB396625856E9893FF3BFD82C657B63F04]
// - EI[:1AF40977153D0FABAB9803BF33EDEBA8EB420CC5(#02828B237FD4A208B4F34553AC66B5EB5996ABE0)] // VI[type:D4EB8490A382914D14B504ABFA2AF0F5144253FA]
// - EI[nil]

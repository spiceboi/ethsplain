package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// Splain sthoeuhstohu
type Splain struct {
	Tokens []Token
}

// Token why can't you just be happy linter
type Token struct {
	Hex  string
	Text string
	More string
}

type field int

var (
	NONCE     field = 0
	GAS_PRICE field = 1
	GAS_LIMIT field = 2
	RECIPIENT field = 3
	VALUE     field = 4
	DATA      field = 5
	SIG_V     field = 6
	SIG_R     field = 7
	SIG_S     field = 8
)

func main() {
	parse()
}

func parse() []byte {
	//http.HandleFunc("/", tokenize)
	//log.Fatal(httpListenAndServe(":8080", nil))

	str := strings.TrimPrefix(data, "0x")
	tx := &types.Transaction{}
	buf, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("start")
	r := bytes.NewReader(buf)
	s := rlp.NewStream(r, 0)
	fmt.Println("predecode")
	err = tx.DecodeRLP(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx:", tx)

	// Because we rely on geth to parse the tx earlier
	// we can make this parser really dumb because we know it is in
	// standard tx format

	//tx.To

	splain := Splain{}
	//var tok Token

	tx.Nonce()
	//tmp := make([]byte, 10)
	encNonce, err := rlp.EncodeToBytes(tx.Nonce())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HEX; ", hex.EncodeToString(encNonce))

	// We find the overall rlp prefix by reading until we find the nonce
	//prefix := buf[:bytes.Index(buf, encNonce)]
	//fmt.Println("PREFIX:", hex.EncodeToString(prefix))
	//splain.addNode(prefix)
	//addRLPNode(&splain, prefix)

	splain.addNode(tx.Nonce(), NONCE)
	splain.addNode(tx.GasPrice().Bytes(), GAS_PRICE)
	splain.addNode(tx.Gas(), GAS_LIMIT)
	splain.addNode(tx.To().Bytes(), RECIPIENT)
	splain.addNode(tx.Value().Bytes(), VALUE)
	splain.addNode(tx.Data(), DATA)
	sigV, sigR, sigS := tx.RawSignatureValues()
	splain.addNode(sigV.Bytes(), SIG_V)
	splain.addNode(sigR.Bytes(), SIG_R)
	splain.addNode(sigS.Bytes(), SIG_S)
	out, _ := json.MarshalIndent(splain, "", "	")
	fmt.Println(string(out))

	concat := ""
	for _, tok := range splain.Tokens {
		concat += tok.Hex
	}
	fmt.Println("concat", concat)

	return out

	// convert value to bytes
	// rlp encode it
	// if it is larger than originally, pass the prefix to rlpExplain

	// start at beginning of raw
	// 2 walking pointers
	// one at start
	// walk until we get a match on the rlp encoding of the first field
	// th

}

func (s *Splain) addNode(val interface{}, f field) {

	enc, err := rlp.EncodeToBytes(val)
	if err != nil {
		log.Fatal(err)
	}
	i := addRLPNode(s, enc)

	// add the value node skipping however long the prefix was
	var tok Token
	tok.Hex = Hex(enc[i:])

	// construct the explanatory text
	var txt, more string
	switch f {
	case NONCE:
		txt, more = nonceInfo(val)
	default:
		txt = "NOT IMPLEMENTED"
		more = "Not IMPLEMENTED"

	}
	tok.Text = txt
	tok.More = more

	// Edgcase for when the prefix tells us the data length of the next argument is zero
	// we don't want to add a node for no data
	if len(tok.Hex) > 0 {
		s.Tokens = append(s.Tokens, tok)
	}

}

func nonceInfo(val interface{}) (string, string) {

	i, _ := val.(uint64)
	txt := fmt.Sprintf("Nonce: %d", i)
	more := "The nonce is a sequence number issued my the transaction creator used to prevent message replay. The nonce of each transaction of an account must be exactly 1 greater than the previous nonce used. The Ethereum yellow paper defines the nonce as 'A scalar value equal to the number of transactions sent from this address or, in the case of accounts with associated code, the number of contract-creations made by this account"

	return txt, more
}

// if there is a rlp length prefix add a node for it, else do nothing.
// Return how many bytes the prefix took
func addRLPNode(s *Splain, enc []byte) int {
	length := len(enc)
	if length == 0 {
		log.Fatal("Unable to decode length of val:", enc)
	}

	var node Token

	prefix := enc[0]
	// This is a single byte value that is its own rlp encoding so no node to add
	if prefix <= 0x7F {
		fmt.Println("Single Byte", enc)
		return 0
	}
	// "string" value of length 0-55
	if prefix <= 0xB7 && length > int(prefix-0x80) {
		node.Hex = Hex([]byte{prefix})
		node.Text = "RLP Length Prefix. The next field is an RLP 'string' of length FIXME"
		node.More = "Specific RLP Rule being used"
		s.Tokens = append(s.Tokens, node)
		return 1
	}
	// "string" value of length > 55
	if prefix < 0xC0 {
		// prefix tells us the length of the length of the field
		l := prefix - 0xb7
		flen := enc[1 : 1+l]
		node.Hex = Hex(append([]byte{prefix}, flen...))
		node.Text = "RLP Length Prefix. The next field is an RLP 'string' of length FIIXXX"
		node.More = "Stuff about the length of the length turtles all the way down"
		return 1 + len(flen)

	}

	log.Fatal("Not Implemented")

	return -1
}

// Hex how do i fix my linter plx halp
func Hex(b []byte) string {
	return hex.EncodeToString(b)
}

func rlpExplain(buf []byte) string {

	return ""
}

var data = "0xf8aa0185012a05f2008327c50e9435fb136cbadbc168910b66a9f7c40b03e4bd467f80b8441e9a695000000000000000000000000035fb136cbadbc168910b66a9f7c40b03e4bd467f000000000000000000000000000000000000000000000000000000003b9aca0026a00320143282b77654f3eedf2c6d384346a4be52c902f6603227f8f0220d30aa35a076ea8a4947327f33e149ec928efd6efa9e49aafe89a189abae7aad599c5feef2"

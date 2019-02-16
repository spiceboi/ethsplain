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

	splain.addNode(tx.Nonce())
	splain.addNode(tx.GasPrice().Bytes())
	splain.addNode(tx.Gas())
	splain.addNode(tx.To().Bytes())
	splain.addNode(tx.Value().Bytes())
	splain.addNode(tx.Data())
	sigV, sigR, sigS := tx.RawSignatureValues()
	splain.addNode(sigV.Bytes())
	splain.addNode(sigR.Bytes())
	splain.addNode(sigS.Bytes())
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

func (s *Splain) addNode(val interface{}) {

	enc, err := rlp.EncodeToBytes(val)
	if err != nil {
		log.Fatal(err)
	}
	i := addRLPNode(s, enc)

	// add the value node skipping however long the prefix was
	var tok Token
	tok.Hex = Hex(enc[i:])
	tok.Text = "I need to inject this somehow"
	tok.More = "I also need to inject this"

	// Edgcase for when the prefix tells us the data length of the next argument is zero
	// we don't want to add a node for no data
	if len(tok.Hex) > 0 {
		s.Tokens = append(s.Tokens, tok)
	}

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

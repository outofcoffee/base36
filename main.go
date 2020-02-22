package main

import (
	"github.com/martinlindhe/base36"

	"flag"
	"fmt"
)

func main() {
	encodePtr := flag.String("encode", "", "Text to encode.")
	decodePtr := flag.String("decode", "", "Text to decode.")
    flag.Parse()

	if encodePtr != nil && len(*encodePtr) > 0 {
		fmt.Println(base36.EncodeBytes([]byte(*encodePtr)))
	} else if decodePtr != nil && len(*decodePtr) > 0 {
		fmt.Println(string(base36.DecodeToBytes(*decodePtr)))
	} else {
		flag.Usage()
	}
}

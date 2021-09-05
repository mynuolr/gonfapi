package gonfapi

import (
	"encoding/binary"
	"fmt"
)

type INT [4]byte

func (i INT) Get() int32 {
	return int32(binary.LittleEndian.Uint32(i[:4]))
}
func (i INT) Set(in int32) {
	fmt.Printf("% x\n", uint32(0x1234567)) // prints 00000000 11111101
	binary.LittleEndian.PutUint32(i[:], uint32(in))
	printAsBinary(i[:])
	binary.BigEndian.PutUint32(i[:], uint32(in))
	printAsBinary(i[:])
	fmt.Println()
}

type ULING [4]byte

func printAsBinary(bytes []byte) {

	for i := 0; i < len(bytes); i++ {
		for j := 0; j < 8; j++ {
			zeroOrOne := bytes[i] >> (7 - j) & 1
			fmt.Printf("%c", '0'+zeroOrOne)
		}
		fmt.Printf(" %p\n", &bytes[i])
	}
	fmt.Println()
}

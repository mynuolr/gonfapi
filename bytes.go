package gonfapi

import (
	"encoding/binary"
	"fmt"
)

type INT [4]byte

func (i INT) Get() int32 {
	return int32(binary.LittleEndian.Uint32(i[:4]))
}
func (i *INT) Set(in int32) {
	binary.LittleEndian.PutUint32(i[:], uint32(in))
	printAsBinary(i[:])
}

type ULONG [4]byte

func (i ULONG) Get() uint32 {
	return binary.LittleEndian.Uint32(i[:4])
}
func (i *ULONG) Set(in uint32) {
	binary.LittleEndian.PutUint32(i[:], in)
	printAsBinary(i[:])
}

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

package gonfapi

import (
	"encoding/binary"
	"fmt"
)

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

type INT16 [2]byte

func (i INT16) Get() int16 {
	return int16(binary.LittleEndian.Uint16(i[:]))
}
func (i *INT16) Set(in int16) {
	binary.LittleEndian.PutUint16(i[:], uint16(in))
	printAsBinary(i[:])
}

type INT32 [4]byte

func (i INT32) Get() int32 {
	return int32(binary.LittleEndian.Uint32(i[:]))
}
func (i *INT32) Set(in int32) {
	binary.LittleEndian.PutUint32(i[:], uint32(in))
	printAsBinary(i[:])
}

type UINT16 [2]byte

func (i UINT16) Get() uint16 {
	return binary.LittleEndian.Uint16(i[:])
}
func (i *UINT16) Set(in uint16) {
	binary.LittleEndian.PutUint16(i[:], in)
	printAsBinary(i[:])
}

type UINT32 [4]byte

func (i UINT32) Get() uint32 {
	return binary.LittleEndian.Uint32(i[:])
}
func (i *UINT32) Set(in uint32) {
	binary.LittleEndian.PutUint32(i[:], in)
	printAsBinary(i[:])
}

type WChar_t [2]byte

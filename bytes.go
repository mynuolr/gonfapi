package gonfapi

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

var hostByteOrder binary.ByteOrder

func init() {
	var i int32 = 0x01020304
	if *(*byte)(unsafe.Pointer(&i)) == 0x04 {
		hostByteOrder = binary.LittleEndian
	} else {
		hostByteOrder = binary.BigEndian
	}

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

type INT16 [2]byte

func (i INT16) Get() int16 {
	return int16(hostByteOrder.Uint16(i[:]))
}
func (i *INT16) Set(in int16) {
	hostByteOrder.PutUint16(i[:], uint16(in))
	printAsBinary(i[:])
}

type INT32 [4]byte

func (i INT32) Get() int32 {
	return int32(hostByteOrder.Uint32(i[:]))
}
func (i *INT32) Set(in int32) {
	hostByteOrder.PutUint32(i[:], uint32(in))
	printAsBinary(i[:])
}

type UINT16 [2]byte

func (i UINT16) Get() uint16 {
	return hostByteOrder.Uint16(i[:])
}
func (i *UINT16) Set(in uint16) {
	hostByteOrder.PutUint16(i[:], in)
	printAsBinary(i[:])
}

type UINT32 [4]byte

func (i UINT32) Get() uint32 {
	return hostByteOrder.Uint32(i[:])
}
func (i *UINT32) Set(in uint32) {
	hostByteOrder.PutUint32(i[:], in)
	printAsBinary(i[:])
}

type UINT64 [8]byte

func (i *UINT64) Get() uint64 {
	return hostByteOrder.Uint64(i[:])
}
func (i *UINT64) Set(in uint64) {
	hostByteOrder.PutUint64(i[:], in)
	printAsBinary(i[:])
}

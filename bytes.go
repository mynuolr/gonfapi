//gonfapi Package sort provides primitives for sorting slices and user-defined
//collections.
package gonfapi

import (
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"unsafe"

	"github.com/mynuolr/gonfapi/basetype"
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

type INT16 = basetype.INT16

type INT32 = basetype.INT32

type UINT16 = basetype.UINT16

type UINT32 = basetype.UINT32

type UINT64 = basetype.UINT64

//sockaddr_in4/6
type SockaddrInx struct {
	Family      UINT16   //AF_INT or AF_INT6. LittleEndian
	Port        UINT16   //Port. BigEndian
	Data1       [4]byte  //ipv4 Adder,ipv6 is zero. BigEndian
	Data2       [16]byte //ipv6 Adder,ipv4 is zero. BigEndian
	IPV6ScopeId UINT32   //ipv6 scope id
}

var emptyBytes16 = make([]byte, 16)

func (s *SockaddrInx) String() string {
	_, ip := s.GetIP()
	return fmt.Sprintf("[%s]:%d", ip, s.GetPort())
}
func (s *SockaddrInx) SetIP(v4 bool, ip net.IP) {
	if v4 {
		s.Family.Set(AF_INET)
		copy(s.Data2[:], emptyBytes16)
		copy(s.Data1[:], ip.To4())
		s.IPV6ScopeId.Set(0)
	} else {
		s.Family.Set(AF_INET6)
		copy(s.Data1[:], emptyBytes16)
		copy(s.Data2[:], ip.To16())
	}
}
func (s *SockaddrInx) GetIP() (v4 bool, ip net.IP) {
	if !s.IsIpv6() {
		return true, net.IP(s.Data1[:])
	} else {
		return false, net.IP(s.Data2[:])
	}
}
func (s *SockaddrInx) IsIpv6() bool {
	return AF_INET6 == s.Family.Get()
}
func (s *SockaddrInx) GetPort() uint16 {
	return s.Port.BigEndianGet()
}
func (s *SockaddrInx) SetPort(p uint16) {
	s.Port.BigEndianSet(p)
}

// IP Addres
//
// |0000|0000|0000|0000|
//
// |ipv4|
//
// |------ ipv6 -------|
type IpAddress [16]byte

func (s *IpAddress) SetIP(v4 bool, ip net.IP) {
	if v4 {
		copy(s[:], emptyBytes16)
		copy(s[:4], ip.To4())
	} else {
		copy(s[:], ip.To16())
	}
}
func (s *IpAddress) GetIP(v4 bool) (ip net.IP) {
	if v4 {
		return net.IP(s[:4])
	} else {
		return net.IP(s[:])
	}
}

//指针转到数组切片
func PtrToBytes(b *byte, len int) (data []byte) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(b))
	sh.Cap = len
	sh.Len = len
	return
}

//指针转到SockaddrInx
func PtrToAddress(b *byte) *SockaddrInx {
	return (*SockaddrInx)(unsafe.Pointer(b))
}

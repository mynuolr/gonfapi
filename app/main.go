package main

import (
	"fmt"
	"unsafe"

	"github.com/mynuolr/gonfapi"
)

type cc2 struct {
	x  gonfapi.NF_RULE
	xs sd
}
type sd struct {
	a, b, c, d byte
}

func main() {
	t()
	var x = new(gonfapi.NF_RULE_EX)
	fmt.Println(x.GetProtocolId())
	x.SetProtocolID(305419896)
	fmt.Println(x.GetProtocolId())
	x.SetDirection(gonfapi.D_BOTH)
	fmt.Println(x.GetDirection())
	x.SetLocalPort(3306)
	x.SetRemotePort(20)
	x.SetProcessName("test")
	fmt.Println(x)
	fmt.Println(unsafe.Sizeof(*x))
	fmt.Println(x.GetProcessName())

}
func t() {
	var sf = gonfapi.NF_DATA{}
	fmt.Println(sf)
	sf.Code.Set(0x1234567)
	fmt.Printf("%b\n", sf.Code)
}

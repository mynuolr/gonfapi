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
	var x = new(gonfapi.NF_RULE)

	fmt.Println(x)
	fmt.Println(unsafe.Sizeof(*x))

}
func t() {
	var sf = gonfapi.NF_DATA{}
	fmt.Println(sf)
	sf.Code.Set(0x1234567)
	fmt.Printf("%b\n", sf.Code)
}

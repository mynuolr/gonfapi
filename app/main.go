package main

import (
	"fmt"

	"github.com/mynuolr/gonfapi"
)

func main() {
	api := &gonfapi.NFApi{}
	fmt.Println(api.Load("nfapi.dll"))

	fmt.Println(api.NfRegisterDriverEx("nf2", "."))
	api.NfUnRegisterDriver("nf2")
}

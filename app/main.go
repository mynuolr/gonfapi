package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mynuolr/gonfapi"
)

var api = &gonfapi.NFApi{}

func main() {

	fmt.Println(api.Load("nfapi.dll"))
	rule := &gonfapi.NF_RULE{}
	rule.RemotePort.Set(6003)
	fmt.Println()
	wd, _ := os.Getwd()
	fmt.Println(api.NfRegisterDriverEx("nfapi2", wd))
	api.NfFree()
	api.NfAddRule(rule, true)
	ev := &gonfapi.NF_EventHandler{}
	cb := e{}
	ev.Build(&cb)
	api.NfInit("nfapi2", ev)
	api.NfAdjustProcessPriviledges()
	defer func() {
		api.NfFree()
	}()
	<-time.After(time.Hour)
	//api.NfUnRegisterDriver("nfapi2")
}

type e struct{}

func (e *e) ThreadStart() uintptr {
	fmt.Println("Strt")
	return 0
}

func (e *e) ThreadEnd() uintptr {
	fmt.Println("Stop")
	return 0
}

func (e *e) TcpConnectRequest(id uint64, pConnInfo *gonfapi.NF_TCP_CONN_INFO) uintptr {
	fmt.Println("TcpConnectRequest", id, pConnInfo)
	return 0
}

func (e *e) TcpConnected(id uint64, pConnInfo *gonfapi.NF_TCP_CONN_INFO) uintptr {
	fmt.Println("TcpConnected", id, pConnInfo)
	return 0
}

func (e *e) TcpClosed(id uint64, pConnInfo *gonfapi.NF_TCP_CONN_INFO) uintptr {
	fmt.Println("TcpClosed", id, pConnInfo)
	return 0
}

func (e *e) TcpReceive(id uint64, buf uintptr, len int32) uintptr {
	fmt.Println("TcpReceive", id, buf, len)
	return 0
}
func (e *e) TcpSend(id uint64, buf uintptr, len int32) uintptr {
	fmt.Println("TcpSend", id, buf, len)
	return 0
}
func (e *e) TcpCanReceive(id uint64) uintptr {
	fmt.Println("TcpCanReceive", id)
	return 0
}

func (e *e) TcpCanSend(id uint64) uintptr {
	fmt.Println("TcpCanSend", id)
	return 0
}

func (e *e) UdpCreated(id uint64, pConnInfo *gonfapi.NF_UDP_CONN_INFO) uintptr {
	fmt.Println("UdpCreated", id, pConnInfo)
	fmt.Println("UdpCreated ProcessID", pConnInfo.ProcessId.Get(), pConnInfo.ProcessId)
	fmt.Println(pConnInfo.IpFamily.Get() == gonfapi.AF_INET6)
	fmt.Println(api.NfGetProcessNameFromKernel(pConnInfo.ProcessId.Get()))
	return 0
}

func (e *e) UdpConnectRequest(id uint64, pConnInfo *gonfapi.NF_UDP_CONN_REQUEST) uintptr {
	fmt.Println("UdpConnectRequest", id, pConnInfo)
	return 0
}

func (e *e) UdpClosed(id uint64, pConnInfo *gonfapi.NF_UDP_CONN_INFO) uintptr {
	fmt.Println("UdpClosed", id, pConnInfo)
	return 0
}

func (e *e) UdpReceive(id uint64, remoteAddress uintptr, buf uintptr, len int32, options *gonfapi.NF_UDP_OPTIONS) uintptr {
	fmt.Println("UdpReceive", id, remoteAddress, buf, len, options)
	return 0
}

func (e *e) UdpSend(id uint64, remoteAddress uintptr, buf uintptr, len int32, options *gonfapi.NF_UDP_OPTIONS) uintptr {
	fmt.Println("UdpSend", id, remoteAddress, buf, len, options)
	return 0
}

func (e *e) UdpCanReceive(id uint64) uintptr {
	fmt.Println("UdpCanReceive", id)
	return 0
}

func (e *e) UdpCanSend(id uint64) uintptr {
	fmt.Println("UdpCanSend", id)
	return 0
}

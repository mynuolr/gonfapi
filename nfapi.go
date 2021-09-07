package gonfapi

import (
	"syscall"
	"unsafe"
)

type NFApi struct {
	dll                          *syscall.LazyDLL
	nf_init                      *syscall.LazyProc
	nf_free                      *syscall.LazyProc
	nf_registerDriver            *syscall.LazyProc
	nf_registerDriverEx          *syscall.LazyProc
	nf_unRegisterDriver          *syscall.LazyProc
	nf_tcpSetConnectionState     *syscall.LazyProc
	nf_tcpPostSend               *syscall.LazyProc
	nf_tcpPostReceive            *syscall.LazyProc
	nf_tcpClose                  *syscall.LazyProc
	nf_setTCPTimeout             *syscall.LazyProc
	nf_tcpDisableFiltering       *syscall.LazyProc
	nf_udpSetConnectionState     *syscall.LazyProc
	nf_udpPostSend               *syscall.LazyProc
	nf_udpPostReceive            *syscall.LazyProc
	nf_udpDisableFiltering       *syscall.LazyProc
	nf_ipPostSend                *syscall.LazyProc
	nf_ipPostReceive             *syscall.LazyProc
	nf_addRule                   *syscall.LazyProc
	nf_deleteRules               *syscall.LazyProc
	nf_setRules                  *syscall.LazyProc
	nf_addRuleEx                 *syscall.LazyProc
	nf_setRulesEx                *syscall.LazyProc
	nf_getConnCount              *syscall.LazyProc
	nf_tcpSetSockOpt             *syscall.LazyProc
	nf_getProcessNameA           *syscall.LazyProc
	nf_getProcessNameW           *syscall.LazyProc
	nf_getProcessNameFromKernel  *syscall.LazyProc
	nf_adjustProcessPriviledges  *syscall.LazyProc
	nf_tcpIsProxy                *syscall.LazyProc
	nf_setOptions                *syscall.LazyProc
	nf_completeTCPConnectRequest *syscall.LazyProc
	nf_completeUDPConnectRequest *syscall.LazyProc
	nf_getTCPConnInfo            *syscall.LazyProc
	nf_getUDPConnInfo            *syscall.LazyProc
	nf_setIPEventHandler         *syscall.LazyProc
	nf_addFlowCtl                *syscall.LazyProc
	nf_deleteFlowCtl             *syscall.LazyProc
	nf_setTCPFlowCtl             *syscall.LazyProc
	nf_setUDPFlowCtl             *syscall.LazyProc
	nf_modifyFlowCtl             *syscall.LazyProc
	nf_getFlowCtlStat            *syscall.LazyProc
	nf_getTCPStat                *syscall.LazyProc
	nf_getUDPStat                *syscall.LazyProc
	nf_addBindingRule            *syscall.LazyProc
	nf_deleteBindingRules        *syscall.LazyProc
	nf_getDriverType             *syscall.LazyProc
}

func (a *NFApi) Load(dll string) error {
	a.dll = syscall.NewLazyDLL(dll)
	e := a.dll.Load()
	if e != nil {
		return e
	}
	a.nf_init = a.dll.NewProc("nf_init")
	a.nf_free = a.dll.NewProc("nf_free")
	a.nf_registerDriver = a.dll.NewProc("nf_registerDriver")
	a.nf_registerDriverEx = a.dll.NewProc("nf_registerDriverEx")
	a.nf_unRegisterDriver = a.dll.NewProc("nf_unRegisterDriver")

	a.nf_tcpSetConnectionState = a.dll.NewProc("nf_tcpSetConnectionState")
	a.nf_tcpPostSend = a.dll.NewProc("nf_tcpPostSend")
	a.nf_tcpPostReceive = a.dll.NewProc("nf_tcpPostReceive")
	a.nf_tcpClose = a.dll.NewProc("nf_tcpClose")
	a.nf_setTCPTimeout = a.dll.NewProc("nf_setTCPTimeout")
	a.nf_tcpDisableFiltering = a.dll.NewProc("nf_tcpDisableFiltering")

	a.nf_udpSetConnectionState = a.dll.NewProc("nf_udpSetConnectionState")
	a.nf_udpPostSend = a.dll.NewProc("nf_udpPostSend")
	a.nf_udpPostReceive = a.dll.NewProc("nf_udpPostReceive")
	a.nf_udpDisableFiltering = a.dll.NewProc("nf_udpDisableFiltering")

	a.nf_ipPostSend = a.dll.NewProc("nf_ipPostSend")
	a.nf_ipPostReceive = a.dll.NewProc("nf_ipPostReceive")

	a.nf_addRule = a.dll.NewProc("nf_addRule")
	a.nf_deleteRules = a.dll.NewProc("nf_deleteRules")
	a.nf_setRules = a.dll.NewProc("nf_setRules")
	a.nf_addRuleEx = a.dll.NewProc("nf_addRuleEx")
	a.nf_setRulesEx = a.dll.NewProc("nf_setRulesEx")

	a.nf_getConnCount = a.dll.NewProc("nf_getConnCount")
	a.nf_tcpSetSockOpt = a.dll.NewProc("nf_tcpSetSockOpt")

	a.nf_getProcessNameA = a.dll.NewProc("nf_getProcessNameA")
	a.nf_getProcessNameW = a.dll.NewProc("nf_getProcessNameW")
	a.nf_getProcessNameFromKernel = a.dll.NewProc("nf_getProcessNameFromKernel")
	a.nf_adjustProcessPriviledges = a.dll.NewProc("nf_adjustProcessPriviledges")
	a.nf_tcpIsProxy = a.dll.NewProc("nf_tcpIsProxy")
	a.nf_setOptions = a.dll.NewProc("nf_setOptions")
	a.nf_completeTCPConnectRequest = a.dll.NewProc("nf_completeTCPConnectRequest")
	a.nf_completeUDPConnectRequest = a.dll.NewProc("nf_completeUDPConnectRequest")
	a.nf_getTCPConnInfo = a.dll.NewProc("nf_getTCPConnInfo")
	a.nf_getUDPConnInfo = a.dll.NewProc("nf_getUDPConnInfo")

	a.nf_setIPEventHandler = a.dll.NewProc("nf_setIPEventHandler")
	a.nf_addFlowCtl = a.dll.NewProc("nf_addFlowCtl")
	a.nf_deleteFlowCtl = a.dll.NewProc("nf_deleteFlowCtl")
	a.nf_setTCPFlowCtl = a.dll.NewProc("nf_setTCPFlowCtl")
	a.nf_setUDPFlowCtl = a.dll.NewProc("nf_setUDPFlowCtl")
	a.nf_modifyFlowCtl = a.dll.NewProc("nf_modifyFlowCtl")
	a.nf_getFlowCtlStat = a.dll.NewProc("nf_getFlowCtlStat")
	a.nf_getTCPStat = a.dll.NewProc("nf_getTCPStat")
	a.nf_getUDPStat = a.dll.NewProc("nf_getUDPStat")
	a.nf_addBindingRule = a.dll.NewProc("nf_addBindingRule")
	a.nf_deleteBindingRules = a.dll.NewProc("nf_deleteBindingRules")
	a.nf_getDriverType = a.dll.NewProc("nf_getDriverType")
	return nil
}
func (a NFApi) NfInit(driverName string, Ev *NF_EventHandler) {
	sp, err := syscall.BytePtrFromString(driverName)
	if err != nil {
		return
	}
	a.nf_init.Call(uintptr(unsafe.Pointer(sp)), uintptr(unsafe.Pointer(Ev)))
}
func (a NFApi) NfFree() {
	a.nf_free.Call()
}
func (a NFApi) NfRegisterDriver(driverName string) (int, error) {
	sp, err := syscall.BytePtrFromString(driverName)
	if err != nil {
		return 0, err
	}
	_, r, err := a.nf_registerDriver.Call(uintptr(unsafe.Pointer(sp)))
	return int(r), err
}
func (a NFApi) NfRegisterDriverEx(driverName string, path string) (int, error) {
	sp, err := syscall.BytePtrFromString(driverName)
	if err != nil {
		return 0, err
	}
	pathp, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}

	_, r, err := a.nf_registerDriverEx.Call(uintptr(unsafe.Pointer(sp)), uintptr(unsafe.Pointer(pathp)))
	return int(r), err
}
func (a NFApi) NfUnRegisterDriver(driverName string) {
	sp, err := syscall.BytePtrFromString(driverName)
	if err != nil {
		return
	}
	a.nf_unRegisterDriver.Call(uintptr(unsafe.Pointer(sp)))
}
func (a NFApi) NfTcpSetConnectionState(id uint64, suspended bool) {
	var suspend int32 = 0
	if suspended {
		suspend = 1
	}
	a.nf_tcpSetConnectionState.Call(uintptr(id), uintptr(suspend))
}

func (a NFApi) NfTcpPostSend(id uint64, bufer []byte) {
	ptr := unsafe.Pointer(&bufer)
	len := len(bufer)
	a.nf_tcpPostSend.Call(uintptr(id), uintptr(ptr), uintptr(len))
}
func (a NFApi) NfTcpPostReceive(id uint64, bufer []byte) {
	ptr := unsafe.Pointer(&bufer)
	len := len(bufer)
	a.nf_tcpPostReceive.Call(uintptr(id), uintptr(ptr), uintptr(len))
}
func (a NFApi) NfTcpClose(id uint64) {
	a.nf_tcpClose.Call(uintptr(id))
}
func (a NFApi) NfSetTCPTimeout(id uint32) {
	a.nf_tcpClose.Call(uintptr(id))
}
func (a NFApi) NfTcpDisableFiltering(id uint64) {
	a.nf_tcpDisableFiltering.Call(uintptr(id))
}

// UDP
func (a NFApi) NfUdpSetConnectionState(id uint64, suspended bool) {
	var suspend int32 = 0
	if suspended {
		suspend = 1
	}
	a.nf_udpSetConnectionState.Call(uintptr(id), uintptr(suspend))
}
func (a NFApi) NfUdpPostSend(id uint64, remoteAddress []uint8, buf []byte, option *NF_UDP_OPTIONS) {
	a.nf_udpPostSend.Call(
		uintptr(id),
		uintptr(unsafe.Pointer(&remoteAddress)),
		uintptr(unsafe.Pointer(&buf)),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(option)),
	)
}
func (a NFApi) NfUdpPostReceive(id uint64, remoteAddress []uint8, buf []byte, option *NF_UDP_OPTIONS) {
	a.nf_udpPostReceive.Call(
		uintptr(id),
		uintptr(unsafe.Pointer(&remoteAddress)),
		uintptr(unsafe.Pointer(&buf)),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(option)),
	)
}
func (a NFApi) NfUdpDisableFiltering(id uint64) {
	a.nf_udpDisableFiltering.Call(uintptr(id))
}

//IP
func (a NFApi) NfIpPostSend(buf []byte, option *NF_IP_PACKET_OPTIONS) {
	a.nf_ipPostSend.Call(
		uintptr(unsafe.Pointer(&buf)),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(option)),
	)
}
func (a NFApi) NfIpPostReceive(buf []byte, option *NF_IP_PACKET_OPTIONS) {
	a.nf_ipPostReceive.Call(
		uintptr(unsafe.Pointer(&buf)),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(option)),
	)
}

// Rule
func (a NFApi) NfAddRule(rule *NF_RULE, ToHead bool) {
	var h int32 = 0
	if ToHead {
		h = 1
	}
	a.nf_addRule.Call(uintptr(unsafe.Pointer(rule)), uintptr(h))
}
func (a NFApi) NfDeleteRules() {
	a.nf_deleteRules.Call()
}
func (a NFApi) NfSetRules(rule []NF_RULE) {

	a.nf_setRules.Call(uintptr(unsafe.Pointer(&rule)), uintptr(int32(len(rule))))
}
func (a NFApi) NfAddRuleEx(rule *NF_RULE_EX, ToHead bool) {
	var h int32 = 0
	if ToHead {
		h = 1
	}
	a.nf_addRuleEx.Call(uintptr(unsafe.Pointer(rule)), uintptr(h))
}

func (a NFApi) NfSetRulesEx(rule []NF_RULE_EX) {
	a.nf_setRulesEx.Call(uintptr(unsafe.Pointer(&rule)), uintptr(int32(len(rule))))
}

//
// Debug routine
//
func (a NFApi) NfGetConnCount() {
	a.nf_getConnCount.Call()
}
func (a NFApi) NfTcpSetSockOpt(id uint64, optname int32, optval []byte) {
	a.nf_tcpSetSockOpt.Call(
		uintptr(id),
		uintptr(optname),
		uintptr(unsafe.Pointer(&optname)),
		uintptr(int32(len(optval))),
	)
}
func (a NFApi) NfGetProcessNameW(processId uint16) string {
	buf := [260]uint16{}
	a.nf_getProcessNameW.Call(uintptr(processId), uintptr(unsafe.Pointer(&buf)), uintptr(uint16(260)))
	return syscall.UTF16ToString(buf[:])
}

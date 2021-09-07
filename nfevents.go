package gonfapi

import "syscall"

type EventHandler interface {
	ThreadStart() uintptr
	ThreadEnd() uintptr
	TcpConnectRequest(id uint64, pConnInfo *NF_TCP_CONN_INFO) uintptr
	TcpConnected(id uint64, pConnInfo *NF_TCP_CONN_INFO) uintptr
	TcpClosed(id uint64, pConnInfo *NF_TCP_CONN_INFO) uintptr
	TcpReceive(id uint64, buf uintptr, len int32) uintptr
	TcpSend(id uint64, buf uintptr, len int32) uintptr
	TcpCanReceive(id uint64) uintptr
	TcpCanSend(id uint64) uintptr
	UdpCreated(id uint64, pConnInfo *NF_UDP_CONN_INFO) uintptr
	UdpConnectRequest(id uint64, pConnInfo *NF_UDP_CONN_REQUEST) uintptr
	UdpClosed(id uint64, pConnInfo *NF_UDP_CONN_INFO) uintptr
	UdpReceive(id uint64, remoteAddress uintptr, buf uintptr, len int32, options *NF_UDP_OPTIONS) uintptr
	UdpSend(id uint64, remoteAddress uintptr, buf uintptr, len int32, options *NF_UDP_OPTIONS) uintptr
	UdpCanReceive(id uint64) uintptr
	UdpCanSend(id uint64) uintptr
}
type NF_EventHandler struct {
	ThreadStart       uintptr
	ThreadEnd         uintptr
	TcpConnectRequest uintptr
	TcpConnected      uintptr
	TcpClosed         uintptr
	TcpReceive        uintptr
	TcpSend           uintptr
	TcpCanReceive     uintptr
	TcpCanSend        uintptr
	UdpCreated        uintptr
	UdpConnectRequest uintptr
	UdpClosed         uintptr
	UdpReceive        uintptr
	UdpSend           uintptr
	UdpCanReceive     uintptr
	UdpCanSend        uintptr
}

func (eh *NF_EventHandler) Build(e EventHandler) {
	eh.ThreadStart = syscall.NewCallbackCDecl(e.ThreadStart)
	eh.ThreadEnd = syscall.NewCallbackCDecl(e.ThreadEnd)
	eh.TcpConnectRequest = syscall.NewCallbackCDecl(e.TcpConnectRequest)
	eh.TcpConnected = syscall.NewCallbackCDecl(e.TcpConnected)
	eh.TcpClosed = syscall.NewCallbackCDecl(e.TcpClosed)
	eh.TcpReceive = syscall.NewCallbackCDecl(e.TcpReceive)
	eh.TcpSend = syscall.NewCallbackCDecl(e.TcpSend)
	eh.TcpCanReceive = syscall.NewCallbackCDecl(e.TcpCanReceive)
	eh.TcpCanSend = syscall.NewCallbackCDecl(e.TcpCanSend)
	eh.UdpCreated = syscall.NewCallbackCDecl(e.UdpCreated)
	eh.UdpConnectRequest = syscall.NewCallbackCDecl(e.UdpConnectRequest)
	eh.UdpClosed = syscall.NewCallbackCDecl(e.UdpClosed)
	eh.UdpReceive = syscall.NewCallbackCDecl(e.UdpReceive)
	eh.UdpSend = syscall.NewCallbackCDecl(e.UdpSend)
	eh.UdpCanReceive = syscall.NewCallbackCDecl(e.UdpCanReceive)
	eh.UdpCanSend = syscall.NewCallbackCDecl(e.UdpCanSend)
}

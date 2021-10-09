package gonfapi

import "golang.org/x/sys/windows"

//事件回调接口 所有返回值为0
type EventHandler interface {
	ThreadStart() uintptr
	ThreadEnd() uintptr
	TcpConnectRequest(id uint64, pConnInfo *NF_TCP_CONN_INFO) uintptr
	TcpConnected(id uint64, pConnInfo *NF_TCP_CONN_INFO) uintptr
	TcpClosed(id uint64, pConnInfo *NF_TCP_CONN_INFO) uintptr
	TcpReceive(id uint64, buf *byte, len int32) uintptr
	TcpSend(id uint64, buf *byte, len int32) uintptr
	TcpCanReceive(id uint64) uintptr
	TcpCanSend(id uint64) uintptr
	UdpCreated(id uint64, pConnInfo *NF_UDP_CONN_INFO) uintptr
	UdpConnectRequest(id uint64, pConnInfo *NF_UDP_CONN_REQUEST) uintptr
	UdpClosed(id uint64, pConnInfo *NF_UDP_CONN_INFO) uintptr
	UdpReceive(id uint64, remoteAddress *SockaddrInx, buf *byte, len int32, options *NF_UDP_OPTIONS) uintptr
	UdpSend(id uint64, remoteAddress *SockaddrInx, buf *byte, len int32, options *NF_UDP_OPTIONS) uintptr
	UdpCanReceive(id uint64) uintptr
	UdpCanSend(id uint64) uintptr
}

//NF_EventHandler 传递到dll的结构体 所有字段皆为回调参数指针
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

//赋值EventHandler的指针
func (eh *NF_EventHandler) Build(e EventHandler) {
	eh.ThreadStart = windows.NewCallbackCDecl(e.ThreadStart)
	eh.ThreadEnd = windows.NewCallbackCDecl(e.ThreadEnd)
	eh.TcpConnectRequest = windows.NewCallbackCDecl(e.TcpConnectRequest)
	eh.TcpConnected = windows.NewCallbackCDecl(e.TcpConnected)
	eh.TcpClosed = windows.NewCallbackCDecl(e.TcpClosed)
	eh.TcpReceive = windows.NewCallbackCDecl(e.TcpReceive)
	eh.TcpSend = windows.NewCallbackCDecl(e.TcpSend)
	eh.TcpCanReceive = windows.NewCallbackCDecl(e.TcpCanReceive)
	eh.TcpCanSend = windows.NewCallbackCDecl(e.TcpCanSend)
	eh.UdpCreated = windows.NewCallbackCDecl(e.UdpCreated)
	eh.UdpConnectRequest = windows.NewCallbackCDecl(e.UdpConnectRequest)
	eh.UdpClosed = windows.NewCallbackCDecl(e.UdpClosed)
	eh.UdpReceive = windows.NewCallbackCDecl(e.UdpReceive)
	eh.UdpSend = windows.NewCallbackCDecl(e.UdpSend)
	eh.UdpCanReceive = windows.NewCallbackCDecl(e.UdpCanReceive)
	eh.UdpCanSend = windows.NewCallbackCDecl(e.UdpCanSend)
}

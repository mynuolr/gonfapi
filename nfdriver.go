package gonfapi

import (
	"encoding/binary"
	"reflect"
	"syscall"
	"unsafe"
)

// C enum and #define is 4Bytes
const (
	TCP_PACKET_BUF_SIZE int32 = 8192
	UDP_PACKET_BUF_SIZE int32 = 2 * 65536
)

type DataCode int32

const (
	TCP_CONNECTED DataCode = iota
	TCP_CLOSED
	TCP_RECEIVE
	TCP_SEND
	TCP_CAN_RECEIVE
	TCP_CAN_SEND
	TCP_REQ_SUSPEND
	TCP_REQ_RESUME
	//UDP
	UDP_CREATED
	UDP_CLOSED
	UDP_RECEIVE
	UDP_SEND
	UDP_CAN_RECEIVE
	UDP_CAN_SEND
	UDP_REQ_SUSPEND
	UDP_REQ_RESUME
	//REQ RULE
	REQ_ADD_HEAD_RULE
	REQ_ADD_TAIL_RULE
	REQ_DELETE_RULES
	//CONNECT
	TCP_CONNECT_REQUEST
	UDP_CONNECT_REQUEST
	//other
	TCP_DISABLE_USER_MODE_FILTERING
	UDP_DISABLE_USER_MODE_FILTERING

	REQ_SET_TCP_OPT
	REQ_IS_PROXY

	TCP_REINJECT
	TCP_REMOVE_CLOSED
	TCP_DEFERRED_DISCONNECT

	IP_RECEIVE
	IP_SEND
	TCP_RECEIVE_PUSH
)

type DIRECTION int32

const (
	D_IN   DIRECTION = 1
	D_OUT  DIRECTION = 2
	D_BOTH DIRECTION = 3
)

type FILTERING_FLAG uint32

const (
	NF_ALLOW                       FILTERING_FLAG = 0   // Allow the activity without filtering transmitted packets
	NF_BLOCK                       FILTERING_FLAG = 1   // Block the activity
	NF_FILTER                      FILTERING_FLAG = 2   // Filter the transmitted packets
	NF_SUSPENDED                   FILTERING_FLAG = 4   // Suspend receives from server and sends from client
	NF_OFFLINE                     FILTERING_FLAG = 8   // Emulate establishing a TCP connection with remote server
	NF_INDICATE_CONNECT_REQUESTS   FILTERING_FLAG = 16  // Indicate outgoing connect requests to API
	NF_DISABLE_REDIRECT_PROTECTION FILTERING_FLAG = 32  // Disable blocking indicating connect requests for outgoing connections of local proxies
	NF_PEND_CONNECT_REQUEST        FILTERING_FLAG = 64  // Pend outgoing connect request to complete it later using nf_complete(TCP|UDP)ConnectRequest
	NF_FILTER_AS_IP_PACKETS        FILTERING_FLAG = 128 // Indicate the traffic as IP packets via ipSend/ipReceive
	NF_READONLY                    FILTERING_FLAG = 256 // Don't block the IP packets and indicate them to ipSend/ipReceive only for monitoring
	NF_CONTROL_FLOW                FILTERING_FLAG = 512
)

const (
	MAX_ADDRESS_LENGTH    = 28
	MAX_IP_ADDRESS_LENGTH = 16
	AF_INET               = 2
	AF_INET6              = 23
)

// NF_RULE
type NF_RULE [83]byte

func (r *NF_RULE) GetProtocol() int32 {
	return int32(binary.LittleEndian.Uint32(r[:4]))
}
func (r *NF_RULE) SetProtocol(i int32) {
	binary.LittleEndian.PutUint32(r[:4], uint32(i))
}
func (r *NF_RULE) GetProtocolId() uint32 {
	return binary.LittleEndian.Uint32(r[4:8])
}
func (r *NF_RULE) SetProtocolID(i uint32) {
	binary.LittleEndian.PutUint32(r[4:8], i)
}
func (r *NF_RULE) GetDirection() DIRECTION {
	return DIRECTION(uint(r[8]))

}
func (r *NF_RULE) SetDirection(i DIRECTION) {
	r[8] = byte(uint(i))
}
func (r *NF_RULE) GetLocalPort() uint16 {
	return binary.LittleEndian.Uint16(r[9:11])

}
func (r *NF_RULE) SetLocalPort(i uint16) {
	binary.LittleEndian.PutUint16(r[9:11], i)
}
func (r *NF_RULE) GetRemotePort() uint16 {
	return binary.LittleEndian.Uint16(r[11:13])

}
func (r *NF_RULE) SetRemotePort(i uint16) {
	binary.LittleEndian.PutUint16(r[11:13], i)
}
func (r *NF_RULE) GetIpFamily() uint16 {
	return binary.LittleEndian.Uint16(r[13:15])

}
func (r *NF_RULE) SetIpFamily(i uint16) {
	binary.LittleEndian.PutUint16(r[13:15], i)
}
func (r *NF_RULE) GetLocalIpAddress() []byte {
	return r[15 : 15+MAX_IP_ADDRESS_LENGTH]

}
func (r *NF_RULE) GetLocalIpAddressMask() []byte {
	return r[31 : 31+MAX_IP_ADDRESS_LENGTH]

}
func (r *NF_RULE) GetRemoteIpAddress() []byte {
	return r[47 : 47+MAX_IP_ADDRESS_LENGTH]

}
func (r *NF_RULE) GetRemoteIpAddressMask() []byte {
	return r[63 : 63+MAX_IP_ADDRESS_LENGTH]
}
func (r NF_RULE) GetFilteringFlag() uint32 {
	return binary.LittleEndian.Uint32(r[79:])
}
func (r NF_RULE) SetFilteringFlag(i uint32) {
	binary.LittleEndian.PutUint32(r[79:], i)
}

// NF_PORT_RANGE
type NF_PORT_RANGE struct {
	valueLow  [2]byte
	valueHigh [2]byte
}

func (n *NF_PORT_RANGE) GetValueLow() uint16 {
	return binary.LittleEndian.Uint16(n.valueLow[:])
}
func (n *NF_PORT_RANGE) GetValueHigh() uint16 {
	return binary.LittleEndian.Uint16(n.valueHigh[:])
}
func (n *NF_PORT_RANGE) SetValueLow(i uint16) {
	binary.LittleEndian.PutUint16(n.valueLow[:], i)
}
func (n *NF_PORT_RANGE) SetValueHigh(i uint16) {
	binary.LittleEndian.PutUint16(n.valueHigh[:], i)
}

//NF_RULE_EX
type NF_RULE_EX struct {
	NF_RULE
	processName         [520]byte
	LocalPortRange      NF_PORT_RANGE
	RemotePortRange     NF_PORT_RANGE
	redirectTo          [28]byte
	localProxyProcessId [4]byte
}

func (n *NF_RULE_EX) GetProcessName() string {
	//dec := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	return syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(&n.processName[0])))
}
func (n *NF_RULE_EX) SetProcessName(s string) {
	//dec := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	var si, _ = syscall.UTF16FromString(s)
	l := len(si) * 2
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&si))
	sh.Cap = l
	sh.Len = l
	copy(n.processName[:], *(*[]byte)(unsafe.Pointer(&sh)))

}
func (n *NF_RULE_EX) GetRedirectTo() []byte {
	return n.redirectTo[:]
}
func (n *NF_RULE_EX) GetLocalProxyProcessId() uint32 {
	return binary.LittleEndian.Uint32(n.localProxyProcessId[:])
}

/**
*	TCP connection properties UNALIGNED
**/
type NF_TCP_CONN_INFO struct {
	FilteringFlag FILTERING_FLAG
	ProcessId     uint32
	Direction     uint8
	IpFamily      uint16
	LocalAddress  [MAX_ADDRESS_LENGTH]byte
	RemoteAddress [MAX_ADDRESS_LENGTH]byte
}

/**
*	UDP endpoint properties UNALIGNED
**/
type NF_UDP_CONN_INFO struct {
	ProcessId    uint32
	IpFamily     uint16
	LocalAddress [MAX_ADDRESS_LENGTH]byte
}

/**
*	UDP TDI_CONNECT request properties UNALIGNED
**/
type NF_UDP_CONN_REQUEST struct {
	FilteringFlag FILTERING_FLAG
	ProcessId     uint32
	IpFamily      uint16
	LocalAddress  [MAX_ADDRESS_LENGTH]byte
	RemoteAddress [MAX_ADDRESS_LENGTH]byte
}

/**
*	UDP options UNALIGNED
**/
type NF_UDP_OPTIONS struct {
	Flags         uint32
	OptionsLength int32
	Options       [1]byte //Options of variable size
}

// IP
type NF_IP_FLAG uint32

const (
	NFIF_NONE NF_IP_FLAG = iota
	NFIF_READONLY
)

/**
*	IP options
**/
type NF_IP_PACKET_OPTIONS struct {
	ip_family         [2]byte
	ipHeaderSize      [4]byte
	compartmentId     [4]byte
	interfaceIndex    [4]byte
	subInterfaceIndex [4]byte
	flags             [4]byte
}

func (n *NF_IP_PACKET_OPTIONS) SetIpFamily(ip uint16) {
	binary.LittleEndian.PutUint16(n.ip_family[:], ip)
}
func (n *NF_IP_PACKET_OPTIONS) GetIpFamily() uint16 {
	return binary.LittleEndian.Uint16(n.ip_family[:])
}
func (n *NF_IP_PACKET_OPTIONS) SetIpHeaderSize(ip uint32) {
	binary.LittleEndian.PutUint32(n.ipHeaderSize[:], ip)
}
func (n *NF_IP_PACKET_OPTIONS) GetIpHeaderSize() uint32 {
	return binary.LittleEndian.Uint32(n.ipHeaderSize[:])
}
func (n *NF_IP_PACKET_OPTIONS) SetCompartmentId(id uint32) {
	binary.LittleEndian.PutUint32(n.compartmentId[:], id)
}
func (n *NF_IP_PACKET_OPTIONS) GetCompartmentId() uint32 {
	return binary.LittleEndian.Uint32(n.compartmentId[:])
}
func (n *NF_IP_PACKET_OPTIONS) SetInterfaceIndex(idx uint32) {
	binary.LittleEndian.PutUint32(n.interfaceIndex[:], idx)
}
func (n *NF_IP_PACKET_OPTIONS) GetInterfaceIndex() uint32 {
	return binary.LittleEndian.Uint32(n.interfaceIndex[:])
}
func (n *NF_IP_PACKET_OPTIONS) SetSubInterfaceIndex(idx uint32) {
	binary.LittleEndian.PutUint32(n.subInterfaceIndex[:], idx)
}
func (n *NF_IP_PACKET_OPTIONS) GetSubInterfaceIndex() uint32 {
	return binary.LittleEndian.Uint32(n.subInterfaceIndex[:])
}
func (n *NF_IP_PACKET_OPTIONS) SetFlags(flags uint32) {
	binary.LittleEndian.PutUint32(n.flags[:], flags)
}
func (n *NF_IP_PACKET_OPTIONS) GetFlags() uint32 {
	return binary.LittleEndian.Uint32(n.flags[:])
}

type NF_DATA struct {
	Code INT
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: darknet/darknet.proto

package darknet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DarknetFlow struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	InBytes              uint64   `protobuf:"varint,2,opt,name=in_bytes,json=inBytes,proto3" json:"in_bytes,omitempty"`
	InPkts               uint64   `protobuf:"varint,3,opt,name=in_pkts,json=inPkts,proto3" json:"in_pkts,omitempty"`
	InputPort            uint32   `protobuf:"varint,4,opt,name=input_port,json=inputPort,proto3" json:"input_port,omitempty"`
	Ipv4DstAddr          uint32   `protobuf:"varint,5,opt,name=ipv4_dst_addr,json=ipv4DstAddr,proto3" json:"ipv4_dst_addr,omitempty"`
	Ipv4SrcAddr          uint32   `protobuf:"varint,6,opt,name=ipv4_src_addr,json=ipv4SrcAddr,proto3" json:"ipv4_src_addr,omitempty"`
	L4DstPort            uint32   `protobuf:"varint,7,opt,name=l4_dst_port,json=l4DstPort,proto3" json:"l4_dst_port,omitempty"`
	L4SrcPort            uint32   `protobuf:"varint,8,opt,name=l4_src_port,json=l4SrcPort,proto3" json:"l4_src_port,omitempty"`
	OutputPort           uint32   `protobuf:"varint,9,opt,name=output_port,json=outputPort,proto3" json:"output_port,omitempty"`
	Protocol             uint32   `protobuf:"varint,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	TcpFlags             uint32   `protobuf:"varint,11,opt,name=tcp_flags,json=tcpFlags,proto3" json:"tcp_flags,omitempty"`
	VlanIn               uint32   `protobuf:"varint,12,opt,name=vlan_in,json=vlanIn,proto3" json:"vlan_in,omitempty"`
	VlanOut              uint32   `protobuf:"varint,13,opt,name=vlan_out,json=vlanOut,proto3" json:"vlan_out,omitempty"`
	SampleRate           uint32   `protobuf:"varint,14,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	PacketId             uint64   `protobuf:"varint,15,opt,name=packet_id,json=packetId,proto3" json:"packet_id,omitempty"`
	DeviceId             uint32   `protobuf:"varint,16,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DarknetFlow) Reset()         { *m = DarknetFlow{} }
func (m *DarknetFlow) String() string { return proto.CompactTextString(m) }
func (*DarknetFlow) ProtoMessage()    {}
func (*DarknetFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c79ed737219ef73, []int{0}
}

func (m *DarknetFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DarknetFlow.Unmarshal(m, b)
}
func (m *DarknetFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DarknetFlow.Marshal(b, m, deterministic)
}
func (m *DarknetFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DarknetFlow.Merge(m, src)
}
func (m *DarknetFlow) XXX_Size() int {
	return xxx_messageInfo_DarknetFlow.Size(m)
}
func (m *DarknetFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_DarknetFlow.DiscardUnknown(m)
}

var xxx_messageInfo_DarknetFlow proto.InternalMessageInfo

func (m *DarknetFlow) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DarknetFlow) GetInBytes() uint64 {
	if m != nil {
		return m.InBytes
	}
	return 0
}

func (m *DarknetFlow) GetInPkts() uint64 {
	if m != nil {
		return m.InPkts
	}
	return 0
}

func (m *DarknetFlow) GetInputPort() uint32 {
	if m != nil {
		return m.InputPort
	}
	return 0
}

func (m *DarknetFlow) GetIpv4DstAddr() uint32 {
	if m != nil {
		return m.Ipv4DstAddr
	}
	return 0
}

func (m *DarknetFlow) GetIpv4SrcAddr() uint32 {
	if m != nil {
		return m.Ipv4SrcAddr
	}
	return 0
}

func (m *DarknetFlow) GetL4DstPort() uint32 {
	if m != nil {
		return m.L4DstPort
	}
	return 0
}

func (m *DarknetFlow) GetL4SrcPort() uint32 {
	if m != nil {
		return m.L4SrcPort
	}
	return 0
}

func (m *DarknetFlow) GetOutputPort() uint32 {
	if m != nil {
		return m.OutputPort
	}
	return 0
}

func (m *DarknetFlow) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *DarknetFlow) GetTcpFlags() uint32 {
	if m != nil {
		return m.TcpFlags
	}
	return 0
}

func (m *DarknetFlow) GetVlanIn() uint32 {
	if m != nil {
		return m.VlanIn
	}
	return 0
}

func (m *DarknetFlow) GetVlanOut() uint32 {
	if m != nil {
		return m.VlanOut
	}
	return 0
}

func (m *DarknetFlow) GetSampleRate() uint32 {
	if m != nil {
		return m.SampleRate
	}
	return 0
}

func (m *DarknetFlow) GetPacketId() uint64 {
	if m != nil {
		return m.PacketId
	}
	return 0
}

func (m *DarknetFlow) GetDeviceId() uint32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

type DarknetFlows struct {
	Flow                 []*DarknetFlow `protobuf:"bytes,1,rep,name=flow,proto3" json:"flow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DarknetFlows) Reset()         { *m = DarknetFlows{} }
func (m *DarknetFlows) String() string { return proto.CompactTextString(m) }
func (*DarknetFlows) ProtoMessage()    {}
func (*DarknetFlows) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c79ed737219ef73, []int{1}
}

func (m *DarknetFlows) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DarknetFlows.Unmarshal(m, b)
}
func (m *DarknetFlows) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DarknetFlows.Marshal(b, m, deterministic)
}
func (m *DarknetFlows) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DarknetFlows.Merge(m, src)
}
func (m *DarknetFlows) XXX_Size() int {
	return xxx_messageInfo_DarknetFlows.Size(m)
}
func (m *DarknetFlows) XXX_DiscardUnknown() {
	xxx_messageInfo_DarknetFlows.DiscardUnknown(m)
}

var xxx_messageInfo_DarknetFlows proto.InternalMessageInfo

func (m *DarknetFlows) GetFlow() []*DarknetFlow {
	if m != nil {
		return m.Flow
	}
	return nil
}

func init() {
	proto.RegisterType((*DarknetFlow)(nil), "darknet.DarknetFlow")
	proto.RegisterType((*DarknetFlows)(nil), "darknet.DarknetFlows")
}

func init() { proto.RegisterFile("darknet/darknet.proto", fileDescriptor_7c79ed737219ef73) }

var fileDescriptor_7c79ed737219ef73 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x24, 0xe4, 0xcf, 0x6c, 0x03, 0xc8, 0x02, 0x61, 0xfe, 0xaf, 0x72, 0xda, 0x53,
	0x91, 0x80, 0x03, 0x57, 0x50, 0x55, 0x29, 0x27, 0xaa, 0xed, 0x03, 0x58, 0xae, 0xed, 0x22, 0x6b,
	0x1d, 0xdb, 0xb2, 0x67, 0x53, 0xf1, 0xd4, 0xbc, 0x42, 0xe5, 0x71, 0x9b, 0xe6, 0xb4, 0xfb, 0x7d,
	0xbf, 0x4f, 0x33, 0xe3, 0xb1, 0xe1, 0x8d, 0x96, 0x69, 0xf4, 0x06, 0xbf, 0x3e, 0x7c, 0xcf, 0x63,
	0x0a, 0x18, 0xd8, 0xf2, 0x41, 0x6e, 0xff, 0xcf, 0xa0, 0xbd, 0xa8, 0xff, 0x97, 0x2e, 0xdc, 0xb1,
	0x8f, 0xb0, 0x46, 0xbb, 0x37, 0x19, 0xe5, 0x3e, 0xf2, 0xa6, 0x6b, 0xfa, 0xd9, 0xf0, 0x64, 0xb0,
	0x77, 0xb0, 0xb2, 0x5e, 0xdc, 0xfc, 0x43, 0x93, 0xf9, 0xb3, 0xae, 0xe9, 0xe7, 0xc3, 0xd2, 0xfa,
	0xdf, 0x45, 0xb2, 0xb7, 0xb0, 0xb4, 0x5e, 0xc4, 0x11, 0x33, 0x9f, 0x11, 0x59, 0x58, 0x7f, 0x35,
	0x62, 0x66, 0x9f, 0x00, 0xac, 0x8f, 0x13, 0x8a, 0x18, 0x12, 0xf2, 0x79, 0xd7, 0xf4, 0x9b, 0x61,
	0x4d, 0xce, 0x55, 0x48, 0xc8, 0xb6, 0xb0, 0xb1, 0xf1, 0xf0, 0x43, 0xe8, 0x8c, 0x42, 0x6a, 0x9d,
	0xf8, 0x73, 0x4a, 0xb4, 0xc5, 0xbc, 0xc8, 0xf8, 0x4b, 0xeb, 0x74, 0xcc, 0xe4, 0xa4, 0x6a, 0x66,
	0xf1, 0x94, 0xb9, 0x4e, 0x8a, 0x32, 0x9f, 0xa1, 0x75, 0xb5, 0x0a, 0xf5, 0x59, 0xd6, 0x3e, 0xae,
	0xd4, 0xa0, 0x3e, 0x95, 0x97, 0x0a, 0xc4, 0x57, 0x8f, 0xfc, 0x3a, 0x29, 0xe2, 0x5f, 0xa0, 0x0d,
	0x13, 0x1e, 0xe7, 0x5c, 0x13, 0x87, 0x6a, 0x51, 0xe0, 0x3d, 0xac, 0x68, 0x77, 0x2a, 0x38, 0x0e,
	0x44, 0x8f, 0x9a, 0x7d, 0x80, 0x35, 0xaa, 0x28, 0x6e, 0x9d, 0xfc, 0x9b, 0x79, 0x5b, 0x21, 0xaa,
	0x78, 0x59, 0x74, 0xd9, 0xcc, 0xc1, 0x49, 0x2f, 0xac, 0xe7, 0x67, 0x84, 0x16, 0x45, 0xee, 0x7c,
	0xd9, 0x26, 0x81, 0x30, 0x21, 0xdf, 0x10, 0xa1, 0xe0, 0x9f, 0x89, 0xa6, 0xc9, 0x72, 0x1f, 0x9d,
	0x11, 0x49, 0xa2, 0xe1, 0x2f, 0xea, 0x34, 0xd5, 0x1a, 0x24, 0x9a, 0xd2, 0x31, 0x4a, 0x35, 0x1a,
	0x14, 0x56, 0xf3, 0x97, 0xb4, 0xf0, 0x55, 0x35, 0x76, 0xba, 0x40, 0x6d, 0x0e, 0x56, 0x99, 0x02,
	0x5f, 0xd5, 0x71, 0xaa, 0xb1, 0xd3, 0xdb, 0x9f, 0x70, 0x76, 0x72, 0xe1, 0x99, 0xf5, 0x30, 0xbf,
	0x75, 0xe1, 0x8e, 0x37, 0xdd, 0xac, 0x6f, 0xbf, 0xbd, 0x3e, 0x7f, 0x7c, 0x28, 0x27, 0xa1, 0x81,
	0x12, 0x37, 0x0b, 0x3a, 0xef, 0xf7, 0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x7b, 0xee, 0x30,
	0x54, 0x02, 0x00, 0x00,
}

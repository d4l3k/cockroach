// Code generated by protoc-gen-gogo.
// source: cockroach/rpc/heartbeat.proto
// DO NOT EDIT!

/*
	Package rpc is a generated protocol buffer package.

	It is generated from these files:
		cockroach/rpc/heartbeat.proto

	It has these top-level messages:
		RemoteOffset
		PingRequest
		PingResponse
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// RemoteOffset keeps track of this client's estimate of its offset from a
// remote server. Uncertainty is the maximum error in the reading of this
// offset, so that the real offset should be in the interval
// [Offset - Uncertainty, Offset + Uncertainty]. If the last heartbeat timed
// out, Offset = 0.
//
// Offset and Uncertainty are measured using the remote clock reading technique
// described in http://se.inf.tu-dresden.de/pubs/papers/SRDS1994.pdf, page 6.
type RemoteOffset struct {
	// The estimated offset from the remote server, in nanoseconds.
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// The maximum error of the measured offset, in nanoseconds.
	Uncertainty int64 `protobuf:"varint,2,opt,name=uncertainty" json:"uncertainty"`
	// Measurement time, in nanoseconds from unix epoch.
	MeasuredAt int64 `protobuf:"varint,3,opt,name=measured_at,json=measuredAt" json:"measured_at"`
}

func (m *RemoteOffset) Reset()                    { *m = RemoteOffset{} }
func (*RemoteOffset) ProtoMessage()               {}
func (*RemoteOffset) Descriptor() ([]byte, []int) { return fileDescriptorHeartbeat, []int{0} }

// A PingRequest specifies the string to echo in response.
// Fields are exported so that they will be serialized in the rpc call.
type PingRequest struct {
	// Echo this string with PingResponse.
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping"`
	// The last offset the client measured with the server.
	Offset RemoteOffset `protobuf:"bytes,2,opt,name=offset" json:"offset"`
	// The address of the client.
	Addr string `protobuf:"bytes,3,opt,name=addr" json:"addr"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptorHeartbeat, []int{1} }

// A PingResponse contains the echoed ping request string.
type PingResponse struct {
	// An echo of value sent with PingRequest.
	Pong       string `protobuf:"bytes,1,opt,name=pong" json:"pong"`
	ServerTime int64  `protobuf:"varint,2,opt,name=server_time,json=serverTime" json:"server_time"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptorHeartbeat, []int{2} }

func init() {
	proto.RegisterType((*RemoteOffset)(nil), "cockroach.rpc.RemoteOffset")
	proto.RegisterType((*PingRequest)(nil), "cockroach.rpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "cockroach.rpc.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Heartbeat service

type HeartbeatClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/cockroach.rpc.Heartbeat/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Heartbeat service

type HeartbeatServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.rpc.Heartbeat/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.rpc.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Heartbeat_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorHeartbeat,
}

func (m *RemoteOffset) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoteOffset) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.Offset))
	data[i] = 0x10
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.Uncertainty))
	data[i] = 0x18
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.MeasuredAt))
	return i, nil
}

func (m *PingRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PingRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintHeartbeat(data, i, uint64(len(m.Ping)))
	i += copy(data[i:], m.Ping)
	data[i] = 0x12
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.Offset.Size()))
	n1, err := m.Offset.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintHeartbeat(data, i, uint64(len(m.Addr)))
	i += copy(data[i:], m.Addr)
	return i, nil
}

func (m *PingResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PingResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintHeartbeat(data, i, uint64(len(m.Pong)))
	i += copy(data[i:], m.Pong)
	data[i] = 0x10
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.ServerTime))
	return i, nil
}

func encodeFixed64Heartbeat(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Heartbeat(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHeartbeat(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RemoteOffset) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHeartbeat(uint64(m.Offset))
	n += 1 + sovHeartbeat(uint64(m.Uncertainty))
	n += 1 + sovHeartbeat(uint64(m.MeasuredAt))
	return n
}

func (m *PingRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ping)
	n += 1 + l + sovHeartbeat(uint64(l))
	l = m.Offset.Size()
	n += 1 + l + sovHeartbeat(uint64(l))
	l = len(m.Addr)
	n += 1 + l + sovHeartbeat(uint64(l))
	return n
}

func (m *PingResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Pong)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.ServerTime))
	return n
}

func sovHeartbeat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeartbeat(x uint64) (n int) {
	return sovHeartbeat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteOffset) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoteOffset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteOffset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uncertainty", wireType)
			}
			m.Uncertainty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Uncertainty |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasuredAt", wireType)
			}
			m.MeasuredAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MeasuredAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PingRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ping = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Offset.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PingResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pong = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			m.ServerTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ServerTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHeartbeat(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeartbeat
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHeartbeat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeartbeat
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHeartbeat(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHeartbeat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeartbeat   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorHeartbeat = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x4f, 0xce,
	0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x2f, 0x2a, 0x48, 0xd6, 0xcf, 0x48, 0x4d, 0x2c, 0x2a, 0x49,
	0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x4b, 0xeb, 0x01, 0xa5,
	0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x32, 0xfa, 0x20, 0x16, 0x44, 0x91, 0x52, 0x23, 0x23,
	0x17, 0x4f, 0x50, 0x6a, 0x6e, 0x7e, 0x49, 0xaa, 0x7f, 0x5a, 0x5a, 0x71, 0x6a, 0x89, 0x90, 0x0c,
	0x17, 0x5b, 0x3e, 0x98, 0x25, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xec, 0xc4, 0x72, 0xe2, 0x9e, 0x3c,
	0x43, 0x10, 0x54, 0x4c, 0x48, 0x8d, 0x8b, 0xbb, 0x34, 0x2f, 0x39, 0xb5, 0xa8, 0x24, 0x31, 0x33,
	0xaf, 0xa4, 0x52, 0x82, 0x09, 0x49, 0x09, 0xb2, 0x84, 0x90, 0x2a, 0x17, 0x77, 0x6e, 0x6a, 0x62,
	0x71, 0x69, 0x51, 0x6a, 0x4a, 0x7c, 0x62, 0x89, 0x04, 0x33, 0x92, 0x3a, 0x2e, 0x98, 0x84, 0x63,
	0x89, 0x15, 0xcb, 0x8c, 0x05, 0xf2, 0x0c, 0x4a, 0x35, 0x5c, 0xdc, 0x01, 0x99, 0x79, 0xe9, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0x2c, 0x05, 0x40, 0x2e, 0xd8, 0x7e, 0x4e,
	0xa8, 0x26, 0xb0, 0x88, 0x90, 0x25, 0xdc, 0x6d, 0x20, 0x8b, 0xb9, 0x8d, 0xa4, 0xf5, 0x50, 0xbc,
	0xa8, 0x87, 0xec, 0x11, 0x34, 0x87, 0x03, 0x0d, 0x4d, 0x4c, 0x49, 0x29, 0x02, 0xbb, 0x04, 0x6e,
	0x28, 0x48, 0x44, 0xc9, 0x9f, 0x8b, 0x07, 0x62, 0x7b, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xd8,
	0xfa, 0x7c, 0x0c, 0xeb, 0x81, 0x22, 0x20, 0x4f, 0x15, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0xc5, 0x97,
	0x64, 0xe6, 0xa6, 0xa2, 0x78, 0x9e, 0x0b, 0x22, 0x11, 0x02, 0x14, 0x37, 0xf2, 0xe3, 0xe2, 0xf4,
	0x80, 0x45, 0x85, 0x90, 0x23, 0x17, 0x0b, 0xc8, 0x74, 0x21, 0x29, 0x34, 0xa7, 0x22, 0x79, 0x58,
	0x4a, 0x1a, 0xab, 0x1c, 0xc4, 0x39, 0x4a, 0x0c, 0x4e, 0xb2, 0x27, 0x1e, 0xca, 0x31, 0x9c, 0x78,
	0x24, 0xc7, 0x78, 0x01, 0x88, 0x6f, 0x00, 0xf1, 0x03, 0x20, 0x9e, 0xf0, 0x58, 0x8e, 0x21, 0x8a,
	0x19, 0xa8, 0x3a, 0x82, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x33, 0x9e, 0xfd, 0x08, 0x02,
	0x00, 0x00,
}

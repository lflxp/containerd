/*
Copyright 2019 The containerd Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

/*
Package api_v1 is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	LoadImageRequest
	LoadImageResponse
*/
package api_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LoadImageRequest struct {
	// FilePath is the absolute path of docker image tarball.
	FilePath string `protobuf:"bytes,1,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
}

func (m *LoadImageRequest) Reset()                    { *m = LoadImageRequest{} }
func (*LoadImageRequest) ProtoMessage()               {}
func (*LoadImageRequest) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

func (m *LoadImageRequest) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

type LoadImageResponse struct {
	// Images have been loaded.
	Images []string `protobuf:"bytes,1,rep,name=Images" json:"Images,omitempty"`
}

func (m *LoadImageResponse) Reset()                    { *m = LoadImageResponse{} }
func (*LoadImageResponse) ProtoMessage()               {}
func (*LoadImageResponse) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *LoadImageResponse) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadImageRequest)(nil), "api.v1.LoadImageRequest")
	proto.RegisterType((*LoadImageResponse)(nil), "api.v1.LoadImageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CRIPluginService service

type CRIPluginServiceClient interface {
	// LoadImage loads a image into containerd.
	LoadImage(ctx context.Context, in *LoadImageRequest, opts ...grpc.CallOption) (*LoadImageResponse, error)
}

type cRIPluginServiceClient struct {
	cc *grpc.ClientConn
}

func NewCRIPluginServiceClient(cc *grpc.ClientConn) CRIPluginServiceClient {
	return &cRIPluginServiceClient{cc}
}

func (c *cRIPluginServiceClient) LoadImage(ctx context.Context, in *LoadImageRequest, opts ...grpc.CallOption) (*LoadImageResponse, error) {
	out := new(LoadImageResponse)
	err := grpc.Invoke(ctx, "/api.v1.CRIPluginService/LoadImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CRIPluginService service

type CRIPluginServiceServer interface {
	// LoadImage loads a image into containerd.
	LoadImage(context.Context, *LoadImageRequest) (*LoadImageResponse, error)
}

func RegisterCRIPluginServiceServer(s *grpc.Server, srv CRIPluginServiceServer) {
	s.RegisterService(&_CRIPluginService_serviceDesc, srv)
}

func _CRIPluginService_LoadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRIPluginServiceServer).LoadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.CRIPluginService/LoadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRIPluginServiceServer).LoadImage(ctx, req.(*LoadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CRIPluginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.CRIPluginService",
	HandlerType: (*CRIPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadImage",
			Handler:    _CRIPluginService_LoadImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *LoadImageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadImageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FilePath) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.FilePath)))
		i += copy(dAtA[i:], m.FilePath)
	}
	return i, nil
}

func (m *LoadImageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadImageResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Images) > 0 {
		for _, s := range m.Images {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoadImageRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *LoadImageResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Images) > 0 {
		for _, s := range m.Images {
			l = len(s)
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LoadImageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LoadImageRequest{`,
		`FilePath:` + fmt.Sprintf("%v", this.FilePath) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LoadImageResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LoadImageResponse{`,
		`Images:` + fmt.Sprintf("%v", this.Images) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LoadImageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadImageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadImageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *LoadImageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadImageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadImageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Images", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Images = append(m.Images, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipApi(dAtA[start:])
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
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x31, 0xcb, 0x0c, 0xa5, 0x74, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xd2, 0x49,
	0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xb4, 0x29, 0xe9, 0x71, 0x09, 0xf8, 0xe4, 0x27,
	0xa6, 0x78, 0xe6, 0x26, 0xa6, 0xa7, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71,
	0x71, 0xb8, 0x65, 0xe6, 0xa4, 0x06, 0x24, 0x96, 0x64, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xf9, 0x4a, 0xda, 0x5c, 0x82, 0x48, 0xea, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4,
	0xb8, 0xd8, 0xc0, 0x02, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x51, 0x18,
	0x97, 0x80, 0x73, 0x90, 0x67, 0x40, 0x4e, 0x69, 0x7a, 0x66, 0x5e, 0x70, 0x6a, 0x51, 0x59, 0x66,
	0x72, 0xaa, 0x90, 0x13, 0x17, 0x27, 0xdc, 0x00, 0x21, 0x09, 0x3d, 0x88, 0xab, 0xf5, 0xd0, 0xdd,
	0x20, 0x25, 0x89, 0x45, 0x06, 0x62, 0x9b, 0x12, 0x83, 0x93, 0xcc, 0x89, 0x87, 0x72, 0x8c, 0x37,
	0x1e, 0xca, 0x31, 0x34, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xfb, 0xcc, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x6f, 0xec, 0xf4, 0x1d, 0x01, 0x00, 0x00,
}

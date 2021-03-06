// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/3scale-istio-adapter/config/config.proto

/*
	Package config is a generated protocol buffer package.

	3scale adapter configuration

	It is generated from these files:
		mixer/adapter/3scale-istio-adapter/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

// 3scale adapter configuration
type Params struct {
	// Unique identification string for the service/api - optional - can be specified as label in workload
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Url to connect to 3scale system
	SystemUrl string `protobuf:"bytes,2,opt,name=system_url,json=systemUrl,proto3" json:"system_url,omitempty"`
	// Access token used to authenticate against the 3scale APIs
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Url to connect to 3scale backend - optional - override system config to use internal dns etc..
	BackendUrl string `protobuf:"bytes,4,opt,name=backend_url,json=backendUrl,proto3" json:"backend_url,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Params) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Params) GetSystemUrl() string {
	if m != nil {
		return m.SystemUrl
	}
	return ""
}

func (m *Params) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Params) GetBackendUrl() string {
	if m != nil {
		return m.BackendUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "adapter.threescale.config.Params")
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ServiceId != that1.ServiceId {
		return false
	}
	if this.SystemUrl != that1.SystemUrl {
		return false
	}
	if this.AccessToken != that1.AccessToken {
		return false
	}
	if this.BackendUrl != that1.BackendUrl {
		return false
	}
	return true
}
func (this *Params) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&config.Params{")
	s = append(s, "ServiceId: "+fmt.Sprintf("%#v", this.ServiceId)+",\n")
	s = append(s, "SystemUrl: "+fmt.Sprintf("%#v", this.SystemUrl)+",\n")
	s = append(s, "AccessToken: "+fmt.Sprintf("%#v", this.AccessToken)+",\n")
	s = append(s, "BackendUrl: "+fmt.Sprintf("%#v", this.BackendUrl)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ServiceId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ServiceId)))
		i += copy(dAtA[i:], m.ServiceId)
	}
	if len(m.SystemUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SystemUrl)))
		i += copy(dAtA[i:], m.SystemUrl)
	}
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	if len(m.BackendUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.BackendUrl)))
		i += copy(dAtA[i:], m.BackendUrl)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.SystemUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.BackendUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`ServiceId:` + fmt.Sprintf("%v", this.ServiceId) + `,`,
		`SystemUrl:` + fmt.Sprintf("%v", this.SystemUrl) + `,`,
		`AccessToken:` + fmt.Sprintf("%v", this.AccessToken) + `,`,
		`BackendUrl:` + fmt.Sprintf("%v", this.BackendUrl) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackendUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BackendUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/adapter/3scale-istio-adapter/config/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x73, 0x01, 0x45, 0xaa, 0xcb, 0x14, 0x31, 0x04, 0x24, 0x2e, 0x3f, 0x13, 0x4b, 0x93,
	0xa1, 0x88, 0x07, 0x60, 0x63, 0x43, 0x88, 0x2e, 0x2c, 0x91, 0xeb, 0x5c, 0x82, 0xd5, 0x24, 0xae,
	0xae, 0x0d, 0x82, 0x8d, 0x99, 0x89, 0xc7, 0xe0, 0x51, 0x18, 0x3b, 0x32, 0x12, 0xb3, 0x30, 0xf6,
	0x11, 0x50, 0x9d, 0x30, 0xd9, 0xfa, 0xce, 0x77, 0xae, 0x74, 0xc4, 0x45, 0xa3, 0x9f, 0x89, 0x73,
	0x59, 0xca, 0xa5, 0x23, 0xce, 0xa7, 0x56, 0xc9, 0x9a, 0x26, 0xda, 0x3a, 0x6d, 0x26, 0xff, 0x50,
	0x99, 0xf6, 0x5e, 0x57, 0xc3, 0x93, 0x2d, 0xd9, 0x38, 0x93, 0xec, 0x0f, 0x61, 0xe6, 0x1e, 0x98,
	0x28, 0xb4, 0xb2, 0x5e, 0x38, 0xd8, 0xab, 0x4c, 0x65, 0x82, 0x95, 0x6f, 0x7e, 0x7d, 0xe1, 0xf4,
	0x0d, 0x44, 0x7c, 0x2d, 0x59, 0x36, 0x36, 0x39, 0x14, 0xc2, 0x12, 0x3f, 0x69, 0x45, 0x85, 0x2e,
	0x53, 0x38, 0x86, 0xb3, 0xd1, 0xcd, 0x68, 0x20, 0x57, 0x65, 0x88, 0x5f, 0xac, 0xa3, 0xa6, 0x78,
	0xe4, 0x3a, 0xdd, 0x1a, 0xe2, 0x40, 0x66, 0x5c, 0x27, 0x27, 0x62, 0x57, 0x2a, 0x45, 0xd6, 0x16,
	0xce, 0x2c, 0xa8, 0x4d, 0xb7, 0x83, 0x30, 0xee, 0xd9, 0xed, 0x06, 0x25, 0x47, 0x62, 0x3c, 0x97,
	0x6a, 0x41, 0x6d, 0x19, 0x4e, 0xec, 0x04, 0x43, 0x0c, 0x68, 0xc6, 0xf5, 0xe5, 0xf9, 0xaa, 0xc3,
	0xe8, 0xab, 0xc3, 0x68, 0xdd, 0x21, 0xbc, 0x7a, 0x84, 0x0f, 0x8f, 0xf0, 0xe9, 0x11, 0x56, 0x1e,
	0xe1, 0xdb, 0x23, 0xfc, 0x7a, 0x8c, 0xd6, 0x1e, 0xe1, 0xfd, 0x07, 0xa3, 0xbb, 0xb8, 0x1f, 0x36,
	0x8f, 0xc3, 0x92, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x6f, 0x2d, 0xbe, 0x34, 0x01,
	0x00, 0x00,
}

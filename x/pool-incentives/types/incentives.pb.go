// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/pool-incentives/v1beta1/incentives.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	// minted_denom is the denomination of the coin expected to be minted by the
	// minting module. Pool-incentives module doesn’t actually mint the coin
	// itself, but rather manages the distribution of coins that matches the
	// defined minted_denom.
	MintedDenom string `protobuf:"bytes,1,opt,name=minted_denom,json=mintedDenom,proto3" json:"minted_denom,omitempty" yaml:"minted_denom"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8153bad03e553d1, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintedDenom() string {
	if m != nil {
		return m.MintedDenom
	}
	return ""
}

type LockableDurationsInfo struct {
	LockableDurations []time.Duration `protobuf:"bytes,1,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
}

func (m *LockableDurationsInfo) Reset()         { *m = LockableDurationsInfo{} }
func (m *LockableDurationsInfo) String() string { return proto.CompactTextString(m) }
func (*LockableDurationsInfo) ProtoMessage()    {}
func (*LockableDurationsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8153bad03e553d1, []int{1}
}
func (m *LockableDurationsInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockableDurationsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockableDurationsInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockableDurationsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockableDurationsInfo.Merge(m, src)
}
func (m *LockableDurationsInfo) XXX_Size() int {
	return m.Size()
}
func (m *LockableDurationsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LockableDurationsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LockableDurationsInfo proto.InternalMessageInfo

func (m *LockableDurationsInfo) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

type DistrInfo struct {
	TotalWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total_weight,json=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_weight" yaml:"total_weight"`
	Records     []DistrRecord                          `protobuf:"bytes,2,rep,name=records,proto3" json:"records"`
}

func (m *DistrInfo) Reset()         { *m = DistrInfo{} }
func (m *DistrInfo) String() string { return proto.CompactTextString(m) }
func (*DistrInfo) ProtoMessage()    {}
func (*DistrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8153bad03e553d1, []int{2}
}
func (m *DistrInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistrInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrInfo.Merge(m, src)
}
func (m *DistrInfo) XXX_Size() int {
	return m.Size()
}
func (m *DistrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DistrInfo proto.InternalMessageInfo

func (m *DistrInfo) GetRecords() []DistrRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type DistrRecord struct {
	GaugeId uint64                                 `protobuf:"varint,1,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty" yaml:"gauge_id"`
	Weight  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"weight"`
}

func (m *DistrRecord) Reset()         { *m = DistrRecord{} }
func (m *DistrRecord) String() string { return proto.CompactTextString(m) }
func (*DistrRecord) ProtoMessage()    {}
func (*DistrRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8153bad03e553d1, []int{3}
}
func (m *DistrRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistrRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistrRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistrRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrRecord.Merge(m, src)
}
func (m *DistrRecord) XXX_Size() int {
	return m.Size()
}
func (m *DistrRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DistrRecord proto.InternalMessageInfo

func (m *DistrRecord) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

type PoolToGauge struct {
	PoolId   uint64        `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	GaugeId  uint64        `protobuf:"varint,2,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty" yaml:"gauge"`
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration" yaml:"duration"`
}

func (m *PoolToGauge) Reset()         { *m = PoolToGauge{} }
func (m *PoolToGauge) String() string { return proto.CompactTextString(m) }
func (*PoolToGauge) ProtoMessage()    {}
func (*PoolToGauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8153bad03e553d1, []int{4}
}
func (m *PoolToGauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolToGauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolToGauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolToGauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolToGauge.Merge(m, src)
}
func (m *PoolToGauge) XXX_Size() int {
	return m.Size()
}
func (m *PoolToGauge) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolToGauge.DiscardUnknown(m)
}

var xxx_messageInfo_PoolToGauge proto.InternalMessageInfo

func (m *PoolToGauge) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PoolToGauge) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

func (m *PoolToGauge) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

type PoolToGauges struct {
	PoolToGauge []PoolToGauge `protobuf:"bytes,2,rep,name=pool_to_gauge,json=poolToGauge,proto3" json:"pool_to_gauge"`
}

func (m *PoolToGauges) Reset()         { *m = PoolToGauges{} }
func (m *PoolToGauges) String() string { return proto.CompactTextString(m) }
func (*PoolToGauges) ProtoMessage()    {}
func (*PoolToGauges) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8153bad03e553d1, []int{5}
}
func (m *PoolToGauges) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolToGauges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolToGauges.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolToGauges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolToGauges.Merge(m, src)
}
func (m *PoolToGauges) XXX_Size() int {
	return m.Size()
}
func (m *PoolToGauges) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolToGauges.DiscardUnknown(m)
}

var xxx_messageInfo_PoolToGauges proto.InternalMessageInfo

func (m *PoolToGauges) GetPoolToGauge() []PoolToGauge {
	if m != nil {
		return m.PoolToGauge
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.poolincentives.v1beta1.Params")
	proto.RegisterType((*LockableDurationsInfo)(nil), "osmosis.poolincentives.v1beta1.LockableDurationsInfo")
	proto.RegisterType((*DistrInfo)(nil), "osmosis.poolincentives.v1beta1.DistrInfo")
	proto.RegisterType((*DistrRecord)(nil), "osmosis.poolincentives.v1beta1.DistrRecord")
	proto.RegisterType((*PoolToGauge)(nil), "osmosis.poolincentives.v1beta1.PoolToGauge")
	proto.RegisterType((*PoolToGauges)(nil), "osmosis.poolincentives.v1beta1.PoolToGauges")
}

func init() {
	proto.RegisterFile("osmosis/pool-incentives/v1beta1/incentives.proto", fileDescriptor_a8153bad03e553d1)
}

var fileDescriptor_a8153bad03e553d1 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x6d, 0x94, 0xb4, 0xe7, 0xf0, 0xcb, 0x05, 0x91, 0x16, 0xc9, 0x8e, 0x4e, 0x02,
	0x55, 0x8a, 0x62, 0x13, 0x18, 0x90, 0xc2, 0x16, 0x05, 0x50, 0x80, 0xa1, 0xb2, 0xa8, 0x90, 0x58,
	0x22, 0x3b, 0xbe, 0x3a, 0x56, 0x6d, 0xbf, 0xc8, 0x77, 0x09, 0xf4, 0x3f, 0x40, 0x62, 0x61, 0xec,
	0xd8, 0xff, 0x83, 0x8d, 0xa9, 0x63, 0x47, 0xc4, 0x60, 0x50, 0xb2, 0x30, 0xe7, 0x2f, 0x40, 0x3e,
	0x9f, 0xc9, 0x89, 0x4a, 0x15, 0x4c, 0xb9, 0xe7, 0x77, 0xdf, 0xf7, 0x3e, 0xdf, 0xf7, 0x4e, 0xc1,
	0x0f, 0x81, 0xc5, 0xc0, 0x42, 0x66, 0x4f, 0x01, 0xa2, 0x4e, 0x98, 0x8c, 0x69, 0xc2, 0xc3, 0x39,
	0x65, 0xf6, 0xbc, 0xeb, 0x51, 0xee, 0x76, 0xed, 0xf5, 0x27, 0x6b, 0x9a, 0x02, 0x07, 0xdd, 0x90,
	0x0a, 0x2b, 0x57, 0x28, 0x59, 0x29, 0xd8, 0xbb, 0x1d, 0x40, 0x00, 0xe2, 0xaa, 0x9d, 0x9f, 0x0a,
	0xd5, 0x9e, 0x11, 0x00, 0x04, 0x11, 0xb5, 0x45, 0xe4, 0xcd, 0x8e, 0x6c, 0x7f, 0x96, 0xba, 0x3c,
	0x84, 0xa4, 0xc8, 0x93, 0x97, 0xb8, 0x76, 0xe0, 0xa6, 0x6e, 0xcc, 0xf4, 0x1e, 0x6e, 0xc4, 0x61,
	0xc2, 0xa9, 0x3f, 0xf2, 0x69, 0x02, 0x71, 0x13, 0xb5, 0xd0, 0xfe, 0x76, 0xff, 0xee, 0x2a, 0x33,
	0x77, 0x4e, 0xdc, 0x38, 0xea, 0x11, 0x35, 0x4b, 0x1c, 0xad, 0x08, 0x07, 0x79, 0xd4, 0xab, 0x9e,
	0x9e, 0x99, 0x15, 0xf2, 0x11, 0xe1, 0x3b, 0xaf, 0x61, 0x7c, 0xec, 0x7a, 0x11, 0x1d, 0xc8, 0x36,
	0x6c, 0x98, 0x1c, 0x81, 0x0e, 0x58, 0x8f, 0x64, 0x62, 0x54, 0x02, 0xb0, 0x26, 0x6a, 0x6d, 0xee,
	0x6b, 0x8f, 0x76, 0xad, 0x02, 0xd1, 0x2a, 0x11, 0xad, 0x52, 0xdb, 0xbf, 0x7f, 0x9e, 0x99, 0x95,
	0x55, 0x66, 0xee, 0x16, 0x00, 0x97, 0x4b, 0x90, 0xd3, 0x1f, 0x26, 0x72, 0x6e, 0x45, 0x7f, 0x37,
	0x25, 0x5f, 0x11, 0xde, 0x1e, 0x84, 0x8c, 0xa7, 0xa2, 0xfd, 0x04, 0x37, 0x38, 0x70, 0x37, 0x1a,
	0xbd, 0xa7, 0x61, 0x30, 0xe1, 0xd2, 0xda, 0xb3, 0xbc, 0xfa, 0xf7, 0xcc, 0x7c, 0x10, 0x84, 0x7c,
	0x32, 0xf3, 0xac, 0x31, 0xc4, 0xf6, 0x58, 0x0c, 0x59, 0xfe, 0x74, 0x98, 0x7f, 0x6c, 0xf3, 0x93,
	0x29, 0x65, 0xd6, 0x30, 0xe1, 0xeb, 0x41, 0xa8, 0xb5, 0x88, 0xa3, 0x89, 0xf0, 0xad, 0x88, 0xf4,
	0x57, 0xb8, 0x9e, 0xd2, 0x31, 0xa4, 0x3e, 0x6b, 0x6e, 0x08, 0x77, 0x6d, 0xeb, 0xea, 0xb5, 0x59,
	0x82, 0xd2, 0x11, 0x9a, 0x7e, 0x35, 0x27, 0x72, 0xca, 0x0a, 0xe4, 0x13, 0xc2, 0x9a, 0x92, 0xd6,
	0x2d, 0xbc, 0x15, 0xb8, 0xb3, 0x80, 0x8e, 0x42, 0x5f, 0x58, 0xa8, 0xf6, 0x77, 0x56, 0x99, 0x79,
	0xa3, 0x80, 0x2a, 0x33, 0xc4, 0xa9, 0x8b, 0xe3, 0xd0, 0xd7, 0x9f, 0xe3, 0x9a, 0x34, 0xbc, 0x21,
	0x0c, 0x5b, 0xff, 0x67, 0xd8, 0x91, 0xea, 0x5e, 0xf5, 0xd7, 0x99, 0x89, 0xc8, 0x17, 0x84, 0xb5,
	0x03, 0x80, 0xe8, 0x0d, 0xbc, 0xc8, 0xeb, 0xeb, 0x6d, 0x5c, 0xcf, 0x2d, 0xad, 0x61, 0xf4, 0x55,
	0x66, 0x5e, 0x2f, 0x60, 0x64, 0x82, 0x38, 0xb5, 0xfc, 0x34, 0xf4, 0xf5, 0xb6, 0x82, 0xbe, 0x21,
	0x6e, 0xdf, 0x5c, 0x65, 0x66, 0x43, 0x41, 0x57, 0xb8, 0x1d, 0xbc, 0x55, 0x6e, 0xb8, 0xb9, 0xd9,
	0x42, 0x57, 0xbf, 0x91, 0x7b, 0xf2, 0x8d, 0xc8, 0x31, 0x94, 0xc2, 0xe2, 0x65, 0xfc, 0xa9, 0x43,
	0x28, 0x6e, 0x28, 0xf0, 0x4c, 0x3f, 0xc4, 0xd7, 0x04, 0x24, 0x87, 0x91, 0x68, 0xfb, 0xaf, 0xeb,
	0x52, 0x8a, 0xc8, 0x75, 0x69, 0x53, 0xe5, 0xd3, 0xe1, 0xf9, 0xc2, 0x40, 0x17, 0x0b, 0x03, 0xfd,
	0x5c, 0x18, 0xe8, 0xf3, 0xd2, 0xa8, 0x5c, 0x2c, 0x8d, 0xca, 0xb7, 0xa5, 0x51, 0x79, 0xf7, 0x54,
	0x19, 0xba, 0xec, 0xd1, 0x89, 0x5c, 0x8f, 0x95, 0x81, 0x3d, 0xef, 0x3e, 0xb1, 0x3f, 0x5c, 0xfa,
	0x3b, 0x10, 0xdb, 0xf0, 0x6a, 0xc2, 0xf7, 0xe3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x7b,
	0x7b, 0x56, 0x36, 0x04, 0x00, 0x00,
}

func (this *DistrRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DistrRecord)
	if !ok {
		that2, ok := that.(DistrRecord)
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
	if this.GaugeId != that1.GaugeId {
		return false
	}
	if !this.Weight.Equal(that1.Weight) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintedDenom) > 0 {
		i -= len(m.MintedDenom)
		copy(dAtA[i:], m.MintedDenom)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.MintedDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LockableDurationsInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockableDurationsInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockableDurationsInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintIncentives(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DistrInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistrInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistrInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentives(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentives(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DistrRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistrRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistrRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentives(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GaugeId != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolToGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolToGauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolToGauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintIncentives(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if m.GaugeId != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolId != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolToGauges) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolToGauges) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolToGauges) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolToGauge) > 0 {
		for iNdEx := len(m.PoolToGauge) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolToGauge[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentives(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentives(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentives(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintedDenom)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	return n
}

func (m *LockableDurationsInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	return n
}

func (m *DistrInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalWeight.Size()
	n += 1 + l + sovIncentives(uint64(l))
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	return n
}

func (m *DistrRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GaugeId != 0 {
		n += 1 + sovIncentives(uint64(m.GaugeId))
	}
	l = m.Weight.Size()
	n += 1 + l + sovIncentives(uint64(l))
	return n
}

func (m *PoolToGauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovIncentives(uint64(m.PoolId))
	}
	if m.GaugeId != 0 {
		n += 1 + sovIncentives(uint64(m.GaugeId))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovIncentives(uint64(l))
	return n
}

func (m *PoolToGauges) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolToGauge) > 0 {
		for _, e := range m.PoolToGauge {
			l = e.Size()
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	return n
}

func sovIncentives(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentives(x uint64) (n int) {
	return sovIncentives(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintedDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintedDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *LockableDurationsInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LockableDurationsInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockableDurationsInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *DistrInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistrInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistrInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, DistrRecord{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *DistrRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DistrRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistrRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *PoolToGauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolToGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolToGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func (m *PoolToGauges) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolToGauges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolToGauges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolToGauge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolToGauge = append(m.PoolToGauge, PoolToGauge{})
			if err := m.PoolToGauge[len(m.PoolToGauge)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
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
func skipIncentives(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentives
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
					return 0, ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIncentives
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
			if length < 0 {
				return 0, ErrInvalidLengthIncentives
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentives
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentives
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentives        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentives          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentives = fmt.Errorf("proto: unexpected end of group")
)

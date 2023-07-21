// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/metoken/v1/genesis.proto

package metoken

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState defines the x/metoken module's genesis state.
type GenesisState struct {
	Params                Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Registry              []Index         `protobuf:"bytes,2,rep,name=registry,proto3" json:"registry"`
	Balances              []IndexBalances `protobuf:"bytes,3,rep,name=balances,proto3" json:"balances"`
	NextRebalancingTime   time.Time       `protobuf:"bytes,4,opt,name=next_rebalancing_time,json=nextRebalancingTime,proto3,stdtime" json:"next_rebalancing_time"`
	NextInterestClaimTime time.Time       `protobuf:"bytes,5,opt,name=next_interest_claim_time,json=nextInterestClaimTime,proto3,stdtime" json:"next_interest_claim_time"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df2a396d6481bf7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRegistry() []Index {
	if m != nil {
		return m.Registry
	}
	return nil
}

func (m *GenesisState) GetBalances() []IndexBalances {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *GenesisState) GetNextRebalancingTime() time.Time {
	if m != nil {
		return m.NextRebalancingTime
	}
	return time.Time{}
}

func (m *GenesisState) GetNextInterestClaimTime() time.Time {
	if m != nil {
		return m.NextInterestClaimTime
	}
	return time.Time{}
}

// IndexBalances is the state of an Index, containing its meToken supply and all underlying asset balances.
type IndexBalances struct {
	MetokenSupply types.Coin     `protobuf:"bytes,1,opt,name=metoken_supply,json=metokenSupply,proto3" json:"metoken_supply"`
	AssetBalances []AssetBalance `protobuf:"bytes,2,rep,name=asset_balances,json=assetBalances,proto3" json:"asset_balances"`
}

func (m *IndexBalances) Reset()         { *m = IndexBalances{} }
func (m *IndexBalances) String() string { return proto.CompactTextString(m) }
func (*IndexBalances) ProtoMessage()    {}
func (*IndexBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df2a396d6481bf7, []int{1}
}
func (m *IndexBalances) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexBalances.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexBalances.Merge(m, src)
}
func (m *IndexBalances) XXX_Size() int {
	return m.Size()
}
func (m *IndexBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexBalances.DiscardUnknown(m)
}

var xxx_messageInfo_IndexBalances proto.InternalMessageInfo

func (m *IndexBalances) GetMetokenSupply() types.Coin {
	if m != nil {
		return m.MetokenSupply
	}
	return types.Coin{}
}

func (m *IndexBalances) GetAssetBalances() []AssetBalance {
	if m != nil {
		return m.AssetBalances
	}
	return nil
}

// AssetBalance tracks how much of a single asset is held in leverage, reserves, fees and interest account.
type AssetBalance struct {
	Denom     string                `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Leveraged cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=leveraged,proto3,customtype=cosmossdk.io/math.Int" json:"leveraged"`
	Reserved  cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=reserved,proto3,customtype=cosmossdk.io/math.Int" json:"reserved"`
	Fees      cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=fees,proto3,customtype=cosmossdk.io/math.Int" json:"fees"`
	Interest  cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=interest,proto3,customtype=cosmossdk.io/math.Int" json:"interest"`
}

func (m *AssetBalance) Reset()         { *m = AssetBalance{} }
func (m *AssetBalance) String() string { return proto.CompactTextString(m) }
func (*AssetBalance) ProtoMessage()    {}
func (*AssetBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_5df2a396d6481bf7, []int{2}
}
func (m *AssetBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetBalance.Merge(m, src)
}
func (m *AssetBalance) XXX_Size() int {
	return m.Size()
}
func (m *AssetBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AssetBalance proto.InternalMessageInfo

func (m *AssetBalance) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "umeenetwork.umee.metoken.v1.GenesisState")
	proto.RegisterType((*IndexBalances)(nil), "umeenetwork.umee.metoken.v1.IndexBalances")
	proto.RegisterType((*AssetBalance)(nil), "umeenetwork.umee.metoken.v1.AssetBalance")
}

func init() { proto.RegisterFile("umee/metoken/v1/genesis.proto", fileDescriptor_5df2a396d6481bf7) }

var fileDescriptor_5df2a396d6481bf7 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6e, 0x13, 0x31,
	0x14, 0xc6, 0x33, 0x4d, 0x5a, 0xa5, 0xee, 0x9f, 0x85, 0x69, 0xa5, 0x21, 0xa8, 0x93, 0x2a, 0x6c,
	0x5a, 0x24, 0x6c, 0xa5, 0x88, 0x05, 0x62, 0xd5, 0x04, 0x81, 0x22, 0xb1, 0x40, 0x29, 0x42, 0x08,
	0x09, 0x45, 0x9e, 0xe4, 0x75, 0x6a, 0x25, 0xb6, 0xa3, 0xb1, 0x33, 0xa4, 0x57, 0x60, 0xd5, 0x13,
	0x70, 0x05, 0xae, 0xd1, 0x65, 0x97, 0x88, 0x45, 0x41, 0xc9, 0x45, 0x90, 0x3d, 0x9e, 0xb4, 0x08,
	0x29, 0x2d, 0xbb, 0xb1, 0xfd, 0x7d, 0x3f, 0x3f, 0x7f, 0xef, 0x0d, 0xda, 0x9b, 0x08, 0x00, 0x2a,
	0xc0, 0xa8, 0x21, 0x48, 0x9a, 0x35, 0x69, 0x02, 0x12, 0x34, 0xd7, 0x64, 0x9c, 0x2a, 0xa3, 0xf0,
	0x23, 0x7b, 0x2c, 0xc1, 0x7c, 0x51, 0xe9, 0x90, 0xd8, 0x6f, 0xe2, 0xa5, 0x24, 0x6b, 0xd6, 0xa2,
	0xbe, 0xd2, 0x42, 0x69, 0x1a, 0x33, 0x0d, 0x34, 0x6b, 0xc6, 0x60, 0x58, 0x93, 0xf6, 0x15, 0x97,
	0xb9, 0xb9, 0xb6, 0x93, 0xa8, 0x44, 0xb9, 0x4f, 0x6a, 0xbf, 0xfc, 0x6e, 0x3d, 0x51, 0x2a, 0x19,
	0x01, 0x75, 0xab, 0x78, 0x72, 0x4a, 0x0d, 0x17, 0xa0, 0x0d, 0x13, 0x63, 0x2f, 0xf8, 0xa7, 0xa4,
	0xe2, 0x4a, 0x77, 0xdc, 0xf8, 0x56, 0x46, 0x9b, 0x6f, 0xf2, 0x22, 0x4f, 0x0c, 0x33, 0x80, 0x8f,
	0xd1, 0xda, 0x98, 0xa5, 0x4c, 0xe8, 0x30, 0xd8, 0x0f, 0x0e, 0x36, 0x8e, 0x1e, 0x93, 0x25, 0x45,
	0x93, 0x77, 0x4e, 0xda, 0xaa, 0x5c, 0x5e, 0xd7, 0x4b, 0x5d, 0x6f, 0xc4, 0xaf, 0x50, 0x35, 0x85,
	0x84, 0x6b, 0x93, 0x9e, 0x87, 0x2b, 0xfb, 0xe5, 0x83, 0x8d, 0xa3, 0xc6, 0x52, 0x48, 0x47, 0x0e,
	0x60, 0xea, 0x19, 0x0b, 0x27, 0x7e, 0x8b, 0xaa, 0x31, 0x1b, 0x31, 0xd9, 0x07, 0x1d, 0x96, 0x1d,
	0xe5, 0xc9, 0x3d, 0x28, 0xde, 0x51, 0xd0, 0x0a, 0x02, 0xfe, 0x88, 0x76, 0x25, 0x4c, 0x4d, 0x2f,
	0x85, 0x7c, 0x8b, 0xcb, 0xa4, 0x67, 0xa3, 0x0a, 0x2b, 0xee, 0x95, 0x35, 0x92, 0xe7, 0x48, 0x8a,
	0x1c, 0xc9, 0xfb, 0x22, 0xc7, 0x56, 0xd5, 0xa2, 0x2e, 0x7e, 0xd5, 0x83, 0xee, 0x03, 0x8b, 0xe8,
	0xde, 0x10, 0xac, 0x06, 0x7f, 0x46, 0xa1, 0x23, 0x73, 0x69, 0x20, 0x05, 0x6d, 0x7a, 0xfd, 0x11,
	0xe3, 0x22, 0x87, 0xaf, 0xfe, 0x07, 0xdc, 0xd5, 0xd7, 0xf1, 0x90, 0xb6, 0x65, 0x58, 0x55, 0xe3,
	0x7b, 0x80, 0xb6, 0xfe, 0x7a, 0x1a, 0x7e, 0x8d, 0xb6, 0xfd, 0xb3, 0x7b, 0x7a, 0x32, 0x1e, 0x8f,
	0xce, 0x7d, 0xa7, 0x1e, 0x92, 0x7c, 0x82, 0x88, 0x9d, 0x20, 0xe2, 0x27, 0x88, 0xb4, 0x15, 0x97,
	0x3e, 0x8d, 0x2d, 0x6f, 0x3b, 0x71, 0x2e, 0xfc, 0x01, 0x6d, 0x33, 0xad, 0xc1, 0xf4, 0x16, 0x31,
	0xe7, 0xcd, 0x3a, 0x5c, 0x1a, 0xf3, 0xb1, 0xb5, 0xf8, 0x5a, 0x0a, 0x2e, 0xbb, 0xb5, 0xa7, 0x1b,
	0x5f, 0x57, 0xd0, 0xe6, 0x6d, 0x15, 0xde, 0x41, 0xab, 0x03, 0x90, 0x4a, 0xb8, 0x3a, 0xd7, 0xbb,
	0xf9, 0x02, 0xbf, 0x44, 0xeb, 0x23, 0xc8, 0x20, 0x65, 0x09, 0x0c, 0xc2, 0x15, 0x7b, 0xd2, 0xda,
	0xb3, 0xb8, 0x9f, 0xd7, 0xf5, 0xdd, 0xfc, 0x21, 0x7a, 0x30, 0x24, 0x5c, 0x51, 0xc1, 0xcc, 0x19,
	0xe9, 0x48, 0xd3, 0xbd, 0xd1, 0xe3, 0x17, 0x76, 0xc4, 0x34, 0xa4, 0x19, 0x0c, 0xc2, 0xf2, 0x7d,
	0xbc, 0x0b, 0x39, 0x6e, 0xa2, 0xca, 0x29, 0x80, 0x76, 0x8d, 0xbf, 0xd3, 0xe6, 0xa4, 0xf6, 0xb6,
	0xa2, 0xbb, 0xae, 0xa5, 0x77, 0xdf, 0x56, 0xc8, 0x5b, 0xed, 0xcb, 0x59, 0x14, 0x5c, 0xcd, 0xa2,
	0xe0, 0xf7, 0x2c, 0x0a, 0x2e, 0xe6, 0x51, 0xe9, 0x6a, 0x1e, 0x95, 0x7e, 0xcc, 0xa3, 0xd2, 0xa7,
	0xc3, 0x84, 0x9b, 0xb3, 0x49, 0x4c, 0xfa, 0x4a, 0x50, 0x1b, 0xf2, 0x53, 0x9f, 0xb8, 0x5b, 0xd0,
	0xec, 0x39, 0x9d, 0x16, 0xbf, 0x6a, 0xbc, 0xe6, 0x06, 0xe7, 0xd9, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x29, 0x77, 0x0b, 0x39, 0x5f, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.NextInterestClaimTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.NextInterestClaimTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.NextRebalancingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.NextRebalancingTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Registry) > 0 {
		for iNdEx := len(m.Registry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IndexBalances) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexBalances) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexBalances) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetBalances) > 0 {
		for iNdEx := len(m.AssetBalances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetBalances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.MetokenSupply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AssetBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Interest.Size()
		i -= size
		if _, err := m.Interest.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Fees.Size()
		i -= size
		if _, err := m.Fees.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Reserved.Size()
		i -= size
		if _, err := m.Reserved.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Leveraged.Size()
		i -= size
		if _, err := m.Leveraged.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Registry) > 0 {
		for _, e := range m.Registry {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.NextRebalancingTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.NextInterestClaimTime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *IndexBalances) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MetokenSupply.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.AssetBalances) > 0 {
		for _, e := range m.AssetBalances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *AssetBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Leveraged.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Reserved.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Fees.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Interest.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = append(m.Registry, Index{})
			if err := m.Registry[len(m.Registry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, IndexBalances{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRebalancingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.NextRebalancingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextInterestClaimTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.NextInterestClaimTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *IndexBalances) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: IndexBalances: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexBalances: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetokenSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MetokenSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetBalances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetBalances = append(m.AssetBalances, AssetBalance{})
			if err := m.AssetBalances[len(m.AssetBalances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *AssetBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: AssetBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leveraged", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Leveraged.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserved.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Interest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
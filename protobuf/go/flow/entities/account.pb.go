// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/account.proto

package entities

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

type Account struct {
	Address              []byte        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              uint64        `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Code                 []byte        `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Keys                 []*AccountKey `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf0d6f6360ab9ab0, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Account) GetKeys() []*AccountKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AccountKey struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	SignAlgo             uint32   `protobuf:"varint,3,opt,name=sign_algo,json=signAlgo,proto3" json:"sign_algo,omitempty"`
	HashAlgo             uint32   `protobuf:"varint,4,opt,name=hash_algo,json=hashAlgo,proto3" json:"hash_algo,omitempty"`
	Weight               uint32   `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	SequenceNumber       uint32   `protobuf:"varint,6,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountKey) Reset()         { *m = AccountKey{} }
func (m *AccountKey) String() string { return proto.CompactTextString(m) }
func (*AccountKey) ProtoMessage()    {}
func (*AccountKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf0d6f6360ab9ab0, []int{1}
}

func (m *AccountKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountKey.Unmarshal(m, b)
}
func (m *AccountKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountKey.Marshal(b, m, deterministic)
}
func (m *AccountKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountKey.Merge(m, src)
}
func (m *AccountKey) XXX_Size() int {
	return xxx_messageInfo_AccountKey.Size(m)
}
func (m *AccountKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountKey.DiscardUnknown(m)
}

var xxx_messageInfo_AccountKey proto.InternalMessageInfo

func (m *AccountKey) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AccountKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *AccountKey) GetSignAlgo() uint32 {
	if m != nil {
		return m.SignAlgo
	}
	return 0
}

func (m *AccountKey) GetHashAlgo() uint32 {
	if m != nil {
		return m.HashAlgo
	}
	return 0
}

func (m *AccountKey) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *AccountKey) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "entities.Account")
	proto.RegisterType((*AccountKey)(nil), "entities.AccountKey")
}

func init() { proto.RegisterFile("flow/entities/account.proto", fileDescriptor_bf0d6f6360ab9ab0) }

var fileDescriptor_bf0d6f6360ab9ab0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x6a, 0xc3, 0x30,
	0x10, 0x45, 0x71, 0xe3, 0x38, 0xc9, 0x34, 0x69, 0x41, 0x84, 0x22, 0x08, 0x05, 0x93, 0x4d, 0xbd,
	0x72, 0xa0, 0x3d, 0x41, 0xd6, 0x85, 0x2e, 0x74, 0x01, 0x23, 0xcb, 0x53, 0x5b, 0xc4, 0x95, 0x52,
	0xcb, 0x26, 0x35, 0xbd, 0x5a, 0x0f, 0x57, 0x3c, 0x8a, 0xc8, 0x4e, 0xff, 0xbd, 0x19, 0xbe, 0x18,
	0xd8, 0x7d, 0xb6, 0xf6, 0x72, 0x40, 0xd3, 0xeb, 0x5e, 0xa3, 0x3b, 0x48, 0xa5, 0xec, 0x60, 0xfa,
	0xfc, 0xdc, 0xd9, 0xde, 0xb2, 0x65, 0xe0, 0xfb, 0x5f, 0x58, 0x1c, 0xbd, 0x62, 0x1c, 0x16, 0xb2,
	0xaa, 0x3a, 0x74, 0x8e, 0x47, 0x69, 0x94, 0xad, 0x45, 0x88, 0x93, 0x29, 0x65, 0x2b, 0x8d, 0x42,
	0x7e, 0x97, 0x46, 0x59, 0x2c, 0x42, 0x64, 0x0c, 0x62, 0x65, 0x2b, 0xe4, 0x33, 0x5a, 0xa0, 0x37,
	0xcb, 0x20, 0x3e, 0xe1, 0xe8, 0x78, 0x9c, 0xce, 0xb2, 0xfb, 0xd7, 0x6d, 0x1e, 0xba, 0xf2, 0x6b,
	0xd1, 0x3b, 0x8e, 0x82, 0x26, 0xf6, 0x7f, 0x11, 0xc0, 0x0d, 0xb2, 0x2d, 0xcc, 0xb5, 0xa9, 0xf0,
	0x87, 0xea, 0x37, 0xc2, 0x07, 0xf6, 0x0c, 0x70, 0x1e, 0xca, 0x56, 0xab, 0xe2, 0x84, 0x23, 0xf5,
	0xaf, 0xc5, 0xca, 0x93, 0x69, 0x69, 0x07, 0x2b, 0xa7, 0x6b, 0x53, 0xc8, 0xb6, 0xb6, 0xf4, 0x8d,
	0x8d, 0x58, 0x4e, 0xe0, 0xd8, 0xd6, 0x76, 0x92, 0x8d, 0x74, 0x8d, 0x97, 0xb1, 0x97, 0x13, 0x20,
	0xf9, 0x04, 0xc9, 0x05, 0x75, 0xdd, 0xf4, 0x7c, 0x4e, 0xe6, 0x9a, 0xd8, 0x0b, 0x3c, 0x3a, 0xfc,
	0x1e, 0xd0, 0x28, 0x2c, 0xcc, 0xf0, 0x55, 0x62, 0xc7, 0x13, 0x1a, 0x78, 0x08, 0xf8, 0x83, 0x68,
	0x99, 0xd0, 0x31, 0xdf, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xb1, 0x00, 0xd8, 0x6b, 0x01,
	0x00, 0x00,
}
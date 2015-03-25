// Code generated by protoc-gen-go.
// source: bwmf.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	bwmf.proto

It has these top-level messages:
	Request
	Response
	MatrixShard
	Row
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// Request for block matrix data at a given epoch.
type Request struct {
	Epoch uint64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

// Response of the block matrix data, with the associated block id.
type Response struct {
	Id    int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Shard *MatrixShard `protobuf:"bytes,2,opt,name=shard" json:"shard,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetShard() *MatrixShard {
	if m != nil {
		return m.Shard
	}
	return nil
}

// sharded matrix data
type MatrixShard struct {
	Rows []*Row `protobuf:"bytes,1,rep,name=rows" json:"rows,omitempty"`
}

func (m *MatrixShard) Reset()         { *m = MatrixShard{} }
func (m *MatrixShard) String() string { return proto1.CompactTextString(m) }
func (*MatrixShard) ProtoMessage()    {}

func (m *MatrixShard) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Row struct {
	Row map[uint32]float32 `protobuf:"bytes,1,rep,name=row" json:"row,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto1.CompactTextString(m) }
func (*Row) ProtoMessage()    {}

func (m *Row) GetRow() map[uint32]float32 {
	if m != nil {
		return m.Row
	}
	return nil
}

func init() {
}

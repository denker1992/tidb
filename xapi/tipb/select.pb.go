// Code generated by protoc-gen-go.
// source: select.proto
// DO NOT EDIT!

package tipb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// KeyRange is the encoded index key range, low is closed, high is open. (low <= x < high)
type KeyRange struct {
	Low              []byte `protobuf:"bytes,1,opt,name=low" json:"low,omitempty"`
	High             []byte `protobuf:"bytes,2,opt,name=high" json:"high,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *KeyRange) GetLow() []byte {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *KeyRange) GetHigh() []byte {
	if m != nil {
		return m.High
	}
	return nil
}

// ByItem type for group by and order by.
type ByItem struct {
	Expr             *Expr  `protobuf:"bytes,1,opt,name=expr" json:"expr,omitempty"`
	Desc             *bool  `protobuf:"varint,2,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ByItem) Reset()                    { *m = ByItem{} }
func (m *ByItem) String() string            { return proto.CompactTextString(m) }
func (*ByItem) ProtoMessage()               {}
func (*ByItem) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ByItem) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *ByItem) GetDesc() bool {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return false
}

// SelectRequest works like a simplified select statement.
type SelectRequest struct {
	// transaction start timestamp.
	StartTs *int64 `protobuf:"varint,1,opt,name=start_ts" json:"start_ts,omitempty"`
	// If table_info is not null, it represents a table scan, index_info would be null.
	TableInfo *TableInfo `protobuf:"bytes,2,opt,name=table_info" json:"table_info,omitempty"`
	// If index_info is not null, it represents an index scan, table_info would be null.
	IndexInfo *IndexInfo `protobuf:"bytes,3,opt,name=index_info" json:"index_info,omitempty"`
	// fields to be selected, fields type can be column reference for simple scan.
	// or aggregation function. If no fields specified, only handle will be returned.
	Fields []*Expr `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
	// disjoint handle ranges to be scanned.
	Ranges []*KeyRange `protobuf:"bytes,5,rep,name=ranges" json:"ranges,omitempty"`
	// handle points to be looked up.
	Points [][]byte `protobuf:"bytes,6,rep,name=points" json:"points,omitempty"`
	// distinct result.
	Distinct *bool `protobuf:"varint,7,opt,name=distinct" json:"distinct,omitempty"`
	// where condition.
	Where *Expr `protobuf:"bytes,8,opt,name=where" json:"where,omitempty"`
	// group by clause.
	GroupBy []*ByItem `protobuf:"bytes,9,rep,name=group_by" json:"group_by,omitempty"`
	// having clause.
	Having *Expr `protobuf:"bytes,10,opt,name=having" json:"having,omitempty"`
	// order by clause.
	OrderBy []*ByItem `protobuf:"bytes,11,rep,name=order_by" json:"order_by,omitempty"`
	// limit the result to be returned;
	Limit            *int64 `protobuf:"varint,12,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SelectRequest) Reset()                    { *m = SelectRequest{} }
func (m *SelectRequest) String() string            { return proto.CompactTextString(m) }
func (*SelectRequest) ProtoMessage()               {}
func (*SelectRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *SelectRequest) GetStartTs() int64 {
	if m != nil && m.StartTs != nil {
		return *m.StartTs
	}
	return 0
}

func (m *SelectRequest) GetTableInfo() *TableInfo {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *SelectRequest) GetIndexInfo() *IndexInfo {
	if m != nil {
		return m.IndexInfo
	}
	return nil
}

func (m *SelectRequest) GetFields() []*Expr {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SelectRequest) GetRanges() []*KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

func (m *SelectRequest) GetPoints() [][]byte {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *SelectRequest) GetDistinct() bool {
	if m != nil && m.Distinct != nil {
		return *m.Distinct
	}
	return false
}

func (m *SelectRequest) GetWhere() *Expr {
	if m != nil {
		return m.Where
	}
	return nil
}

func (m *SelectRequest) GetGroupBy() []*ByItem {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *SelectRequest) GetHaving() *Expr {
	if m != nil {
		return m.Having
	}
	return nil
}

func (m *SelectRequest) GetOrderBy() []*ByItem {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *SelectRequest) GetLimit() int64 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

// values are all in text format.
type Row struct {
	Handle           []byte `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Data             []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *Row) GetHandle() []byte {
	if m != nil {
		return m.Handle
	}
	return nil
}

func (m *Row) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Error struct {
	Code             *int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg              *string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// Response for SelectRequest.
type SelectResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// Result rows.
	Rows             []*Row `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SelectResponse) Reset()                    { *m = SelectResponse{} }
func (m *SelectResponse) String() string            { return proto.CompactTextString(m) }
func (*SelectResponse) ProtoMessage()               {}
func (*SelectResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *SelectResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SelectResponse) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "tipb.KeyRange")
	proto.RegisterType((*ByItem)(nil), "tipb.ByItem")
	proto.RegisterType((*SelectRequest)(nil), "tipb.SelectRequest")
	proto.RegisterType((*Row)(nil), "tipb.Row")
	proto.RegisterType((*Error)(nil), "tipb.Error")
	proto.RegisterType((*SelectResponse)(nil), "tipb.SelectResponse")
}

var fileDescriptor2 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x25, 0xf5, 0xc7, 0x9c, 0x6b, 0x37, 0x2b, 0x86, 0x31, 0x35, 0x0f, 0xa5, 0xb8, 0x0c, 0xf6,
	0x14, 0x46, 0x7f, 0x42, 0x21, 0x0f, 0x61, 0x2f, 0x23, 0xdb, 0x7b, 0x50, 0xac, 0x5b, 0x5b, 0x60,
	0x4b, 0x9e, 0xa4, 0x2e, 0xcd, 0xaf, 0xdc, 0x5f, 0x9a, 0xae, 0xe4, 0x7d, 0xb0, 0xf6, 0xd1, 0xe7,
	0xe3, 0xfa, 0x1c, 0x1d, 0xa8, 0x2c, 0x0e, 0xd8, 0xba, 0xcd, 0x64, 0xb4, 0xd3, 0x75, 0xea, 0xe4,
	0x74, 0x5c, 0x5f, 0xe1, 0xf3, 0x64, 0xd0, 0x5a, 0xa9, 0x55, 0xc4, 0xd7, 0x95, 0x6d, 0x7b, 0x1c,
	0x79, 0xfc, 0x6a, 0x3e, 0x40, 0xf1, 0x19, 0xcf, 0x7b, 0xae, 0x3a, 0xac, 0x4b, 0x48, 0x06, 0x7d,
	0x62, 0x8b, 0xdb, 0xc5, 0xc7, 0xaa, 0xae, 0x20, 0xed, 0x65, 0xd7, 0xb3, 0x0b, 0xfa, 0x6a, 0x3e,
	0x41, 0xfe, 0x70, 0xde, 0x39, 0x1c, 0x6b, 0x06, 0x29, 0x9d, 0x0c, 0xaa, 0xf2, 0x1e, 0x36, 0xf4,
	0x97, 0xcd, 0xd6, 0x23, 0xe4, 0x10, 0x68, 0xdb, 0xe0, 0x28, 0x9a, 0x9f, 0x17, 0x70, 0xf9, 0x35,
	0xe4, 0xd9, 0xe3, 0xf7, 0x27, 0xb4, 0xae, 0xbe, 0x82, 0xc2, 0x3a, 0x6e, 0xdc, 0xc1, 0xd9, 0xe0,
	0x4e, 0xea, 0x3b, 0x00, 0xc7, 0x8f, 0x03, 0x1e, 0xa4, 0x7a, 0xd4, 0xc1, 0x57, 0xde, 0xbf, 0x8d,
	0x17, 0xbf, 0x11, 0xbe, 0xf3, 0x30, 0x89, 0xa4, 0x12, 0xf8, 0x1c, 0x45, 0xc9, 0xbf, 0xa2, 0x1d,
	0xe1, 0x41, 0xb4, 0x86, 0xfc, 0x51, 0xe2, 0x20, 0x2c, 0x4b, 0x6f, 0x93, 0xff, 0x72, 0xdd, 0x40,
	0x6e, 0xa8, 0x9f, 0x65, 0x59, 0xe0, 0x56, 0x91, 0xfb, 0x53, 0x7b, 0x05, 0xf9, 0xa4, 0xa5, 0xf2,
	0xa9, 0x72, 0xcf, 0x57, 0x94, 0x53, 0x48, 0xeb, 0xa4, 0x6a, 0x1d, 0x7b, 0x43, 0x5d, 0xea, 0x6b,
	0xc8, 0x4e, 0x3d, 0x1a, 0x64, 0xc5, 0x8b, 0xd2, 0x37, 0x50, 0x74, 0x46, 0x3f, 0x4d, 0x87, 0xe3,
	0x99, 0x2d, 0xc3, 0xf9, 0x2a, 0xb2, 0xf3, 0x73, 0xf9, 0x60, 0x3d, 0xff, 0x21, 0x55, 0xc7, 0xe0,
	0x35, 0xaf, 0x36, 0x02, 0x0d, 0x79, 0xcb, 0x57, 0xbc, 0x97, 0x90, 0x0d, 0x72, 0x94, 0x8e, 0x55,
	0xf4, 0x5a, 0xcd, 0x1d, 0x24, 0x7b, 0x7d, 0xa2, 0xb8, 0x3d, 0x57, 0x62, 0xc0, 0xbf, 0x43, 0x09,
	0xee, 0xf8, 0x3c, 0x54, 0x03, 0xd9, 0xd6, 0x18, 0x1d, 0xd6, 0x68, 0xb5, 0x88, 0xa2, 0x8c, 0xa6,
	0x1d, 0x6d, 0x17, 0x34, 0xcb, 0x66, 0x0b, 0xab, 0xdf, 0xcb, 0xd8, 0x49, 0x2b, 0x8b, 0x3e, 0x65,
	0x86, 0xe4, 0x9a, 0x57, 0x2d, 0xe7, 0x90, 0xe1, 0xd0, 0x7b, 0x48, 0x8d, 0x3e, 0x59, 0xef, 0xa5,
	0x84, 0xcb, 0x48, 0xf9, 0x20, 0x0f, 0xd7, 0xf0, 0xae, 0xd5, 0xe3, 0x66, 0xf2, 0xe5, 0x5a, 0x3e,
	0x79, 0x5c, 0x1c, 0x03, 0xf9, 0x65, 0xf1, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xa1, 0xaf, 0xe2,
	0x8a, 0x02, 0x00, 0x00,
}

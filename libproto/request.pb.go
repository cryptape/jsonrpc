// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package libproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BlockTag int32

const (
	BlockTag_Latest   BlockTag = 0
	BlockTag_Earliest BlockTag = 1
)

var BlockTag_name = map[int32]string{
	0: "Latest",
	1: "Earliest",
}
var BlockTag_value = map[string]int32{
	"Latest":   0,
	"Earliest": 1,
}

func (x BlockTag) String() string {
	return proto.EnumName(BlockTag_name, int32(x))
}
func (BlockTag) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type Call struct {
	From   []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To     []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Height string `protobuf:"bytes,4,opt,name=height" json:"height,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Call) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Call) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Call) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Call) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

type Request struct {
	RequestId []byte `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Req:
	//	*Request_BlockNumber
	//	*Request_BlockByHash
	//	*Request_BlockByHeight
	//	*Request_Transaction
	//	*Request_Height
	//	*Request_Peercount
	//	*Request_Call
	//	*Request_Filter
	//	*Request_TransactionReceipt
	//	*Request_TransactionCount
	//	*Request_Code
	//	*Request_NewFilter
	//	*Request_NewBlockFilter
	//	*Request_UninstallFilter
	//	*Request_FilterChanges
	//	*Request_FilterLogs
	//	*Request_UnTx
	//	*Request_BatchReq
	Req isRequest_Req `protobuf_oneof:"req"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type isRequest_Req interface {
	isRequest_Req()
}

type Request_BlockNumber struct {
	BlockNumber bool `protobuf:"varint,2,opt,name=block_number,json=blockNumber,oneof"`
}
type Request_BlockByHash struct {
	BlockByHash string `protobuf:"bytes,3,opt,name=block_by_hash,json=blockByHash,oneof"`
}
type Request_BlockByHeight struct {
	BlockByHeight string `protobuf:"bytes,4,opt,name=block_by_height,json=blockByHeight,oneof"`
}
type Request_Transaction struct {
	Transaction []byte `protobuf:"bytes,5,opt,name=transaction,proto3,oneof"`
}
type Request_Height struct {
	Height uint64 `protobuf:"varint,6,opt,name=height,oneof"`
}
type Request_Peercount struct {
	Peercount bool `protobuf:"varint,7,opt,name=peercount,oneof"`
}
type Request_Call struct {
	Call *Call `protobuf:"bytes,8,opt,name=call,oneof"`
}
type Request_Filter struct {
	Filter string `protobuf:"bytes,9,opt,name=filter,oneof"`
}
type Request_TransactionReceipt struct {
	TransactionReceipt []byte `protobuf:"bytes,10,opt,name=transaction_receipt,json=transactionReceipt,proto3,oneof"`
}
type Request_TransactionCount struct {
	TransactionCount string `protobuf:"bytes,11,opt,name=transaction_count,json=transactionCount,oneof"`
}
type Request_Code struct {
	Code string `protobuf:"bytes,12,opt,name=code,oneof"`
}
type Request_NewFilter struct {
	NewFilter string `protobuf:"bytes,13,opt,name=new_filter,json=newFilter,oneof"`
}
type Request_NewBlockFilter struct {
	NewBlockFilter bool `protobuf:"varint,14,opt,name=new_block_filter,json=newBlockFilter,oneof"`
}
type Request_UninstallFilter struct {
	UninstallFilter uint64 `protobuf:"varint,15,opt,name=uninstall_filter,json=uninstallFilter,oneof"`
}
type Request_FilterChanges struct {
	FilterChanges uint64 `protobuf:"varint,16,opt,name=filter_changes,json=filterChanges,oneof"`
}
type Request_FilterLogs struct {
	FilterLogs uint64 `protobuf:"varint,17,opt,name=filter_logs,json=filterLogs,oneof"`
}
type Request_UnTx struct {
	UnTx *UnverifiedTransaction `protobuf:"bytes,18,opt,name=un_tx,json=unTx,oneof"`
}
type Request_BatchReq struct {
	BatchReq *BatchRequest `protobuf:"bytes,19,opt,name=batch_req,json=batchReq,oneof"`
}

func (*Request_BlockNumber) isRequest_Req()        {}
func (*Request_BlockByHash) isRequest_Req()        {}
func (*Request_BlockByHeight) isRequest_Req()      {}
func (*Request_Transaction) isRequest_Req()        {}
func (*Request_Height) isRequest_Req()             {}
func (*Request_Peercount) isRequest_Req()          {}
func (*Request_Call) isRequest_Req()               {}
func (*Request_Filter) isRequest_Req()             {}
func (*Request_TransactionReceipt) isRequest_Req() {}
func (*Request_TransactionCount) isRequest_Req()   {}
func (*Request_Code) isRequest_Req()               {}
func (*Request_NewFilter) isRequest_Req()          {}
func (*Request_NewBlockFilter) isRequest_Req()     {}
func (*Request_UninstallFilter) isRequest_Req()    {}
func (*Request_FilterChanges) isRequest_Req()      {}
func (*Request_FilterLogs) isRequest_Req()         {}
func (*Request_UnTx) isRequest_Req()               {}
func (*Request_BatchReq) isRequest_Req()           {}

func (m *Request) GetReq() isRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *Request) GetRequestId() []byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *Request) GetBlockNumber() bool {
	if x, ok := m.GetReq().(*Request_BlockNumber); ok {
		return x.BlockNumber
	}
	return false
}

func (m *Request) GetBlockByHash() string {
	if x, ok := m.GetReq().(*Request_BlockByHash); ok {
		return x.BlockByHash
	}
	return ""
}

func (m *Request) GetBlockByHeight() string {
	if x, ok := m.GetReq().(*Request_BlockByHeight); ok {
		return x.BlockByHeight
	}
	return ""
}

func (m *Request) GetTransaction() []byte {
	if x, ok := m.GetReq().(*Request_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (m *Request) GetHeight() uint64 {
	if x, ok := m.GetReq().(*Request_Height); ok {
		return x.Height
	}
	return 0
}

func (m *Request) GetPeercount() bool {
	if x, ok := m.GetReq().(*Request_Peercount); ok {
		return x.Peercount
	}
	return false
}

func (m *Request) GetCall() *Call {
	if x, ok := m.GetReq().(*Request_Call); ok {
		return x.Call
	}
	return nil
}

func (m *Request) GetFilter() string {
	if x, ok := m.GetReq().(*Request_Filter); ok {
		return x.Filter
	}
	return ""
}

func (m *Request) GetTransactionReceipt() []byte {
	if x, ok := m.GetReq().(*Request_TransactionReceipt); ok {
		return x.TransactionReceipt
	}
	return nil
}

func (m *Request) GetTransactionCount() string {
	if x, ok := m.GetReq().(*Request_TransactionCount); ok {
		return x.TransactionCount
	}
	return ""
}

func (m *Request) GetCode() string {
	if x, ok := m.GetReq().(*Request_Code); ok {
		return x.Code
	}
	return ""
}

func (m *Request) GetNewFilter() string {
	if x, ok := m.GetReq().(*Request_NewFilter); ok {
		return x.NewFilter
	}
	return ""
}

func (m *Request) GetNewBlockFilter() bool {
	if x, ok := m.GetReq().(*Request_NewBlockFilter); ok {
		return x.NewBlockFilter
	}
	return false
}

func (m *Request) GetUninstallFilter() uint64 {
	if x, ok := m.GetReq().(*Request_UninstallFilter); ok {
		return x.UninstallFilter
	}
	return 0
}

func (m *Request) GetFilterChanges() uint64 {
	if x, ok := m.GetReq().(*Request_FilterChanges); ok {
		return x.FilterChanges
	}
	return 0
}

func (m *Request) GetFilterLogs() uint64 {
	if x, ok := m.GetReq().(*Request_FilterLogs); ok {
		return x.FilterLogs
	}
	return 0
}

func (m *Request) GetUnTx() *UnverifiedTransaction {
	if x, ok := m.GetReq().(*Request_UnTx); ok {
		return x.UnTx
	}
	return nil
}

func (m *Request) GetBatchReq() *BatchRequest {
	if x, ok := m.GetReq().(*Request_BatchReq); ok {
		return x.BatchReq
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_BlockNumber)(nil),
		(*Request_BlockByHash)(nil),
		(*Request_BlockByHeight)(nil),
		(*Request_Transaction)(nil),
		(*Request_Height)(nil),
		(*Request_Peercount)(nil),
		(*Request_Call)(nil),
		(*Request_Filter)(nil),
		(*Request_TransactionReceipt)(nil),
		(*Request_TransactionCount)(nil),
		(*Request_Code)(nil),
		(*Request_NewFilter)(nil),
		(*Request_NewBlockFilter)(nil),
		(*Request_UninstallFilter)(nil),
		(*Request_FilterChanges)(nil),
		(*Request_FilterLogs)(nil),
		(*Request_UnTx)(nil),
		(*Request_BatchReq)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// req
	switch x := m.Req.(type) {
	case *Request_BlockNumber:
		t := uint64(0)
		if x.BlockNumber {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Request_BlockByHash:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.BlockByHash)
	case *Request_BlockByHeight:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.BlockByHeight)
	case *Request_Transaction:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Transaction)
	case *Request_Height:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Height))
	case *Request_Peercount:
		t := uint64(0)
		if x.Peercount {
			t = 1
		}
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Request_Call:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Call); err != nil {
			return err
		}
	case *Request_Filter:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Filter)
	case *Request_TransactionReceipt:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.TransactionReceipt)
	case *Request_TransactionCount:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TransactionCount)
	case *Request_Code:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Code)
	case *Request_NewFilter:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NewFilter)
	case *Request_NewBlockFilter:
		t := uint64(0)
		if x.NewBlockFilter {
			t = 1
		}
		b.EncodeVarint(14<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Request_UninstallFilter:
		b.EncodeVarint(15<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.UninstallFilter))
	case *Request_FilterChanges:
		b.EncodeVarint(16<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.FilterChanges))
	case *Request_FilterLogs:
		b.EncodeVarint(17<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.FilterLogs))
	case *Request_UnTx:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UnTx); err != nil {
			return err
		}
	case *Request_BatchReq:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BatchReq); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Request.Req has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 2: // req.block_number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_BlockNumber{x != 0}
		return true, err
	case 3: // req.block_by_hash
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Req = &Request_BlockByHash{x}
		return true, err
	case 4: // req.block_by_height
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Req = &Request_BlockByHeight{x}
		return true, err
	case 5: // req.transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Req = &Request_Transaction{x}
		return true, err
	case 6: // req.height
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_Height{x}
		return true, err
	case 7: // req.peercount
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_Peercount{x != 0}
		return true, err
	case 8: // req.call
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Call)
		err := b.DecodeMessage(msg)
		m.Req = &Request_Call{msg}
		return true, err
	case 9: // req.filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Req = &Request_Filter{x}
		return true, err
	case 10: // req.transaction_receipt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Req = &Request_TransactionReceipt{x}
		return true, err
	case 11: // req.transaction_count
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Req = &Request_TransactionCount{x}
		return true, err
	case 12: // req.code
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Req = &Request_Code{x}
		return true, err
	case 13: // req.new_filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Req = &Request_NewFilter{x}
		return true, err
	case 14: // req.new_block_filter
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_NewBlockFilter{x != 0}
		return true, err
	case 15: // req.uninstall_filter
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_UninstallFilter{x}
		return true, err
	case 16: // req.filter_changes
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_FilterChanges{x}
		return true, err
	case 17: // req.filter_logs
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Req = &Request_FilterLogs{x}
		return true, err
	case 18: // req.un_tx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnverifiedTransaction)
		err := b.DecodeMessage(msg)
		m.Req = &Request_UnTx{msg}
		return true, err
	case 19: // req.batch_req
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BatchRequest)
		err := b.DecodeMessage(msg)
		m.Req = &Request_BatchReq{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// req
	switch x := m.Req.(type) {
	case *Request_BlockNumber:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *Request_BlockByHash:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BlockByHash)))
		n += len(x.BlockByHash)
	case *Request_BlockByHeight:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BlockByHeight)))
		n += len(x.BlockByHeight)
	case *Request_Transaction:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Transaction)))
		n += len(x.Transaction)
	case *Request_Height:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Height))
	case *Request_Peercount:
		n += proto.SizeVarint(7<<3 | proto.WireVarint)
		n += 1
	case *Request_Call:
		s := proto.Size(x.Call)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Filter:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Filter)))
		n += len(x.Filter)
	case *Request_TransactionReceipt:
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TransactionReceipt)))
		n += len(x.TransactionReceipt)
	case *Request_TransactionCount:
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TransactionCount)))
		n += len(x.TransactionCount)
	case *Request_Code:
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Code)))
		n += len(x.Code)
	case *Request_NewFilter:
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.NewFilter)))
		n += len(x.NewFilter)
	case *Request_NewBlockFilter:
		n += proto.SizeVarint(14<<3 | proto.WireVarint)
		n += 1
	case *Request_UninstallFilter:
		n += proto.SizeVarint(15<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.UninstallFilter))
	case *Request_FilterChanges:
		n += proto.SizeVarint(16<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.FilterChanges))
	case *Request_FilterLogs:
		n += proto.SizeVarint(17<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.FilterLogs))
	case *Request_UnTx:
		s := proto.Size(x.UnTx)
		n += proto.SizeVarint(18<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_BatchReq:
		s := proto.Size(x.BatchReq)
		n += proto.SizeVarint(19<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BatchRequest struct {
	NewTxRequests []*Request `protobuf:"bytes,1,rep,name=new_tx_requests,json=newTxRequests" json:"new_tx_requests,omitempty"`
}

func (m *BatchRequest) Reset()                    { *m = BatchRequest{} }
func (m *BatchRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()               {}
func (*BatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *BatchRequest) GetNewTxRequests() []*Request {
	if m != nil {
		return m.NewTxRequests
	}
	return nil
}

func init() {
	proto.RegisterType((*Call)(nil), "Call")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*BatchRequest)(nil), "BatchRequest")
	proto.RegisterEnum("BlockTag", BlockTag_name, BlockTag_value)
}

func init() { proto.RegisterFile("request.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x93, 0x61, 0x6f, 0xd3, 0x3e,
	0x10, 0xc6, 0x93, 0x2d, 0xed, 0x92, 0x4b, 0xd3, 0x66, 0xde, 0x5f, 0x93, 0xf5, 0x47, 0x40, 0x29,
	0x93, 0x88, 0x06, 0xab, 0x60, 0x7c, 0x01, 0xd4, 0x09, 0x14, 0xa4, 0x89, 0x17, 0x51, 0xe1, 0x6d,
	0xe4, 0xa4, 0x5e, 0x13, 0x91, 0x39, 0x9b, 0xe3, 0xd0, 0xee, 0x7b, 0xf2, 0x81, 0x90, 0xed, 0xdb,
	0x96, 0x77, 0xb9, 0xe7, 0x7e, 0xf6, 0x3d, 0xf7, 0xc8, 0x81, 0x48, 0xf2, 0xfb, 0x9e, 0x77, 0x6a,
	0x79, 0x27, 0x5b, 0xd5, 0xfe, 0x1f, 0x17, 0x4d, 0x5b, 0xfe, 0x2e, 0x2b, 0x56, 0x0b, 0xab, 0x2c,
	0x7e, 0x81, 0x77, 0xc5, 0x9a, 0x86, 0x10, 0xf0, 0x6e, 0x64, 0x7b, 0x4b, 0xdd, 0xb9, 0x9b, 0x4c,
	0x32, 0xf3, 0x4d, 0xa6, 0x70, 0xa0, 0x5a, 0x7a, 0x60, 0x94, 0x03, 0xd5, 0x6a, 0x66, 0xc3, 0x14,
	0xa3, 0x87, 0x96, 0xd1, 0xdf, 0xe4, 0x14, 0xc6, 0x15, 0xaf, 0xb7, 0x95, 0xa2, 0xde, 0xdc, 0x4d,
	0x82, 0x0c, 0xab, 0xc5, 0xdf, 0x11, 0x1c, 0x65, 0x76, 0x36, 0x79, 0x09, 0x80, 0x36, 0xf2, 0x7a,
	0x83, 0x13, 0x02, 0x54, 0xbe, 0x6f, 0xc8, 0x5b, 0x98, 0x18, 0x5b, 0xb9, 0xe8, 0x6f, 0x0b, 0x2e,
	0xcd, 0x40, 0x3f, 0x75, 0xb2, 0xd0, 0xa8, 0x3f, 0x8c, 0x48, 0xce, 0x20, 0xb2, 0x50, 0xf1, 0x90,
	0x57, 0xac, 0xab, 0x8c, 0x89, 0xe0, 0x89, 0x5a, 0x3d, 0xa4, 0xac, 0xab, 0x48, 0x02, 0xb3, 0x67,
	0x6a, 0x60, 0x2b, 0x75, 0xb2, 0xe8, 0x91, 0x33, 0x32, 0x59, 0x40, 0xa8, 0x24, 0x13, 0x1d, 0x2b,
	0x55, 0xdd, 0x0a, 0x3a, 0xd2, 0xa6, 0xf4, 0x6d, 0x03, 0x91, 0xd0, 0xa7, 0xdd, 0xc6, 0x73, 0x37,
	0xf1, 0x52, 0xe7, 0x71, 0x3b, 0xf2, 0x0a, 0x82, 0x3b, 0xce, 0x65, 0xd9, 0xf6, 0x42, 0xd1, 0x23,
	0xf4, 0xfb, 0x2c, 0x91, 0x17, 0xe0, 0x95, 0xac, 0x69, 0xa8, 0x3f, 0x77, 0x93, 0xf0, 0x72, 0xb4,
	0xd4, 0x11, 0xa7, 0x4e, 0x66, 0x44, 0x7d, 0xed, 0x4d, 0xdd, 0x28, 0x2e, 0x69, 0x80, 0xde, 0xb0,
	0x26, 0x9f, 0xe0, 0x64, 0x30, 0x3f, 0x97, 0xbc, 0xe4, 0xf5, 0x9d, 0xa2, 0x80, 0xe6, 0xc8, 0xa0,
	0x99, 0xd9, 0x1e, 0xb9, 0x80, 0xe3, 0xe1, 0x11, 0xeb, 0x28, 0xc4, 0x7b, 0xe3, 0x41, 0xeb, 0xca,
	0x18, 0xfb, 0x0f, 0xbc, 0xb2, 0xdd, 0x70, 0x3a, 0x41, 0xc2, 0x54, 0xe4, 0x35, 0x80, 0xe0, 0xbb,
	0x1c, 0x5d, 0x45, 0xd8, 0x0b, 0x04, 0xdf, 0x7d, 0xb3, 0xc6, 0xce, 0x21, 0xd6, 0x80, 0xcd, 0x16,
	0xb1, 0x29, 0xae, 0x3d, 0x15, 0x7c, 0xb7, 0xd2, 0x0d, 0x64, 0xdf, 0x43, 0xdc, 0x8b, 0x5a, 0x74,
	0x8a, 0x35, 0xcd, 0x23, 0x3b, 0xc3, 0xfc, 0x66, 0x4f, 0x1d, 0x84, 0xdf, 0xc1, 0xd4, 0x22, 0x79,
	0x59, 0x31, 0xb1, 0xe5, 0x1d, 0x8d, 0x11, 0x8d, 0xac, 0x7e, 0x65, 0x65, 0xf2, 0x06, 0x42, 0x04,
	0x9b, 0x76, 0xdb, 0xd1, 0x63, 0xa4, 0xc0, 0x8a, 0xd7, 0xed, 0xb6, 0x23, 0x17, 0x30, 0xea, 0x45,
	0xae, 0xf6, 0x94, 0x98, 0xd4, 0x4f, 0x97, 0x3f, 0xc5, 0x1f, 0x2e, 0xeb, 0x9b, 0x9a, 0x6f, 0xd6,
	0xcf, 0x39, 0xe8, 0xa5, 0x7b, 0xb1, 0xde, 0x93, 0x0f, 0x10, 0x14, 0x4c, 0x95, 0x55, 0x2e, 0xf9,
	0x3d, 0x3d, 0x31, 0x47, 0xa2, 0xe5, 0x4a, 0x2b, 0xf8, 0x6e, 0x53, 0x27, 0xf3, 0x0b, 0xac, 0x57,
	0x23, 0x38, 0x94, 0xfc, 0x7e, 0xf1, 0x05, 0x26, 0x43, 0x84, 0x7c, 0x84, 0x99, 0x0e, 0x46, 0xed,
	0x73, 0x7c, 0xcf, 0x1d, 0x75, 0xe7, 0x87, 0x49, 0x78, 0xe9, 0x2f, 0x11, 0xc9, 0x22, 0xc1, 0x77,
	0xeb, 0x3d, 0x56, 0xdd, 0xf9, 0x19, 0xf8, 0x26, 0xad, 0x35, 0xdb, 0x12, 0x80, 0xf1, 0x35, 0x53,
	0xbc, 0x53, 0xb1, 0x43, 0x26, 0xe0, 0x7f, 0x65, 0xb2, 0xa9, 0x75, 0xe5, 0x16, 0x63, 0xf3, 0x77,
	0x7e, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x16, 0xfc, 0xfc, 0x37, 0xc0, 0x03, 0x00, 0x00,
}

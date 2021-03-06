// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	v1.proto

It has these top-level messages:
	Type
	FullType
	UserFn
	DynFn
	Fn
	CustomCoder
	MultiEdge
	InjectPayload
	TransformPayload
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Kind is mostly identical to reflect.TypeKind, expect we handle certain
// types specially, such as "error".
type Type_Kind int32

const (
	Type_INVALID Type_Kind = 0
	// Primitive.
	Type_BOOL    Type_Kind = 1
	Type_INT     Type_Kind = 2
	Type_INT8    Type_Kind = 3
	Type_INT16   Type_Kind = 4
	Type_INT32   Type_Kind = 5
	Type_INT64   Type_Kind = 6
	Type_UINT    Type_Kind = 7
	Type_UINT8   Type_Kind = 8
	Type_UINT16  Type_Kind = 9
	Type_UINT32  Type_Kind = 10
	Type_UINT64  Type_Kind = 11
	Type_STRING  Type_Kind = 12
	Type_FLOAT32 Type_Kind = 13
	Type_FLOAT64 Type_Kind = 14
	// Non-primitive types.
	Type_SLICE    Type_Kind = 20
	Type_STRUCT   Type_Kind = 21
	Type_FUNC     Type_Kind = 22
	Type_CHAN     Type_Kind = 23
	Type_PTR      Type_Kind = 24
	Type_SPECIAL  Type_Kind = 25
	Type_EXTERNAL Type_Kind = 26
)

var Type_Kind_name = map[int32]string{
	0:  "INVALID",
	1:  "BOOL",
	2:  "INT",
	3:  "INT8",
	4:  "INT16",
	5:  "INT32",
	6:  "INT64",
	7:  "UINT",
	8:  "UINT8",
	9:  "UINT16",
	10: "UINT32",
	11: "UINT64",
	12: "STRING",
	13: "FLOAT32",
	14: "FLOAT64",
	20: "SLICE",
	21: "STRUCT",
	22: "FUNC",
	23: "CHAN",
	24: "PTR",
	25: "SPECIAL",
	26: "EXTERNAL",
}
var Type_Kind_value = map[string]int32{
	"INVALID":  0,
	"BOOL":     1,
	"INT":      2,
	"INT8":     3,
	"INT16":    4,
	"INT32":    5,
	"INT64":    6,
	"UINT":     7,
	"UINT8":    8,
	"UINT16":   9,
	"UINT32":   10,
	"UINT64":   11,
	"STRING":   12,
	"FLOAT32":  13,
	"FLOAT64":  14,
	"SLICE":    20,
	"STRUCT":   21,
	"FUNC":     22,
	"CHAN":     23,
	"PTR":      24,
	"SPECIAL":  25,
	"EXTERNAL": 26,
}

func (x Type_Kind) String() string {
	return proto.EnumName(Type_Kind_name, int32(x))
}
func (Type_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// ChanDir matches reflect.ChanDir.
type Type_ChanDir int32

const (
	Type_RECV Type_ChanDir = 0
	Type_SEND Type_ChanDir = 1
	Type_BOTH Type_ChanDir = 2
)

var Type_ChanDir_name = map[int32]string{
	0: "RECV",
	1: "SEND",
	2: "BOTH",
}
var Type_ChanDir_value = map[string]int32{
	"RECV": 0,
	"SEND": 1,
	"BOTH": 2,
}

func (x Type_ChanDir) String() string {
	return proto.EnumName(Type_ChanDir_name, int32(x))
}
func (Type_ChanDir) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Type_Special int32

const (
	Type_ILLEGAL Type_Special = 0
	// Go
	Type_ERROR   Type_Special = 1
	Type_CONTEXT Type_Special = 2
	Type_TYPE    Type_Special = 3
	// Beam
	Type_EVENTTIME     Type_Special = 10
	Type_KV            Type_Special = 11
	Type_COGBK         Type_Special = 13
	Type_WINDOWEDVALUE Type_Special = 14
	Type_T             Type_Special = 15
	Type_U             Type_Special = 16
	Type_V             Type_Special = 17
	Type_W             Type_Special = 18
	Type_X             Type_Special = 19
	Type_Y             Type_Special = 20
	Type_Z             Type_Special = 21
)

var Type_Special_name = map[int32]string{
	0:  "ILLEGAL",
	1:  "ERROR",
	2:  "CONTEXT",
	3:  "TYPE",
	10: "EVENTTIME",
	11: "KV",
	13: "COGBK",
	14: "WINDOWEDVALUE",
	15: "T",
	16: "U",
	17: "V",
	18: "W",
	19: "X",
	20: "Y",
	21: "Z",
}
var Type_Special_value = map[string]int32{
	"ILLEGAL":       0,
	"ERROR":         1,
	"CONTEXT":       2,
	"TYPE":          3,
	"EVENTTIME":     10,
	"KV":            11,
	"COGBK":         13,
	"WINDOWEDVALUE": 14,
	"T":             15,
	"U":             16,
	"V":             17,
	"W":             18,
	"X":             19,
	"Y":             20,
	"Z":             21,
}

func (x Type_Special) String() string {
	return proto.EnumName(Type_Special_name, int32(x))
}
func (Type_Special) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

type MultiEdge_Inbound_InputKind int32

const (
	MultiEdge_Inbound_INVALID   MultiEdge_Inbound_InputKind = 0
	MultiEdge_Inbound_MAIN      MultiEdge_Inbound_InputKind = 1
	MultiEdge_Inbound_SINGLETON MultiEdge_Inbound_InputKind = 2
	MultiEdge_Inbound_SLICE     MultiEdge_Inbound_InputKind = 3
	MultiEdge_Inbound_MAP       MultiEdge_Inbound_InputKind = 4
	MultiEdge_Inbound_MULTIMAP  MultiEdge_Inbound_InputKind = 5
	MultiEdge_Inbound_ITER      MultiEdge_Inbound_InputKind = 6
	MultiEdge_Inbound_REITER    MultiEdge_Inbound_InputKind = 7
)

var MultiEdge_Inbound_InputKind_name = map[int32]string{
	0: "INVALID",
	1: "MAIN",
	2: "SINGLETON",
	3: "SLICE",
	4: "MAP",
	5: "MULTIMAP",
	6: "ITER",
	7: "REITER",
}
var MultiEdge_Inbound_InputKind_value = map[string]int32{
	"INVALID":   0,
	"MAIN":      1,
	"SINGLETON": 2,
	"SLICE":     3,
	"MAP":       4,
	"MULTIMAP":  5,
	"ITER":      6,
	"REITER":    7,
}

func (x MultiEdge_Inbound_InputKind) String() string {
	return proto.EnumName(MultiEdge_Inbound_InputKind_name, int32(x))
}
func (MultiEdge_Inbound_InputKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0, 0}
}

// Type represents a serializable reflect.Type.
type Type struct {
	// (Required) Type kind.
	Kind Type_Kind `protobuf:"varint,1,opt,name=kind,enum=v1.Type_Kind" json:"kind,omitempty"`
	// (Optional) Element type (if SLICE, PTR or CHAN)
	Element *Type `protobuf:"bytes,2,opt,name=element" json:"element,omitempty"`
	// (Optional) Fields (if STRUCT).
	Fields []*Type_StructField `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
	// (Optional) Parameter types (if FUNC).
	ParameterTypes []*Type `protobuf:"bytes,4,rep,name=parameter_types,json=parameterTypes" json:"parameter_types,omitempty"`
	// (Optional) Return types (if FUNC).
	ReturnTypes []*Type `protobuf:"bytes,5,rep,name=return_types,json=returnTypes" json:"return_types,omitempty"`
	// (Optional) Is variadic (if FUNC).
	IsVariadic bool `protobuf:"varint,6,opt,name=is_variadic,json=isVariadic" json:"is_variadic,omitempty"`
	// (Optional) Channel direction (if CHAN).
	ChanDir Type_ChanDir `protobuf:"varint,7,opt,name=chan_dir,json=chanDir,enum=v1.Type_ChanDir" json:"chan_dir,omitempty"`
	// (Optional) Special type (if SPECIAL)
	Special Type_Special `protobuf:"varint,8,opt,name=special,enum=v1.Type_Special" json:"special,omitempty"`
	// (Optional) Key for external types.
	// External types are types that are not directly serialized using
	// the types above, but rather indirectly serialized.  The wire format
	// holds a lookup key into a registry to reify the types in a worker from a
	// registry. The main usage of external serialization is to preserve
	// methods attached to types.
	ExternalKey string `protobuf:"bytes,9,opt,name=external_key,json=externalKey" json:"external_key,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Type) GetKind() Type_Kind {
	if m != nil {
		return m.Kind
	}
	return Type_INVALID
}

func (m *Type) GetElement() *Type {
	if m != nil {
		return m.Element
	}
	return nil
}

func (m *Type) GetFields() []*Type_StructField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetParameterTypes() []*Type {
	if m != nil {
		return m.ParameterTypes
	}
	return nil
}

func (m *Type) GetReturnTypes() []*Type {
	if m != nil {
		return m.ReturnTypes
	}
	return nil
}

func (m *Type) GetIsVariadic() bool {
	if m != nil {
		return m.IsVariadic
	}
	return false
}

func (m *Type) GetChanDir() Type_ChanDir {
	if m != nil {
		return m.ChanDir
	}
	return Type_RECV
}

func (m *Type) GetSpecial() Type_Special {
	if m != nil {
		return m.Special
	}
	return Type_ILLEGAL
}

func (m *Type) GetExternalKey() string {
	if m != nil {
		return m.ExternalKey
	}
	return ""
}

// StructField matches reflect.StructField.
type Type_StructField struct {
	Name      string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PkgPath   string  `protobuf:"bytes,2,opt,name=pkg_path,json=pkgPath" json:"pkg_path,omitempty"`
	Type      *Type   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tag       string  `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Offset    int64   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Index     []int32 `protobuf:"varint,6,rep,packed,name=index" json:"index,omitempty"`
	Anonymous bool    `protobuf:"varint,7,opt,name=anonymous" json:"anonymous,omitempty"`
}

func (m *Type_StructField) Reset()                    { *m = Type_StructField{} }
func (m *Type_StructField) String() string            { return proto.CompactTextString(m) }
func (*Type_StructField) ProtoMessage()               {}
func (*Type_StructField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Type_StructField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type_StructField) GetPkgPath() string {
	if m != nil {
		return m.PkgPath
	}
	return ""
}

func (m *Type_StructField) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type_StructField) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Type_StructField) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Type_StructField) GetIndex() []int32 {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Type_StructField) GetAnonymous() bool {
	if m != nil {
		return m.Anonymous
	}
	return false
}

// FullType represents a serialized typex.FullType
type FullType struct {
	Type       *Type       `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Components []*FullType `protobuf:"bytes,2,rep,name=components" json:"components,omitempty"`
}

func (m *FullType) Reset()                    { *m = FullType{} }
func (m *FullType) String() string            { return proto.CompactTextString(m) }
func (*FullType) ProtoMessage()               {}
func (*FullType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FullType) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *FullType) GetComponents() []*FullType {
	if m != nil {
		return m.Components
	}
	return nil
}

// UserFn represents a serialized function reference. The
// implementation is notably not serialized and must be present (and
// somehow discoverable from the symbol name) on the decoding side.
type UserFn struct {
	// (Required) Symbol name of function.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Function type.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *UserFn) Reset()                    { *m = UserFn{} }
func (m *UserFn) String() string            { return proto.CompactTextString(m) }
func (*UserFn) ProtoMessage()               {}
func (*UserFn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserFn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserFn) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// DynFn represents a serialized function generator.
type DynFn struct {
	// (Required) Name of the generated function.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Type of the generated function.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Required) Input to generator.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// (Required) Symbol name of generator (of type []byte ->
	// []reflect.Value -> []reflect.Value).
	Gen string `protobuf:"bytes,4,opt,name=gen" json:"gen,omitempty"`
}

func (m *DynFn) Reset()                    { *m = DynFn{} }
func (m *DynFn) String() string            { return proto.CompactTextString(m) }
func (*DynFn) ProtoMessage()               {}
func (*DynFn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DynFn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DynFn) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DynFn) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DynFn) GetGen() string {
	if m != nil {
		return m.Gen
	}
	return ""
}

// Fn represents a serialized function reference or struct.
type Fn struct {
	// (Optional) Function reference.
	Fn *UserFn `protobuf:"bytes,1,opt,name=fn" json:"fn,omitempty"`
	// (Optional) Struct type.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Optional) JSON-serialized value, if struct.
	Opt string `protobuf:"bytes,3,opt,name=opt" json:"opt,omitempty"`
	// (Optional) Function generator, if dynamic function.
	Dynfn *DynFn `protobuf:"bytes,4,opt,name=dynfn" json:"dynfn,omitempty"`
}

func (m *Fn) Reset()                    { *m = Fn{} }
func (m *Fn) String() string            { return proto.CompactTextString(m) }
func (*Fn) ProtoMessage()               {}
func (*Fn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Fn) GetFn() *UserFn {
	if m != nil {
		return m.Fn
	}
	return nil
}

func (m *Fn) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Fn) GetOpt() string {
	if m != nil {
		return m.Opt
	}
	return ""
}

func (m *Fn) GetDynfn() *DynFn {
	if m != nil {
		return m.Dynfn
	}
	return nil
}

// CustomCoder
type CustomCoder struct {
	// (Required) Name of the coder. For informational purposes only.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Concrete type being coded.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Required) Encoding function.
	Enc *UserFn `protobuf:"bytes,3,opt,name=enc" json:"enc,omitempty"`
	// (Required) Decoding function.
	Dec *UserFn `protobuf:"bytes,4,opt,name=dec" json:"dec,omitempty"`
}

func (m *CustomCoder) Reset()                    { *m = CustomCoder{} }
func (m *CustomCoder) String() string            { return proto.CompactTextString(m) }
func (*CustomCoder) ProtoMessage()               {}
func (*CustomCoder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CustomCoder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomCoder) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CustomCoder) GetEnc() *UserFn {
	if m != nil {
		return m.Enc
	}
	return nil
}

func (m *CustomCoder) GetDec() *UserFn {
	if m != nil {
		return m.Dec
	}
	return nil
}

// MultiEdge represents a partly-serialized MultiEdge. It does not include
// node information, because runners manipulate the graph structure.
type MultiEdge struct {
	Fn       *Fn                   `protobuf:"bytes,1,opt,name=fn" json:"fn,omitempty"`
	Opcode   string                `protobuf:"bytes,4,opt,name=opcode" json:"opcode,omitempty"`
	Inbound  []*MultiEdge_Inbound  `protobuf:"bytes,2,rep,name=inbound" json:"inbound,omitempty"`
	Outbound []*MultiEdge_Outbound `protobuf:"bytes,3,rep,name=outbound" json:"outbound,omitempty"`
}

func (m *MultiEdge) Reset()                    { *m = MultiEdge{} }
func (m *MultiEdge) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge) ProtoMessage()               {}
func (*MultiEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MultiEdge) GetFn() *Fn {
	if m != nil {
		return m.Fn
	}
	return nil
}

func (m *MultiEdge) GetOpcode() string {
	if m != nil {
		return m.Opcode
	}
	return ""
}

func (m *MultiEdge) GetInbound() []*MultiEdge_Inbound {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *MultiEdge) GetOutbound() []*MultiEdge_Outbound {
	if m != nil {
		return m.Outbound
	}
	return nil
}

type MultiEdge_Inbound struct {
	Kind MultiEdge_Inbound_InputKind `protobuf:"varint,1,opt,name=kind,enum=v1.MultiEdge_Inbound_InputKind" json:"kind,omitempty"`
	Type *FullType                   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Inbound) Reset()                    { *m = MultiEdge_Inbound{} }
func (m *MultiEdge_Inbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Inbound) ProtoMessage()               {}
func (*MultiEdge_Inbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *MultiEdge_Inbound) GetKind() MultiEdge_Inbound_InputKind {
	if m != nil {
		return m.Kind
	}
	return MultiEdge_Inbound_INVALID
}

func (m *MultiEdge_Inbound) GetType() *FullType {
	if m != nil {
		return m.Type
	}
	return nil
}

type MultiEdge_Outbound struct {
	Type *FullType `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Outbound) Reset()                    { *m = MultiEdge_Outbound{} }
func (m *MultiEdge_Outbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Outbound) ProtoMessage()               {}
func (*MultiEdge_Outbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *MultiEdge_Outbound) GetType() *FullType {
	if m != nil {
		return m.Type
	}
	return nil
}

// InjectPayload is the payload for the built-in Inject function.
type InjectPayload struct {
	N int32 `protobuf:"varint,1,opt,name=n" json:"n,omitempty"`
}

func (m *InjectPayload) Reset()                    { *m = InjectPayload{} }
func (m *InjectPayload) String() string            { return proto.CompactTextString(m) }
func (*InjectPayload) ProtoMessage()               {}
func (*InjectPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *InjectPayload) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

// TransformPayload represents the full payload for transforms, both
// user defined and built-in.
type TransformPayload struct {
	// urn is included here. It is also present in the model pipeline, but
	// not when submitting through Dataflow yet.
	Urn    string         `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Edge   *MultiEdge     `protobuf:"bytes,2,opt,name=edge" json:"edge,omitempty"`
	Inject *InjectPayload `protobuf:"bytes,3,opt,name=inject" json:"inject,omitempty"`
}

func (m *TransformPayload) Reset()                    { *m = TransformPayload{} }
func (m *TransformPayload) String() string            { return proto.CompactTextString(m) }
func (*TransformPayload) ProtoMessage()               {}
func (*TransformPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TransformPayload) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *TransformPayload) GetEdge() *MultiEdge {
	if m != nil {
		return m.Edge
	}
	return nil
}

func (m *TransformPayload) GetInject() *InjectPayload {
	if m != nil {
		return m.Inject
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "v1.Type")
	proto.RegisterType((*Type_StructField)(nil), "v1.Type.StructField")
	proto.RegisterType((*FullType)(nil), "v1.FullType")
	proto.RegisterType((*UserFn)(nil), "v1.UserFn")
	proto.RegisterType((*DynFn)(nil), "v1.DynFn")
	proto.RegisterType((*Fn)(nil), "v1.Fn")
	proto.RegisterType((*CustomCoder)(nil), "v1.CustomCoder")
	proto.RegisterType((*MultiEdge)(nil), "v1.MultiEdge")
	proto.RegisterType((*MultiEdge_Inbound)(nil), "v1.MultiEdge.Inbound")
	proto.RegisterType((*MultiEdge_Outbound)(nil), "v1.MultiEdge.Outbound")
	proto.RegisterType((*InjectPayload)(nil), "v1.InjectPayload")
	proto.RegisterType((*TransformPayload)(nil), "v1.TransformPayload")
	proto.RegisterEnum("v1.Type_Kind", Type_Kind_name, Type_Kind_value)
	proto.RegisterEnum("v1.Type_ChanDir", Type_ChanDir_name, Type_ChanDir_value)
	proto.RegisterEnum("v1.Type_Special", Type_Special_name, Type_Special_value)
	proto.RegisterEnum("v1.MultiEdge_Inbound_InputKind", MultiEdge_Inbound_InputKind_name, MultiEdge_Inbound_InputKind_value)
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1091 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x0e, 0x45, 0x89, 0x87, 0x91, 0xe4, 0x6c, 0xf6, 0x77, 0xfc, 0x33, 0x46, 0x8a, 0x28, 0xbc,
	0xa9, 0xda, 0x18, 0x2a, 0x2c, 0x1b, 0x46, 0xd1, 0x3b, 0x45, 0xa2, 0x1c, 0xc2, 0x34, 0x25, 0xac,
	0x28, 0xd9, 0xe9, 0x8d, 0xc0, 0x88, 0x2b, 0x99, 0xb5, 0xb4, 0x64, 0x49, 0xca, 0x88, 0xd0, 0x57,
	0x29, 0xfa, 0x1c, 0x7d, 0x87, 0x3e, 0x54, 0x8a, 0xe5, 0x41, 0xb6, 0x5c, 0x17, 0x05, 0x72, 0xb5,
	0xb3, 0x33, 0xdf, 0xb7, 0x73, 0xd8, 0xd9, 0x21, 0x41, 0xb9, 0x3b, 0x6e, 0x85, 0x51, 0x90, 0x04,
	0xb8, 0x74, 0x77, 0xac, 0x7f, 0x91, 0xa1, 0xec, 0x6c, 0x42, 0x8a, 0xdf, 0x42, 0xf9, 0xd6, 0x67,
	0x9e, 0x26, 0x34, 0x84, 0xe6, 0x5e, 0xbb, 0xde, 0xba, 0x3b, 0x6e, 0x71, 0x7d, 0xeb, 0xc2, 0x67,
	0x1e, 0x49, 0x4d, 0x58, 0x07, 0x99, 0x2e, 0xe9, 0x8a, 0xb2, 0x44, 0x2b, 0x35, 0x84, 0x66, 0xb5,
	0xad, 0x14, 0x28, 0x52, 0x18, 0xf0, 0x11, 0x48, 0x73, 0x9f, 0x2e, 0xbd, 0x58, 0x13, 0x1b, 0x62,
	0xb3, 0xda, 0xde, 0xdf, 0x1e, 0x34, 0x4a, 0xa2, 0xf5, 0x2c, 0xe9, 0x73, 0x23, 0xc9, 0x31, 0xf8,
	0x18, 0x9e, 0x87, 0x6e, 0xe4, 0xae, 0x68, 0x42, 0xa3, 0x69, 0xb2, 0x09, 0x69, 0xac, 0x95, 0x53,
	0xda, 0xfd, 0xc9, 0x7b, 0x5b, 0x00, 0xdf, 0xc6, 0xf8, 0x1d, 0xd4, 0x22, 0x9a, 0xac, 0x23, 0x96,
	0xe3, 0x2b, 0x8f, 0xf0, 0xd5, 0xcc, 0x9a, 0x81, 0xdf, 0x40, 0xd5, 0x8f, 0xa7, 0x77, 0x6e, 0xe4,
	0xbb, 0x9e, 0x3f, 0xd3, 0xa4, 0x86, 0xd0, 0x54, 0x08, 0xf8, 0xf1, 0x24, 0xd7, 0xe0, 0x77, 0xa0,
	0xcc, 0x6e, 0x5c, 0x36, 0xf5, 0xfc, 0x48, 0x93, 0xd3, 0xcc, 0xd1, 0x36, 0xe0, 0xee, 0x8d, 0xcb,
	0x7a, 0x7e, 0x44, 0xe4, 0x59, 0x26, 0xe0, 0xef, 0x41, 0x8e, 0x43, 0x3a, 0xf3, 0xdd, 0xa5, 0xa6,
	0x3c, 0xc2, 0x8e, 0x32, 0x3d, 0x29, 0x00, 0xf8, 0x2d, 0xd4, 0xe8, 0xe7, 0x84, 0x46, 0xcc, 0x5d,
	0x4e, 0x6f, 0xe9, 0x46, 0x53, 0x1b, 0x42, 0x53, 0x25, 0xd5, 0x42, 0x77, 0x41, 0x37, 0x87, 0x7f,
	0x0a, 0x50, 0x7d, 0x50, 0x14, 0x8c, 0xa1, 0xcc, 0xdc, 0x15, 0x4d, 0x6f, 0x40, 0x25, 0xa9, 0x8c,
	0x5f, 0x81, 0x12, 0xde, 0x2e, 0xa6, 0xa1, 0x9b, 0xdc, 0xa4, 0x35, 0x57, 0x89, 0x1c, 0xde, 0x2e,
	0x86, 0x6e, 0x72, 0x83, 0x5f, 0x43, 0x99, 0x57, 0x40, 0x13, 0x1f, 0x5d, 0x45, 0xaa, 0xc5, 0x08,
	0xc4, 0xc4, 0x5d, 0x68, 0xe5, 0x94, 0xc3, 0x45, 0x7c, 0x00, 0x52, 0x30, 0x9f, 0xc7, 0x34, 0xd1,
	0x2a, 0x0d, 0xa1, 0x29, 0x92, 0x7c, 0x87, 0xf7, 0xa1, 0xe2, 0x33, 0x8f, 0x7e, 0xd6, 0xa4, 0x86,
	0xd8, 0xac, 0x90, 0x6c, 0x83, 0x5f, 0x83, 0xea, 0xb2, 0x80, 0x6d, 0x56, 0xc1, 0x3a, 0x4e, 0x2b,
	0xa3, 0x90, 0x7b, 0x85, 0xfe, 0x45, 0x80, 0x32, 0x6f, 0x0c, 0x5c, 0x05, 0xd9, 0xb4, 0x27, 0x1d,
	0xcb, 0xec, 0xa1, 0x67, 0x58, 0x81, 0xf2, 0xfb, 0xc1, 0xc0, 0x42, 0x02, 0x96, 0x41, 0x34, 0x6d,
	0x07, 0x95, 0xb8, 0xca, 0xb4, 0x9d, 0x1f, 0x91, 0x88, 0x55, 0xa8, 0x98, 0xb6, 0x73, 0x7c, 0x86,
	0xca, 0xb9, 0x78, 0xd2, 0x46, 0x95, 0x5c, 0x3c, 0x3b, 0x45, 0x12, 0x87, 0x8e, 0x39, 0x49, 0xe6,
	0xca, 0x71, 0xca, 0x52, 0x30, 0x80, 0x34, 0xce, 0x68, 0x6a, 0x21, 0x9f, 0xb4, 0x11, 0x14, 0xf2,
	0xd9, 0x29, 0xaa, 0x72, 0x79, 0xe4, 0x10, 0xd3, 0x3e, 0x47, 0x35, 0x1e, 0x4f, 0xdf, 0x1a, 0x74,
	0x38, 0xa8, 0xbe, 0xdd, 0x9c, 0x9d, 0xa2, 0x3d, 0x7e, 0xe8, 0xc8, 0x32, 0xbb, 0x06, 0xda, 0xcf,
	0x09, 0xe3, 0xae, 0x83, 0x5e, 0x72, 0xaf, 0xfd, 0xb1, 0xdd, 0x45, 0x07, 0x5c, 0xea, 0x7e, 0xe8,
	0xd8, 0xe8, 0xff, 0x3c, 0xfa, 0xa1, 0x43, 0x90, 0xc6, 0x0f, 0x18, 0x0d, 0x8d, 0xae, 0xd9, 0xb1,
	0xd0, 0x2b, 0x5c, 0x03, 0xc5, 0xb8, 0x76, 0x0c, 0x62, 0x77, 0x2c, 0x74, 0xa8, 0x7f, 0x0b, 0x72,
	0xde, 0x1f, 0x9c, 0x48, 0x8c, 0xee, 0x24, 0x2b, 0xc0, 0xc8, 0xb0, 0x7b, 0x48, 0xc8, 0x4a, 0xe1,
	0x7c, 0x40, 0x25, 0xfd, 0x0f, 0x01, 0xe4, 0xbc, 0x3b, 0xd2, 0x6a, 0x59, 0x96, 0x71, 0xde, 0xb1,
	0xd0, 0x33, 0x1e, 0x90, 0x41, 0xc8, 0x80, 0x20, 0x81, 0xeb, 0xbb, 0x03, 0xdb, 0x31, 0xae, 0xf3,
	0x92, 0x39, 0x1f, 0x87, 0x06, 0x12, 0x71, 0x1d, 0x54, 0x63, 0x62, 0xd8, 0x8e, 0x63, 0x5e, 0x1a,
	0x08, 0xb0, 0x04, 0xa5, 0x8b, 0x09, 0xaa, 0x72, 0x62, 0x77, 0x70, 0xfe, 0xfe, 0x02, 0xd5, 0xf1,
	0x0b, 0xa8, 0x5f, 0x99, 0x76, 0x6f, 0x70, 0x65, 0xf4, 0x26, 0x1d, 0x6b, 0x6c, 0xa0, 0x3d, 0x5c,
	0x01, 0xc1, 0x41, 0xcf, 0xf9, 0x32, 0x46, 0x88, 0x2f, 0x13, 0xf4, 0x82, 0x2f, 0x57, 0x08, 0xf3,
	0xe5, 0x1a, 0xfd, 0x8f, 0x2f, 0x1f, 0xd1, 0x3e, 0x5f, 0x7e, 0x46, 0x2f, 0xf5, 0x09, 0x28, 0xfd,
	0xf5, 0x72, 0x99, 0x0e, 0x81, 0xa2, 0xa7, 0x84, 0x27, 0x7b, 0xea, 0x08, 0x60, 0x16, 0xac, 0xc2,
	0x80, 0x51, 0x96, 0xc4, 0x5a, 0x29, 0x7d, 0x78, 0x35, 0x8e, 0x29, 0xf8, 0xe4, 0x81, 0x5d, 0xff,
	0x09, 0xa4, 0x71, 0x4c, 0xa3, 0x3e, 0x7b, 0xb2, 0xb1, 0x0b, 0x4f, 0xa5, 0xa7, 0x3c, 0xe9, 0x53,
	0xa8, 0xf4, 0x36, 0xec, 0x6b, 0xa8, 0x9c, 0xe1, 0xb9, 0x89, 0x9b, 0x3e, 0x8b, 0x1a, 0x49, 0x65,
	0xfe, 0x18, 0x16, 0x94, 0x15, 0x8f, 0x61, 0x41, 0x99, 0xfe, 0x2b, 0x94, 0xfa, 0x0c, 0x1f, 0x42,
	0x69, 0xce, 0xf2, 0x64, 0x81, 0x9f, 0x93, 0x05, 0x4c, 0x4a, 0x73, 0xf6, 0x1f, 0x5e, 0x10, 0x88,
	0x41, 0x98, 0xa4, 0x4e, 0x54, 0xc2, 0x45, 0xfc, 0x06, 0x2a, 0xde, 0x86, 0xcd, 0x33, 0x2f, 0xd5,
	0xb6, 0xca, 0x09, 0x69, 0x0e, 0x24, 0xd3, 0xeb, 0xbf, 0x41, 0xb5, 0xbb, 0x8e, 0x93, 0x60, 0xd5,
	0x0d, 0x3c, 0x1a, 0x7d, 0x45, 0x66, 0xaf, 0x41, 0xa4, 0x6c, 0x96, 0xbf, 0xf7, 0x87, 0xe1, 0x72,
	0x35, 0xb7, 0x7a, 0x74, 0x96, 0x7b, 0xdf, 0xb1, 0x7a, 0x74, 0xa6, 0xff, 0x2e, 0x82, 0x7a, 0xb9,
	0x5e, 0x26, 0xbe, 0xe1, 0x2d, 0x28, 0x3e, 0x78, 0x90, 0xb7, 0x94, 0x5e, 0x60, 0x96, 0x33, 0x1f,
	0x11, 0xe1, 0x2c, 0xf0, 0x68, 0x5e, 0xaa, 0x7c, 0x87, 0x7f, 0x00, 0xd9, 0x67, 0x9f, 0x82, 0x35,
	0xf3, 0xf2, 0x5b, 0x7f, 0xc9, 0x49, 0xdb, 0xf3, 0x5a, 0x66, 0x66, 0x24, 0x05, 0x0a, 0xb7, 0x41,
	0x09, 0xd6, 0x49, 0xc6, 0xc8, 0xbe, 0x03, 0x07, 0xbb, 0x8c, 0x41, 0x6e, 0x25, 0x5b, 0xdc, 0xe1,
	0x5f, 0x02, 0xc8, 0xf9, 0x41, 0xf8, 0x64, 0xe7, 0x63, 0xf4, 0xe6, 0x49, 0x6f, 0x2d, 0x93, 0x85,
	0xeb, 0xe4, 0xc1, 0xe7, 0xa9, 0xb1, 0x53, 0xbd, 0xdd, 0xc6, 0xcc, 0xda, 0xca, 0x07, 0x75, 0x4b,
	0xfa, 0xc7, 0xe8, 0xba, 0xec, 0x98, 0x36, 0x12, 0xf8, 0xa3, 0x1b, 0x99, 0xf6, 0xb9, 0x65, 0x38,
	0x03, 0x1b, 0x95, 0xee, 0xc7, 0x86, 0xc8, 0xc7, 0xc2, 0x65, 0x67, 0x88, 0xca, 0x7c, 0x12, 0x5c,
	0x8e, 0x2d, 0xc7, 0xe4, 0xbb, 0x4a, 0x3a, 0xe2, 0x1c, 0x83, 0x20, 0x89, 0xcf, 0x15, 0x62, 0xa4,
	0xb2, 0x7c, 0x78, 0x04, 0x4a, 0x91, 0xe3, 0x36, 0x30, 0xe1, 0x5f, 0x03, 0xfb, 0x06, 0xea, 0x26,
	0xfb, 0x85, 0xce, 0x92, 0xa1, 0xbb, 0x59, 0x06, 0xae, 0x87, 0x6b, 0x20, 0x64, 0x17, 0x54, 0x21,
	0x02, 0xd3, 0x23, 0x40, 0x4e, 0xe4, 0xb2, 0x78, 0x1e, 0x44, 0xab, 0x02, 0x81, 0x40, 0x5c, 0x47,
	0x2c, 0x6f, 0x1f, 0x2e, 0xf2, 0x2f, 0x38, 0xf5, 0x16, 0x45, 0xfe, 0xf5, 0x9d, 0xa2, 0x91, 0xd4,
	0x84, 0xbf, 0x03, 0xc9, 0x4f, 0xfd, 0xe4, 0x5d, 0xf4, 0x82, 0x83, 0x76, 0x3c, 0x93, 0x1c, 0xf0,
	0x49, 0x4a, 0xff, 0x11, 0x4e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x68, 0x2d, 0xd6, 0x2f,
	0x08, 0x00, 0x00,
}

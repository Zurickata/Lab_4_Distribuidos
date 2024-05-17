// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/dir_merc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Mensaje que representa a un mercenario
type Mercenario struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Id     int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Mercenario) Reset() {
	*x = Mercenario{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dir_merc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mercenario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mercenario) ProtoMessage() {}

func (x *Mercenario) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dir_merc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mercenario.ProtoReflect.Descriptor instead.
func (*Mercenario) Descriptor() ([]byte, []int) {
	return file_proto_dir_merc_proto_rawDescGZIP(), []int{0}
}

func (x *Mercenario) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Mercenario) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Mensaje que representa un piso y la decisión tomada
type Piso struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numero     int32       `protobuf:"varint,1,opt,name=numero,proto3" json:"numero,omitempty"`
	Decision   string      `protobuf:"bytes,2,opt,name=decision,proto3" json:"decision,omitempty"`
	Mercenario *Mercenario `protobuf:"bytes,3,opt,name=mercenario,proto3" json:"mercenario,omitempty"`
}

func (x *Piso) Reset() {
	*x = Piso{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dir_merc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Piso) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Piso) ProtoMessage() {}

func (x *Piso) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dir_merc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Piso.ProtoReflect.Descriptor instead.
func (*Piso) Descriptor() ([]byte, []int) {
	return file_proto_dir_merc_proto_rawDescGZIP(), []int{1}
}

func (x *Piso) GetNumero() int32 {
	if x != nil {
		return x.Numero
	}
	return 0
}

func (x *Piso) GetDecision() string {
	if x != nil {
		return x.Decision
	}
	return ""
}

func (x *Piso) GetMercenario() *Mercenario {
	if x != nil {
		return x.Mercenario
	}
	return nil
}

// Mensaje de respuesta general
type Respuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	Exito   bool   `protobuf:"varint,2,opt,name=exito,proto3" json:"exito,omitempty"`
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dir_merc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dir_merc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuesta.ProtoReflect.Descriptor instead.
func (*Respuesta) Descriptor() ([]byte, []int) {
	return file_proto_dir_merc_proto_rawDescGZIP(), []int{2}
}

func (x *Respuesta) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *Respuesta) GetExito() bool {
	if x != nil {
		return x.Exito
	}
	return false
}

// Mensaje que representa el monto en el doshbank
type Monto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monto int64 `protobuf:"varint,1,opt,name=monto,proto3" json:"monto,omitempty"`
}

func (x *Monto) Reset() {
	*x = Monto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dir_merc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monto) ProtoMessage() {}

func (x *Monto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dir_merc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monto.ProtoReflect.Descriptor instead.
func (*Monto) Descriptor() ([]byte, []int) {
	return file_proto_dir_merc_proto_rawDescGZIP(), []int{3}
}

func (x *Monto) GetMonto() int64 {
	if x != nil {
		return x.Monto
	}
	return 0
}

// Mensaje vacío utilizado para solicitudes que no requieren parámetros
type Vacio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Vacio) Reset() {
	*x = Vacio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dir_merc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vacio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vacio) ProtoMessage() {}

func (x *Vacio) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dir_merc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vacio.ProtoReflect.Descriptor instead.
func (*Vacio) Descriptor() ([]byte, []int) {
	return file_proto_dir_merc_proto_rawDescGZIP(), []int{4}
}

// Mensaje para iniciar el piso
type Iniciar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *Iniciar) Reset() {
	*x = Iniciar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dir_merc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iniciar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iniciar) ProtoMessage() {}

func (x *Iniciar) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dir_merc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iniciar.ProtoReflect.Descriptor instead.
func (*Iniciar) Descriptor() ([]byte, []int) {
	return file_proto_dir_merc_proto_rawDescGZIP(), []int{5}
}

func (x *Iniciar) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_proto_dir_merc_proto protoreflect.FileDescriptor

var file_proto_dir_merc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x5f, 0x6d, 0x65, 0x72, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x22, 0x34, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x04, 0x50, 0x69, 0x73, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x3b, 0x0a, 0x09,
	0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x05, 0x4d, 0x6f, 0x6e,
	0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x56, 0x61, 0x63, 0x69,
	0x6f, 0x22, 0x23, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x32, 0xb6, 0x02, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e,
	0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x12, 0x39, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x50,
	0x69, 0x73, 0x6f, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x2e, 0x50, 0x69, 0x73, 0x6f, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x44, 0x6f, 0x73, 0x68, 0x42, 0x61, 0x6e,
	0x6b, 0x12, 0x11, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x56,
	0x61, 0x63, 0x69, 0x6f, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x63, 0x69,
	0x61, 0x72, 0x50, 0x69, 0x73, 0x6f, 0x12, 0x11, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x2e, 0x56, 0x61, 0x63, 0x69, 0x6f, 0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x12, 0x3f,
	0x0a, 0x0e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x72, 0x4d, 0x75, 0x65, 0x72, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x4d, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x42,
	0x19, 0x5a, 0x17, 0x4c, 0x41, 0x42, 0x5f, 0x34, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55,
	0x49, 0x44, 0x4f, 0x53, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_dir_merc_proto_rawDescOnce sync.Once
	file_proto_dir_merc_proto_rawDescData = file_proto_dir_merc_proto_rawDesc
)

func file_proto_dir_merc_proto_rawDescGZIP() []byte {
	file_proto_dir_merc_proto_rawDescOnce.Do(func() {
		file_proto_dir_merc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dir_merc_proto_rawDescData)
	})
	return file_proto_dir_merc_proto_rawDescData
}

var file_proto_dir_merc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_dir_merc_proto_goTypes = []interface{}{
	(*Mercenario)(nil), // 0: mercenario.Mercenario
	(*Piso)(nil),       // 1: mercenario.Piso
	(*Respuesta)(nil),  // 2: mercenario.Respuesta
	(*Monto)(nil),      // 3: mercenario.Monto
	(*Vacio)(nil),      // 4: mercenario.Vacio
	(*Iniciar)(nil),    // 5: mercenario.Iniciar
}
var file_proto_dir_merc_proto_depIdxs = []int32{
	0, // 0: mercenario.Piso.mercenario:type_name -> mercenario.Mercenario
	0, // 1: mercenario.Director.Preparacion:input_type -> mercenario.Mercenario
	1, // 2: mercenario.Director.DecisionesPiso:input_type -> mercenario.Piso
	4, // 3: mercenario.Director.ConsultarDoshBank:input_type -> mercenario.Vacio
	4, // 4: mercenario.Director.IniciarPiso:input_type -> mercenario.Vacio
	0, // 5: mercenario.Director.InformarMuerte:input_type -> mercenario.Mercenario
	2, // 6: mercenario.Director.Preparacion:output_type -> mercenario.Respuesta
	2, // 7: mercenario.Director.DecisionesPiso:output_type -> mercenario.Respuesta
	3, // 8: mercenario.Director.ConsultarDoshBank:output_type -> mercenario.Monto
	5, // 9: mercenario.Director.IniciarPiso:output_type -> mercenario.Iniciar
	2, // 10: mercenario.Director.InformarMuerte:output_type -> mercenario.Respuesta
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_dir_merc_proto_init() }
func file_proto_dir_merc_proto_init() {
	if File_proto_dir_merc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dir_merc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mercenario); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_dir_merc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Piso); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_dir_merc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respuesta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_dir_merc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_dir_merc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vacio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_dir_merc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iniciar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_dir_merc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dir_merc_proto_goTypes,
		DependencyIndexes: file_proto_dir_merc_proto_depIdxs,
		MessageInfos:      file_proto_dir_merc_proto_msgTypes,
	}.Build()
	File_proto_dir_merc_proto = out.File
	file_proto_dir_merc_proto_rawDesc = nil
	file_proto_dir_merc_proto_goTypes = nil
	file_proto_dir_merc_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: basket.proto

package basket

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

type NewProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *NewProduct) Reset() {
	*x = NewProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProduct) ProtoMessage() {}

func (x *NewProduct) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProduct.ProtoReflect.Descriptor instead.
func (*NewProduct) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{0}
}

func (x *NewProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *NewProduct) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Quantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Quantity) Reset() {
	*x = Quantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quantity) ProtoMessage() {}

func (x *Quantity) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quantity.ProtoReflect.Descriptor instead.
func (*Quantity) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{1}
}

func (x *Quantity) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Quantity) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price    float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity int64   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Product `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{4}
}

func (x *Products) GetItems() []*Product {
	if x != nil {
		return x.Items
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{5}
}

var File_basket_proto protoreflect.FileDescriptor

var file_basket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x22, 0x47, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x45, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x31, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0xc3, 0x01, 0x0a, 0x06, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a,
	0x10, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x2f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x49, 0x64, 0x1a,
	0x0c, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x11, 0x5a,
	0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basket_proto_rawDescOnce sync.Once
	file_basket_proto_rawDescData = file_basket_proto_rawDesc
)

func file_basket_proto_rawDescGZIP() []byte {
	file_basket_proto_rawDescOnce.Do(func() {
		file_basket_proto_rawDescData = protoimpl.X.CompressGZIP(file_basket_proto_rawDescData)
	})
	return file_basket_proto_rawDescData
}

var file_basket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_basket_proto_goTypes = []interface{}{
	(*NewProduct)(nil), // 0: basket.NewProduct
	(*Quantity)(nil),   // 1: basket.Quantity
	(*Id)(nil),         // 2: basket.Id
	(*Product)(nil),    // 3: basket.Product
	(*Products)(nil),   // 4: basket.Products
	(*Void)(nil),       // 5: basket.Void
}
var file_basket_proto_depIdxs = []int32{
	3, // 0: basket.Products.items:type_name -> basket.Product
	0, // 1: basket.Basket.AddProduct:input_type -> basket.NewProduct
	5, // 2: basket.Basket.GetProducts:input_type -> basket.Void
	1, // 3: basket.Basket.UpdateProduct:input_type -> basket.Quantity
	2, // 4: basket.Basket.RemoveProduct:input_type -> basket.Id
	5, // 5: basket.Basket.AddProduct:output_type -> basket.Void
	4, // 6: basket.Basket.GetProducts:output_type -> basket.Products
	5, // 7: basket.Basket.UpdateProduct:output_type -> basket.Void
	5, // 8: basket.Basket.RemoveProduct:output_type -> basket.Void
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_basket_proto_init() }
func file_basket_proto_init() {
	if File_basket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProduct); i {
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
		file_basket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quantity); i {
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
		file_basket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_basket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_basket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
		file_basket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_basket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basket_proto_goTypes,
		DependencyIndexes: file_basket_proto_depIdxs,
		MessageInfos:      file_basket_proto_msgTypes,
	}.Build()
	File_basket_proto = out.File
	file_basket_proto_rawDesc = nil
	file_basket_proto_goTypes = nil
	file_basket_proto_depIdxs = nil
}

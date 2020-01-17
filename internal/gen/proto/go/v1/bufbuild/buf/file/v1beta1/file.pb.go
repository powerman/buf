// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bufbuild/buf/file/v1beta1/file.proto

package filev1beta1

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

// FileSet is a set of files.
type FileSet struct {
	// files are the files that make up the file set.
	//
	// All files must have unique paths for a FileSet to be valid.
	Files                []*File  `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileSet) Reset()         { *m = FileSet{} }
func (m *FileSet) String() string { return proto.CompactTextString(m) }
func (*FileSet) ProtoMessage()    {}
func (*FileSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e66da8bbfd29a12, []int{0}
}

func (m *FileSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSet.Unmarshal(m, b)
}
func (m *FileSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSet.Marshal(b, m, deterministic)
}
func (m *FileSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSet.Merge(m, src)
}
func (m *FileSet) XXX_Size() int {
	return xxx_messageInfo_FileSet.Size(m)
}
func (m *FileSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSet.DiscardUnknown(m)
}

var xxx_messageInfo_FileSet proto.InternalMessageInfo

func (m *FileSet) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// File is an individual file.
type File struct {
	// path is the path of the file.
	//
	// This path must be relative and use '/' as the separator character .
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// content is the file content.
	//
	// This can potentially be empty.
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e66da8bbfd29a12, []int{1}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *File) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// FileAnnotation is an annotation for a specific file location.
type FileAnnotation struct {
	// path is the path of the file.
	//
	// This path must be relative and use '/' as the separator character .
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// start_line is the starting line.
	//
	// If the starting line is not known, this will be 0.
	StartLine uint32 `protobuf:"varint,2,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	// start_column is the starting column.
	//
	// If the starting column is not known, this will be 0.
	StartColumn uint32 `protobuf:"varint,3,opt,name=start_column,json=startColumn,proto3" json:"start_column,omitempty"`
	// end_line is the ending line.
	//
	// If the ending line is not known, this will be 0.
	// If the ending line is the same as the starting line, this will be explicitly set
	// to the same value as start_line.
	EndLine uint32 `protobuf:"varint,4,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
	// end_column is the ending column.
	//
	// If the ending column is not known, this will be 0.
	// If the ending column is the same as the starting column, this will be explicitly set
	// to the same value as start_column.
	EndColumn uint32 `protobuf:"varint,5,opt,name=end_column,json=endColumn,proto3" json:"end_column,omitempty"`
	// type is the type of annotation, typically and ID representing a failure type.
	Type string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	// message is the message of the annotation.
	Message              string   `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileAnnotation) Reset()         { *m = FileAnnotation{} }
func (m *FileAnnotation) String() string { return proto.CompactTextString(m) }
func (*FileAnnotation) ProtoMessage()    {}
func (*FileAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e66da8bbfd29a12, []int{2}
}

func (m *FileAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileAnnotation.Unmarshal(m, b)
}
func (m *FileAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileAnnotation.Marshal(b, m, deterministic)
}
func (m *FileAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileAnnotation.Merge(m, src)
}
func (m *FileAnnotation) XXX_Size() int {
	return xxx_messageInfo_FileAnnotation.Size(m)
}
func (m *FileAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_FileAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_FileAnnotation proto.InternalMessageInfo

func (m *FileAnnotation) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileAnnotation) GetStartLine() uint32 {
	if m != nil {
		return m.StartLine
	}
	return 0
}

func (m *FileAnnotation) GetStartColumn() uint32 {
	if m != nil {
		return m.StartColumn
	}
	return 0
}

func (m *FileAnnotation) GetEndLine() uint32 {
	if m != nil {
		return m.EndLine
	}
	return 0
}

func (m *FileAnnotation) GetEndColumn() uint32 {
	if m != nil {
		return m.EndColumn
	}
	return 0
}

func (m *FileAnnotation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FileAnnotation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FileSet)(nil), "bufbuild.buf.file.v1beta1.FileSet")
	proto.RegisterType((*File)(nil), "bufbuild.buf.file.v1beta1.File")
	proto.RegisterType((*FileAnnotation)(nil), "bufbuild.buf.file.v1beta1.FileAnnotation")
}

func init() {
	proto.RegisterFile("bufbuild/buf/file/v1beta1/file.proto", fileDescriptor_7e66da8bbfd29a12)
}

var fileDescriptor_7e66da8bbfd29a12 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbb, 0x4e, 0xeb, 0x40,
	0x10, 0x86, 0xe5, 0x13, 0x27, 0x3e, 0x99, 0x24, 0x14, 0x5b, 0x6d, 0x8a, 0x08, 0x63, 0x51, 0xb8,
	0xb2, 0x15, 0x2e, 0x3d, 0x17, 0x89, 0x8a, 0xca, 0x74, 0x34, 0xc8, 0x1b, 0x8f, 0x61, 0x25, 0x67,
	0xd6, 0x8a, 0xc7, 0x48, 0xbc, 0x24, 0xcf, 0x84, 0x76, 0x6c, 0x77, 0xd0, 0xcd, 0x7f, 0xf9, 0x46,
	0xbb, 0x03, 0x97, 0xa6, 0xaf, 0x4d, 0x6f, 0x9b, 0x2a, 0x37, 0x7d, 0x9d, 0xd7, 0xb6, 0xc1, 0xfc,
	0x73, 0x6f, 0x90, 0xcb, 0xbd, 0x88, 0xac, 0x3d, 0x39, 0x76, 0x6a, 0x3b, 0xb5, 0x32, 0xd3, 0xd7,
	0x99, 0x04, 0x63, 0x2b, 0xb9, 0x83, 0xe8, 0xc9, 0x36, 0xf8, 0x82, 0xac, 0x6e, 0x61, 0xee, 0xa3,
	0x4e, 0x07, 0xf1, 0x2c, 0x5d, 0x5d, 0x9d, 0x67, 0x7f, 0x52, 0x99, 0x47, 0x8a, 0xa1, 0x9d, 0xdc,
	0x40, 0xe8, 0xa5, 0x52, 0x10, 0xb6, 0x25, 0x7f, 0xe8, 0x20, 0x0e, 0xd2, 0x65, 0x21, 0xb3, 0xd2,
	0x10, 0x1d, 0x1c, 0x31, 0x12, 0xeb, 0x7f, 0x71, 0x90, 0xae, 0x8b, 0x49, 0x26, 0xdf, 0x01, 0x9c,
	0x79, 0xec, 0x9e, 0xc8, 0x71, 0xc9, 0xd6, 0xd1, 0xaf, 0x0b, 0x76, 0x00, 0x1d, 0x97, 0x27, 0x7e,
	0x6b, 0x2c, 0xa1, 0xec, 0xd8, 0x14, 0x4b, 0x71, 0x9e, 0x2d, 0xa1, 0xba, 0x80, 0xf5, 0x10, 0x1f,
	0x5c, 0xd3, 0x1f, 0x49, 0xcf, 0xa4, 0xb0, 0x12, 0xef, 0x51, 0x2c, 0xb5, 0x85, 0xff, 0x48, 0xd5,
	0xc0, 0x87, 0x12, 0x47, 0x48, 0x95, 0xd0, 0x3b, 0x00, 0x1f, 0x8d, 0xec, 0x7c, 0x58, 0x8e, 0x54,
	0x8d, 0xa4, 0x82, 0x90, 0xbf, 0x5a, 0xd4, 0x8b, 0xe1, 0x3d, 0x7e, 0xf6, 0x1f, 0x3a, 0x62, 0xd7,
	0x95, 0xef, 0xa8, 0x23, 0xb1, 0x27, 0xf9, 0xb0, 0x79, 0x5d, 0xf9, 0x7b, 0x8c, 0x17, 0x32, 0x0b,
	0xb9, 0xfc, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x26, 0x70, 0xc4, 0xa1, 0x01, 0x00,
	0x00,
}

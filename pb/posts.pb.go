// Code generated by protoc-gen-go. DO NOT EDIT.
// source: posts.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Post struct {
	Slug      string                     `protobuf:"bytes,1,opt,name=slug" json:"slug,omitempty"`
	VendorID  *ID                        `protobuf:"bytes,2,opt,name=vendorID" json:"vendorID,omitempty"`
	Title     string                     `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	LongForm  string                     `protobuf:"bytes,4,opt,name=longForm" json:"longForm,omitempty"`
	Images    []*Post_Image              `protobuf:"bytes,5,rep,name=images" json:"images,omitempty"`
	Tags      []string                   `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetVendorID() *ID {
	if m != nil {
		return m.VendorID
	}
	return nil
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetLongForm() string {
	if m != nil {
		return m.LongForm
	}
	return ""
}

func (m *Post) GetImages() []*Post_Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Post) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Post_Image struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original" json:"original,omitempty"`
	Large    string `protobuf:"bytes,3,opt,name=large" json:"large,omitempty"`
	Medium   string `protobuf:"bytes,4,opt,name=medium" json:"medium,omitempty"`
	Small    string `protobuf:"bytes,5,opt,name=small" json:"small,omitempty"`
	Tiny     string `protobuf:"bytes,6,opt,name=tiny" json:"tiny,omitempty"`
}

func (m *Post_Image) Reset()                    { *m = Post_Image{} }
func (m *Post_Image) String() string            { return proto.CompactTextString(m) }
func (*Post_Image) ProtoMessage()               {}
func (*Post_Image) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

func (m *Post_Image) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Post_Image) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *Post_Image) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

func (m *Post_Image) GetMedium() string {
	if m != nil {
		return m.Medium
	}
	return ""
}

func (m *Post_Image) GetSmall() string {
	if m != nil {
		return m.Small
	}
	return ""
}

func (m *Post_Image) GetTiny() string {
	if m != nil {
		return m.Tiny
	}
	return ""
}

type SignedPost struct {
	Post      *Post  `protobuf:"bytes,1,opt,name=post" json:"post,omitempty"`
	Hash      string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedPost) Reset()                    { *m = SignedPost{} }
func (m *SignedPost) String() string            { return proto.CompactTextString(m) }
func (*SignedPost) ProtoMessage()               {}
func (*SignedPost) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *SignedPost) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *SignedPost) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SignedPost) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "Post")
	proto.RegisterType((*Post_Image)(nil), "Post.Image")
	proto.RegisterType((*SignedPost)(nil), "SignedPost")
}

func init() { proto.RegisterFile("posts.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xf3, 0xc7, 0x66, 0x22, 0x08, 0x8b, 0xc8, 0x1a, 0x84, 0x86, 0x7a, 0xe9, 0x29,
	0x85, 0x7a, 0xf1, 0x2c, 0x45, 0xe8, 0x4d, 0xa2, 0x17, 0xbd, 0x6d, 0xdb, 0xed, 0x76, 0x61, 0xb3,
	0x1b, 0xb2, 0x1b, 0xc1, 0x4f, 0xe2, 0x77, 0xf5, 0x24, 0x3b, 0xf9, 0xd3, 0xdb, 0xbc, 0x99, 0x37,
	0xc9, 0xfb, 0xcd, 0x42, 0xd6, 0x18, 0xeb, 0x6c, 0xd9, 0xb4, 0xc6, 0x99, 0x7c, 0x21, 0x8c, 0x11,
	0x8a, 0xaf, 0x51, 0xed, 0xbb, 0xd3, 0xda, 0xc9, 0x9a, 0x5b, 0xc7, 0xea, 0x66, 0x30, 0xdc, 0x1c,
	0x8c, 0x76, 0x2d, 0x3b, 0x8c, 0x1b, 0xcb, 0xbf, 0x19, 0x44, 0x6f, 0xc6, 0x3a, 0x42, 0x20, 0xb2,
	0xaa, 0x13, 0x34, 0x28, 0x82, 0x55, 0x5a, 0x61, 0x4d, 0x16, 0x30, 0xff, 0xe6, 0xfa, 0x68, 0xda,
	0xdd, 0x96, 0xce, 0x8a, 0x60, 0x95, 0x6d, 0xc2, 0x72, 0xb7, 0xad, 0xa6, 0x26, 0xb9, 0x85, 0xd8,
	0x49, 0xa7, 0x38, 0x0d, 0x71, 0xab, 0x17, 0x24, 0x87, 0xb9, 0x32, 0x5a, 0xbc, 0x9a, 0xb6, 0xa6,
	0x11, 0x0e, 0x26, 0x4d, 0x1e, 0x21, 0x91, 0x35, 0x13, 0xdc, 0xd2, 0xb8, 0x08, 0x57, 0xd9, 0x26,
	0x2b, 0xfd, 0xdf, 0xcb, 0x9d, 0xef, 0x55, 0xc3, 0xc8, 0x67, 0x71, 0x4c, 0x58, 0x9a, 0x14, 0xa1,
	0xcf, 0xe2, 0x6b, 0xf2, 0x0c, 0xe9, 0x04, 0x43, 0xaf, 0x30, 0x4c, 0x5e, 0xf6, 0xb8, 0xe5, 0x88,
	0x5b, 0x7e, 0x8c, 0x8e, 0xea, 0x62, 0xce, 0x7f, 0x03, 0x88, 0xf1, 0xfb, 0x3e, 0xd8, 0x49, 0x2a,
	0xae, 0x59, 0xcd, 0x07, 0xce, 0x49, 0xfb, 0x99, 0x69, 0xa5, 0x90, 0x9a, 0x29, 0x64, 0x4d, 0xab,
	0x49, 0x7b, 0x4c, 0xc5, 0x5a, 0x31, 0x61, 0xa2, 0x20, 0x77, 0x90, 0xd4, 0xfc, 0x28, 0xbb, 0x11,
	0x72, 0x50, 0xde, 0x6d, 0x6b, 0xa6, 0x14, 0x8d, 0x7b, 0x37, 0x0a, 0x64, 0x92, 0xfa, 0x87, 0x26,
	0xfd, 0x7d, 0x7d, 0xbd, 0xfc, 0x04, 0x78, 0x97, 0x42, 0xf3, 0x23, 0xbe, 0xc0, 0x3d, 0x44, 0xfe,
	0x2d, 0x31, 0x59, 0xb6, 0x89, 0xf1, 0x30, 0x15, 0xb6, 0xfc, 0xf2, 0x99, 0xd9, 0xf3, 0x10, 0x0c,
	0x6b, 0xf2, 0x00, 0xa9, 0x95, 0x42, 0x33, 0xd7, 0xb5, 0x7d, 0xb0, 0xeb, 0xea, 0xd2, 0x78, 0x89,
	0xbe, 0x66, 0xcd, 0x7e, 0x9f, 0xe0, 0x65, 0x9e, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x98,
	0xdd, 0x6f, 0x25, 0x02, 0x00, 0x00,
}

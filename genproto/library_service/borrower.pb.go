// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: borrower.proto

package library_service

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

type Borrower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId     string `protobuf:"bytes,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowDate string `protobuf:"bytes,4,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	CreatedAt  string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Borrower) Reset() {
	*x = Borrower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Borrower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Borrower) ProtoMessage() {}

func (x *Borrower) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Borrower.ProtoReflect.Descriptor instead.
func (*Borrower) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{0}
}

func (x *Borrower) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Borrower) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Borrower) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *Borrower) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *Borrower) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

func (x *Borrower) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Borrower) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Borrower) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BorrowerCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId string `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *BorrowerCreateReq) Reset() {
	*x = BorrowerCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerCreateReq) ProtoMessage() {}

func (x *BorrowerCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerCreateReq.ProtoReflect.Descriptor instead.
func (*BorrowerCreateReq) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{1}
}

func (x *BorrowerCreateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BorrowerCreateReq) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type BorrowerRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User       *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Book       *BookRes `protobuf:"bytes,3,opt,name=book,proto3" json:"book,omitempty"`
	BorrowDate string   `protobuf:"bytes,4,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string   `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
}

func (x *BorrowerRes) Reset() {
	*x = BorrowerRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerRes) ProtoMessage() {}

func (x *BorrowerRes) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerRes.ProtoReflect.Descriptor instead.
func (*BorrowerRes) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{2}
}

func (x *BorrowerRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BorrowerRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BorrowerRes) GetBook() *BookRes {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *BorrowerRes) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *BorrowerRes) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type BorrowedBooksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BorrowedBooksReq) Reset() {
	*x = BorrowedBooksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowedBooksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowedBooksReq) ProtoMessage() {}

func (x *BorrowedBooksReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowedBooksReq.ProtoReflect.Descriptor instead.
func (*BorrowedBooksReq) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{3}
}

func (x *BorrowedBooksReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BorrowedBooksRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*BookRes `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BorrowedBooksRes) Reset() {
	*x = BorrowedBooksRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowedBooksRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowedBooksRes) ProtoMessage() {}

func (x *BorrowedBooksRes) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowedBooksRes.ProtoReflect.Descriptor instead.
func (*BorrowedBooksRes) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{4}
}

func (x *BorrowedBooksRes) GetBooks() []*BookRes {
	if x != nil {
		return x.Books
	}
	return nil
}

type BorrowerUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId     string `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowDate string `protobuf:"bytes,3,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string `protobuf:"bytes,4,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
}

func (x *BorrowerUpdate) Reset() {
	*x = BorrowerUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerUpdate) ProtoMessage() {}

func (x *BorrowerUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerUpdate.ProtoReflect.Descriptor instead.
func (*BorrowerUpdate) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{5}
}

func (x *BorrowerUpdate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BorrowerUpdate) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BorrowerUpdate) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *BorrowerUpdate) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type BorrowerUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *GetByIdReq     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdateBorrower *BorrowerUpdate `protobuf:"bytes,2,opt,name=updateBorrower,proto3" json:"updateBorrower,omitempty"`
}

func (x *BorrowerUpdateReq) Reset() {
	*x = BorrowerUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerUpdateReq) ProtoMessage() {}

func (x *BorrowerUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerUpdateReq.ProtoReflect.Descriptor instead.
func (*BorrowerUpdateReq) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{6}
}

func (x *BorrowerUpdateReq) GetId() *GetByIdReq {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *BorrowerUpdateReq) GetUpdateBorrower() *BorrowerUpdate {
	if x != nil {
		return x.UpdateBorrower
	}
	return nil
}

type BorrowerGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId     string  `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowDate string  `protobuf:"bytes,3,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string  `protobuf:"bytes,4,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	Filter     *Filter `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *BorrowerGetAllReq) Reset() {
	*x = BorrowerGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerGetAllReq) ProtoMessage() {}

func (x *BorrowerGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerGetAllReq.ProtoReflect.Descriptor instead.
func (*BorrowerGetAllReq) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{7}
}

func (x *BorrowerGetAllReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BorrowerGetAllReq) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BorrowerGetAllReq) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *BorrowerGetAllReq) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

func (x *BorrowerGetAllReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type BorrowerGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Borrowers []*BorrowerRes `protobuf:"bytes,1,rep,name=Borrowers,proto3" json:"Borrowers,omitempty"`
}

func (x *BorrowerGetAllRes) Reset() {
	*x = BorrowerGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerGetAllRes) ProtoMessage() {}

func (x *BorrowerGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerGetAllRes.ProtoReflect.Descriptor instead.
func (*BorrowerGetAllRes) Descriptor() ([]byte, []int) {
	return file_borrower_proto_rawDescGZIP(), []int{8}
}

func (x *BorrowerGetAllRes) GetBorrowers() []*BorrowerRes {
	if x != nil {
		return x.Borrowers
	}
	return nil
}

var File_borrower_proto protoreflect.FileDescriptor

var file_borrower_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a,
	0x0b, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x52,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x10, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x89, 0x01, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x47, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x22, 0xb8, 0x01, 0x0a, 0x11,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x42,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x09, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x32, 0x90, 0x05, 0x0a, 0x0f, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x15, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x5a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_borrower_proto_rawDescOnce sync.Once
	file_borrower_proto_rawDescData = file_borrower_proto_rawDesc
)

func file_borrower_proto_rawDescGZIP() []byte {
	file_borrower_proto_rawDescOnce.Do(func() {
		file_borrower_proto_rawDescData = protoimpl.X.CompressGZIP(file_borrower_proto_rawDescData)
	})
	return file_borrower_proto_rawDescData
}

var file_borrower_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_borrower_proto_goTypes = []interface{}{
	(*Borrower)(nil),          // 0: library_service.Borrower
	(*BorrowerCreateReq)(nil), // 1: library_service.BorrowerCreateReq
	(*BorrowerRes)(nil),       // 2: library_service.BorrowerRes
	(*BorrowedBooksReq)(nil),  // 3: library_service.BorrowedBooksReq
	(*BorrowedBooksRes)(nil),  // 4: library_service.BorrowedBooksRes
	(*BorrowerUpdate)(nil),    // 5: library_service.BorrowerUpdate
	(*BorrowerUpdateReq)(nil), // 6: library_service.BorrowerUpdateReq
	(*BorrowerGetAllReq)(nil), // 7: library_service.BorrowerGetAllReq
	(*BorrowerGetAllRes)(nil), // 8: library_service.BorrowerGetAllRes
	(*User)(nil),              // 9: library_service.User
	(*BookRes)(nil),           // 10: library_service.BookRes
	(*GetByIdReq)(nil),        // 11: library_service.GetByIdReq
	(*Filter)(nil),            // 12: library_service.Filter
	(*Void)(nil),              // 13: library_service.Void
}
var file_borrower_proto_depIdxs = []int32{
	9,  // 0: library_service.BorrowerRes.user:type_name -> library_service.User
	10, // 1: library_service.BorrowerRes.book:type_name -> library_service.BookRes
	10, // 2: library_service.BorrowedBooksRes.books:type_name -> library_service.BookRes
	11, // 3: library_service.BorrowerUpdateReq.id:type_name -> library_service.GetByIdReq
	5,  // 4: library_service.BorrowerUpdateReq.updateBorrower:type_name -> library_service.BorrowerUpdate
	12, // 5: library_service.BorrowerGetAllReq.filter:type_name -> library_service.Filter
	2,  // 6: library_service.BorrowerGetAllRes.Borrowers:type_name -> library_service.BorrowerRes
	1,  // 7: library_service.BorrowerService.Create:input_type -> library_service.BorrowerCreateReq
	11, // 8: library_service.BorrowerService.Get:input_type -> library_service.GetByIdReq
	7,  // 9: library_service.BorrowerService.GetAll:input_type -> library_service.BorrowerGetAllReq
	6,  // 10: library_service.BorrowerService.Update:input_type -> library_service.BorrowerUpdateReq
	11, // 11: library_service.BorrowerService.Delete:input_type -> library_service.GetByIdReq
	13, // 12: library_service.BorrowerService.GetOverdueBooks:input_type -> library_service.Void
	3,  // 13: library_service.BorrowerService.GetBorrowedBooks:input_type -> library_service.BorrowedBooksReq
	3,  // 14: library_service.BorrowerService.GetBorrowingHistory:input_type -> library_service.BorrowedBooksReq
	2,  // 15: library_service.BorrowerService.Create:output_type -> library_service.BorrowerRes
	2,  // 16: library_service.BorrowerService.Get:output_type -> library_service.BorrowerRes
	8,  // 17: library_service.BorrowerService.GetAll:output_type -> library_service.BorrowerGetAllRes
	2,  // 18: library_service.BorrowerService.Update:output_type -> library_service.BorrowerRes
	13, // 19: library_service.BorrowerService.Delete:output_type -> library_service.Void
	8,  // 20: library_service.BorrowerService.GetOverdueBooks:output_type -> library_service.BorrowerGetAllRes
	4,  // 21: library_service.BorrowerService.GetBorrowedBooks:output_type -> library_service.BorrowedBooksRes
	4,  // 22: library_service.BorrowerService.GetBorrowingHistory:output_type -> library_service.BorrowedBooksRes
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_borrower_proto_init() }
func file_borrower_proto_init() {
	if File_borrower_proto != nil {
		return
	}
	file_common_proto_init()
	file_user_proto_init()
	file_book_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_borrower_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Borrower); i {
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
		file_borrower_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerCreateReq); i {
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
		file_borrower_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerRes); i {
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
		file_borrower_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowedBooksReq); i {
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
		file_borrower_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowedBooksRes); i {
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
		file_borrower_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerUpdate); i {
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
		file_borrower_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerUpdateReq); i {
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
		file_borrower_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerGetAllReq); i {
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
		file_borrower_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerGetAllRes); i {
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
			RawDescriptor: file_borrower_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_borrower_proto_goTypes,
		DependencyIndexes: file_borrower_proto_depIdxs,
		MessageInfos:      file_borrower_proto_msgTypes,
	}.Build()
	File_borrower_proto = out.File
	file_borrower_proto_rawDesc = nil
	file_borrower_proto_goTypes = nil
	file_borrower_proto_depIdxs = nil
}

// Code generated by protoc-gen-go.
// source: github.com/tomrittervg/pond/protos/pond.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Reply_Status int32

const (
	Reply_OK                         Reply_Status = 0
	Reply_PARSE_ERROR                Reply_Status = 1
	Reply_NO_REQUEST                 Reply_Status = 2
	Reply_INTERNAL_ERROR             Reply_Status = 3
	Reply_IDENTITY_ALREADY_KNOWN     Reply_Status = 10
	Reply_OVERLOAD                   Reply_Status = 11
	Reply_NO_SUCH_ADDRESS            Reply_Status = 12
	Reply_DELIVERY_SIGNATURE_INVALID Reply_Status = 13
	Reply_INCORRECT_GENERATION       Reply_Status = 14
	Reply_MAILBOX_FULL               Reply_Status = 15
	Reply_NO_ACCOUNT                 Reply_Status = 16
	Reply_OVER_QUOTA                 Reply_Status = 17
	Reply_FILE_LARGER_THAN_SIZE      Reply_Status = 18
	Reply_FILE_COMPLETE              Reply_Status = 19
	Reply_NO_SUCH_FILE               Reply_Status = 20
	Reply_RESUME_PAST_END_OF_FILE    Reply_Status = 21
	Reply_GENERATION_REVOKED         Reply_Status = 22
	Reply_CANNOT_PARSE_REVOCATION    Reply_Status = 23
	Reply_REGISTRATION_DISABLED      Reply_Status = 24
	Reply_HMAC_KEY_ALREADY_SET       Reply_Status = 25
	Reply_HMAC_NOT_SETUP             Reply_Status = 26
	Reply_HMAC_INCORRECT             Reply_Status = 27
	Reply_HMAC_USED                  Reply_Status = 28
	Reply_HMAC_REVOKED               Reply_Status = 29
)

var Reply_Status_name = map[int32]string{
	0:  "OK",
	1:  "PARSE_ERROR",
	2:  "NO_REQUEST",
	3:  "INTERNAL_ERROR",
	10: "IDENTITY_ALREADY_KNOWN",
	11: "OVERLOAD",
	12: "NO_SUCH_ADDRESS",
	13: "DELIVERY_SIGNATURE_INVALID",
	14: "INCORRECT_GENERATION",
	15: "MAILBOX_FULL",
	16: "NO_ACCOUNT",
	17: "OVER_QUOTA",
	18: "FILE_LARGER_THAN_SIZE",
	19: "FILE_COMPLETE",
	20: "NO_SUCH_FILE",
	21: "RESUME_PAST_END_OF_FILE",
	22: "GENERATION_REVOKED",
	23: "CANNOT_PARSE_REVOCATION",
	24: "REGISTRATION_DISABLED",
	25: "HMAC_KEY_ALREADY_SET",
	26: "HMAC_NOT_SETUP",
	27: "HMAC_INCORRECT",
	28: "HMAC_USED",
	29: "HMAC_REVOKED",
}
var Reply_Status_value = map[string]int32{
	"OK":                         0,
	"PARSE_ERROR":                1,
	"NO_REQUEST":                 2,
	"INTERNAL_ERROR":             3,
	"IDENTITY_ALREADY_KNOWN":     10,
	"OVERLOAD":                   11,
	"NO_SUCH_ADDRESS":            12,
	"DELIVERY_SIGNATURE_INVALID": 13,
	"INCORRECT_GENERATION":       14,
	"MAILBOX_FULL":               15,
	"NO_ACCOUNT":                 16,
	"OVER_QUOTA":                 17,
	"FILE_LARGER_THAN_SIZE":      18,
	"FILE_COMPLETE":              19,
	"NO_SUCH_FILE":               20,
	"RESUME_PAST_END_OF_FILE":    21,
	"GENERATION_REVOKED":         22,
	"CANNOT_PARSE_REVOCATION":    23,
	"REGISTRATION_DISABLED":      24,
	"HMAC_KEY_ALREADY_SET":       25,
	"HMAC_NOT_SETUP":             26,
	"HMAC_INCORRECT":             27,
	"HMAC_USED":                  28,
	"HMAC_REVOKED":               29,
}

func (x Reply_Status) Enum() *Reply_Status {
	p := new(Reply_Status)
	*p = x
	return p
}
func (x Reply_Status) String() string {
	return proto.EnumName(Reply_Status_name, int32(x))
}
func (x Reply_Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Reply_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Status_value, data, "Reply_Status")
	if err != nil {
		return err
	}
	*x = Reply_Status(value)
	return nil
}

type Message_Encoding int32

const (
	Message_RAW  Message_Encoding = 0
	Message_GZIP Message_Encoding = 1
)

var Message_Encoding_name = map[int32]string{
	0: "RAW",
	1: "GZIP",
}
var Message_Encoding_value = map[string]int32{
	"RAW":  0,
	"GZIP": 1,
}

func (x Message_Encoding) Enum() *Message_Encoding {
	p := new(Message_Encoding)
	*p = x
	return p
}
func (x Message_Encoding) String() string {
	return proto.EnumName(Message_Encoding_name, int32(x))
}
func (x Message_Encoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Message_Encoding) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_Encoding_value, data, "Message_Encoding")
	if err != nil {
		return err
	}
	*x = Message_Encoding(value)
	return nil
}

type Request struct {
	NewAccount       *NewAccount       `protobuf:"bytes,1,opt,name=new_account" json:"new_account,omitempty"`
	Deliver          *Delivery         `protobuf:"bytes,2,opt,name=deliver" json:"deliver,omitempty"`
	Fetch            *Fetch            `protobuf:"bytes,3,opt,name=fetch" json:"fetch,omitempty"`
	Upload           *Upload           `protobuf:"bytes,4,opt,name=upload" json:"upload,omitempty"`
	Download         *Download         `protobuf:"bytes,5,opt,name=download" json:"download,omitempty"`
	Revocation       *SignedRevocation `protobuf:"bytes,6,opt,name=revocation" json:"revocation,omitempty"`
	HmacSetup        *HMACSetup        `protobuf:"bytes,7,opt,name=hmac_setup" json:"hmac_setup,omitempty"`
	HmacStrike       *HMACStrike       `protobuf:"bytes,8,opt,name=hmac_strike" json:"hmac_strike,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *Request) Reset()         { *this = Request{} }
func (this *Request) String() string { return proto.CompactTextString(this) }
func (*Request) ProtoMessage()       {}

func (this *Request) GetNewAccount() *NewAccount {
	if this != nil {
		return this.NewAccount
	}
	return nil
}

func (this *Request) GetDeliver() *Delivery {
	if this != nil {
		return this.Deliver
	}
	return nil
}

func (this *Request) GetFetch() *Fetch {
	if this != nil {
		return this.Fetch
	}
	return nil
}

func (this *Request) GetUpload() *Upload {
	if this != nil {
		return this.Upload
	}
	return nil
}

func (this *Request) GetDownload() *Download {
	if this != nil {
		return this.Download
	}
	return nil
}

func (this *Request) GetRevocation() *SignedRevocation {
	if this != nil {
		return this.Revocation
	}
	return nil
}

func (this *Request) GetHmacSetup() *HMACSetup {
	if this != nil {
		return this.HmacSetup
	}
	return nil
}

func (this *Request) GetHmacStrike() *HMACStrike {
	if this != nil {
		return this.HmacStrike
	}
	return nil
}

type Reply struct {
	Status           *Reply_Status       `protobuf:"varint,1,opt,name=status,enum=protos.Reply_Status,def=0" json:"status,omitempty"`
	AccountCreated   *AccountCreated     `protobuf:"bytes,2,opt,name=account_created" json:"account_created,omitempty"`
	Fetched          *Fetched            `protobuf:"bytes,3,opt,name=fetched" json:"fetched,omitempty"`
	Announce         *ServerAnnounce     `protobuf:"bytes,4,opt,name=announce" json:"announce,omitempty"`
	Upload           *UploadReply        `protobuf:"bytes,5,opt,name=upload" json:"upload,omitempty"`
	Download         *DownloadReply      `protobuf:"bytes,6,opt,name=download" json:"download,omitempty"`
	Revocation       *SignedRevocation   `protobuf:"bytes,7,opt,name=revocation" json:"revocation,omitempty"`
	ExtraRevocations []*SignedRevocation `protobuf:"bytes,8,rep,name=extra_revocations" json:"extra_revocations,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (this *Reply) Reset()         { *this = Reply{} }
func (this *Reply) String() string { return proto.CompactTextString(this) }
func (*Reply) ProtoMessage()       {}

const Default_Reply_Status Reply_Status = Reply_OK

func (this *Reply) GetStatus() Reply_Status {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return Default_Reply_Status
}

func (this *Reply) GetAccountCreated() *AccountCreated {
	if this != nil {
		return this.AccountCreated
	}
	return nil
}

func (this *Reply) GetFetched() *Fetched {
	if this != nil {
		return this.Fetched
	}
	return nil
}

func (this *Reply) GetAnnounce() *ServerAnnounce {
	if this != nil {
		return this.Announce
	}
	return nil
}

func (this *Reply) GetUpload() *UploadReply {
	if this != nil {
		return this.Upload
	}
	return nil
}

func (this *Reply) GetDownload() *DownloadReply {
	if this != nil {
		return this.Download
	}
	return nil
}

func (this *Reply) GetRevocation() *SignedRevocation {
	if this != nil {
		return this.Revocation
	}
	return nil
}

func (this *Reply) GetExtraRevocations() []*SignedRevocation {
	if this != nil {
		return this.ExtraRevocations
	}
	return nil
}

type NewAccount struct {
	Generation       *uint32 `protobuf:"fixed32,1,req,name=generation" json:"generation,omitempty"`
	Group            []byte  `protobuf:"bytes,2,req,name=group" json:"group,omitempty"`
	HmacKey          []byte  `protobuf:"bytes,3,opt,name=hmac_key" json:"hmac_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *NewAccount) Reset()         { *this = NewAccount{} }
func (this *NewAccount) String() string { return proto.CompactTextString(this) }
func (*NewAccount) ProtoMessage()       {}

func (this *NewAccount) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *NewAccount) GetGroup() []byte {
	if this != nil {
		return this.Group
	}
	return nil
}

func (this *NewAccount) GetHmacKey() []byte {
	if this != nil {
		return this.HmacKey
	}
	return nil
}

type AccountDetails struct {
	Queue            *uint32 `protobuf:"varint,1,req,name=queue" json:"queue,omitempty"`
	MaxQueue         *uint32 `protobuf:"varint,2,req,name=max_queue" json:"max_queue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *AccountDetails) Reset()         { *this = AccountDetails{} }
func (this *AccountDetails) String() string { return proto.CompactTextString(this) }
func (*AccountDetails) ProtoMessage()       {}

func (this *AccountDetails) GetQueue() uint32 {
	if this != nil && this.Queue != nil {
		return *this.Queue
	}
	return 0
}

func (this *AccountDetails) GetMaxQueue() uint32 {
	if this != nil && this.MaxQueue != nil {
		return *this.MaxQueue
	}
	return 0
}

type AccountCreated struct {
	Details          *AccountDetails `protobuf:"bytes,1,req,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *AccountCreated) Reset()         { *this = AccountCreated{} }
func (this *AccountCreated) String() string { return proto.CompactTextString(this) }
func (*AccountCreated) ProtoMessage()       {}

func (this *AccountCreated) GetDetails() *AccountDetails {
	if this != nil {
		return this.Details
	}
	return nil
}

type Delivery struct {
	To               []byte  `protobuf:"bytes,1,req,name=to" json:"to,omitempty"`
	GroupSignature   []byte  `protobuf:"bytes,2,opt,name=group_signature" json:"group_signature,omitempty"`
	Generation       *uint32 `protobuf:"fixed32,3,opt,name=generation" json:"generation,omitempty"`
	Message          []byte  `protobuf:"bytes,4,req,name=message" json:"message,omitempty"`
	OneTimePublicKey []byte  `protobuf:"bytes,5,opt,name=one_time_public_key" json:"one_time_public_key,omitempty"`
	HmacOfPublicKey  *uint64 `protobuf:"fixed64,6,opt,name=hmac_of_public_key" json:"hmac_of_public_key,omitempty"`
	OneTimeSignature []byte  `protobuf:"bytes,7,opt,name=one_time_signature" json:"one_time_signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Delivery) Reset()         { *this = Delivery{} }
func (this *Delivery) String() string { return proto.CompactTextString(this) }
func (*Delivery) ProtoMessage()       {}

func (this *Delivery) GetTo() []byte {
	if this != nil {
		return this.To
	}
	return nil
}

func (this *Delivery) GetGroupSignature() []byte {
	if this != nil {
		return this.GroupSignature
	}
	return nil
}

func (this *Delivery) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *Delivery) GetMessage() []byte {
	if this != nil {
		return this.Message
	}
	return nil
}

func (this *Delivery) GetOneTimePublicKey() []byte {
	if this != nil {
		return this.OneTimePublicKey
	}
	return nil
}

func (this *Delivery) GetHmacOfPublicKey() uint64 {
	if this != nil && this.HmacOfPublicKey != nil {
		return *this.HmacOfPublicKey
	}
	return 0
}

func (this *Delivery) GetOneTimeSignature() []byte {
	if this != nil {
		return this.OneTimeSignature
	}
	return nil
}

type Fetch struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *Fetch) Reset()         { *this = Fetch{} }
func (this *Fetch) String() string { return proto.CompactTextString(this) }
func (*Fetch) ProtoMessage()       {}

type Fetched struct {
	GroupSignature   []byte          `protobuf:"bytes,1,req,name=group_signature" json:"group_signature,omitempty"`
	Generation       *uint32         `protobuf:"fixed32,2,req,name=generation" json:"generation,omitempty"`
	Message          []byte          `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	Details          *AccountDetails `protobuf:"bytes,4,req,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *Fetched) Reset()         { *this = Fetched{} }
func (this *Fetched) String() string { return proto.CompactTextString(this) }
func (*Fetched) ProtoMessage()       {}

func (this *Fetched) GetGroupSignature() []byte {
	if this != nil {
		return this.GroupSignature
	}
	return nil
}

func (this *Fetched) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *Fetched) GetMessage() []byte {
	if this != nil {
		return this.Message
	}
	return nil
}

func (this *Fetched) GetDetails() *AccountDetails {
	if this != nil {
		return this.Details
	}
	return nil
}

type ServerAnnounce struct {
	Message          *Message `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *ServerAnnounce) Reset()         { *this = ServerAnnounce{} }
func (this *ServerAnnounce) String() string { return proto.CompactTextString(this) }
func (*ServerAnnounce) ProtoMessage()       {}

func (this *ServerAnnounce) GetMessage() *Message {
	if this != nil {
		return this.Message
	}
	return nil
}

type Upload struct {
	Id               *uint64 `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Size             *int64  `protobuf:"varint,2,req,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Upload) Reset()         { *this = Upload{} }
func (this *Upload) String() string { return proto.CompactTextString(this) }
func (*Upload) ProtoMessage()       {}

func (this *Upload) GetId() uint64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Upload) GetSize() int64 {
	if this != nil && this.Size != nil {
		return *this.Size
	}
	return 0
}

type UploadReply struct {
	Resume           *int64 `protobuf:"varint,1,opt,name=resume" json:"resume,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *UploadReply) Reset()         { *this = UploadReply{} }
func (this *UploadReply) String() string { return proto.CompactTextString(this) }
func (*UploadReply) ProtoMessage()       {}

func (this *UploadReply) GetResume() int64 {
	if this != nil && this.Resume != nil {
		return *this.Resume
	}
	return 0
}

type Download struct {
	From             []byte  `protobuf:"bytes,1,req,name=from" json:"from,omitempty"`
	Id               *uint64 `protobuf:"fixed64,2,req,name=id" json:"id,omitempty"`
	Resume           *int64  `protobuf:"varint,3,opt,name=resume" json:"resume,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Download) Reset()         { *this = Download{} }
func (this *Download) String() string { return proto.CompactTextString(this) }
func (*Download) ProtoMessage()       {}

func (this *Download) GetFrom() []byte {
	if this != nil {
		return this.From
	}
	return nil
}

func (this *Download) GetId() uint64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Download) GetResume() int64 {
	if this != nil && this.Resume != nil {
		return *this.Resume
	}
	return 0
}

type DownloadReply struct {
	Size             *int64 `protobuf:"varint,1,req,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *DownloadReply) Reset()         { *this = DownloadReply{} }
func (this *DownloadReply) String() string { return proto.CompactTextString(this) }
func (*DownloadReply) ProtoMessage()       {}

func (this *DownloadReply) GetSize() int64 {
	if this != nil && this.Size != nil {
		return *this.Size
	}
	return 0
}

type SignedRevocation struct {
	Revocation       *SignedRevocation_Revocation `protobuf:"bytes,1,req,name=revocation" json:"revocation,omitempty"`
	Signature        []byte                       `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (this *SignedRevocation) Reset()         { *this = SignedRevocation{} }
func (this *SignedRevocation) String() string { return proto.CompactTextString(this) }
func (*SignedRevocation) ProtoMessage()       {}

func (this *SignedRevocation) GetRevocation() *SignedRevocation_Revocation {
	if this != nil {
		return this.Revocation
	}
	return nil
}

func (this *SignedRevocation) GetSignature() []byte {
	if this != nil {
		return this.Signature
	}
	return nil
}

type SignedRevocation_Revocation struct {
	Generation       *uint32 `protobuf:"fixed32,1,req,name=generation" json:"generation,omitempty"`
	Revocation       []byte  `protobuf:"bytes,2,req,name=revocation" json:"revocation,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *SignedRevocation_Revocation) Reset()         { *this = SignedRevocation_Revocation{} }
func (this *SignedRevocation_Revocation) String() string { return proto.CompactTextString(this) }
func (*SignedRevocation_Revocation) ProtoMessage()       {}

func (this *SignedRevocation_Revocation) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

func (this *SignedRevocation_Revocation) GetRevocation() []byte {
	if this != nil {
		return this.Revocation
	}
	return nil
}

type HMACSetup struct {
	HmacKey          []byte `protobuf:"bytes,1,req,name=hmac_key" json:"hmac_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *HMACSetup) Reset()         { *this = HMACSetup{} }
func (this *HMACSetup) String() string { return proto.CompactTextString(this) }
func (*HMACSetup) ProtoMessage()       {}

func (this *HMACSetup) GetHmacKey() []byte {
	if this != nil {
		return this.HmacKey
	}
	return nil
}

type HMACStrike struct {
	Hmacs            []uint64 `protobuf:"fixed64,1,rep,packed,name=hmacs" json:"hmacs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *HMACStrike) Reset()         { *this = HMACStrike{} }
func (this *HMACStrike) String() string { return proto.CompactTextString(this) }
func (*HMACStrike) ProtoMessage()       {}

func (this *HMACStrike) GetHmacs() []uint64 {
	if this != nil {
		return this.Hmacs
	}
	return nil
}

type KeyExchange struct {
	PublicKey        []byte  `protobuf:"bytes,1,req,name=public_key" json:"public_key,omitempty"`
	IdentityPublic   []byte  `protobuf:"bytes,2,req,name=identity_public" json:"identity_public,omitempty"`
	Server           *string `protobuf:"bytes,3,req,name=server" json:"server,omitempty"`
	Dh               []byte  `protobuf:"bytes,4,req,name=dh" json:"dh,omitempty"`
	Dh1              []byte  `protobuf:"bytes,8,opt,name=dh1" json:"dh1,omitempty"`
	Group            []byte  `protobuf:"bytes,5,req,name=group" json:"group,omitempty"`
	GroupKey         []byte  `protobuf:"bytes,6,req,name=group_key" json:"group_key,omitempty"`
	Generation       *uint32 `protobuf:"varint,7,req,name=generation" json:"generation,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *KeyExchange) Reset()         { *this = KeyExchange{} }
func (this *KeyExchange) String() string { return proto.CompactTextString(this) }
func (*KeyExchange) ProtoMessage()       {}

func (this *KeyExchange) GetPublicKey() []byte {
	if this != nil {
		return this.PublicKey
	}
	return nil
}

func (this *KeyExchange) GetIdentityPublic() []byte {
	if this != nil {
		return this.IdentityPublic
	}
	return nil
}

func (this *KeyExchange) GetServer() string {
	if this != nil && this.Server != nil {
		return *this.Server
	}
	return ""
}

func (this *KeyExchange) GetDh() []byte {
	if this != nil {
		return this.Dh
	}
	return nil
}

func (this *KeyExchange) GetDh1() []byte {
	if this != nil {
		return this.Dh1
	}
	return nil
}

func (this *KeyExchange) GetGroup() []byte {
	if this != nil {
		return this.Group
	}
	return nil
}

func (this *KeyExchange) GetGroupKey() []byte {
	if this != nil {
		return this.GroupKey
	}
	return nil
}

func (this *KeyExchange) GetGeneration() uint32 {
	if this != nil && this.Generation != nil {
		return *this.Generation
	}
	return 0
}

type SignedKeyExchange struct {
	Signed           []byte `protobuf:"bytes,1,req,name=signed" json:"signed,omitempty"`
	Signature        []byte `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *SignedKeyExchange) Reset()         { *this = SignedKeyExchange{} }
func (this *SignedKeyExchange) String() string { return proto.CompactTextString(this) }
func (*SignedKeyExchange) ProtoMessage()       {}

func (this *SignedKeyExchange) GetSigned() []byte {
	if this != nil {
		return this.Signed
	}
	return nil
}

func (this *SignedKeyExchange) GetSignature() []byte {
	if this != nil {
		return this.Signature
	}
	return nil
}

type Message struct {
	Id               *uint64               `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Time             *int64                `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	Body             []byte                `protobuf:"bytes,3,req,name=body" json:"body,omitempty"`
	BodyEncoding     *Message_Encoding     `protobuf:"varint,4,opt,name=body_encoding,enum=protos.Message_Encoding" json:"body_encoding,omitempty"`
	MyNextDh         []byte                `protobuf:"bytes,5,opt,name=my_next_dh" json:"my_next_dh,omitempty"`
	InReplyTo        *uint64               `protobuf:"varint,6,opt,name=in_reply_to" json:"in_reply_to,omitempty"`
	AlsoAck          []uint64              `protobuf:"varint,10,rep,name=also_ack" json:"also_ack,omitempty"`
	Files            []*Message_Attachment `protobuf:"bytes,7,rep,name=files" json:"files,omitempty"`
	DetachedFiles    []*Message_Detachment `protobuf:"bytes,8,rep,name=detached_files" json:"detached_files,omitempty"`
	SupportedVersion *int32                `protobuf:"varint,9,opt,name=supported_version" json:"supported_version,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (this *Message) Reset()         { *this = Message{} }
func (this *Message) String() string { return proto.CompactTextString(this) }
func (*Message) ProtoMessage()       {}

func (this *Message) GetId() uint64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Message) GetTime() int64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *Message) GetBody() []byte {
	if this != nil {
		return this.Body
	}
	return nil
}

func (this *Message) GetBodyEncoding() Message_Encoding {
	if this != nil && this.BodyEncoding != nil {
		return *this.BodyEncoding
	}
	return 0
}

func (this *Message) GetMyNextDh() []byte {
	if this != nil {
		return this.MyNextDh
	}
	return nil
}

func (this *Message) GetInReplyTo() uint64 {
	if this != nil && this.InReplyTo != nil {
		return *this.InReplyTo
	}
	return 0
}

func (this *Message) GetAlsoAck() []uint64 {
	if this != nil {
		return this.AlsoAck
	}
	return nil
}

func (this *Message) GetFiles() []*Message_Attachment {
	if this != nil {
		return this.Files
	}
	return nil
}

func (this *Message) GetDetachedFiles() []*Message_Detachment {
	if this != nil {
		return this.DetachedFiles
	}
	return nil
}

func (this *Message) GetSupportedVersion() int32 {
	if this != nil && this.SupportedVersion != nil {
		return *this.SupportedVersion
	}
	return 0
}

type Message_Attachment struct {
	Filename         *string `protobuf:"bytes,1,req,name=filename" json:"filename,omitempty"`
	Contents         []byte  `protobuf:"bytes,2,req,name=contents" json:"contents,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Message_Attachment) Reset()         { *this = Message_Attachment{} }
func (this *Message_Attachment) String() string { return proto.CompactTextString(this) }
func (*Message_Attachment) ProtoMessage()       {}

func (this *Message_Attachment) GetFilename() string {
	if this != nil && this.Filename != nil {
		return *this.Filename
	}
	return ""
}

func (this *Message_Attachment) GetContents() []byte {
	if this != nil {
		return this.Contents
	}
	return nil
}

type Message_Detachment struct {
	Filename         *string `protobuf:"bytes,1,req,name=filename" json:"filename,omitempty"`
	Size             *uint64 `protobuf:"varint,2,req,name=size" json:"size,omitempty"`
	PaddedSize       *uint64 `protobuf:"varint,3,req,name=padded_size" json:"padded_size,omitempty"`
	ChunkSize        *uint32 `protobuf:"varint,4,req,name=chunk_size" json:"chunk_size,omitempty"`
	Key              []byte  `protobuf:"bytes,5,req,name=key" json:"key,omitempty"`
	Url              *string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Message_Detachment) Reset()         { *this = Message_Detachment{} }
func (this *Message_Detachment) String() string { return proto.CompactTextString(this) }
func (*Message_Detachment) ProtoMessage()       {}

func (this *Message_Detachment) GetFilename() string {
	if this != nil && this.Filename != nil {
		return *this.Filename
	}
	return ""
}

func (this *Message_Detachment) GetSize() uint64 {
	if this != nil && this.Size != nil {
		return *this.Size
	}
	return 0
}

func (this *Message_Detachment) GetPaddedSize() uint64 {
	if this != nil && this.PaddedSize != nil {
		return *this.PaddedSize
	}
	return 0
}

func (this *Message_Detachment) GetChunkSize() uint32 {
	if this != nil && this.ChunkSize != nil {
		return *this.ChunkSize
	}
	return 0
}

func (this *Message_Detachment) GetKey() []byte {
	if this != nil {
		return this.Key
	}
	return nil
}

func (this *Message_Detachment) GetUrl() string {
	if this != nil && this.Url != nil {
		return *this.Url
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.Reply_Status", Reply_Status_name, Reply_Status_value)
	proto.RegisterEnum("protos.Message_Encoding", Message_Encoding_name, Message_Encoding_value)
}

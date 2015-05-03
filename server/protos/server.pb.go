// Code generated by protoc-gen-go.
// source: github.com/tomrittervg/pond/server/protos/server.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Config struct {
	Port              *uint32 `protobuf:"varint,1,req,name=port" json:"port,omitempty"`
	Address           *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	AllowRegistration *bool   `protobuf:"varint,3,opt,name=allow_registration,def=1" json:"allow_registration,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (this *Config) Reset()         { *this = Config{} }
func (this *Config) String() string { return proto.CompactTextString(this) }
func (*Config) ProtoMessage()       {}

const Default_Config_AllowRegistration bool = true

func (this *Config) GetPort() uint32 {
	if this != nil && this.Port != nil {
		return *this.Port
	}
	return 0
}

func (this *Config) GetAddress() string {
	if this != nil && this.Address != nil {
		return *this.Address
	}
	return ""
}

func (this *Config) GetAllowRegistration() bool {
	if this != nil && this.AllowRegistration != nil {
		return *this.AllowRegistration
	}
	return Default_Config_AllowRegistration
}

func init() {
}

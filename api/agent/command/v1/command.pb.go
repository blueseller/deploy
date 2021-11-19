// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: github.com/blueseller/deploy/api/agent/command/v1/command.proto

package command

import (
	context "context"
	types "github.com/blueseller/deploy/api/agent/types/"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CmdType int32

const (
	CmdType_NONE CmdType = 0
	// 要求上报机器状态
	CmdType_STAT_REPORT CmdType = 1
	// 要求进行自我升级
	CmdType_UPDATE CmdType = 2
	// 要求执行一些命令
	CmdType_COMMAND_RUN CmdType = 3
	// 要求执行一些命令
	CmdType_INSTALL_SOFT CmdType = 4
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "NONE",
		1: "STAT_REPORT",
		2: "UPDATE",
		3: "COMMAND_RUN",
		4: "INSTALL_SOFT",
	}
	CmdType_value = map[string]int32{
		"NONE":         0,
		"STAT_REPORT":  1,
		"UPDATE":       2,
		"COMMAND_RUN":  3,
		"INSTALL_SOFT": 4,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_enumTypes[0].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_enumTypes[0]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescGZIP(), []int{0}
}

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId *types.AgentId `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	CmdType CmdType        `protobuf:"varint,2,opt,name=cmd_type,json=cmdType,proto3,enum=agent.command.v1.CmdType" json:"cmd_type,omitempty"`
	Payload []byte         `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Result  *Result        `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescGZIP(), []int{0}
}

func (x *Cmd) GetAgentId() *types.AgentId {
	if x != nil {
		return x.AgentId
	}
	return nil
}

func (x *Cmd) GetCmdType() CmdType {
	if x != nil {
		return x.CmdType
	}
	return CmdType_NONE
}

func (x *Cmd) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Cmd) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *Result) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

var File_github_com_blueseller_deploy_api_agent_command_v1_command_proto protoreflect.FileDescriptor

var file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x75,
	0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x6c, 0x75, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01,
	0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x07, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6d, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x2a, 0x53, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x54, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x53,
	0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x4f, 0x46, 0x54, 0x10, 0x04, 0x32, 0x4f, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6d, 0x64, 0x1a,
	0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6d, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescOnce sync.Once
	file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescData = file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDesc
)

func file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescGZIP() []byte {
	file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescOnce.Do(func() {
		file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescData)
	})
	return file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDescData
}

var file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_goTypes = []interface{}{
	(CmdType)(0),          // 0: agent.command.v1.CmdType
	(*Cmd)(nil),           // 1: agent.command.v1.Cmd
	(*Result)(nil),        // 2: agent.command.v1.Result
	(*types.AgentId)(nil), // 3: agent.types.AgentId
}
var file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_depIdxs = []int32{
	3, // 0: agent.command.v1.Cmd.agent_id:type_name -> agent.types.AgentId
	0, // 1: agent.command.v1.Cmd.cmd_type:type_name -> agent.command.v1.CmdType
	2, // 2: agent.command.v1.Cmd.result:type_name -> agent.command.v1.Result
	1, // 3: agent.command.v1.CommandSerivce.Command:input_type -> agent.command.v1.Cmd
	1, // 4: agent.command.v1.CommandSerivce.Command:output_type -> agent.command.v1.Cmd
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_init() }
func file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_init() {
	if File_github_com_blueseller_deploy_api_agent_command_v1_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
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
		file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_goTypes,
		DependencyIndexes: file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_depIdxs,
		EnumInfos:         file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_enumTypes,
		MessageInfos:      file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_msgTypes,
	}.Build()
	File_github_com_blueseller_deploy_api_agent_command_v1_command_proto = out.File
	file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_rawDesc = nil
	file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_goTypes = nil
	file_github_com_blueseller_deploy_api_agent_command_v1_command_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommandSerivceClient is the client API for CommandSerivce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandSerivceClient interface {
	Command(ctx context.Context, opts ...grpc.CallOption) (CommandSerivce_CommandClient, error)
}

type commandSerivceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandSerivceClient(cc grpc.ClientConnInterface) CommandSerivceClient {
	return &commandSerivceClient{cc}
}

func (c *commandSerivceClient) Command(ctx context.Context, opts ...grpc.CallOption) (CommandSerivce_CommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandSerivce_serviceDesc.Streams[0], "/agent.command.v1.CommandSerivce/Command", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandSerivceCommandClient{stream}
	return x, nil
}

type CommandSerivce_CommandClient interface {
	Send(*Cmd) error
	Recv() (*Cmd, error)
	grpc.ClientStream
}

type commandSerivceCommandClient struct {
	grpc.ClientStream
}

func (x *commandSerivceCommandClient) Send(m *Cmd) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandSerivceCommandClient) Recv() (*Cmd, error) {
	m := new(Cmd)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandSerivceServer is the server API for CommandSerivce service.
type CommandSerivceServer interface {
	Command(CommandSerivce_CommandServer) error
}

// UnimplementedCommandSerivceServer can be embedded to have forward compatible implementations.
type UnimplementedCommandSerivceServer struct {
}

func (*UnimplementedCommandSerivceServer) Command(CommandSerivce_CommandServer) error {
	return status.Errorf(codes.Unimplemented, "method Command not implemented")
}

func RegisterCommandSerivceServer(s *grpc.Server, srv CommandSerivceServer) {
	s.RegisterService(&_CommandSerivce_serviceDesc, srv)
}

func _CommandSerivce_Command_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandSerivceServer).Command(&commandSerivceCommandServer{stream})
}

type CommandSerivce_CommandServer interface {
	Send(*Cmd) error
	Recv() (*Cmd, error)
	grpc.ServerStream
}

type commandSerivceCommandServer struct {
	grpc.ServerStream
}

func (x *commandSerivceCommandServer) Send(m *Cmd) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandSerivceCommandServer) Recv() (*Cmd, error) {
	m := new(Cmd)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CommandSerivce_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.command.v1.CommandSerivce",
	HandlerType: (*CommandSerivceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Command",
			Handler:       _CommandSerivce_Command_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/blueseller/deploy/api/agent/command/v1/command.proto",
}

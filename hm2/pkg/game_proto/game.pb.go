// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: game_proto/game.proto

package game

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

type GameConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int32  `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GameConnect) Reset() {
	*x = GameConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameConnect) ProtoMessage() {}

func (x *GameConnect) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameConnect.ProtoReflect.Descriptor instead.
func (*GameConnect) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameConnect) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *GameConnect) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type KillSomeone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *KillSomeone) Reset() {
	*x = KillSomeone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillSomeone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillSomeone) ProtoMessage() {}

func (x *KillSomeone) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillSomeone.ProtoReflect.Descriptor instead.
func (*KillSomeone) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{1}
}

func (x *KillSomeone) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type EndDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndDay) Reset() {
	*x = EndDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndDay) ProtoMessage() {}

func (x *EndDay) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndDay.ProtoReflect.Descriptor instead.
func (*EndDay) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{2}
}

type GameRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *GameRole) Reset() {
	*x = GameRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRole) ProtoMessage() {}

func (x *GameRole) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRole.ProtoReflect.Descriptor instead.
func (*GameRole) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{3}
}

func (x *GameRole) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type GameStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayNum    int32    `protobuf:"varint,1,opt,name=day_num,json=dayNum,proto3" json:"day_num,omitempty"`
	Time      string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Usernames []string `protobuf:"bytes,3,rep,name=usernames,proto3" json:"usernames,omitempty"`
	Alive     []bool   `protobuf:"varint,4,rep,packed,name=alive,proto3" json:"alive,omitempty"`
	Messages  []string `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GameStatus) Reset() {
	*x = GameStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatus) ProtoMessage() {}

func (x *GameStatus) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatus.ProtoReflect.Descriptor instead.
func (*GameStatus) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{4}
}

func (x *GameStatus) GetDayNum() int32 {
	if x != nil {
		return x.DayNum
	}
	return 0
}

func (x *GameStatus) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *GameStatus) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

func (x *GameStatus) GetAlive() []bool {
	if x != nil {
		return x.Alive
	}
	return nil
}

func (x *GameStatus) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Investigate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Investigate) Reset() {
	*x = Investigate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Investigate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Investigate) ProtoMessage() {}

func (x *Investigate) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Investigate.ProtoReflect.Descriptor instead.
func (*Investigate) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{5}
}

func (x *Investigate) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type InvestigateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mafia bool `protobuf:"varint,1,opt,name=mafia,proto3" json:"mafia,omitempty"`
}

func (x *InvestigateResult) Reset() {
	*x = InvestigateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestigateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestigateResult) ProtoMessage() {}

func (x *InvestigateResult) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestigateResult.ProtoReflect.Descriptor instead.
func (*InvestigateResult) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{6}
}

func (x *InvestigateResult) GetMafia() bool {
	if x != nil {
		return x.Mafia
	}
	return false
}

type InvestigatePublish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer bool `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *InvestigatePublish) Reset() {
	*x = InvestigatePublish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestigatePublish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestigatePublish) ProtoMessage() {}

func (x *InvestigatePublish) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestigatePublish.ProtoReflect.Descriptor instead.
func (*InvestigatePublish) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{7}
}

func (x *InvestigatePublish) GetAnswer() bool {
	if x != nil {
		return x.Answer
	}
	return false
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[8]
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
	return file_game_proto_game_proto_rawDescGZIP(), []int{8}
}

func (x *Result) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GameMessage:
	//
	//	*GameRequest_GameConnect
	//	*GameRequest_KillSomeone
	//	*GameRequest_EndDay
	//	*GameRequest_Investigate
	//	*GameRequest_InvestigatePublish
	GameMessage isGameRequest_GameMessage `protobuf_oneof:"game_message"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{9}
}

func (m *GameRequest) GetGameMessage() isGameRequest_GameMessage {
	if m != nil {
		return m.GameMessage
	}
	return nil
}

func (x *GameRequest) GetGameConnect() *GameConnect {
	if x, ok := x.GetGameMessage().(*GameRequest_GameConnect); ok {
		return x.GameConnect
	}
	return nil
}

func (x *GameRequest) GetKillSomeone() *KillSomeone {
	if x, ok := x.GetGameMessage().(*GameRequest_KillSomeone); ok {
		return x.KillSomeone
	}
	return nil
}

func (x *GameRequest) GetEndDay() *EndDay {
	if x, ok := x.GetGameMessage().(*GameRequest_EndDay); ok {
		return x.EndDay
	}
	return nil
}

func (x *GameRequest) GetInvestigate() *Investigate {
	if x, ok := x.GetGameMessage().(*GameRequest_Investigate); ok {
		return x.Investigate
	}
	return nil
}

func (x *GameRequest) GetInvestigatePublish() *InvestigatePublish {
	if x, ok := x.GetGameMessage().(*GameRequest_InvestigatePublish); ok {
		return x.InvestigatePublish
	}
	return nil
}

type isGameRequest_GameMessage interface {
	isGameRequest_GameMessage()
}

type GameRequest_GameConnect struct {
	GameConnect *GameConnect `protobuf:"bytes,1,opt,name=game_connect,json=gameConnect,proto3,oneof"`
}

type GameRequest_KillSomeone struct {
	KillSomeone *KillSomeone `protobuf:"bytes,2,opt,name=kill_someone,json=killSomeone,proto3,oneof"`
}

type GameRequest_EndDay struct {
	EndDay *EndDay `protobuf:"bytes,3,opt,name=end_day,json=endDay,proto3,oneof"`
}

type GameRequest_Investigate struct {
	Investigate *Investigate `protobuf:"bytes,4,opt,name=investigate,proto3,oneof"`
}

type GameRequest_InvestigatePublish struct {
	InvestigatePublish *InvestigatePublish `protobuf:"bytes,5,opt,name=investigate_publish,json=investigatePublish,proto3,oneof"`
}

func (*GameRequest_GameConnect) isGameRequest_GameMessage() {}

func (*GameRequest_KillSomeone) isGameRequest_GameMessage() {}

func (*GameRequest_EndDay) isGameRequest_GameMessage() {}

func (*GameRequest_Investigate) isGameRequest_GameMessage() {}

func (*GameRequest_InvestigatePublish) isGameRequest_GameMessage() {}

type GameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GameMessage:
	//
	//	*GameResponse_GameRole
	//	*GameResponse_GameStatus
	//	*GameResponse_Result
	//	*GameResponse_InvestigateResult
	GameMessage isGameResponse_GameMessage `protobuf_oneof:"game_message"`
}

func (x *GameResponse) Reset() {
	*x = GameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_game_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResponse) ProtoMessage() {}

func (x *GameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_game_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResponse.ProtoReflect.Descriptor instead.
func (*GameResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_game_proto_rawDescGZIP(), []int{10}
}

func (m *GameResponse) GetGameMessage() isGameResponse_GameMessage {
	if m != nil {
		return m.GameMessage
	}
	return nil
}

func (x *GameResponse) GetGameRole() *GameRole {
	if x, ok := x.GetGameMessage().(*GameResponse_GameRole); ok {
		return x.GameRole
	}
	return nil
}

func (x *GameResponse) GetGameStatus() *GameStatus {
	if x, ok := x.GetGameMessage().(*GameResponse_GameStatus); ok {
		return x.GameStatus
	}
	return nil
}

func (x *GameResponse) GetResult() *Result {
	if x, ok := x.GetGameMessage().(*GameResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *GameResponse) GetInvestigateResult() *InvestigateResult {
	if x, ok := x.GetGameMessage().(*GameResponse_InvestigateResult); ok {
		return x.InvestigateResult
	}
	return nil
}

type isGameResponse_GameMessage interface {
	isGameResponse_GameMessage()
}

type GameResponse_GameRole struct {
	GameRole *GameRole `protobuf:"bytes,1,opt,name=game_role,json=gameRole,proto3,oneof"`
}

type GameResponse_GameStatus struct {
	GameStatus *GameStatus `protobuf:"bytes,2,opt,name=game_status,json=gameStatus,proto3,oneof"`
}

type GameResponse_Result struct {
	Result *Result `protobuf:"bytes,3,opt,name=result,proto3,oneof"`
}

type GameResponse_InvestigateResult struct {
	InvestigateResult *InvestigateResult `protobuf:"bytes,4,opt,name=investigate_result,json=investigateResult,proto3,oneof"`
}

func (*GameResponse_GameRole) isGameResponse_GameMessage() {}

func (*GameResponse_GameStatus) isGameResponse_GameMessage() {}

func (*GameResponse_Result) isGameResponse_GameMessage() {}

func (*GameResponse_InvestigateResult) isGameResponse_GameMessage() {}

var File_game_proto_game_proto protoreflect.FileDescriptor

var file_game_proto_game_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0b, 0x4b, 0x69, 0x6c, 0x6c, 0x53,
	0x6f, 0x6d, 0x65, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x22, 0x1e, 0x0a, 0x08,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x89, 0x01, 0x0a,
	0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x61,
	0x79, 0x4e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x66, 0x69,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x22, 0x2c,
	0x0a, 0x12, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xba,
	0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73,
	0x6f, 0x6d, 0x65, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x53, 0x6f, 0x6d, 0x65, 0x6f, 0x6e, 0x65, 0x48,
	0x00, 0x52, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x6f, 0x6d, 0x65, 0x6f, 0x6e, 0x65, 0x12, 0x27,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x48, 0x00, 0x52,
	0x06, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x48,
	0x00, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x12, 0x4b,
	0x0a, 0x13, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x48, 0x00, 0x52, 0x12, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x67, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x0e, 0x0a, 0x0c, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf4, 0x01, 0x0a, 0x0c,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52,
	0x11, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x3b, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_game_proto_rawDescOnce sync.Once
	file_game_proto_game_proto_rawDescData = file_game_proto_game_proto_rawDesc
)

func file_game_proto_game_proto_rawDescGZIP() []byte {
	file_game_proto_game_proto_rawDescOnce.Do(func() {
		file_game_proto_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_game_proto_rawDescData)
	})
	return file_game_proto_game_proto_rawDescData
}

var file_game_proto_game_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_game_proto_game_proto_goTypes = []interface{}{
	(*GameConnect)(nil),        // 0: game.GameConnect
	(*KillSomeone)(nil),        // 1: game.KillSomeone
	(*EndDay)(nil),             // 2: game.EndDay
	(*GameRole)(nil),           // 3: game.GameRole
	(*GameStatus)(nil),         // 4: game.GameStatus
	(*Investigate)(nil),        // 5: game.Investigate
	(*InvestigateResult)(nil),  // 6: game.InvestigateResult
	(*InvestigatePublish)(nil), // 7: game.InvestigatePublish
	(*Result)(nil),             // 8: game.Result
	(*GameRequest)(nil),        // 9: game.GameRequest
	(*GameResponse)(nil),       // 10: game.GameResponse
}
var file_game_proto_game_proto_depIdxs = []int32{
	0,  // 0: game.GameRequest.game_connect:type_name -> game.GameConnect
	1,  // 1: game.GameRequest.kill_someone:type_name -> game.KillSomeone
	2,  // 2: game.GameRequest.end_day:type_name -> game.EndDay
	5,  // 3: game.GameRequest.investigate:type_name -> game.Investigate
	7,  // 4: game.GameRequest.investigate_publish:type_name -> game.InvestigatePublish
	3,  // 5: game.GameResponse.game_role:type_name -> game.GameRole
	4,  // 6: game.GameResponse.game_status:type_name -> game.GameStatus
	8,  // 7: game.GameResponse.result:type_name -> game.Result
	6,  // 8: game.GameResponse.investigate_result:type_name -> game.InvestigateResult
	9,  // 9: game.Game.Join:input_type -> game.GameRequest
	10, // 10: game.Game.Join:output_type -> game.GameResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_game_proto_game_proto_init() }
func file_game_proto_game_proto_init() {
	if File_game_proto_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameConnect); i {
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
		file_game_proto_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillSomeone); i {
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
		file_game_proto_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndDay); i {
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
		file_game_proto_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRole); i {
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
		file_game_proto_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStatus); i {
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
		file_game_proto_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Investigate); i {
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
		file_game_proto_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestigateResult); i {
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
		file_game_proto_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestigatePublish); i {
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
		file_game_proto_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_game_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
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
		file_game_proto_game_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResponse); i {
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
	file_game_proto_game_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*GameRequest_GameConnect)(nil),
		(*GameRequest_KillSomeone)(nil),
		(*GameRequest_EndDay)(nil),
		(*GameRequest_Investigate)(nil),
		(*GameRequest_InvestigatePublish)(nil),
	}
	file_game_proto_game_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*GameResponse_GameRole)(nil),
		(*GameResponse_GameStatus)(nil),
		(*GameResponse_Result)(nil),
		(*GameResponse_InvestigateResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_game_proto_goTypes,
		DependencyIndexes: file_game_proto_game_proto_depIdxs,
		MessageInfos:      file_game_proto_game_proto_msgTypes,
	}.Build()
	File_game_proto_game_proto = out.File
	file_game_proto_game_proto_rawDesc = nil
	file_game_proto_game_proto_goTypes = nil
	file_game_proto_game_proto_depIdxs = nil
}
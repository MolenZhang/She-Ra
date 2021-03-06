// Code generated by protoc-gen-go.
// source: configData.proto
// DO NOT EDIT!

/*
Package configdata is a generated protocol buffer package.

It is generated from these files:
	configData.proto

It has these top-level messages:
	Job
	CodeManager
	BuildManager
	Git
	Repo
	GitAdvanced
	Ant
	Maven
	Execution
*/
package configdata

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

type CodeManager_Kind int32

const (
	CodeManager_NONE CodeManager_Kind = 0
	CodeManager_GIT  CodeManager_Kind = 1
)

var CodeManager_Kind_name = map[int32]string{
	0: "NONE",
	1: "GIT",
}
var CodeManager_Kind_value = map[string]int32{
	"NONE": 0,
	"GIT":  1,
}

func (x CodeManager_Kind) String() string {
	return proto.EnumName(CodeManager_Kind_name, int32(x))
}
func (CodeManager_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type BuildManager_Kind int32

const (
	BuildManager_NONE BuildManager_Kind = 0
	BuildManager_ANT  BuildManager_Kind = 1
	BuildManager_MVN  BuildManager_Kind = 2
)

var BuildManager_Kind_name = map[int32]string{
	0: "NONE",
	1: "ANT",
	2: "MVN",
}
var BuildManager_Kind_value = map[string]int32{
	"NONE": 0,
	"ANT":  1,
	"MVN":  2,
}

func (x BuildManager_Kind) String() string {
	return proto.EnumName(BuildManager_Kind_name, int32(x))
}
func (BuildManager_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Execution_State int32

const (
	Execution_INIT           Execution_State = 0
	Execution_CODE_PULLING   Execution_State = 1
	Execution_CODE_BUILDING  Execution_State = 2
	Execution_IMAGE_BUILDING Execution_State = 3
	Execution_IMAGE_PUSHING  Execution_State = 4
)

var Execution_State_name = map[int32]string{
	0: "INIT",
	1: "CODE_PULLING",
	2: "CODE_BUILDING",
	3: "IMAGE_BUILDING",
	4: "IMAGE_PUSHING",
}
var Execution_State_value = map[string]int32{
	"INIT":           0,
	"CODE_PULLING":   1,
	"CODE_BUILDING":  2,
	"IMAGE_BUILDING": 3,
	"IMAGE_PUSHING":  4,
}

func (x Execution_State) String() string {
	return proto.EnumName(Execution_State_name, int32(x))
}
func (Execution_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type Execution_Status int32

const (
	Execution_SUCCESS   Execution_Status = 0
	Execution_FAILURE   Execution_Status = 1
	Execution_CANCELLED Execution_Status = 2
	Execution_IGNORED   Execution_Status = 3
)

var Execution_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
	2: "CANCELLED",
	3: "IGNORED",
}
var Execution_Status_value = map[string]int32{
	"SUCCESS":   0,
	"FAILURE":   1,
	"CANCELLED": 2,
	"IGNORED":   3,
}

func (x Execution_Status) String() string {
	return proto.EnumName(Execution_Status_name, int32(x))
}
func (Execution_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 1} }

type Job struct {
	Id                  string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NextExecTime        string        `protobuf:"bytes,2,opt,name=nextExecTime" json:"nextExecTime,omitempty"`
	JdkVersion          string        `protobuf:"bytes,3,opt,name=jdkVersion" json:"jdkVersion,omitempty"`
	BuildImgCmd         string        `protobuf:"bytes,4,opt,name=buildImgCmd" json:"buildImgCmd,omitempty"`
	PushImgCmd          string        `protobuf:"bytes,5,opt,name=pushImgCmd" json:"pushImgCmd,omitempty"`
	CurrentNumber       int32         `protobuf:"varint,6,opt,name=currentNumber" json:"currentNumber,omitempty"`
	MaxExecutionRecords int32         `protobuf:"varint,7,opt,name=maxExecutionRecords" json:"maxExecutionRecords,omitempty"`
	MaxKeepDays         int32         `protobuf:"varint,8,opt,name=maxKeepDays" json:"maxKeepDays,omitempty"`
	CodeManager         *CodeManager  `protobuf:"bytes,9,opt,name=codeManager" json:"codeManager,omitempty"`
	BuildManager        *BuildManager `protobuf:"bytes,10,opt,name=buildManager" json:"buildManager,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetCodeManager() *CodeManager {
	if m != nil {
		return m.CodeManager
	}
	return nil
}

func (m *Job) GetBuildManager() *BuildManager {
	if m != nil {
		return m.BuildManager
	}
	return nil
}

type CodeManager struct {
	Select    CodeManager_Kind `protobuf:"varint,1,opt,name=select,enum=configdata.CodeManager_Kind" json:"select,omitempty"`
	GitConfig *Git             `protobuf:"bytes,2,opt,name=gitConfig" json:"gitConfig,omitempty"`
}

func (m *CodeManager) Reset()                    { *m = CodeManager{} }
func (m *CodeManager) String() string            { return proto.CompactTextString(m) }
func (*CodeManager) ProtoMessage()               {}
func (*CodeManager) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CodeManager) GetGitConfig() *Git {
	if m != nil {
		return m.GitConfig
	}
	return nil
}

type BuildManager struct {
	Select    BuildManager_Kind `protobuf:"varint,1,opt,name=select,enum=configdata.BuildManager_Kind" json:"select,omitempty"`
	AntConfig *Ant              `protobuf:"bytes,2,opt,name=antConfig" json:"antConfig,omitempty"`
	MvnConfig *Maven            `protobuf:"bytes,3,opt,name=mvnConfig" json:"mvnConfig,omitempty"`
}

func (m *BuildManager) Reset()                    { *m = BuildManager{} }
func (m *BuildManager) String() string            { return proto.CompactTextString(m) }
func (*BuildManager) ProtoMessage()               {}
func (*BuildManager) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BuildManager) GetAntConfig() *Ant {
	if m != nil {
		return m.AntConfig
	}
	return nil
}

func (m *BuildManager) GetMvnConfig() *Maven {
	if m != nil {
		return m.MvnConfig
	}
	return nil
}

type Git struct {
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Repo   *Repo  `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
}

func (m *Git) Reset()                    { *m = Git{} }
func (m *Git) String() string            { return proto.CompactTextString(m) }
func (*Git) ProtoMessage()               {}
func (*Git) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Git) GetRepo() *Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

type Repo struct {
	Url         string       `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Credentials string       `protobuf:"bytes,2,opt,name=credentials" json:"credentials,omitempty"`
	Advanced    *GitAdvanced `protobuf:"bytes,3,opt,name=advanced" json:"advanced,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Repo) GetAdvanced() *GitAdvanced {
	if m != nil {
		return m.Advanced
	}
	return nil
}

type GitAdvanced struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Refspec string `protobuf:"bytes,2,opt,name=refspec" json:"refspec,omitempty"`
}

func (m *GitAdvanced) Reset()                    { *m = GitAdvanced{} }
func (m *GitAdvanced) String() string            { return proto.CompactTextString(m) }
func (*GitAdvanced) ProtoMessage()               {}
func (*GitAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Ant struct {
	Version    string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Targets    string `protobuf:"bytes,2,opt,name=targets" json:"targets,omitempty"`
	BuildFile  string `protobuf:"bytes,3,opt,name=buildFile" json:"buildFile,omitempty"`
	Properties string `protobuf:"bytes,4,opt,name=properties" json:"properties,omitempty"`
	Javaopts   string `protobuf:"bytes,5,opt,name=javaopts" json:"javaopts,omitempty"`
}

func (m *Ant) Reset()                    { *m = Ant{} }
func (m *Ant) String() string            { return proto.CompactTextString(m) }
func (*Ant) ProtoMessage()               {}
func (*Ant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Maven struct {
	Version           string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Goals             string `protobuf:"bytes,2,opt,name=goals" json:"goals,omitempty"`
	Pom               string `protobuf:"bytes,3,opt,name=pom" json:"pom,omitempty"`
	Properties        string `protobuf:"bytes,4,opt,name=properties" json:"properties,omitempty"`
	Jvmopts           string `protobuf:"bytes,5,opt,name=jvmopts" json:"jvmopts,omitempty"`
	SettingFile       string `protobuf:"bytes,6,opt,name=settingFile" json:"settingFile,omitempty"`
	GlobalSettingFile string `protobuf:"bytes,7,opt,name=globalSettingFile" json:"globalSettingFile,omitempty"`
}

func (m *Maven) Reset()                    { *m = Maven{} }
func (m *Maven) String() string            { return proto.CompactTextString(m) }
func (*Maven) ProtoMessage()               {}
func (*Maven) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Execution struct {
	Job       string           `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Number    int32            `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	LogFile   string           `protobuf:"bytes,3,opt,name=logFile" json:"logFile,omitempty"`
	Progress  Execution_State  `protobuf:"varint,4,opt,name=progress,enum=configdata.Execution_State" json:"progress,omitempty"`
	EndStatus Execution_Status `protobuf:"varint,5,opt,name=endStatus,enum=configdata.Execution_Status" json:"endStatus,omitempty"`
}

func (m *Execution) Reset()                    { *m = Execution{} }
func (m *Execution) String() string            { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()               {}
func (*Execution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Job)(nil), "configdata.Job")
	proto.RegisterType((*CodeManager)(nil), "configdata.CodeManager")
	proto.RegisterType((*BuildManager)(nil), "configdata.BuildManager")
	proto.RegisterType((*Git)(nil), "configdata.Git")
	proto.RegisterType((*Repo)(nil), "configdata.Repo")
	proto.RegisterType((*GitAdvanced)(nil), "configdata.GitAdvanced")
	proto.RegisterType((*Ant)(nil), "configdata.Ant")
	proto.RegisterType((*Maven)(nil), "configdata.Maven")
	proto.RegisterType((*Execution)(nil), "configdata.Execution")
	proto.RegisterEnum("configdata.CodeManager_Kind", CodeManager_Kind_name, CodeManager_Kind_value)
	proto.RegisterEnum("configdata.BuildManager_Kind", BuildManager_Kind_name, BuildManager_Kind_value)
	proto.RegisterEnum("configdata.Execution_State", Execution_State_name, Execution_State_value)
	proto.RegisterEnum("configdata.Execution_Status", Execution_Status_name, Execution_Status_value)
}

func init() { proto.RegisterFile("configData.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x8e, 0xe3, 0x34,
	0x14, 0x9e, 0x34, 0xfd, 0x3d, 0x9d, 0x19, 0x32, 0x06, 0xb1, 0x01, 0x16, 0x34, 0x8a, 0xf6, 0x62,
	0x2e, 0xa0, 0xa0, 0x59, 0x10, 0xe2, 0x47, 0x48, 0xdd, 0xb4, 0x5b, 0xc2, 0x76, 0x32, 0xab, 0x74,
	0xba, 0x77, 0x08, 0xb9, 0x89, 0x37, 0x9b, 0xd2, 0xd8, 0x91, 0xe3, 0x56, 0xc3, 0x43, 0xc0, 0x0d,
	0x4f, 0xc5, 0x35, 0x8f, 0xc0, 0x8b, 0x20, 0x3b, 0x4e, 0xe3, 0xce, 0xee, 0xce, 0xde, 0xe5, 0xfb,
	0xce, 0x77, 0xec, 0xef, 0xd8, 0xe7, 0x38, 0xe0, 0xc4, 0x8c, 0xbe, 0xcc, 0xd2, 0x09, 0x16, 0x78,
	0x54, 0x70, 0x26, 0x18, 0x82, 0x8a, 0x49, 0xb0, 0xc0, 0xde, 0x5f, 0x36, 0xd8, 0xbf, 0xb0, 0x15,
	0x3a, 0x85, 0x56, 0x96, 0xb8, 0xd6, 0xb9, 0x75, 0x31, 0x88, 0x5a, 0x59, 0x82, 0x3c, 0x38, 0xa6,
	0xe4, 0x56, 0x4c, 0x6f, 0x49, 0x7c, 0x93, 0xe5, 0xc4, 0x6d, 0xa9, 0xc8, 0x01, 0x87, 0x3e, 0x03,
	0x58, 0x27, 0xbf, 0xbf, 0x20, 0xbc, 0xcc, 0x18, 0x75, 0x6d, 0xa5, 0x30, 0x18, 0x74, 0x0e, 0xc3,
	0xd5, 0x36, 0xdb, 0x24, 0x41, 0x9e, 0xfa, 0x79, 0xe2, 0xb6, 0x95, 0xc0, 0xa4, 0xe4, 0x0a, 0xc5,
	0xb6, 0x7c, 0xa5, 0x05, 0x9d, 0x6a, 0x85, 0x86, 0x41, 0x8f, 0xe0, 0x24, 0xde, 0x72, 0x4e, 0xa8,
	0x08, 0xb7, 0xf9, 0x8a, 0x70, 0xb7, 0x7b, 0x6e, 0x5d, 0x74, 0xa2, 0x43, 0x12, 0x7d, 0x05, 0xef,
	0xe7, 0xf8, 0x56, 0xda, 0xda, 0x8a, 0x8c, 0xd1, 0x88, 0xc4, 0x8c, 0x27, 0xa5, 0xdb, 0x53, 0xda,
	0x37, 0x85, 0xa4, 0xb3, 0x1c, 0xdf, 0x3e, 0x23, 0xa4, 0x98, 0xe0, 0x3f, 0x4a, 0xb7, 0xaf, 0x94,
	0x26, 0x85, 0xbe, 0x83, 0x61, 0xcc, 0x12, 0x72, 0x85, 0x29, 0x4e, 0x09, 0x77, 0x07, 0xe7, 0xd6,
	0xc5, 0xf0, 0xf2, 0xc1, 0xa8, 0x39, 0xb9, 0x91, 0xdf, 0x84, 0x23, 0x53, 0x8b, 0x7e, 0x84, 0x63,
	0x55, 0x63, 0x9d, 0x0b, 0x2a, 0xd7, 0x35, 0x73, 0x9f, 0x18, 0xf1, 0xe8, 0x40, 0xed, 0xfd, 0x69,
	0xc1, 0xd0, 0x58, 0x1a, 0x7d, 0x0d, 0xdd, 0x92, 0x6c, 0x48, 0x2c, 0xd4, 0xe5, 0x9c, 0x5e, 0x3e,
	0x7c, 0x8b, 0x87, 0xd1, 0xb3, 0x8c, 0x26, 0x91, 0xd6, 0xa2, 0x2f, 0x60, 0x90, 0x66, 0xc2, 0x57,
	0x4a, 0x75, 0x77, 0xc3, 0xcb, 0xf7, 0xcc, 0xc4, 0x59, 0x26, 0xa2, 0x46, 0xe1, 0x7d, 0x04, 0x6d,
	0x99, 0x8e, 0xfa, 0xd0, 0x0e, 0xaf, 0xc3, 0xa9, 0x73, 0x84, 0x7a, 0x60, 0xcf, 0x82, 0x1b, 0xc7,
	0xf2, 0xfe, 0xb1, 0xe0, 0xd8, 0xb4, 0x8b, 0xbe, 0xb9, 0x63, 0xe8, 0xd3, 0xb7, 0x15, 0xf6, 0x9a,
	0x23, 0x4c, 0xef, 0x71, 0x34, 0xa6, 0x22, 0x6a, 0x14, 0xe8, 0x4b, 0x18, 0xe4, 0x3b, 0xaa, 0xe5,
	0xb6, 0x92, 0x9f, 0x99, 0xf2, 0x2b, 0xbc, 0x23, 0x34, 0x6a, 0x34, 0x9e, 0xf7, 0xa6, 0x12, 0xc6,
	0xe1, 0x8d, 0x63, 0xc9, 0x8f, 0xab, 0x17, 0xa1, 0xd3, 0xf2, 0x7c, 0xb0, 0x67, 0x99, 0x40, 0x1f,
	0x42, 0x77, 0xc5, 0x31, 0x8d, 0x5f, 0xe9, 0x7e, 0xd7, 0x08, 0x3d, 0x82, 0x36, 0x27, 0x05, 0xd3,
	0xee, 0x1c, 0x73, 0xbb, 0x88, 0x14, 0x2c, 0x52, 0x51, 0x8f, 0x41, 0x5b, 0x22, 0xe4, 0x80, 0xbd,
	0xe5, 0x1b, 0xbd, 0x84, 0xfc, 0x94, 0x5d, 0x15, 0x73, 0x92, 0x10, 0x2a, 0x32, 0xbc, 0x29, 0xf5,
	0xc8, 0x98, 0x14, 0x7a, 0x0c, 0x7d, 0x9c, 0xec, 0x30, 0x8d, 0x49, 0xa2, 0x8b, 0x7a, 0x70, 0xe7,
	0x56, 0xc6, 0x3a, 0x1c, 0xed, 0x85, 0xde, 0x0f, 0x30, 0x34, 0x02, 0x08, 0x41, 0x9b, 0xe2, 0x9c,
	0xe8, 0x8d, 0xd5, 0x37, 0x72, 0xa1, 0xc7, 0xc9, 0xcb, 0xb2, 0x20, 0xb1, 0xde, 0xb5, 0x86, 0xde,
	0xdf, 0x16, 0xd8, 0x63, 0x2a, 0xa4, 0x62, 0xa7, 0x07, 0xb5, 0x4a, 0xac, 0xa1, 0x8c, 0x08, 0xcc,
	0x53, 0x22, 0x6a, 0xc7, 0x35, 0x44, 0x0f, 0x61, 0xa0, 0x5a, 0xf3, 0x69, 0xb6, 0x21, 0x7a, 0xbc,
	0x1b, 0x42, 0xcd, 0x2e, 0x67, 0x05, 0xe1, 0x22, 0x23, 0xa5, 0x1e, 0x6e, 0x83, 0x41, 0x1f, 0x43,
	0x7f, 0x8d, 0x77, 0x98, 0x15, 0xa2, 0xd4, 0x93, 0xbd, 0xc7, 0xde, 0xbf, 0x16, 0x74, 0xd4, 0x0d,
	0xde, 0xe3, 0xeb, 0x03, 0xe8, 0xa4, 0xac, 0x39, 0xc7, 0x0a, 0xc8, 0x53, 0x2f, 0x58, 0xae, 0xdd,
	0xc8, 0xcf, 0x77, 0xfa, 0x70, 0xa1, 0xb7, 0xde, 0xe5, 0x86, 0x8d, 0x1a, 0xca, 0xfb, 0x2a, 0x89,
	0x10, 0x19, 0x4d, 0x55, 0x85, 0xdd, 0xea, 0xbe, 0x0c, 0x0a, 0x7d, 0x0e, 0x67, 0xe9, 0x86, 0xad,
	0xf0, 0x66, 0x61, 0xe8, 0x7a, 0x4a, 0xf7, 0x7a, 0xc0, 0xfb, 0xaf, 0x05, 0x83, 0xfd, 0x53, 0x23,
	0x9d, 0xae, 0xd9, 0xaa, 0xee, 0x8f, 0x35, 0x5b, 0xc9, 0xbe, 0xa3, 0xd5, 0x33, 0xd6, 0x52, 0x0f,
	0x8e, 0x46, 0xd2, 0xe1, 0x86, 0xa5, 0xc6, 0x29, 0xd7, 0x10, 0x7d, 0x0b, 0xfd, 0x82, 0xb3, 0x94,
	0x93, 0xb2, 0xaa, 0xec, 0xf4, 0xf2, 0x13, 0xb3, 0x5f, 0xf6, 0x9b, 0x8d, 0x16, 0x02, 0x0b, 0x12,
	0xed, 0xc5, 0xe8, 0x7b, 0x18, 0x10, 0x9a, 0x48, 0x76, 0x5b, 0x95, 0x7d, 0xe7, 0xe1, 0x38, 0xcc,
	0xdc, 0x96, 0x51, 0x23, 0xf7, 0x7e, 0x85, 0x8e, 0x5a, 0x4e, 0x8e, 0x52, 0x10, 0x06, 0x37, 0xce,
	0x11, 0x72, 0xe0, 0xd8, 0xbf, 0x9e, 0x4c, 0x7f, 0x7b, 0xbe, 0x9c, 0xcf, 0x83, 0x70, 0xe6, 0x58,
	0xe8, 0x0c, 0x4e, 0x14, 0xf3, 0x64, 0x19, 0xcc, 0x27, 0x92, 0x6a, 0x21, 0x04, 0xa7, 0xc1, 0xd5,
	0x78, 0x66, 0x70, 0xb6, 0x94, 0x55, 0xdc, 0xf3, 0xe5, 0xe2, 0x67, 0x49, 0xb5, 0xbd, 0x9f, 0xa0,
	0x5b, 0x6d, 0x84, 0x86, 0xd0, 0x5b, 0x2c, 0x7d, 0x7f, 0xba, 0x58, 0x38, 0x47, 0x12, 0x3c, 0x1d,
	0x07, 0xf3, 0x65, 0x34, 0x75, 0x2c, 0x74, 0x02, 0x03, 0x7f, 0x1c, 0xfa, 0xd3, 0xf9, 0x7c, 0x3a,
	0x71, 0x5a, 0x32, 0x16, 0xcc, 0xc2, 0xeb, 0x68, 0x3a, 0x71, 0xec, 0x55, 0x57, 0xfd, 0xc4, 0x1e,
	0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x84, 0x26, 0x8d, 0x84, 0xd8, 0x06, 0x00, 0x00,
}

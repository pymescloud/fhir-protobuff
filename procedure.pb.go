// Code generated by protoc-gen-go. DO NOT EDIT.
// source: procedure.proto

package buffer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Procedure struct {
	Status          string           `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	PerformedPeriod *PerformedPeriod `protobuf:"bytes,2,opt,name=performedPeriod" json:"performedPeriod,omitempty"`
	Code            *Code            `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	ResourceType    string           `protobuf:"bytes,4,opt,name=resourceType" json:"resourceType,omitempty"`
	Text            *Text            `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
	FollowUp        []*FollowUp      `protobuf:"bytes,6,rep,name=followUp" json:"followUp,omitempty"`
	Performer       []*PPerformer    `protobuf:"bytes,7,rep,name=performer" json:"performer,omitempty"`
	BodySite        []*BodySite      `protobuf:"bytes,8,rep,name=bodySite" json:"bodySite,omitempty"`
	Context         *Context         `protobuf:"bytes,9,opt,name=context" json:"context,omitempty"`
	Report          []*Report        `protobuf:"bytes,10,rep,name=report" json:"report,omitempty"`
	ReasonCode      []*ReasonCode    `protobuf:"bytes,11,rep,name=reasonCode" json:"reasonCode,omitempty"`
	Outcome         *Outcome         `protobuf:"bytes,12,opt,name=outcome" json:"outcome,omitempty"`
	Id              string           `protobuf:"bytes,13,opt,name=id" json:"id,omitempty"`
	Subject         *Subject         `protobuf:"bytes,14,opt,name=subject" json:"subject,omitempty"`
	Meta            *Meta            `protobuf:"bytes,15,opt,name=meta" json:"meta,omitempty"`
}

func (m *Procedure) Reset()                    { *m = Procedure{} }
func (m *Procedure) String() string            { return proto.CompactTextString(m) }
func (*Procedure) ProtoMessage()               {}
func (*Procedure) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *Procedure) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Procedure) GetPerformedPeriod() *PerformedPeriod {
	if m != nil {
		return m.PerformedPeriod
	}
	return nil
}

func (m *Procedure) GetCode() *Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Procedure) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Procedure) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Procedure) GetFollowUp() []*FollowUp {
	if m != nil {
		return m.FollowUp
	}
	return nil
}

func (m *Procedure) GetPerformer() []*PPerformer {
	if m != nil {
		return m.Performer
	}
	return nil
}

func (m *Procedure) GetBodySite() []*BodySite {
	if m != nil {
		return m.BodySite
	}
	return nil
}

func (m *Procedure) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Procedure) GetReport() []*Report {
	if m != nil {
		return m.Report
	}
	return nil
}

func (m *Procedure) GetReasonCode() []*ReasonCode {
	if m != nil {
		return m.ReasonCode
	}
	return nil
}

func (m *Procedure) GetOutcome() *Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (m *Procedure) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Procedure) GetSubject() *Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Procedure) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Outcome struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Outcome) Reset()                    { *m = Outcome{} }
func (m *Outcome) String() string            { return proto.CompactTextString(m) }
func (*Outcome) ProtoMessage()               {}
func (*Outcome) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *Outcome) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ReasonCode struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *ReasonCode) Reset()                    { *m = ReasonCode{} }
func (m *ReasonCode) String() string            { return proto.CompactTextString(m) }
func (*ReasonCode) ProtoMessage()               {}
func (*ReasonCode) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *ReasonCode) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Report struct {
	Display   string `protobuf:"bytes,1,opt,name=display" json:"display,omitempty"`
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *Report) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Report) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type PPerformer struct {
	Role  *Role  `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Actor *Actor `protobuf:"bytes,2,opt,name=actor" json:"actor,omitempty"`
}

func (m *PPerformer) Reset()                    { *m = PPerformer{} }
func (m *PPerformer) String() string            { return proto.CompactTextString(m) }
func (*PPerformer) ProtoMessage()               {}
func (*PPerformer) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *PPerformer) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *PPerformer) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

type FollowUp struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *FollowUp) Reset()                    { *m = FollowUp{} }
func (m *FollowUp) String() string            { return proto.CompactTextString(m) }
func (*FollowUp) ProtoMessage()               {}
func (*FollowUp) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *FollowUp) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PerformedPeriod struct {
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *PerformedPeriod) Reset()                    { *m = PerformedPeriod{} }
func (m *PerformedPeriod) String() string            { return proto.CompactTextString(m) }
func (*PerformedPeriod) ProtoMessage()               {}
func (*PerformedPeriod) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *PerformedPeriod) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *PerformedPeriod) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func init() {
	proto.RegisterType((*Procedure)(nil), "buffer.Procedure")
	proto.RegisterType((*Outcome)(nil), "buffer.Outcome")
	proto.RegisterType((*ReasonCode)(nil), "buffer.ReasonCode")
	proto.RegisterType((*Report)(nil), "buffer.Report")
	proto.RegisterType((*PPerformer)(nil), "buffer.PPerformer")
	proto.RegisterType((*FollowUp)(nil), "buffer.FollowUp")
	proto.RegisterType((*PerformedPeriod)(nil), "buffer.PerformedPeriod")
}

func init() { proto.RegisterFile("procedure.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xdd, 0x8a, 0xdb, 0x3c,
	0x10, 0x25, 0x7f, 0x4e, 0x3c, 0xc9, 0xc6, 0x8b, 0xf8, 0xf8, 0x2a, 0x4a, 0x5b, 0x8c, 0x0b, 0x25,
	0x85, 0x12, 0x4a, 0x7a, 0xd5, 0xbb, 0x6e, 0x0b, 0xbd, 0x2b, 0x0d, 0xca, 0xf6, 0x01, 0x1c, 0x6b,
	0x0c, 0x2e, 0xb6, 0xc7, 0x8c, 0x65, 0xba, 0x79, 0xe6, 0xbe, 0x44, 0xb1, 0x6c, 0xc5, 0xeb, 0x65,
	0xef, 0xac, 0x73, 0xce, 0x1c, 0x9f, 0xd1, 0x8c, 0x20, 0xa8, 0x98, 0x12, 0xd4, 0x0d, 0xe3, 0xbe,
	0x62, 0x32, 0x24, 0xbc, 0x73, 0x93, 0xa6, 0xc8, 0x2f, 0x37, 0x09, 0x15, 0x05, 0x95, 0x1d, 0x1a,
	0xfd, 0x9d, 0x83, 0x7f, 0x74, 0x4a, 0xf1, 0x3f, 0x78, 0xb5, 0x89, 0x4d, 0x53, 0xcb, 0x49, 0x38,
	0xd9, 0xf9, 0xaa, 0x3f, 0x89, 0x3b, 0x08, 0x2a, 0xe4, 0x94, 0xb8, 0x40, 0x7d, 0x44, 0xce, 0x48,
	0xcb, 0x69, 0x38, 0xd9, 0xad, 0x0f, 0x2f, 0xf6, 0x9d, 0xeb, 0xfe, 0x38, 0xa6, 0xd5, 0x53, 0xbd,
	0x08, 0x61, 0x9e, 0x90, 0x46, 0x39, 0xb3, 0x75, 0x1b, 0x57, 0xf7, 0x8d, 0x34, 0x2a, 0xcb, 0x88,
	0x08, 0x36, 0x8c, 0x35, 0x35, 0x9c, 0xe0, 0xfd, 0xa5, 0x42, 0x39, 0xb7, 0x11, 0x46, 0x58, 0xeb,
	0x62, 0xf0, 0xc1, 0xc8, 0xc5, 0xd8, 0xe5, 0x1e, 0x1f, 0x8c, 0xb2, 0x8c, 0xf8, 0x00, 0xab, 0x94,
	0xf2, 0x9c, 0xfe, 0xfc, 0xaa, 0xa4, 0x17, 0xce, 0x76, 0xeb, 0xc3, 0xad, 0x53, 0x7d, 0xef, 0x71,
	0x75, 0x55, 0x88, 0x8f, 0xe0, 0xbb, 0xa0, 0x2c, 0x97, 0x56, 0x2e, 0xae, 0x2d, 0xb9, 0x9e, 0x58,
	0x0d, 0xa2, 0xd6, 0xff, 0x4c, 0xfa, 0x72, 0xca, 0x0c, 0xca, 0xd5, 0xd8, 0xff, 0x6b, 0x8f, 0xab,
	0xab, 0x42, 0xbc, 0x87, 0x65, 0x42, 0xa5, 0x8d, 0xec, 0xdb, 0xc8, 0xc1, 0xd0, 0xb8, 0x85, 0x95,
	0xe3, 0xc5, 0x3b, 0xf0, 0x18, 0x2b, 0x62, 0x23, 0xc1, 0xda, 0x6e, 0x9d, 0x52, 0x59, 0x54, 0xf5,
	0xac, 0x38, 0x00, 0x30, 0xc6, 0x35, 0x95, 0xed, 0xd5, 0xc9, 0xf5, 0x38, 0xb3, 0xba, 0x32, 0xea,
	0x91, 0xaa, 0x8d, 0x41, 0x8d, 0x49, 0xa8, 0x40, 0xb9, 0x19, 0xc7, 0xf8, 0xd9, 0xc1, 0xca, 0xf1,
	0x62, 0x0b, 0xd3, 0x4c, 0xcb, 0x1b, 0x7b, 0xf7, 0xd3, 0x4c, 0xb7, 0xa5, 0x75, 0x73, 0xfe, 0x8d,
	0x89, 0x91, 0xdb, 0x71, 0xe9, 0xa9, 0x83, 0x95, 0xe3, 0xdb, 0xe1, 0x14, 0x68, 0x62, 0x19, 0x8c,
	0x87, 0xf3, 0x03, 0x4d, 0xac, 0x2c, 0x13, 0xbd, 0x86, 0x65, 0xff, 0x43, 0x21, 0xfa, 0x49, 0x76,
	0x8b, 0x66, 0xbf, 0xa3, 0x10, 0x60, 0x68, 0xe0, 0x59, 0xc5, 0x17, 0xf0, 0xba, 0xeb, 0x10, 0x12,
	0x96, 0x3a, 0xab, 0xab, 0x3c, 0xbe, 0xf4, 0x02, 0x77, 0x14, 0xaf, 0xc0, 0x67, 0x4c, 0x91, 0xb1,
	0x4c, 0xd0, 0xae, 0xa9, 0xaf, 0x06, 0x20, 0x3a, 0x01, 0x0c, 0x83, 0x6d, 0x23, 0x33, 0xe5, 0x68,
	0x2d, 0x1e, 0x45, 0x56, 0x94, 0xa3, 0xb2, 0x8c, 0x78, 0x0b, 0x8b, 0x38, 0x31, 0xc4, 0xfd, 0xc2,
	0xdf, 0x38, 0xc9, 0x5d, 0x0b, 0xaa, 0x8e, 0x8b, 0xde, 0xc0, 0xca, 0x2d, 0xd7, 0xb3, 0xb1, 0x3f,
	0x43, 0xf0, 0xe4, 0x81, 0x88, 0xff, 0x60, 0x51, 0x9b, 0x98, 0x9d, 0xae, 0x3b, 0x88, 0x5b, 0x98,
	0x61, 0xa9, 0xfb, 0xd4, 0xed, 0xe7, 0xd9, 0xb3, 0xef, 0xf4, 0xd3, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x5a, 0x0b, 0x52, 0xd0, 0x03, 0x00, 0x00,
}

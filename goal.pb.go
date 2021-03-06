// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goal.proto

package buffer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Goal struct {
	Status           string              `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Category         []*Category         `protobuf:"bytes,2,rep,name=category" json:"category,omitempty"`
	OutcomeReference []*OutcomeReference `protobuf:"bytes,3,rep,name=outcomeReference" json:"outcomeReference,omitempty"`
	Addresses        []*Addresses        `protobuf:"bytes,4,rep,name=addresses" json:"addresses,omitempty"`
	Description      *Description        `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	StartDate        string              `protobuf:"bytes,6,opt,name=startDate" json:"startDate,omitempty"`
	ResourceType     string              `protobuf:"bytes,7,opt,name=resourceType" json:"resourceType,omitempty"`
	Text             *Text               `protobuf:"bytes,8,opt,name=text" json:"text,omitempty"`
	ExpressedBy      *ExpressedBy        `protobuf:"bytes,9,opt,name=expressedBy" json:"expressedBy,omitempty"`
	StatusReason     string              `protobuf:"bytes,10,opt,name=statusReason" json:"statusReason,omitempty"`
	Priority         *Priority           `protobuf:"bytes,11,opt,name=priority" json:"priority,omitempty"`
	Target           *Target             `protobuf:"bytes,12,opt,name=target" json:"target,omitempty"`
	StatusDate       string              `protobuf:"bytes,13,opt,name=statusDate" json:"statusDate,omitempty"`
	Identifier       []*Identifier       `protobuf:"bytes,14,rep,name=identifier" json:"identifier,omitempty"`
	Id               string              `protobuf:"bytes,15,opt,name=id" json:"id,omitempty"`
	Subject          *Subject            `protobuf:"bytes,16,opt,name=subject" json:"subject,omitempty"`
	Meta             *Meta               `protobuf:"bytes,17,opt,name=meta" json:"meta,omitempty"`
}

func (m *Goal) Reset()                    { *m = Goal{} }
func (m *Goal) String() string            { return proto.CompactTextString(m) }
func (*Goal) ProtoMessage()               {}
func (*Goal) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *Goal) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Goal) GetCategory() []*Category {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *Goal) GetOutcomeReference() []*OutcomeReference {
	if m != nil {
		return m.OutcomeReference
	}
	return nil
}

func (m *Goal) GetAddresses() []*Addresses {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Goal) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Goal) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Goal) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Goal) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Goal) GetExpressedBy() *ExpressedBy {
	if m != nil {
		return m.ExpressedBy
	}
	return nil
}

func (m *Goal) GetStatusReason() string {
	if m != nil {
		return m.StatusReason
	}
	return ""
}

func (m *Goal) GetPriority() *Priority {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *Goal) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Goal) GetStatusDate() string {
	if m != nil {
		return m.StatusDate
	}
	return ""
}

func (m *Goal) GetIdentifier() []*Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Goal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Goal) GetSubject() *Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Goal) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Target struct {
	DetailRange *DetailRange `protobuf:"bytes,1,opt,name=detailRange" json:"detailRange,omitempty"`
	DueDate     string       `protobuf:"bytes,2,opt,name=dueDate" json:"dueDate,omitempty"`
	Measure     *Measure     `protobuf:"bytes,3,opt,name=measure" json:"measure,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *Target) GetDetailRange() *DetailRange {
	if m != nil {
		return m.DetailRange
	}
	return nil
}

func (m *Target) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *Target) GetMeasure() *Measure {
	if m != nil {
		return m.Measure
	}
	return nil
}

type Measure struct {
	Coding []*Coding `protobuf:"bytes,1,rep,name=coding" json:"coding,omitempty"`
}

func (m *Measure) Reset()                    { *m = Measure{} }
func (m *Measure) String() string            { return proto.CompactTextString(m) }
func (*Measure) ProtoMessage()               {}
func (*Measure) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *Measure) GetCoding() []*Coding {
	if m != nil {
		return m.Coding
	}
	return nil
}

type DetailRange struct {
	High *High `protobuf:"bytes,1,opt,name=high" json:"high,omitempty"`
	Low  *Low  `protobuf:"bytes,2,opt,name=low" json:"low,omitempty"`
}

func (m *DetailRange) Reset()                    { *m = DetailRange{} }
func (m *DetailRange) String() string            { return proto.CompactTextString(m) }
func (*DetailRange) ProtoMessage()               {}
func (*DetailRange) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *DetailRange) GetHigh() *High {
	if m != nil {
		return m.High
	}
	return nil
}

func (m *DetailRange) GetLow() *Low {
	if m != nil {
		return m.Low
	}
	return nil
}

type Low struct {
	Code   string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Value  int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Unit   string `protobuf:"bytes,3,opt,name=unit" json:"unit,omitempty"`
	System string `protobuf:"bytes,4,opt,name=system" json:"system,omitempty"`
}

func (m *Low) Reset()                    { *m = Low{} }
func (m *Low) String() string            { return proto.CompactTextString(m) }
func (*Low) ProtoMessage()               {}
func (*Low) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *Low) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Low) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Low) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Low) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

type High struct {
	Code   string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Value  int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Unit   string `protobuf:"bytes,3,opt,name=unit" json:"unit,omitempty"`
	System string `protobuf:"bytes,4,opt,name=system" json:"system,omitempty"`
}

func (m *High) Reset()                    { *m = High{} }
func (m *High) String() string            { return proto.CompactTextString(m) }
func (*High) ProtoMessage()               {}
func (*High) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *High) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *High) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *High) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *High) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

type ExpressedBy struct {
	Display   string `protobuf:"bytes,1,opt,name=display" json:"display,omitempty"`
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *ExpressedBy) Reset()                    { *m = ExpressedBy{} }
func (m *ExpressedBy) String() string            { return proto.CompactTextString(m) }
func (*ExpressedBy) ProtoMessage()               {}
func (*ExpressedBy) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *ExpressedBy) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *ExpressedBy) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type Description struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Description) Reset()                    { *m = Description{} }
func (m *Description) String() string            { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()               {}
func (*Description) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *Description) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type OutcomeReference struct {
	Display   string `protobuf:"bytes,1,opt,name=display" json:"display,omitempty"`
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *OutcomeReference) Reset()                    { *m = OutcomeReference{} }
func (m *OutcomeReference) String() string            { return proto.CompactTextString(m) }
func (*OutcomeReference) ProtoMessage()               {}
func (*OutcomeReference) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *OutcomeReference) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *OutcomeReference) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func init() {
	proto.RegisterType((*Goal)(nil), "buffer.Goal")
	proto.RegisterType((*Target)(nil), "buffer.Target")
	proto.RegisterType((*Measure)(nil), "buffer.Measure")
	proto.RegisterType((*DetailRange)(nil), "buffer.DetailRange")
	proto.RegisterType((*Low)(nil), "buffer.Low")
	proto.RegisterType((*High)(nil), "buffer.High")
	proto.RegisterType((*ExpressedBy)(nil), "buffer.ExpressedBy")
	proto.RegisterType((*Description)(nil), "buffer.Description")
	proto.RegisterType((*OutcomeReference)(nil), "buffer.OutcomeReference")
}

func init() { proto.RegisterFile("goal.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0x3e, 0x9a, 0x34, 0xb3, 0xa1, 0x4d, 0x0d, 0x42, 0x16, 0x02, 0x14, 0xf6, 0x80, 0x8a,
	0x84, 0x8a, 0x08, 0xe2, 0x07, 0x40, 0x5b, 0xf1, 0xa1, 0x16, 0x90, 0xe9, 0x0d, 0x2e, 0xee, 0xee,
	0x64, 0x6b, 0x94, 0xac, 0x23, 0xdb, 0x4b, 0x9b, 0x3b, 0x3f, 0x8c, 0x9f, 0x86, 0x3c, 0x6b, 0x6f,
	0x36, 0xe5, 0x86, 0xc4, 0xcd, 0x7e, 0xef, 0x79, 0xe6, 0xcd, 0x8c, 0x6d, 0x80, 0x42, 0xcb, 0xc5,
	0xd1, 0xca, 0x68, 0xa7, 0xd9, 0xe0, 0xb2, 0x9a, 0xcf, 0xd1, 0x3c, 0x18, 0x67, 0x7a, 0xb9, 0xd4,
	0x65, 0x8d, 0xa6, 0xbf, 0x77, 0xa0, 0xff, 0x4e, 0xcb, 0x05, 0xbb, 0x0f, 0x03, 0xeb, 0xa4, 0xab,
	0x2c, 0xef, 0x4c, 0x3b, 0x87, 0x23, 0x11, 0x76, 0xec, 0x39, 0xec, 0x66, 0xd2, 0x61, 0xa1, 0xcd,
	0x9a, 0x77, 0xa7, 0xbd, 0xc3, 0x64, 0x36, 0x39, 0xaa, 0x23, 0x1d, 0x1d, 0x07, 0x5c, 0x34, 0x0a,
	0x76, 0x02, 0x13, 0x5d, 0xb9, 0x4c, 0x2f, 0x51, 0xe0, 0x1c, 0x0d, 0x96, 0x19, 0xf2, 0x1e, 0x9d,
	0xe2, 0xf1, 0xd4, 0xe7, 0x5b, 0xbc, 0xf8, 0xeb, 0x04, 0x7b, 0x01, 0x23, 0x99, 0xe7, 0x06, 0xad,
	0x45, 0xcb, 0xfb, 0x74, 0xfc, 0x20, 0x1e, 0x7f, 0x13, 0x09, 0xb1, 0xd1, 0xb0, 0xd7, 0x90, 0xe4,
	0x68, 0x33, 0xa3, 0x56, 0x4e, 0xe9, 0x92, 0xef, 0x4c, 0x3b, 0x87, 0xc9, 0xec, 0x6e, 0x3c, 0x72,
	0xb2, 0xa1, 0x44, 0x5b, 0xc7, 0x1e, 0xc2, 0xc8, 0x3a, 0x69, 0xdc, 0x89, 0x74, 0xc8, 0x07, 0x54,
	0xf6, 0x06, 0x60, 0x29, 0x8c, 0x0d, 0x5a, 0x5d, 0x99, 0x0c, 0x2f, 0xd6, 0x2b, 0xe4, 0x43, 0x12,
	0x6c, 0x61, 0x6c, 0x0a, 0x7d, 0x87, 0x37, 0x8e, 0xef, 0x52, 0xc6, 0x71, 0xcc, 0x78, 0x81, 0x37,
	0x4e, 0x10, 0xe3, 0xad, 0xe1, 0xcd, 0x8a, 0x7c, 0xe6, 0x6f, 0xd7, 0x7c, 0xb4, 0x6d, 0xed, 0x74,
	0x43, 0x89, 0xb6, 0xce, 0x27, 0xaf, 0x07, 0x20, 0x50, 0x5a, 0x5d, 0x72, 0xa8, 0x93, 0xb7, 0x31,
	0x3f, 0x9a, 0x95, 0x51, 0xda, 0x28, 0xb7, 0xe6, 0x09, 0xc5, 0x6d, 0x46, 0xf3, 0x25, 0xe0, 0xa2,
	0x51, 0xb0, 0xa7, 0x30, 0x70, 0xd2, 0x14, 0xe8, 0xf8, 0x98, 0xb4, 0x7b, 0x8d, 0x59, 0x42, 0x45,
	0x60, 0xd9, 0x63, 0x80, 0x3a, 0x0b, 0x75, 0xe5, 0x0e, 0xe5, 0x6d, 0x21, 0x6c, 0x06, 0xa0, 0x72,
	0x2c, 0x9d, 0x9a, 0x2b, 0x34, 0x7c, 0x8f, 0xa6, 0xc3, 0x62, 0xac, 0x0f, 0x0d, 0x23, 0x5a, 0x2a,
	0xb6, 0x07, 0x5d, 0x95, 0xf3, 0x7d, 0x8a, 0xd5, 0x55, 0x39, 0x7b, 0x06, 0x43, 0x5b, 0x5d, 0xfe,
	0xc0, 0xcc, 0xf1, 0x09, 0x99, 0xd9, 0x8f, 0x01, 0xbe, 0xd6, 0xb0, 0x88, 0xbc, 0xef, 0xf0, 0x12,
	0x9d, 0xe4, 0x07, 0xdb, 0x1d, 0x3e, 0x47, 0x27, 0x05, 0x31, 0xe9, 0xaf, 0x0e, 0x0c, 0xea, 0x1a,
	0xea, 0x7b, 0xe0, 0xa4, 0x5a, 0x08, 0x59, 0x16, 0x48, 0x37, 0x79, 0xeb, 0x1e, 0x34, 0x94, 0x68,
	0xeb, 0x18, 0x87, 0x61, 0x5e, 0x21, 0xd5, 0xdb, 0x25, 0x8f, 0x71, 0xeb, 0x8d, 0x2e, 0x51, 0xda,
	0xca, 0xf8, 0x6b, 0xbc, 0x65, 0xf4, 0xbc, 0x86, 0x45, 0xe4, 0xd3, 0x97, 0x30, 0x0c, 0x98, 0x6f,
	0x75, 0xa6, 0x73, 0x55, 0x16, 0xbc, 0x43, 0xed, 0x69, 0x5a, 0x7d, 0x4c, 0xa8, 0x08, 0x6c, 0xfa,
	0x09, 0x92, 0x96, 0x27, 0x5f, 0xea, 0x95, 0x2a, 0xae, 0x82, 0xed, 0xa6, 0xd4, 0xf7, 0xaa, 0xb8,
	0x12, 0xc4, 0xb0, 0x47, 0xd0, 0x5b, 0xe8, 0x6b, 0x32, 0x99, 0xcc, 0x92, 0x28, 0x38, 0xd3, 0xd7,
	0xc2, 0xe3, 0xe9, 0x37, 0xe8, 0x9d, 0xe9, 0x6b, 0xc6, 0xa0, 0x9f, 0xe9, 0x1c, 0xc3, 0x43, 0xa6,
	0x35, 0xbb, 0x07, 0x3b, 0x3f, 0xe5, 0xa2, 0xaa, 0x0b, 0xec, 0x89, 0x7a, 0xe3, 0x95, 0x55, 0xa9,
	0x1c, 0xd5, 0x36, 0x12, 0xb4, 0xa6, 0x8f, 0x60, 0x6d, 0x1d, 0x2e, 0x79, 0x3f, 0x7c, 0x04, 0xb4,
	0x4b, 0xbf, 0x43, 0xdf, 0x3b, 0xf9, 0x4f, 0xd1, 0x4f, 0x21, 0x69, 0xbd, 0x05, 0x9a, 0x88, 0xb2,
	0xab, 0x85, 0x5c, 0x87, 0x3c, 0x71, 0xeb, 0xdf, 0xac, 0x69, 0xbe, 0x96, 0x7a, 0x5a, 0x1b, 0x20,
	0x7d, 0xe2, 0x3b, 0xba, 0x79, 0xe0, 0x2c, 0x3c, 0xcf, 0xe0, 0xd5, 0xaf, 0xd3, 0x8f, 0x30, 0xb9,
	0xfd, 0x05, 0xfd, 0x6b, 0xba, 0xcb, 0x01, 0x7d, 0xa2, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x0c, 0x80, 0x30, 0x68, 0x05, 0x00, 0x00,
}

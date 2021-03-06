// Code generated by protoc-gen-go. DO NOT EDIT.
// source: allergyIntolerance.proto

/*
Package buffer is a generated protocol buffer package.

It is generated from these files:
	allergyIntolerance.proto
	bundle.proto
	careplan.proto
	common.proto
	condition.proto
	diagnosticReport.proto
	encounter.proto
	fhirServer.proto
	goal.proto
	immunization.proto
	medication.proto
	medicationRequest.proto
	observation.proto
	organization.proto
	patient.proto
	practitioner.proto
	procedure.proto

It has these top-level messages:
	AllergyIntolerance
	Recorder
	Reaction
	Manifestation
	ExposureRoute
	Substance
	Bundle
	Entry
	CarePlan
	CareTeam
	Activity
	Detail
	DailyAmount
	ProductReference
	CContained
	CParticipant
	Member
	CGoal
	Meta
	Qualification
	Subject
	Stage
	Type
	Patient
	EncounterRef
	Requester
	Performer
	Context
	Category
	Issuer
	Code
	PractitionerRole
	HealthcareService
	Period
	Role
	Location
	Organization
	Name
	Identifier
	Assigner
	Contact
	Address
	Relationship
	Extension
	Telecom
	Communication
	Language
	Coding
	Text
	ManagingOrganization
	Summary
	Asserter
	BodySite
	Severity
	Evidence
	Link
	Reason
	BasedOn
	Note
	Contained
	Actor
	Priority
	Addresses
	ConditionData
	DiagnosticReport
	Result
	PresentedForm
	DiagnosticPerformer
	Encounter
	PartOf
	EpisodeOfCare
	Appointment
	ServiceProvider
	Diagnosis
	Condition
	Class
	IncomingReferral
	Hospitalization
	Destination
	ReAdmission
	DietPreference
	SpecialArrangement
	SpecialCourtesy
	AdmitSource
	Origin
	Participant
	Individual
	StatusHistory
	Account
	ID
	QueryString
	ReferenceIDs
	FHIRUpdateResource
	Goal
	Target
	Measure
	DetailRange
	Low
	High
	ExpressedBy
	Description
	OutcomeReference
	Immunization
	ImmuEncounter
	VaccineCode
	Medication
	MedicationRequest
	DosageInstruction
	Method
	Timing
	Repeat
	BoundsPeriod
	DoseQuantity
	Route
	MRequester
	Agent
	OnBehalfOf
	MedicationReference
	Substitution
	DispenseRequest
	Quantity
	ExpectedSupplyDuration
	ValidityPeriod
	GroupIdentifier
	Definition
	Observation
	Component
	ValueQuantity
	OrganizationData
	Endpoint
	PatientData
	MaritalStatus
	PractitionerData
	Photo
	Procedure
	Outcome
	ReasonCode
	Report
	PPerformer
	FollowUp
	PerformedPeriod
*/
package buffer

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

type AllergyIntolerance struct {
	Status         string        `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Category       string        `protobuf:"bytes,2,opt,name=category" json:"category,omitempty"`
	Code           *Code         `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	Patient        *Patient      `protobuf:"bytes,4,opt,name=patient" json:"patient,omitempty"`
	Criticality    string        `protobuf:"bytes,5,opt,name=criticality" json:"criticality,omitempty"`
	Asserter       *Asserter     `protobuf:"bytes,6,opt,name=asserter" json:"asserter,omitempty"`
	ResourceType   string        `protobuf:"bytes,7,opt,name=resourceType" json:"resourceType,omitempty"`
	Text           *Text         `protobuf:"bytes,8,opt,name=text" json:"text,omitempty"`
	Reaction       []*Reaction   `protobuf:"bytes,9,rep,name=reaction" json:"reaction,omitempty"`
	LastOccurrence string        `protobuf:"bytes,10,opt,name=lastOccurrence" json:"lastOccurrence,omitempty"`
	Onset          string        `protobuf:"bytes,11,opt,name=onset" json:"onset,omitempty"`
	Note           []*Note       `protobuf:"bytes,12,rep,name=note" json:"note,omitempty"`
	AttestedDate   string        `protobuf:"bytes,13,opt,name=attestedDate" json:"attestedDate,omitempty"`
	Recorder       *Recorder     `protobuf:"bytes,14,opt,name=recorder" json:"recorder,omitempty"`
	Identifier     []*Identifier `protobuf:"bytes,15,rep,name=identifier" json:"identifier,omitempty"`
	Type           string        `protobuf:"bytes,16,opt,name=type" json:"type,omitempty"`
	Id             string        `protobuf:"bytes,17,opt,name=id" json:"id,omitempty"`
	Meta           *Meta         `protobuf:"bytes,18,opt,name=meta" json:"meta,omitempty"`
}

func (m *AllergyIntolerance) Reset()                    { *m = AllergyIntolerance{} }
func (m *AllergyIntolerance) String() string            { return proto.CompactTextString(m) }
func (*AllergyIntolerance) ProtoMessage()               {}
func (*AllergyIntolerance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllergyIntolerance) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AllergyIntolerance) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *AllergyIntolerance) GetCode() *Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *AllergyIntolerance) GetPatient() *Patient {
	if m != nil {
		return m.Patient
	}
	return nil
}

func (m *AllergyIntolerance) GetCriticality() string {
	if m != nil {
		return m.Criticality
	}
	return ""
}

func (m *AllergyIntolerance) GetAsserter() *Asserter {
	if m != nil {
		return m.Asserter
	}
	return nil
}

func (m *AllergyIntolerance) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *AllergyIntolerance) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *AllergyIntolerance) GetReaction() []*Reaction {
	if m != nil {
		return m.Reaction
	}
	return nil
}

func (m *AllergyIntolerance) GetLastOccurrence() string {
	if m != nil {
		return m.LastOccurrence
	}
	return ""
}

func (m *AllergyIntolerance) GetOnset() string {
	if m != nil {
		return m.Onset
	}
	return ""
}

func (m *AllergyIntolerance) GetNote() []*Note {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *AllergyIntolerance) GetAttestedDate() string {
	if m != nil {
		return m.AttestedDate
	}
	return ""
}

func (m *AllergyIntolerance) GetRecorder() *Recorder {
	if m != nil {
		return m.Recorder
	}
	return nil
}

func (m *AllergyIntolerance) GetIdentifier() []*Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *AllergyIntolerance) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AllergyIntolerance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AllergyIntolerance) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Recorder struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *Recorder) Reset()                    { *m = Recorder{} }
func (m *Recorder) String() string            { return proto.CompactTextString(m) }
func (*Recorder) ProtoMessage()               {}
func (*Recorder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Recorder) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type Reaction struct {
	Substance     *Substance       `protobuf:"bytes,1,opt,name=substance" json:"substance,omitempty"`
	Severity      string           `protobuf:"bytes,2,opt,name=severity" json:"severity,omitempty"`
	Certainty     string           `protobuf:"bytes,3,opt,name=certainty" json:"certainty,omitempty"`
	ExposureRoute *ExposureRoute   `protobuf:"bytes,4,opt,name=exposureRoute" json:"exposureRoute,omitempty"`
	Onset         string           `protobuf:"bytes,5,opt,name=onset" json:"onset,omitempty"`
	Manifestation []*Manifestation `protobuf:"bytes,6,rep,name=manifestation" json:"manifestation,omitempty"`
	Description   string           `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
}

func (m *Reaction) Reset()                    { *m = Reaction{} }
func (m *Reaction) String() string            { return proto.CompactTextString(m) }
func (*Reaction) ProtoMessage()               {}
func (*Reaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Reaction) GetSubstance() *Substance {
	if m != nil {
		return m.Substance
	}
	return nil
}

func (m *Reaction) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *Reaction) GetCertainty() string {
	if m != nil {
		return m.Certainty
	}
	return ""
}

func (m *Reaction) GetExposureRoute() *ExposureRoute {
	if m != nil {
		return m.ExposureRoute
	}
	return nil
}

func (m *Reaction) GetOnset() string {
	if m != nil {
		return m.Onset
	}
	return ""
}

func (m *Reaction) GetManifestation() []*Manifestation {
	if m != nil {
		return m.Manifestation
	}
	return nil
}

func (m *Reaction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Manifestation struct {
	Coding []*Coding `protobuf:"bytes,1,rep,name=coding" json:"coding,omitempty"`
}

func (m *Manifestation) Reset()                    { *m = Manifestation{} }
func (m *Manifestation) String() string            { return proto.CompactTextString(m) }
func (*Manifestation) ProtoMessage()               {}
func (*Manifestation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Manifestation) GetCoding() []*Coding {
	if m != nil {
		return m.Coding
	}
	return nil
}

type ExposureRoute struct {
	Coding []*Coding `protobuf:"bytes,1,rep,name=coding" json:"coding,omitempty"`
}

func (m *ExposureRoute) Reset()                    { *m = ExposureRoute{} }
func (m *ExposureRoute) String() string            { return proto.CompactTextString(m) }
func (*ExposureRoute) ProtoMessage()               {}
func (*ExposureRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExposureRoute) GetCoding() []*Coding {
	if m != nil {
		return m.Coding
	}
	return nil
}

type Substance struct {
	Coding []*Coding `protobuf:"bytes,1,rep,name=coding" json:"coding,omitempty"`
}

func (m *Substance) Reset()                    { *m = Substance{} }
func (m *Substance) String() string            { return proto.CompactTextString(m) }
func (*Substance) ProtoMessage()               {}
func (*Substance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Substance) GetCoding() []*Coding {
	if m != nil {
		return m.Coding
	}
	return nil
}

func init() {
	proto.RegisterType((*AllergyIntolerance)(nil), "buffer.AllergyIntolerance")
	proto.RegisterType((*Recorder)(nil), "buffer.Recorder")
	proto.RegisterType((*Reaction)(nil), "buffer.Reaction")
	proto.RegisterType((*Manifestation)(nil), "buffer.Manifestation")
	proto.RegisterType((*ExposureRoute)(nil), "buffer.ExposureRoute")
	proto.RegisterType((*Substance)(nil), "buffer.Substance")
}

func init() { proto.RegisterFile("allergyIntolerance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x1a, 0x3d,
	0x14, 0x15, 0x04, 0x26, 0xcc, 0xe5, 0x27, 0x89, 0xf5, 0x7d, 0x95, 0x15, 0x75, 0x81, 0x58, 0x44,
	0x54, 0xaa, 0xa8, 0x44, 0x16, 0x5d, 0x64, 0x15, 0xb5, 0x5d, 0x64, 0x91, 0xb6, 0x72, 0xf3, 0x02,
	0xc6, 0x73, 0x41, 0x96, 0xc0, 0x46, 0xf6, 0x9d, 0x8a, 0x79, 0x80, 0x3e, 0x4b, 0x5f, 0xb3, 0xb2,
	0xe7, 0x87, 0x99, 0x76, 0x93, 0x1d, 0x3e, 0xe7, 0xdc, 0x7b, 0x7c, 0xef, 0x1c, 0x03, 0x5c, 0xee,
	0xf7, 0xe8, 0x76, 0xc5, 0x93, 0x21, 0xbb, 0x47, 0x27, 0x8d, 0xc2, 0xd5, 0xd1, 0x59, 0xb2, 0x2c,
	0xd9, 0xe4, 0xdb, 0x2d, 0xba, 0xdb, 0x89, 0xb2, 0x87, 0x83, 0x35, 0x25, 0xba, 0xf8, 0x35, 0x04,
	0xf6, 0xf8, 0x4f, 0x09, 0x7b, 0x03, 0x89, 0x27, 0x49, 0xb9, 0xe7, 0xbd, 0x79, 0x6f, 0x99, 0x8a,
	0xea, 0xc4, 0x6e, 0x61, 0xa4, 0x24, 0xe1, 0xce, 0xba, 0x82, 0xf7, 0x23, 0xd3, 0x9c, 0xd9, 0x1c,
	0x06, 0xca, 0x66, 0xc8, 0x2f, 0xe6, 0xbd, 0xe5, 0x78, 0x3d, 0x59, 0x95, 0x7e, 0xab, 0x4f, 0x36,
	0x43, 0x11, 0x19, 0xf6, 0x0e, 0x2e, 0x8f, 0x92, 0x34, 0x1a, 0xe2, 0x83, 0x28, 0xba, 0xaa, 0x45,
	0xdf, 0x4b, 0x58, 0xd4, 0x3c, 0x9b, 0xc3, 0x58, 0x39, 0x4d, 0x5a, 0xc9, 0xbd, 0xa6, 0x82, 0x0f,
	0xa3, 0x57, 0x1b, 0x62, 0xef, 0x61, 0x24, 0xbd, 0x47, 0x47, 0xe8, 0x78, 0x12, 0xbb, 0x5d, 0xd7,
	0xdd, 0x1e, 0x2b, 0x5c, 0x34, 0x0a, 0xb6, 0x80, 0x89, 0x43, 0x6f, 0x73, 0xa7, 0xf0, 0xa5, 0x38,
	0x22, 0xbf, 0x8c, 0x0d, 0x3b, 0x58, 0x18, 0x80, 0xf0, 0x44, 0x7c, 0xd4, 0x1d, 0xe0, 0x05, 0x4f,
	0x24, 0x22, 0x13, 0x3c, 0x1d, 0x4a, 0x45, 0xda, 0x1a, 0x9e, 0xce, 0x2f, 0xda, 0x9e, 0xa2, 0xc2,
	0x45, 0xa3, 0x60, 0x77, 0x30, 0xdb, 0x4b, 0x4f, 0xdf, 0x94, 0xca, 0x9d, 0x43, 0xa3, 0x90, 0x43,
	0x74, 0xfd, 0x0b, 0x65, 0xff, 0xc1, 0xd0, 0x1a, 0x8f, 0xc4, 0xc7, 0x91, 0x2e, 0x0f, 0xe1, 0x36,
	0xc6, 0x12, 0xf2, 0x49, 0xf4, 0x69, 0x6e, 0xf3, 0xd5, 0x12, 0x8a, 0xc8, 0x84, 0x99, 0x24, 0x11,
	0x7a, 0xc2, 0xec, 0xb3, 0x24, 0xe4, 0xd3, 0x72, 0xa6, 0x36, 0x56, 0xde, 0x58, 0x59, 0x97, 0xa1,
	0xe3, 0xb3, 0xee, 0x96, 0x44, 0x85, 0x8b, 0x46, 0xc1, 0xd6, 0x00, 0x3a, 0x43, 0x43, 0x7a, 0xab,
	0xd1, 0xf1, 0xab, 0xe8, 0xcc, 0x6a, 0xfd, 0x53, 0xc3, 0x88, 0x96, 0x8a, 0x31, 0x18, 0x50, 0xd8,
	0xe8, 0x75, 0x74, 0x8f, 0xbf, 0xd9, 0x0c, 0xfa, 0x3a, 0xe3, 0x37, 0x11, 0xe9, 0xeb, 0x2c, 0xcc,
	0x72, 0x40, 0x92, 0x9c, 0x75, 0x37, 0xfb, 0x8c, 0x24, 0x45, 0x64, 0x16, 0x4b, 0x18, 0xd5, 0xf7,
	0x61, 0x6f, 0x21, 0x75, 0xb8, 0xc5, 0x72, 0x65, 0x65, 0xfe, 0xce, 0xc0, 0xe2, 0x77, 0x3f, 0x48,
	0xab, 0x15, 0x7f, 0x80, 0xd4, 0xe7, 0x1b, 0x4f, 0xb2, 0x96, 0x8e, 0xd7, 0x37, 0x75, 0xf7, 0x1f,
	0x35, 0x21, 0xce, 0x9a, 0x10, 0x60, 0x8f, 0x3f, 0xd1, 0x85, 0x50, 0x55, 0x01, 0xae, 0xcf, 0xc1,
	0x57, 0xa1, 0x23, 0xa9, 0x0d, 0x15, 0x31, 0xc5, 0xa9, 0x38, 0x03, 0xec, 0x01, 0xa6, 0x78, 0x3a,
	0x5a, 0x9f, 0x3b, 0x14, 0x36, 0x27, 0xac, 0x22, 0xfc, 0x7f, 0x6d, 0xf7, 0xa5, 0x4d, 0x8a, 0xae,
	0xf6, 0xfc, 0x89, 0x87, 0xed, 0x4f, 0xfc, 0x00, 0xd3, 0x83, 0x34, 0x7a, 0x8b, 0xe1, 0x75, 0x85,
	0x4c, 0x25, 0x71, 0xe3, 0x4d, 0xcb, 0xe7, 0x36, 0x29, 0xba, 0xda, 0xf0, 0x42, 0x32, 0xf4, 0xca,
	0xe9, 0x63, 0x2c, 0x2d, 0x03, 0xdd, 0x86, 0x16, 0x1f, 0x61, 0xda, 0xe9, 0xc0, 0xee, 0x20, 0x51,
	0x36, 0xd3, 0x66, 0xc7, 0x7b, 0xd1, 0x68, 0xd6, 0x7a, 0xa3, 0xda, 0xec, 0x44, 0xc5, 0x86, 0xc2,
	0xce, 0x34, 0xaf, 0x2e, 0xbc, 0x87, 0xb4, 0xd9, 0xfa, 0x6b, 0x8b, 0x36, 0x49, 0xfc, 0x27, 0xba,
	0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x69, 0x6e, 0xb9, 0xda, 0xbb, 0x04, 0x00, 0x00,
}

package protos

type ProtocolLockDB struct {
	ProtocolId int32  `json:",omitempty,omitzero"`
	BeginDate  MxTime `json:",omitempty,omitzero"`
	EndDate    MxTime `json:",omitempty,omitzero"`
	CreateDate MxTime `json:",omitempty,omitzero"`
}

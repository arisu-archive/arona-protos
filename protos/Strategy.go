package protos

type Strategy struct {
	PlayAnimation bool `json:",omitempty,omitzero"`
	Activated bool `json:",omitempty,omitzero"`
	Values []int32 `json:",omitempty,omitzero"`
	Index int32 `json:",omitempty,omitzero"`
}

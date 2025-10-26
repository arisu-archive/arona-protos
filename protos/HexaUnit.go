package protos

type HexaUnit struct {
	SkillCardHand SkillCardHand `json:",omitempty,omitzero"`
	PlayAnimation bool `json:",omitempty,omitzero"`
}

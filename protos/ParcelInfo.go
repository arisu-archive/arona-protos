package protos

type ParcelInfo struct {
	Key         ParcelKeyPair
	Amount      int64 `json:",omitempty,omitzero"`
	Multiplier  BasisPoint
	Probability BasisPoint
}

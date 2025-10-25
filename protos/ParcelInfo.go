package protos

type ParcelInfo struct {
	Key         ParcelKeyPair
	Amount      int64
	Multiplier  BasisPoint
	Probability BasisPoint
}

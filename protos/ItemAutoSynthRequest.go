package protos

type ItemAutoSynthRequest struct {
	RequestPacket
	Protocol      Protocol
	TargetParcels []ParcelKeyPair
}

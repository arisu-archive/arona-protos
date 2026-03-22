package protos

type EventContentConcentrationRoundSkipResponse struct {
	ResponsePacket
	SaveDB         EventContentConcentrationSaveDB
	ParcelResultDB ParcelResultDB
}

package protos

type EventContentConcentrationRoundCompleteResponse struct {
	ResponsePacket
	SaveDB         EventContentConcentrationSaveDB
	ParcelResultDB ParcelResultDB
}

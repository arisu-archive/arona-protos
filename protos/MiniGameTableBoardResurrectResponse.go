package protos

type MiniGameTableBoardResurrectResponse struct {
	ResponsePacket
	Protocol       Protocol
	PlayerDB       TBGPlayerDB
	ParcelResultDB ParcelResultDB
}

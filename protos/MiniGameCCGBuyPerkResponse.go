package protos

type MiniGameCCGBuyPerkResponse struct {
	ResponsePacket
	Protocol                  Protocol
	Perks                     []int64
	ParcelResultDB            ParcelResultDB
	EventContentCollectionDBs []EventContentCollectionDB
}

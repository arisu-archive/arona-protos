package protos

type GachaSimulateCheatResponse struct {
	ResponsePacket
	Protocol            Protocol
	CharacterIdAndCount map[int64]int32
	SimulationCount     int64
	GoodsUniqueId       int64
	GoodsDevName        string
}

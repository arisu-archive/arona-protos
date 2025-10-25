package protos

type ConquestMainStoryConquerResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	ConquestTileDB ConquestTileDB
	ConquestInfoDB ConquestInfoDB
	DisplayInfos   []ConquestDisplayInfo
}

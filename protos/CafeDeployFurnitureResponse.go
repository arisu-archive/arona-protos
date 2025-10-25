package protos

type CafeDeployFurnitureResponse struct {
	ResponsePacket
	Protocol             Protocol
	CafeDB               CafeDB
	NewFurnitureServerId int64
	ChangedFurnitureDBs  []FurnitureDB
}

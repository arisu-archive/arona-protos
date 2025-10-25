package protos

type ConquestTakeEventObjectResponse struct {
	ResponsePacket
	Protocol                     Protocol
	ParcelResultDB               ParcelResultDB
	ConquestEventObjectDBWrapper ConquestEventObjectDB
}

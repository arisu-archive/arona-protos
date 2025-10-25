package protos

type CraftSelectNodeResponse struct {
	ResponsePacket
	Protocol       Protocol
	SelectedNodeDB CraftNodeDB
}

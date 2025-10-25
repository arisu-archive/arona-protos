package protos

type CraftSelectNodeRequest struct {
	RequestPacket
	Protocol      Protocol
	SlotId        int64
	LeafNodeIndex int64
}

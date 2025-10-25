package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanConferRequest struct {
	RequestPacket
	Protocol        Protocol
	MemberAccountId int64
	ConferingGrade  flatdata.ClanSocialGrade
}

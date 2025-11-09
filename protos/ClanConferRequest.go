package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanConferRequest struct {
	RequestPacket
	MemberAccountId int64                    `json:",omitempty,omitzero"`
	ConferingGrade  flatdata.ClanSocialGrade `json:",omitempty,omitzero"`
}

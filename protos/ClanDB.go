package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanDB struct {
	ClanDBId                                 int64
	ClanName                                 string
	ClanChannelName                          string
	ClanPresidentNickName                    string
	ClanPresidentRepresentCharacterUniqueId  int64
	ClanPresidentRepresentCharacterCostumeId int64
	ClanNotice                               string
	ClanMemberCount                          int64
	ClanJoinOption                           flatdata.ClanJoinOption
}

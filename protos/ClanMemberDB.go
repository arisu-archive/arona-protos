package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClanMemberDB struct {
	AccountId                   int64
	AccountLevel                int64
	AccountNickName             string
	ClanDBId                    int64
	RepresentCharacterUniqueId  int64
	RepresentCharacterCostumeId int64
	AttendanceCount             int64
	CafeComfortValue            int64
	ClanSocialGrade             flatdata.ClanSocialGrade
	JoinDate                    time.Time
	SocialGradeUpdateTime       time.Time
	LastLoginDate               time.Time
	GameLoginDate               time.Time
	AppliedDate                 time.Time
	AttachmentDB                AccountAttachmentDB
}

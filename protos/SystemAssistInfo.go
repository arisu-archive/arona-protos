package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type SystemAssistInfo struct {
	AssistType       SystemAssistType `json:",omitempty,omitzero"`
	ContentType      flatdata.ContentType
	StageId          int64 `json:",omitempty,omitzero"`
	AssistCharacters []*SystemAssistCharacter
}

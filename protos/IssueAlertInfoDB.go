package protos

import (
	"time"
)

type IssueAlertInfoDB struct {
	IssueAlertId   int32
	IssueAlertType IssueAlertTypeCode
	StartDate      time.Time
	EndDate        time.Time
	DisplayOrder   byte
	PublishId      int32
	Url            string
	Subject        string
}

package repository

import "time"

type AuditRecord struct {
	Action string `yaml:"action"`

	Object string `yaml:"object"`

	Result string `yaml:"result"`

	Time time.Time `yaml:"time"`
}

func NewAuditRecord(
	action string,
	object string,
	result string,
) AuditRecord {

	return AuditRecord{

		Action: action,

		Object: object,

		Result: result,

		Time: time.Now(),
	}
}

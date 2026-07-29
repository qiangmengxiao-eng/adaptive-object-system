package repository

type Transaction struct {
	Events []ObjectEvent

	Audits []AuditRecord

	committed bool
}

func NewTransaction() *Transaction {

	return &Transaction{

		Events: make([]ObjectEvent, 0),

		Audits: make([]AuditRecord, 0),
	}
}

func (t *Transaction) AddEvent(
	event ObjectEvent,
) {

	t.Events =
		append(
			t.Events,
			event,
		)
}

func (t *Transaction) AddAudit(
	audit AuditRecord,
) {

	t.Audits =
		append(
			t.Audits,
			audit,
		)
}

func (t *Transaction) Commit() {

	t.committed = true
}

func (t *Transaction) Rollback() {

	t.Events = nil

	t.Audits = nil

	t.committed = false
}

func (t *Transaction) Committed() bool {

	return t.committed
}

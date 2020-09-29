package ted

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"ted-processor/pkg/domain/infra/errors"
	infra "ted-processor/pkg/domain/infra/logger"
)

type tedPostgresRepo struct {
	db *sql.DB
}

func NewTedPostgresRepo(db *sql.DB) *tedPostgresRepo {
	return &tedPostgresRepo{db: db}
}

func (repo *tedPostgresRepo) StorePreTransaction(t *PreTransaction) (*PreTransaction, error) {
	stmt, err := repo.db.Prepare("INSERT INTO pre_transaction (id, from_account, to_account, value, time, device_type, status, started_at, metadata) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		infra.Logger.Errorw("error to create statement", "error", errors.Wrap(err))
		return nil, err
	}
	defer stmt.Close()

	id, _ := uuid.NewUUID()

	metadata, err := json.Marshal(t.Metadata)

	_, err = stmt.Exec(
		id.String(),
		t.FromAccount,
		t.ToAccount,
		t.Value,
		t.Time,
		t.DeviceType,
		t.Status,
		t.StartedAt,
		metadata,
	)
	if err != nil {
		infra.Logger.Errorw("error to execute statement", "error", errors.Wrap(err))
		return nil, err
	}
	return t, nil
}

func (repo *tedPostgresRepo) StoreTransactionConfirmation(t *TransactionConfirmation) (*TransactionConfirmation, error) {
	stmt, err := repo.db.Prepare(`
		insert into transaction_confirmation (id, from_account, to_account, value, time, device_type, status, end_at, metadata,payment_id) 
		values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`)
	if err != nil {
		infra.Logger.Errorw("error to create statement", "error", errors.Wrap(err))
		return nil, err
	}
	defer stmt.Close()
	id, _ := uuid.NewUUID()

	metadata, err := json.Marshal(t.Metadata)

	_, err = stmt.Exec(
		id.String(),
		t.FromAccount,
		t.ToAccount,
		t.Value,
		t.Time,
		t.DeviceType,
		t.Status,
		t.EndAt,
		metadata,
		t.PaymentId,
	)
	if err != nil {
		infra.Logger.Errorw("error to execute statement", "error", errors.Wrap(err))
		return nil, err
	}
	return t, nil
}

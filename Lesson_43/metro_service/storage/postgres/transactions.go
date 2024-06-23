package postgres

import (
	"database/sql"
	"metro-service/model"
	"github.com/pkg/errors"
)

type TransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{DB: db}
}

func (tr *TransactionRepo) Create(t *model.Transaction) error {
	if t.CardID == "" || t.Amount < 1 || t.Type == "" {
		return errors.New("error: cannot insert empty fields")
	}

	if t.Type == "debit" {
		var currentBalance int
		q := `
		SELECT SUM(CASE WHEN type = 'credit' THEN amount ELSE -amount END) AS balance
		FROM transactions
		WHERE card_id = $1
		GROUP BY card_id`
		err := tr.DB.QueryRow(q, t.CardID).Scan(&currentBalance)
		if err != nil {
			return errors.Wrap(err, "failed to fetch current balance from database")
		}

		if t.Amount > currentBalance {
			return errors.New("error: debit amount exceeds current balance")
		}
	}

	query := "insert into transactions "
	params := []interface{}{t.CardID, t.Amount, t.TerminalID, t.Type}
	if t.ID != "" {
		query += "(id, card_id, amount, terminal_id, type) values($1, $2, $3, $4, $5)"
		params = append([]interface{}{t.ID}, params...)
	} else {
		query += "(card_id, amount, terminal_id, type) values($1, $2, $3, $4)"
	}

	_, err := tr.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert transaction into database")
	}

	return nil
}

func (tr *TransactionRepo) Read(transactionID string) (*model.Transaction, error) {
	t := model.Transaction{ID: transactionID}
	
	row := tr.DB.QueryRow("select card_id, amount, terminal_id, type from transactions where id = $1", transactionID)
	err := row.Scan(&t.CardID, &t.Amount, &t.TerminalID, &t.Type)
	if err != nil {
		return nil, errors.Wrap(err, "transaction not found")
	}

	return &t, nil
}

func (tr *TransactionRepo) Update(t *model.Transaction) error {
	if t.ID == "" && t.CardID == "" && t.Amount < 1 && t.TerminalID == "" && t.Type == "" {
		return errors.New("error: no fields provided for update")
	}

	query := "update cards set"
	var filter []string
	var params = make(map[string]interface{})

	if t.CardID != "" {
		filter = append(filter, "card_id = :card_id")
		params["card_id"] = t.CardID
	}
	if t.Amount > 0 {
		filter = append(filter, "amount = :amount")
		params["amount"] = t.Amount
	}
	if t.TerminalID != "" {
		filter = append(filter, "terminal_id = :terminal_id")
		params["terminal_id"] = t.TerminalID
	}
	if t.Type != "" {
		filter = append(filter, "type = :type")
		params["type"] = t.Type
	}
	filter = append(filter, " where id = :id")
	params["id"] = t.ID

	q, p := ReplaceUpdateParams(query, filter, params)
	_, err := tr.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update transaction")
	}

	return nil
}

func (tr *TransactionRepo) Delete(transactionID string) error {
	_, err := tr.DB.Exec("delete from transactions where id = $1", transactionID)
	if err != nil {
		return errors.Wrap(err, "failed to delete transaction")
	}

	return nil
}

func (tr *TransactionRepo) FetchTransactions(f model.TransactionFilter) ([]model.Transaction, error) {
	query := "select id, card_id, amount, terminal_id, type from transactions where true"
	var filter []string
	var params = make(map[string]interface{})

	if f.CardID != "" {
		filter = append(filter, "card_id = :card_id")
		params["card_id"] = f.CardID
	}
	if f.Amount > 0 {
		filter = append(filter, "amount = :amount")
		params["amount"] = f.Amount
	}
	if f.TerminalID != "" {
		filter = append(filter, "terminal_id = :terminal_id")
		params["terminal_id"] = f.TerminalID
	}
	if f.Type != "" {
		filter = append(filter, "type = :type")
		params["type"] = f.Type
	}
	if f.Limit > 0 {
		filter = append(filter, "LIMIT :limit")
		params["limit"] = f.Limit
	}
	if f.Offset > 0 {
		filter = append(filter, "OFFSET :offset")
		params["offset"] = f.Offset
	}

	q, p := ReplaceReadParams(query, filter, params)
	rows, err := tr.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve transactions")
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		var terminalID sql.NullString
		err = rows.Scan(&t.ID, &t.CardID, &t.Amount, &terminalID, &t.Type)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read transaction")
		}
		if !terminalID.Valid {
			t.TerminalID = ""
		} else {
			t.TerminalID = terminalID.String
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

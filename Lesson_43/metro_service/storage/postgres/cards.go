package postgres

import (
	"database/sql"
	"fmt"
	"metro-service/model"

	"github.com/pkg/errors"
)

type CardRepo struct {
	DB *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{DB: db}
}

func (cr *CardRepo) Create(c *model.Card) error {
	if c.Number == "" || c.UserID == "" {
		return errors.New("error: cannot insert empty fields")
	}
	
	query := "insert into cards "
	params := []interface{}{c.Number, c.UserID}
	if c.ID != "" {
		query += "(id, number, user_id) values($1, $2, $3)"
		params = append([]interface{}{c.ID}, params...)
	} else {
		query += "(number, user_id) values($1, $2)"
	}

	_, err := cr.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert card into database")
	}

	return nil
}

func (cr *CardRepo) Read(cardID string) (*model.Card, error) {
	c := model.Card{ID: cardID}

	row := cr.DB.QueryRow("select number, user_id from cards where id = $1", cardID)
	err := row.Scan(&c.Number, &c.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "card not found")
	}

	return &c, nil
}

func (cr *CardRepo) Update(c *model.Card) error {
	if c.ID == "" && c.Number == "" && c.UserID == "" {
		return errors.New("error: no fields provided for update")
	}

	query := "update cards set"
	var filter []string
	var params = make(map[string]interface{})

	if c.Number != "" {
		filter = append(filter, "number = :number")
		params["number"] = c.Number
	}
	if c.UserID != "" {
		filter = append(filter, "user_id = :user_id")
		params["user_id"] = c.UserID
	}
	filter = append(filter, " where id = :id")
	params["id"] = c.ID

	q, p := ReplaceUpdateParams(query, filter, params)
	_, err := cr.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update card")
	}

	return nil
}

func (cr *CardRepo) Delete(cardID string) error {
	_, err := cr.DB.Exec("delete from cards where id = $1", cardID)
	if err != nil {
		return errors.Wrap(err, "failed to delete card")
	}

	return nil
}

func (cr *CardRepo) FetchCards(f model.CardFilter) ([]model.Card, error) {
	query := "select id, number, user_id from cards where true"
	var filter []string
	var params = make(map[string]interface{})

	if f.UserID != "" {
		filter = append(filter, "user_id = :user_id")
		params["user_id"] = f.UserID
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
	rows, err := cr.DB.Query(q, p...)
	fmt.Println(q, p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve cards")
	}
	defer rows.Close()

	var cards []model.Card
	for rows.Next() {
		var c model.Card
		err = rows.Scan(&c.ID, &c.Number, &c.UserID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read card")
		}
		cards = append(cards, c)
	}

	return cards, nil
}

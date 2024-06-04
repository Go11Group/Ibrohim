package postgres

import (
	"Mod_Index/model"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) ExplainQuery(query string, params []interface{}) error {
	res, err := p.DB.Query("explain (analyze) " + query, params...)
	if err != nil {
        return fmt.Errorf("error explaining query: %v", err)
    }
	defer res.Close()

	for res.Next() {
        var explanation string
        if err := res.Scan(&explanation); err != nil {
            return fmt.Errorf("error scanning row: %v", err)
        }
        fmt.Println(explanation)
    }
    return nil
}

func (p *ProductRepo) GetProduct(filter model.Product) ([]model.Product, error) {
	query := "select * from product where 1=1"
	var params []interface{}
	if filter.ID != uuid.Nil {
		query += " and id = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.ID)
	}
	if filter.Name != "" {
		query += " and name = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.Name)
	}
	if filter.Category != "" {
		query += " and category = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.Category)
	}
	if filter.Cost >= 0 {
		query += " and cost = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.Cost)
	}

	rows, err := p.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []model.Product
	for rows.Next() {
		var pr model.Product
		err = rows.Scan(&pr.ID, &pr.Name, &pr.Category, &pr.Cost)
		if err != nil {
			return nil, err
		}
		products = append(products, pr)
	}
	return products, nil
}

func (p *ProductRepo) CreateProduct(pr model.Product) error {
	_, err := p.DB.Exec("insert into product(name, category, cost) values($1, $2, $3)",
	pr.Name, pr.Category, pr.Cost)
	if err != nil {
		return fmt.Errorf("error creating a product")
	}
	return nil
}

func (p *ProductRepo) CreateIndexID() error {
	_, err := p.DB.Exec("create index product_id_index on product(id);")
	if err != nil {
		return fmt.Errorf("error creating an id index")
	}
	return nil
}

func (p *ProductRepo) CreateIndexNameCategory() error {
	_, err := p.DB.Exec("create index product_name_category_index on product (name, category);")
	if err != nil {
		return fmt.Errorf("error creating a name/category index")
	}
	return nil
}

func (p *ProductRepo) CreateIndexCost() error {
	_, err := p.DB.Exec("create index product_cost_index on product using hash (cost);")
	if err != nil {
		return fmt.Errorf("error creating a cost index")
	}
	return nil
}
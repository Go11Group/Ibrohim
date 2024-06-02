package postgres

import (
	model "User_Product/models"
	"database/sql"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) GetProduct(filter model.Product) ([]model.Product, error) {
	query := "select * from products where 1=1"
	var params []interface{}
	if filter.ID != 0 {
		query += " and id = $1"
		params = append(params, filter.ID)
	}
	if filter.Name != "" {
		query += " and name = $2"
		params = append(params, filter.Name)
	}
	if filter.Description != "" {
		query += " and description = $3"
		params = append(params, filter.Description)
	}
	if filter.Price != 0 {
		query += " and price = $4"
		params = append(params, filter.Price)
	}
	if filter.Stock_quantity != 0 {
		query += "and stock_quantity >= $5"
		params = append(params, filter.Stock_quantity)
	}

	rows, err := p.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	var pr model.Product
	for rows.Next() {
		err = rows.Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price, &pr.Stock_quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, pr)
	}
	return products, nil
}

func (p *ProductRepo) CreateProduct(product model.Product) error {
	_, err := p.DB.Exec("insert into products(name, description, price, stock_quantity) values($1,$2,$3,$4)",
	product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepo) UpdateProduct(product model.Product) error {
	query := "update products set"
	var params []interface{}
	if product.Name != "" {
		query += " name = $1"
		params = append(params, product.Name)
	}
	if product.Description != "" {
		query += " description = $2"
		params = append(params, product.Description)
	}
	if product.Price != 0 {
		query += " price = $3"
		params = append(params, product.Price)
	}
	if product.Stock_quantity != 0 {
		query += " stock_quantity = $4"
		params = append(params, product.Stock_quantity)
	}
	query += " where id = $5"
	params = append(params, product.ID)

	_, err := p.DB.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepo) DeleteProduct(id int) error {
	_, err := p.DB.Exec("delete from products where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
package postgres

import (
	"database/sql"
	"errors"
	"http_pg/model"
	"strconv"
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
	paramIndex := 1
	if filter.ID != 0 {
		query += " and id = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.ID)
		paramIndex++
	}
	if filter.Name != "" {
		query += " and name = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Name)
		paramIndex++
	}
	if filter.Description != "" {
		query += " and description = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Description)
		paramIndex++
	}
	if filter.Price > 0 {
		query += " and price = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Price)
		paramIndex++
	}
	if filter.Stock_quantity > 0 {
		query += "and stock_quantity >= $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Stock_quantity)
		paramIndex++
	}

	rows, err := p.DB.Query(query, params...)
	if err != nil {
		rows.Close()
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
	if product.Name == "" || product.Description == "" || product.Price <= 0 || product.Stock_quantity <= 0 {
		return errors.New("cannot insert empty fields")
	}
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
	paramIndex := 1
	updated := false
	if product.Name != "" {
		query += " name = $" + strconv.Itoa(paramIndex)
		params = append(params, product.Name)
		paramIndex++
		updated = true
	}
	if product.Description != "" {
		query += ", description = $" + strconv.Itoa(paramIndex)
		params = append(params, product.Description)
		paramIndex++
		updated = true
	}
	if product.Price > 0 {
		query += ", price = $" + strconv.Itoa(paramIndex)
		params = append(params, product.Price)
		paramIndex++
		updated = true
	}
	if product.Stock_quantity > 0 {
		query += ", stock_quantity = $" + strconv.Itoa(paramIndex)
		params = append(params, product.Stock_quantity)
		paramIndex++
		updated = true
	}
	if !updated {
		return errors.New("no fields provided for update")
	}
	query += " where id = $" + strconv.Itoa(paramIndex)
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
package postgres

import (
	model "User_Product/models"
	"database/sql"
)

type UserProductRepo struct {
	DB *sql.DB
}

func NewUserProductRepo(db *sql.DB) *UserProductRepo {
	return &UserProductRepo{DB: db}
}

func (up *UserProductRepo) GetUserProducts(userID int) ([]model.Product, error) {
	tr, err := up.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	query := `
	SELECT 	p.id, p.name, p.description, p.price,
            up.quantity
    FROM users u
    JOIN user_products up ON u.id = up.user_id
    JOIN products p ON up.product_id = p.id
	WHERE u.id = $1`
	rows, err := up.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []model.Product
	for rows.Next() {
		var pr model.Product
		err := rows.Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price, &pr.Stock_quantity)
		if err != nil {
            return nil, err
        }
		products = append(products, pr)
	}
	return products, nil
}

func (up *UserProductRepo) AddProductToUser(userID, productID, quantity int) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	
	query := `insert into user_products(user_id, product_id, quantity)
	values($1, $2, $3)
	on conflict (user_id, product_id)
	do update set quantity = user_products.quantity + $3`
	_, err = up.DB.Exec(query, userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (up *UserProductRepo) UpdateProductQuantityForUser(userID, productID, quantity int) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	
	query := "update user_products set quantity = $3 where user_id = $1 and product_id = $2"
	_, err = up.DB.Exec(query, userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (up *UserProductRepo) RemoveProductFromUser(userID, productID int) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	
	_, err = up.DB.Exec("delete from user_products where user_id = $1 and product_id = $2",
	userID, productID)
	if err != nil {
		return err
	}
	return nil
}
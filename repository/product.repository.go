package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := `SELECT id, product_name, price FROM product`
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var products []model.Product
	var product model.Product

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		products = append(products, product)
	}
	rows.Close()

	return products, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) error {
	query := `INSERT INTO product (product_name, price) VALUES ($1, $2)`
	_, err := pr.connection.Exec(query, product.Name, product.Price)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	query := `SELECT id, product_name, price FROM product WHERE id = $1`
	row := pr.connection.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) DeleteProductByID(id int) error {
	query := `DELETE FROM product WHERE id = $1`
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pr *ProductRepository) UpdateProductByID(id int, product model.Product) error {
	query := `UPDATE product SET product_name = $1, price = $2 WHERE id = $3`
	_, err := pr.connection.Exec(query, product.Name, product.Price, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

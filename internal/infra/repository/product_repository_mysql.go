package repository

import (
	"database/sql"
	"github.com/sergiolucena1/kafkaGolan/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

// Create metodo
func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("insert into products (id, name, price) values (?, ?, ?)", product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // espera tudo ser executado abaixo pra rodar depois

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, product.Name, product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

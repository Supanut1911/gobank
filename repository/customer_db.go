package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB

}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db:db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	return nil, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	return nil, nil
}
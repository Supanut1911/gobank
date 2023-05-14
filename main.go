package main

import (
	"bank/repository"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)
func main() {
	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", "localhost", 5432, "postgres", "postgres", "banking", "disable")
	db, err := sqlx.Open("postgres",dataSoruce)
	if err != nil {
		panic(err)
	}

	CustomerRepository := repository.NewCustomerRepositoryDB(db)

	_ = CustomerRepository

	customer,err := CustomerRepository.GetAll()
	if err!= nil {
		panic(err)
	}

	fmt.Println(customer)

}
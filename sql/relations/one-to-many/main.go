package main

import (
	"fmt"
	"github.com/course-testing/sql/relations/one-to-many/entity"
	"github.com/course-testing/sql/relations/one-to-many/repo"
	"github.com/google/uuid"
	"log"
)

func main() {
	db, err := repo.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	var (
		customer entity.Customer
	)

	err = db.Preload("Orders").First(&customer, "id = ?", uuid.MustParse("95285f8f-4880-4258-8712-a622f413bd30")).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", customer)
}

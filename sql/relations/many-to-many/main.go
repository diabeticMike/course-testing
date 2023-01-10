package main

import (
	"fmt"
	"github.com/course-testing/sql/relations/many-to-many/entity"
	"github.com/course-testing/sql/relations/many-to-many/repo"
	"github.com/google/uuid"
	"log"
)

func main() {
	db, err := repo.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	var (
		user entity.User
		team entity.Team
	)

	err = db.Preload("Teams").First(&user, "id = ?", uuid.MustParse("95285f8f-4880-4258-8712-a622f413bd30")).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Preload("Users").First(&team, "id = ?", uuid.MustParse("75dabb89-f079-4510-94e0-cc8d23f67d19")).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", team)
}

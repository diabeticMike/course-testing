package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/course-testing/cache/redis/entity"
	redisRepo "github.com/course-testing/cache/redis/repo/redis"
	sqlRepo "github.com/course-testing/cache/redis/repo/sql"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	rConn, err := redisRepo.NewConn()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rConn.Close()
	sqlConn, err := sqlRepo.NewConn()
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlConn.AutoMigrate(&entity.User{})
	c := ctrl{rConn, sqlConn}
	r := mux.NewRouter()
	r.HandleFunc("/users", c.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", c.GetUsers).Methods(http.MethodGet)

	http.ListenAndServe(":8000", r)
}

type ctrl struct {
	rConn  *redis.Client
	dbConn *gorm.DB
}

func (c *ctrl) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.ID = uuid.New()
	err = c.dbConn.Create(&u).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.rConn.Del("users").Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Created")
}

func (c *ctrl) GetUsers(w http.ResponseWriter, r *http.Request) {
	usersStr, err := c.rConn.Get("users").Result()
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(usersStr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Hit")
		return
	}

	if !errors.Is(err, redis.Nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []entity.User
	err = c.dbConn.Table("users").Find(&users).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usersBody []byte
	usersBody, err = json.Marshal(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.rConn.Set("users", usersBody, time.Minute*2).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Miss")
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func readUsers(filepath string) ([]byte, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	usersInByte, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return usersInByte, nil
}

func writeUsers(filepath string, users []user) error {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Seek(0, 0)
	byteValue, err := json.Marshal(&users)
	if err != nil {
		return err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = f.Write(byteValue)
	if err != nil {
		return err
	}
	return nil
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := readUsers("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "users upload error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setUsers(w http.ResponseWriter, r *http.Request) {
	users := []user{}
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println(err.Error())
		return
	}

	err = writeUsers("users.json", users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	r.HandleFunc("/users", setUsers).Methods(http.MethodPost)
	http.ListenAndServe(":8000", r)
}

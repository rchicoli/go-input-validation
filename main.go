package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var emailValidation = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type Person struct {
	Email string `json:"email"`
}

func Validate(p Person) error {
	if !emailValidation.MatchString(p.Email) {
		return fmt.Errorf("NOT VALID")
	}
	return nil
}

type InputValidation interface {
	Validate(r *http.Request) error
}

var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	//var person *Person
	person := &Person{}
	err := decodeAndValidate(r, person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func decodeAndValidate(r *http.Request, p *Person) error {
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return err
	}
	defer r.Body.Close()
	return Validate(*p)
}

func main() {
	srv := &http.Server{
		Addr:    ":1234",
		Handler: f,
	}
	srv.ListenAndServe()
}

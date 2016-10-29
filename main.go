package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"runtime"
)

// Request
// { "email": "test01@example.com" }

// Response
//{ "validation":
//  {
//    "email": "test01@example.com",
//    "check": "VALID"]
//  }
//}

var emailValidation = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type Person struct {
	Email      string      `json:"email"`
	Check      string      `json:"check"`
	Additional interface{} `json:"additional,omitempty"`
}

type Validation struct {
	Person Person `json:"validation"`
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

	if r.Method != "POST" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HEALTHY"))
		return
	}

	//var person *Person
	person := &Person{}

	err := decodeAndValidate(r, person)
	if err != nil {
		person.Check = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		person.Check = "VALID"
		w.WriteHeader(http.StatusOK)
	}
	var buf []byte
	validation := &Validation{*person}

	buf, _ = json.Marshal(validation)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(buf))
}

func decodeAndValidate(r *http.Request, p *Person) error {
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return err
	}
	defer r.Body.Close()
	return Validate(*p)
}

func main() {
	runtime.GOMAXPROCS(4)

	srv := &http.Server{
		Addr:    ":1234",
		Handler: f,
	}
	srv.ListenAndServe()
}

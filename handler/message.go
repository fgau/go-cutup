package handler

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type Payload struct {
	Cutup string `json:"cutup" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(payload Payload) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(Payload)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate the form data
	errors := ValidateStruct(*payload)
	b, err := json.Marshal(errors)
	if err != nil {
		log.Fatal(err)
	}

	if errors == nil {
		mylist := strings.Split(strings.TrimSuffix(payload.Cutup, "\n"), "\n")

		// randomize/shuffle the array of strings
		rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		rand.Shuffle(len(mylist), func(i, j int) { mylist[i], mylist[j] = mylist[j], mylist[i] })

		payload.Cutup = strings.Join(mylist, "\n")

		// write back the cutup text as json
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(&payload)
	} else {
		// write back the error json
		w.Header().Set("content-type", "application/json")
		w.Write(b)
	}

}

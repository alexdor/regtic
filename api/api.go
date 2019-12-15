// +build api

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexdor/regtic/api/handlers"
	"github.com/julienschmidt/httprouter"
)

func init() {
	go Start()
}

func find(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.URL.Query().Get("name")
	ctx := context.Background()
	headers := map[string]string{}
	res, err := handlers.FindCompanies(ctx, name, headers)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write([]byte(res.Body))
	if err != nil {
		fmt.Println(err)
	}
}
func validate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.URL.Query().Get("id")
	ctx := context.Background()
	headers := map[string]string{}

	res, err := handlers.ValidateCompanyHandler(ctx, id, headers)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(res.Body))
	if err != nil {
		fmt.Println(err)
	}
}
func Start() {
	router := httprouter.New()
	router.GET("/v2/find_companies", find)
	router.GET("/v2/validate_company", validate)

	log.Fatal(http.ListenAndServe(":8090", router))
}

package main

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "log"
  "fmt"
)

type Company struct {
	Name  string `json:"name"`
	Tel   string `json:"tel"`
	Email string `json:"email"`
  Address *Address `json:"address"`
}

type Address struct {
  City string `json:"city"`
  State string `json:"state"`
}

var companies []Company

func DisplayAll(w http.ResponseWriter, r *http.Request){
  fmt.Println("opened localhost")
  json.NewEncoder(w).Encode(companies)
}

func GetACompany(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  for _, company := range companies {
    if company.Name == params["name"] {
      json.NewEncoder(w).Encode(company)
      return
    }
  }
  json.NewEncoder(w).Encode(&Company{})
}

func CreateACompany(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  var company Company
  _ = json.NewDecoder(r.Body).Decode(&company)
  company.Name = params["name"]
  companies = append(companies, company)
  json.NewEncoder(w).Encode(companies)
}

func DeleteACompany(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for i, company := range companies {
    if company.Name == params["name"] {

      companies = append(companies[:i],companies[i+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(companies)
}


func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/companies", http.StatusSeeOther)
}

func main() {
    router := mux.NewRouter()
    companies = append(companies, Company{Name: "apple", Tel: "02040238483", Email:"apple.com", Address: &Address{City: "City X", State: "State X"}})
    companies = append(companies, Company{Name: "google", Tel: "4567893256", Email:"google.com", Address: &Address{City: "City X", State: "State X"}})
    router.HandleFunc("/", index).Methods("GET")
    router.HandleFunc("/companies", DisplayAll).Methods("GET")
    router.HandleFunc("/companies/{name}", GetACompany).Methods("GET")
    router.HandleFunc("/companies/new/{name}/{tel}/{email}", CreateACompany).Methods("GET")
    router.HandleFunc("/companies/delete/{name}", DeleteACompany).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

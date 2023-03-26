package main

import (
	"Users\rahmad novriandi\src\endcoding\json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	Id     int
	Title  string
	Auther string
	Desc   string
}

var employees = []Employee{
	{ID: 1, Title: "Golang", Auther: "Gopher", Desc: "A Book For Go"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployess)
	fmt.println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWrite, r *http.request) {
	w.header().set("Content-Type", "application/json")
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

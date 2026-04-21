package main

import (
	"net/http"

	"expense-tracker/internal/db"
	"expense-tracker/internal/handler"
	"expense-tracker/internal/repository"
	"expense-tracker/internal/service"
)

func main() {

	database := db.InitDB()

	repo := repository.NewSQLiteRepo(database)
	svc := service.NewExpenseService(repo)
	h := &handler.ExpenseHandler{Service: svc}

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h.Create(w, r)
		} else {
			h.Get(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}

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

	// allowed origins (your Netlify frontend)
	allowed := map[string]bool{
		"https://exptrackertool.netlify.app": true,
		"http://localhost:5173":              true,
	}

	// wrap default mux
	handler := CORS(allowed)(http.DefaultServeMux)

	http.ListenAndServe(":8080", handler)
}

func CORS(allowed map[string]bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			origin := r.Header.Get("Origin")

			if allowed[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

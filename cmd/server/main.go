package main

import (
	"log"
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
		log.Println("➡️ /expenses hit", r.Method)
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)

		case http.MethodGet:
			h.Get(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/expenses/summary", func(w http.ResponseWriter, r *http.Request) {

		log.Println("➡️ /expenses/summary hit", r.Method)

		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		h.Summary(w, r)
	})

	allowed := map[string]bool{
		"https://exptrackertool.netlify.app": true,
		"http://localhost:5173":              true,
	}

	corsHandler := CORS(allowed)(http.DefaultServeMux)

	http.ListenAndServe(":8080", corsHandler)
}

func CORS(allowed map[string]bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			origin := r.Header.Get("Origin")

			if origin == "" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else if allowed[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

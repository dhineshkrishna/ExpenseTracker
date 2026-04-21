package handler

import (
	"encoding/json"
	"net/http"

	"expense-tracker/internal/domain"
	"expense-tracker/internal/service"
)

type ExpenseHandler struct {
	Service *service.ExpenseService
}

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {

	var e domain.Expense
	json.NewDecoder(r.Body).Decode(&e)

	exp, err := h.Service.CreateExpense(e)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(exp)
}

func (h *ExpenseHandler) Get(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	category := q.Get("category")
	from := q.Get("from")
	to := q.Get("to")
	sort := q.Get("sort")

	expenses, total, err := h.Service.GetExpenses(category, from, to, sort)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"total":    total,
		"expenses": expenses,
	})
}
func (h *ExpenseHandler) Summary(w http.ResponseWriter, r *http.Request) {

	data, err := h.Service.GetSummary()
	if err != nil {
		http.Error(w, "failed to fetch summary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

package repository

import (
	"database/sql"
	"expense-tracker/internal/domain"
	"strings"
)

type SQLiteRepo struct {
	DB *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{DB: db}
}
func (r *SQLiteRepo) Create(e domain.Expense) error {
	_, err := r.DB.Exec(`
		INSERT OR IGNORE INTO expenses
		(id, amount, category, description, expense_date, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, e.ID, e.Amount, e.Category, e.Description, e.ExpenseDate, e.CreatedAt)

	return err
}

func (r *SQLiteRepo) Get(category, from, to, sort string) ([]domain.Expense, float64, error) {

	conditions := []string{"1=1"}
	args := []interface{}{}

	if category != "" {
		conditions = append(conditions, "category = ?")
		args = append(args, category)
	}

	if from != "" {
		conditions = append(conditions, "expense_date >= ?")
		args = append(args, from)
	}

	if to != "" {
		conditions = append(conditions, "expense_date <= ?")
		args = append(args, to)
	}

	query := `
	SELECT id, amount, category, description, expense_date, created_at
	FROM expenses
	WHERE ` + strings.Join(conditions, " AND ")

	if sort == "date_asc" {
		query += " ORDER BY expense_date ASC"
	} else {
		query += " ORDER BY expense_date DESC"
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var expenses []domain.Expense
	var total float64

	for rows.Next() {
		var e domain.Expense
		rows.Scan(&e.ID, &e.Amount, &e.Category, &e.Description, &e.ExpenseDate, &e.CreatedAt)

		expenses = append(expenses, e)
		total += e.Amount
	}

	return expenses, total, nil
}
func (r *SQLiteRepo) GetSummary() (map[string]float64, error) {

	query := `
		SELECT category, SUM(amount)
		FROM expenses
		GROUP BY category
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)

	for rows.Next() {
		var category string
		var total float64

		err := rows.Scan(&category, &total)
		if err != nil {
			return nil, err
		}

		result[category] = total
	}

	return result, nil
}

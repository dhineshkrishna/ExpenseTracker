📘 Expense Management API — README
Overview

This is a simple Expense Management API designed to work reliably under:

network retries
page refreshes
duplicate submissions


Storage: SQLite (recommended) or in-memory Map


# Expense Tracker API (Go + Clean Architecture)

## Features
- Add expense (retry-safe, UUID-based)
- List expenses
- Filter by category
- Filter by date range
- Sort (newest / oldest)
- Total spending calculation

---

## Run

```bash
go run ./cmd/server


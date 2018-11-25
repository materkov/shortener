package analytics

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type App struct {
	db *sql.DB
}

func NewApp(connStr string) *App {
	db, _ := sql.Open("postgres", connStr)
	return &App{
		db: db,
	}
}

func (a *App) GetTotalClicks(ctx context.Context, urlId int) (result int, err error) {
	err = a.db.QueryRow("SELECT COALESCE(COUNT(*), 0) FROM analytics.click WHERE url_id = $1", urlId).Scan(&result)
	return
}

func (a *App) AddClick(ctx context.Context, urlId int) (err error) {
	_, err = a.db.Exec("INSERT INTO analytics.click(url_id, date) VALUES ($1, NOW())", urlId)
	return
}

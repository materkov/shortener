package redirecter

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type StorePersistent struct {
	db *sql.DB
}

func NewStorePersistent(connStr string) Store {
	db, _ := sql.Open("postgres", connStr)
	return &StorePersistent{db: db}
}

func (m *StorePersistent) GetByKey(ctx context.Context, id int) (url string, err error) {
	err = m.db.QueryRow("SELECT url FROM redirecter.shortened_url WHERE id = $1", id).Scan(&url)
	return
}

func (m *StorePersistent) Create(ctx context.Context, url string) (id int, err error) {
	err = m.db.QueryRow("INSERT INTO redirecter.shortened_url(url) VALUES ($1) RETURNING id", url).Scan(&id)
	return
}

package postgres

import (
	"log"
	"database/sql"
	"com.sjbabadi/snippetbox/pkg/models"
	"errors"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	sql := `INSERT INTO snippets (title, content, created_at, expires_at)
	VALUES($1, $2, now(), now() + $3 * INTERVAL '1 days') RETURNING id`

	var id int
	err := m.DB.QueryRow(sql, title, content, expires).Scan(&id)

	log.Println(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created_at, expires_at FROM snippets
	WHERE expires_at > now() AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}


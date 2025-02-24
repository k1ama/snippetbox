package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}


type SnippetModel struct {
	DB *sql.DB
}

func (m *snippetModel) Insert(title string, content string, expires int) (int, err)) {
	return 0, nil
}

func (m *snippetModel) Get(id int) (*Snippet, err) {
	return nil, nil
}

func (m *snippetModel) Latest() ([]*Snippet, err) {
	return nil, nil
}
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config     *Config
	db         *sql.DB
	studentRep *StudentRep
	groupRep   *GroupRep
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Подключение к базе данных
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Student() *StudentRep {
	if s.studentRep != nil {
		return s.studentRep
	}

	s.studentRep = &StudentRep{
		store: s,
	}

	return s.studentRep
}

func (s *Store) Group() *GroupRep {
	if s.groupRep != nil {
		return s.groupRep
	}

	s.groupRep = &GroupRep{
		store: s,
	}

	return s.groupRep
}

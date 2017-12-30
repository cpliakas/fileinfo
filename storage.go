package fileinfo

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Storage saves information about a file in a SQLite database.
type Storage struct {
	sqlite *sql.DB
}

// NewStorage returns a *Storage with an initialized SQLite database and
// open connection to it.
func NewStorage(dbfile string) (s *Storage, err error) {
	s = &Storage{}

	s.sqlite, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		return
	}

	err = s.init()
	return
}

// Close closes the connection to the SQLite database.
func (s Storage) Close() error {
	return s.sqlite.Close()
}

// Save saves the file info in the SQLite database.
func (s Storage) Save(i *Fileinfo) (err error) {

	f := []string{"name", "type", "size", "hash", "first_bytes", "last_bytes"}
	p := make([]string, len(f))
	for i := range f {
		p[i] = "?"
	}

	query := fmt.Sprintf("INSERT INTO files(%s) VALUES(%s)", strings.Join(f, ", "), strings.Join(p, ", "))
	stmt, err := s.sqlite.Prepare(query)
	if err != nil {
		return
	}

	typ, err := i.Type()
	if err != nil {
		return
	}

	hash, err := i.Hash()
	if err != nil {
		return
	}

	fb, err := i.FirstBytes()
	if err != nil {
		return
	}

	lb, err := i.LastBytes()
	if err != nil {
		return
	}

	_, err = stmt.Exec(i.Name(), typ, i.Size(), hash, fb, lb)
	return
}

// Truncate drops all data from the files table.
func (s Storage) Truncate() (err error) {
	query := "DELETE FROM files"
	_, err = s.sqlite.Exec(query)
	return
}

// init initializes the SQLite database.
func (s Storage) init() (err error) {
	query := `
	CREATE TABLE IF NOT EXISTS files (
		name VARCHAR(255) NOT NULL PRIMARY KEY,
		type VARCHAR(16),
		size INT,
		hash VARCHAR(32),
		first_bytes VARCHAR(64),
		last_bytes VARCHAR(64)
	);
	`
	_, err = s.sqlite.Exec(query)
	return
}

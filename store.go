package gpterm

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/collinvandyck/gpterm/db/query"
	"github.com/collinvandyck/gpterm/lib/errs"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
)

const DBName = "gpterm.db"

func DefaultStorePath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	res := filepath.Join(hd, ".config", "gpterm")
	return res, nil
}

func DefaultDBPath() (string, error) {
	sp, err := DefaultStorePath()
	if err != nil {
		return "", err
	}
	res := filepath.Join(sp, DBName)
	return res, nil
}

type Store struct {
	dir     string
	db      *sql.DB
	queries *query.Queries
}

type StoreOpt func(*Store)

func StoreDir(path string) StoreOpt {
	return func(s *Store) {
		s.dir = path
	}
}

func NewStore(opts ...StoreOpt) (*Store, error) {
	store := &Store{}
	for _, o := range opts {
		o(store)
	}
	if store.dir == "" {
		dir, err := DefaultStorePath()
		if err != nil {
			return nil, err
		}
		store.dir = dir
	}
	err := store.init()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) SetAPIKey(ctx context.Context, key string) error {
	err := s.queries.InsertAPIKey(ctx, key)
	switch {
	case err == nil:
		return nil
	case errs.IsUniqueConstaint(err):
	case err != nil:
		return err
	}
	return s.queries.UpdateAPIKey(ctx, key)
}

func (s *Store) GetAPIKey(ctx context.Context) (string, error) {
	res, err := s.queries.GetAPIKey(ctx)
	switch {
	case errs.IsDBNotFound(err):
	case err != nil:
		return "", err
	}
	return res, nil
}

func (s *Store) Close() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Store) init() error {
	if err := ensureDir(s.dir); err != nil {
		return err
	}
	if err := s.migrate(); err != nil {
		return err
	}
	if err := s.initDB(); err != nil {
		return err
	}
	s.queries = query.New(s.db)
	return nil
}

//go:embed db/migrations/*.sql
var FSMigrations embed.FS

func (s *Store) migrate() error {
	sourceDriver, err := iofs.New(FSMigrations, "db/migrations")
	if err != nil {
		return err
	}
	path := "sqlite3://" + s.DBPath()
	mg, err := migrate.NewWithSourceInstance("iofs", sourceDriver, path)
	if err != nil {
		return err
	}
	err = mg.Up()
	switch {
	case errors.Is(err, migrate.ErrNoChange):
	case err != nil:
		return fmt.Errorf("up: %w", err)
	}
	return nil
}

func (s *Store) initDB() error {
	db, err := sql.Open("sqlite3", s.DBPath())
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) DBPath() string {
	return filepath.Join(s.dir, DBName)
}

func ensureDir(dir string) error {
	info, err := os.Stat(dir)
	switch {
	case os.IsNotExist(err):
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	case err != nil:
		return err
	case info.IsDir():
	case !info.IsDir():
		return fmt.Errorf("%q exists but is a file", dir)
	}
	return nil
}
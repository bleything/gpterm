// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package query

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getAPIKeyStmt, err = db.PrepareContext(ctx, getAPIKey); err != nil {
		return nil, fmt.Errorf("error preparing query GetAPIKey: %w", err)
	}
	if q.getMessagesStmt, err = db.PrepareContext(ctx, getMessages); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessages: %w", err)
	}
	if q.insertAPIKeyStmt, err = db.PrepareContext(ctx, insertAPIKey); err != nil {
		return nil, fmt.Errorf("error preparing query InsertAPIKey: %w", err)
	}
	if q.insertMessageStmt, err = db.PrepareContext(ctx, insertMessage); err != nil {
		return nil, fmt.Errorf("error preparing query InsertMessage: %w", err)
	}
	if q.updateAPIKeyStmt, err = db.PrepareContext(ctx, updateAPIKey); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAPIKey: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getAPIKeyStmt != nil {
		if cerr := q.getAPIKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAPIKeyStmt: %w", cerr)
		}
	}
	if q.getMessagesStmt != nil {
		if cerr := q.getMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessagesStmt: %w", cerr)
		}
	}
	if q.insertAPIKeyStmt != nil {
		if cerr := q.insertAPIKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertAPIKeyStmt: %w", cerr)
		}
	}
	if q.insertMessageStmt != nil {
		if cerr := q.insertMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertMessageStmt: %w", cerr)
		}
	}
	if q.updateAPIKeyStmt != nil {
		if cerr := q.updateAPIKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAPIKeyStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                DBTX
	tx                *sql.Tx
	getAPIKeyStmt     *sql.Stmt
	getMessagesStmt   *sql.Stmt
	insertAPIKeyStmt  *sql.Stmt
	insertMessageStmt *sql.Stmt
	updateAPIKeyStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                tx,
		tx:                tx,
		getAPIKeyStmt:     q.getAPIKeyStmt,
		getMessagesStmt:   q.getMessagesStmt,
		insertAPIKeyStmt:  q.insertAPIKeyStmt,
		insertMessageStmt: q.insertMessageStmt,
		updateAPIKeyStmt:  q.updateAPIKeyStmt,
	}
}

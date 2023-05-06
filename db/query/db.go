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
	if q.countMessagesForConversationStmt, err = db.PrepareContext(ctx, countMessagesForConversation); err != nil {
		return nil, fmt.Errorf("error preparing query CountMessagesForConversation: %w", err)
	}
	if q.createConversationStmt, err = db.PrepareContext(ctx, createConversation); err != nil {
		return nil, fmt.Errorf("error preparing query CreateConversation: %w", err)
	}
	if q.cycleClientConfigStmt, err = db.PrepareContext(ctx, cycleClientConfig); err != nil {
		return nil, fmt.Errorf("error preparing query CycleClientConfig: %w", err)
	}
	if q.getActiveConversationStmt, err = db.PrepareContext(ctx, getActiveConversation); err != nil {
		return nil, fmt.Errorf("error preparing query GetActiveConversation: %w", err)
	}
	if q.getClientConfigStmt, err = db.PrepareContext(ctx, getClientConfig); err != nil {
		return nil, fmt.Errorf("error preparing query GetClientConfig: %w", err)
	}
	if q.getCompletionTokensStmt, err = db.PrepareContext(ctx, getCompletionTokens); err != nil {
		return nil, fmt.Errorf("error preparing query GetCompletionTokens: %w", err)
	}
	if q.getConfigStmt, err = db.PrepareContext(ctx, getConfig); err != nil {
		return nil, fmt.Errorf("error preparing query GetConfig: %w", err)
	}
	if q.getConfigValueStmt, err = db.PrepareContext(ctx, getConfigValue); err != nil {
		return nil, fmt.Errorf("error preparing query GetConfigValue: %w", err)
	}
	if q.getConversationsStmt, err = db.PrepareContext(ctx, getConversations); err != nil {
		return nil, fmt.Errorf("error preparing query GetConversations: %w", err)
	}
	if q.getCredentialStmt, err = db.PrepareContext(ctx, getCredential); err != nil {
		return nil, fmt.Errorf("error preparing query GetCredential: %w", err)
	}
	if q.getLatestMessagesStmt, err = db.PrepareContext(ctx, getLatestMessages); err != nil {
		return nil, fmt.Errorf("error preparing query GetLatestMessages: %w", err)
	}
	if q.getMessagesStmt, err = db.PrepareContext(ctx, getMessages); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessages: %w", err)
	}
	if q.getPreviousMessageForRoleStmt, err = db.PrepareContext(ctx, getPreviousMessageForRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetPreviousMessageForRole: %w", err)
	}
	if q.getPromptTokensStmt, err = db.PrepareContext(ctx, getPromptTokens); err != nil {
		return nil, fmt.Errorf("error preparing query GetPromptTokens: %w", err)
	}
	if q.getTotalTokensStmt, err = db.PrepareContext(ctx, getTotalTokens); err != nil {
		return nil, fmt.Errorf("error preparing query GetTotalTokens: %w", err)
	}
	if q.insertMessageStmt, err = db.PrepareContext(ctx, insertMessage); err != nil {
		return nil, fmt.Errorf("error preparing query InsertMessage: %w", err)
	}
	if q.insertUsageStmt, err = db.PrepareContext(ctx, insertUsage); err != nil {
		return nil, fmt.Errorf("error preparing query InsertUsage: %w", err)
	}
	if q.nextConversationStmt, err = db.PrepareContext(ctx, nextConversation); err != nil {
		return nil, fmt.Errorf("error preparing query NextConversation: %w", err)
	}
	if q.previousConversationStmt, err = db.PrepareContext(ctx, previousConversation); err != nil {
		return nil, fmt.Errorf("error preparing query PreviousConversation: %w", err)
	}
	if q.setConfigValueStmt, err = db.PrepareContext(ctx, setConfigValue); err != nil {
		return nil, fmt.Errorf("error preparing query SetConfigValue: %w", err)
	}
	if q.setSelectedConversationStmt, err = db.PrepareContext(ctx, setSelectedConversation); err != nil {
		return nil, fmt.Errorf("error preparing query SetSelectedConversation: %w", err)
	}
	if q.unsetSelectedConversationStmt, err = db.PrepareContext(ctx, unsetSelectedConversation); err != nil {
		return nil, fmt.Errorf("error preparing query UnsetSelectedConversation: %w", err)
	}
	if q.updateClientConfigStmt, err = db.PrepareContext(ctx, updateClientConfig); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateClientConfig: %w", err)
	}
	if q.updateCredentialStmt, err = db.PrepareContext(ctx, updateCredential); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCredential: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countMessagesForConversationStmt != nil {
		if cerr := q.countMessagesForConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countMessagesForConversationStmt: %w", cerr)
		}
	}
	if q.createConversationStmt != nil {
		if cerr := q.createConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createConversationStmt: %w", cerr)
		}
	}
	if q.cycleClientConfigStmt != nil {
		if cerr := q.cycleClientConfigStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cycleClientConfigStmt: %w", cerr)
		}
	}
	if q.getActiveConversationStmt != nil {
		if cerr := q.getActiveConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getActiveConversationStmt: %w", cerr)
		}
	}
	if q.getClientConfigStmt != nil {
		if cerr := q.getClientConfigStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getClientConfigStmt: %w", cerr)
		}
	}
	if q.getCompletionTokensStmt != nil {
		if cerr := q.getCompletionTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCompletionTokensStmt: %w", cerr)
		}
	}
	if q.getConfigStmt != nil {
		if cerr := q.getConfigStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getConfigStmt: %w", cerr)
		}
	}
	if q.getConfigValueStmt != nil {
		if cerr := q.getConfigValueStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getConfigValueStmt: %w", cerr)
		}
	}
	if q.getConversationsStmt != nil {
		if cerr := q.getConversationsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getConversationsStmt: %w", cerr)
		}
	}
	if q.getCredentialStmt != nil {
		if cerr := q.getCredentialStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCredentialStmt: %w", cerr)
		}
	}
	if q.getLatestMessagesStmt != nil {
		if cerr := q.getLatestMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLatestMessagesStmt: %w", cerr)
		}
	}
	if q.getMessagesStmt != nil {
		if cerr := q.getMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessagesStmt: %w", cerr)
		}
	}
	if q.getPreviousMessageForRoleStmt != nil {
		if cerr := q.getPreviousMessageForRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPreviousMessageForRoleStmt: %w", cerr)
		}
	}
	if q.getPromptTokensStmt != nil {
		if cerr := q.getPromptTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPromptTokensStmt: %w", cerr)
		}
	}
	if q.getTotalTokensStmt != nil {
		if cerr := q.getTotalTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTotalTokensStmt: %w", cerr)
		}
	}
	if q.insertMessageStmt != nil {
		if cerr := q.insertMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertMessageStmt: %w", cerr)
		}
	}
	if q.insertUsageStmt != nil {
		if cerr := q.insertUsageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertUsageStmt: %w", cerr)
		}
	}
	if q.nextConversationStmt != nil {
		if cerr := q.nextConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing nextConversationStmt: %w", cerr)
		}
	}
	if q.previousConversationStmt != nil {
		if cerr := q.previousConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing previousConversationStmt: %w", cerr)
		}
	}
	if q.setConfigValueStmt != nil {
		if cerr := q.setConfigValueStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setConfigValueStmt: %w", cerr)
		}
	}
	if q.setSelectedConversationStmt != nil {
		if cerr := q.setSelectedConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setSelectedConversationStmt: %w", cerr)
		}
	}
	if q.unsetSelectedConversationStmt != nil {
		if cerr := q.unsetSelectedConversationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unsetSelectedConversationStmt: %w", cerr)
		}
	}
	if q.updateClientConfigStmt != nil {
		if cerr := q.updateClientConfigStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateClientConfigStmt: %w", cerr)
		}
	}
	if q.updateCredentialStmt != nil {
		if cerr := q.updateCredentialStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCredentialStmt: %w", cerr)
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
	db                               DBTX
	tx                               *sql.Tx
	countMessagesForConversationStmt *sql.Stmt
	createConversationStmt           *sql.Stmt
	cycleClientConfigStmt            *sql.Stmt
	getActiveConversationStmt        *sql.Stmt
	getClientConfigStmt              *sql.Stmt
	getCompletionTokensStmt          *sql.Stmt
	getConfigStmt                    *sql.Stmt
	getConfigValueStmt               *sql.Stmt
	getConversationsStmt             *sql.Stmt
	getCredentialStmt                *sql.Stmt
	getLatestMessagesStmt            *sql.Stmt
	getMessagesStmt                  *sql.Stmt
	getPreviousMessageForRoleStmt    *sql.Stmt
	getPromptTokensStmt              *sql.Stmt
	getTotalTokensStmt               *sql.Stmt
	insertMessageStmt                *sql.Stmt
	insertUsageStmt                  *sql.Stmt
	nextConversationStmt             *sql.Stmt
	previousConversationStmt         *sql.Stmt
	setConfigValueStmt               *sql.Stmt
	setSelectedConversationStmt      *sql.Stmt
	unsetSelectedConversationStmt    *sql.Stmt
	updateClientConfigStmt           *sql.Stmt
	updateCredentialStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                               tx,
		tx:                               tx,
		countMessagesForConversationStmt: q.countMessagesForConversationStmt,
		createConversationStmt:           q.createConversationStmt,
		cycleClientConfigStmt:            q.cycleClientConfigStmt,
		getActiveConversationStmt:        q.getActiveConversationStmt,
		getClientConfigStmt:              q.getClientConfigStmt,
		getCompletionTokensStmt:          q.getCompletionTokensStmt,
		getConfigStmt:                    q.getConfigStmt,
		getConfigValueStmt:               q.getConfigValueStmt,
		getConversationsStmt:             q.getConversationsStmt,
		getCredentialStmt:                q.getCredentialStmt,
		getLatestMessagesStmt:            q.getLatestMessagesStmt,
		getMessagesStmt:                  q.getMessagesStmt,
		getPreviousMessageForRoleStmt:    q.getPreviousMessageForRoleStmt,
		getPromptTokensStmt:              q.getPromptTokensStmt,
		getTotalTokensStmt:               q.getTotalTokensStmt,
		insertMessageStmt:                q.insertMessageStmt,
		insertUsageStmt:                  q.insertUsageStmt,
		nextConversationStmt:             q.nextConversationStmt,
		previousConversationStmt:         q.previousConversationStmt,
		setConfigValueStmt:               q.setConfigValueStmt,
		setSelectedConversationStmt:      q.setSelectedConversationStmt,
		unsetSelectedConversationStmt:    q.unsetSelectedConversationStmt,
		updateClientConfigStmt:           q.updateClientConfigStmt,
		updateCredentialStmt:             q.updateCredentialStmt,
	}
}

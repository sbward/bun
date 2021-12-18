package bun

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

type AddColumnQuery struct {
	baseQuery

	ifNotExists bool
}

var _ Query = (*AddColumnQuery)(nil)

func NewAddColumnQuery(db IDB) *AddColumnQuery {
	q := &AddColumnQuery{
		baseQuery: baseQuery{
			db: db,
		},
	}
	return q
}

func (q *AddColumnQuery) Model(model interface{}) *AddColumnQuery {
	q.setTableModel(model)
	return q
}

//------------------------------------------------------------------------------

func (q *AddColumnQuery) Table(tables ...string) *AddColumnQuery {
	for _, table := range tables {
		q.addTable(schema.UnsafeIdent(table))
	}
	return q
}

func (q *AddColumnQuery) TableExpr(query string, args ...interface{}) *AddColumnQuery {
	q.addTable(schema.SafeQuery(query, args))
	return q
}

func (q *AddColumnQuery) ModelTableExpr(query string, args ...interface{}) *AddColumnQuery {
	q.modelTableName = schema.SafeQuery(query, args)
	return q
}

//------------------------------------------------------------------------------

func (q *AddColumnQuery) ColumnExpr(query string, args ...interface{}) *AddColumnQuery {
	q.addColumn(schema.SafeQuery(query, args))
	return q
}

func (q *AddColumnQuery) IfNotExists() *AddColumnQuery {
	q.ifNotExists = true
	return q
}

//------------------------------------------------------------------------------

func (q *AddColumnQuery) Operation() string {
	return "ADD COLUMN"
}

func (q *AddColumnQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	if q.err != nil {
		return nil, q.err
	}
	if len(q.columns) != 1 {
		return nil, fmt.Errorf("bun: AddColumnQuery requires exactly one column")
	}

	b = append(b, "ALTER TABLE "...)

	b, err = q.appendFirstTable(fmter, b)
	if err != nil {
		return nil, err
	}

	b = append(b, " ADD "...)

	if q.ifNotExists {
		b = append(b, "IF NOT EXISTS "...)
	}

	b, err = q.columns[0].AppendQuery(fmter, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//------------------------------------------------------------------------------

func (q *AddColumnQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	queryBytes, err := q.AppendQuery(q.db.Formatter(), q.db.makeQueryBytes())
	if err != nil {
		return nil, err
	}

	query := internal.String(queryBytes)
	return q.exec(ctx, q, query)
}

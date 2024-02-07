//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var FijoyTransaction = newFijoyTransactionTable("public", "fijoy_transaction", "")

type fijoyTransactionTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnString
	TransactionType postgres.ColumnString
	Amount          postgres.ColumnFloat
	FromAccountID   postgres.ColumnString
	ToAccountID     postgres.ColumnString
	UserID          postgres.ColumnString
	WorkspaceID     postgres.ColumnString
	Datetime        postgres.ColumnTimestampz
	Note            postgres.ColumnString
	CategoryID      postgres.ColumnString
	PayeeName       postgres.ColumnString
	PayerName       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FijoyTransactionTable struct {
	fijoyTransactionTable

	EXCLUDED fijoyTransactionTable
}

// AS creates new FijoyTransactionTable with assigned alias
func (a FijoyTransactionTable) AS(alias string) *FijoyTransactionTable {
	return newFijoyTransactionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FijoyTransactionTable with assigned schema name
func (a FijoyTransactionTable) FromSchema(schemaName string) *FijoyTransactionTable {
	return newFijoyTransactionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FijoyTransactionTable with assigned table prefix
func (a FijoyTransactionTable) WithPrefix(prefix string) *FijoyTransactionTable {
	return newFijoyTransactionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FijoyTransactionTable with assigned table suffix
func (a FijoyTransactionTable) WithSuffix(suffix string) *FijoyTransactionTable {
	return newFijoyTransactionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFijoyTransactionTable(schemaName, tableName, alias string) *FijoyTransactionTable {
	return &FijoyTransactionTable{
		fijoyTransactionTable: newFijoyTransactionTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newFijoyTransactionTableImpl("", "excluded", ""),
	}
}

func newFijoyTransactionTableImpl(schemaName, tableName, alias string) fijoyTransactionTable {
	var (
		IDColumn              = postgres.StringColumn("id")
		TransactionTypeColumn = postgres.StringColumn("transaction_type")
		AmountColumn          = postgres.FloatColumn("amount")
		FromAccountIDColumn   = postgres.StringColumn("from_account_id")
		ToAccountIDColumn     = postgres.StringColumn("to_account_id")
		UserIDColumn          = postgres.StringColumn("user_id")
		WorkspaceIDColumn     = postgres.StringColumn("workspace_id")
		DatetimeColumn        = postgres.TimestampzColumn("datetime")
		NoteColumn            = postgres.StringColumn("note")
		CategoryIDColumn      = postgres.StringColumn("category_id")
		PayeeNameColumn       = postgres.StringColumn("payee_name")
		PayerNameColumn       = postgres.StringColumn("payer_name")
		allColumns            = postgres.ColumnList{IDColumn, TransactionTypeColumn, AmountColumn, FromAccountIDColumn, ToAccountIDColumn, UserIDColumn, WorkspaceIDColumn, DatetimeColumn, NoteColumn, CategoryIDColumn, PayeeNameColumn, PayerNameColumn}
		mutableColumns        = postgres.ColumnList{TransactionTypeColumn, AmountColumn, FromAccountIDColumn, ToAccountIDColumn, UserIDColumn, WorkspaceIDColumn, DatetimeColumn, NoteColumn, CategoryIDColumn, PayeeNameColumn, PayerNameColumn}
	)

	return fijoyTransactionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		TransactionType: TransactionTypeColumn,
		Amount:          AmountColumn,
		FromAccountID:   FromAccountIDColumn,
		ToAccountID:     ToAccountIDColumn,
		UserID:          UserIDColumn,
		WorkspaceID:     WorkspaceIDColumn,
		Datetime:        DatetimeColumn,
		Note:            NoteColumn,
		CategoryID:      CategoryIDColumn,
		PayeeName:       PayeeNameColumn,
		PayerName:       PayerNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

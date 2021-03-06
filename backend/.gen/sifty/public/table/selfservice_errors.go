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

var SelfserviceErrors = newSelfserviceErrorsTable("public", "selfservice_errors", "")

type selfserviceErrorsTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnString
	Errors    postgres.ColumnString
	SeenAt    postgres.ColumnTimestamp
	WasSeen   postgres.ColumnBool
	CreatedAt postgres.ColumnTimestamp
	UpdatedAt postgres.ColumnTimestamp
	CsrfToken postgres.ColumnString
	Nid       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SelfserviceErrorsTable struct {
	selfserviceErrorsTable

	EXCLUDED selfserviceErrorsTable
}

// AS creates new SelfserviceErrorsTable with assigned alias
func (a SelfserviceErrorsTable) AS(alias string) *SelfserviceErrorsTable {
	return newSelfserviceErrorsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SelfserviceErrorsTable with assigned schema name
func (a SelfserviceErrorsTable) FromSchema(schemaName string) *SelfserviceErrorsTable {
	return newSelfserviceErrorsTable(schemaName, a.TableName(), a.Alias())
}

func newSelfserviceErrorsTable(schemaName, tableName, alias string) *SelfserviceErrorsTable {
	return &SelfserviceErrorsTable{
		selfserviceErrorsTable: newSelfserviceErrorsTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newSelfserviceErrorsTableImpl("", "excluded", ""),
	}
}

func newSelfserviceErrorsTableImpl(schemaName, tableName, alias string) selfserviceErrorsTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		ErrorsColumn    = postgres.StringColumn("errors")
		SeenAtColumn    = postgres.TimestampColumn("seen_at")
		WasSeenColumn   = postgres.BoolColumn("was_seen")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		UpdatedAtColumn = postgres.TimestampColumn("updated_at")
		CsrfTokenColumn = postgres.StringColumn("csrf_token")
		NidColumn       = postgres.StringColumn("nid")
		allColumns      = postgres.ColumnList{IDColumn, ErrorsColumn, SeenAtColumn, WasSeenColumn, CreatedAtColumn, UpdatedAtColumn, CsrfTokenColumn, NidColumn}
		mutableColumns  = postgres.ColumnList{ErrorsColumn, SeenAtColumn, WasSeenColumn, CreatedAtColumn, UpdatedAtColumn, CsrfTokenColumn, NidColumn}
	)

	return selfserviceErrorsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Errors:    ErrorsColumn,
		SeenAt:    SeenAtColumn,
		WasSeen:   WasSeenColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		CsrfToken: CsrfTokenColumn,
		Nid:       NidColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

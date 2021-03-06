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

var Networks = newNetworksTable("public", "networks", "")

type networksTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnString
	CreatedAt postgres.ColumnTimestamp
	UpdatedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type NetworksTable struct {
	networksTable

	EXCLUDED networksTable
}

// AS creates new NetworksTable with assigned alias
func (a NetworksTable) AS(alias string) *NetworksTable {
	return newNetworksTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new NetworksTable with assigned schema name
func (a NetworksTable) FromSchema(schemaName string) *NetworksTable {
	return newNetworksTable(schemaName, a.TableName(), a.Alias())
}

func newNetworksTable(schemaName, tableName, alias string) *NetworksTable {
	return &NetworksTable{
		networksTable: newNetworksTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newNetworksTableImpl("", "excluded", ""),
	}
}

func newNetworksTableImpl(schemaName, tableName, alias string) networksTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		UpdatedAtColumn = postgres.TimestampColumn("updated_at")
		allColumns      = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns  = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn}
	)

	return networksTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

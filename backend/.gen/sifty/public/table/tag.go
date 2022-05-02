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

var Tag = newTagTable("public", "tag", "")

type tagTable struct {
	postgres.Table

	//Columns
	ID   postgres.ColumnString
	Name postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TagTable struct {
	tagTable

	EXCLUDED tagTable
}

// AS creates new TagTable with assigned alias
func (a TagTable) AS(alias string) *TagTable {
	return newTagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TagTable with assigned schema name
func (a TagTable) FromSchema(schemaName string) *TagTable {
	return newTagTable(schemaName, a.TableName(), a.Alias())
}

func newTagTable(schemaName, tableName, alias string) *TagTable {
	return &TagTable{
		tagTable: newTagTableImpl(schemaName, tableName, alias),
		EXCLUDED: newTagTableImpl("", "excluded", ""),
	}
}

func newTagTableImpl(schemaName, tableName, alias string) tagTable {
	var (
		IDColumn       = postgres.StringColumn("id")
		NameColumn     = postgres.StringColumn("name")
		allColumns     = postgres.ColumnList{IDColumn, NameColumn}
		mutableColumns = postgres.ColumnList{NameColumn}
	)

	return tagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:   IDColumn,
		Name: NameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

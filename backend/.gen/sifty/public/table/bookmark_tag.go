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

var BookmarkTag = newBookmarkTagTable("public", "bookmark_tag", "")

type bookmarkTagTable struct {
	postgres.Table

	//Columns
	BookmarkID postgres.ColumnString
	TagID      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BookmarkTagTable struct {
	bookmarkTagTable

	EXCLUDED bookmarkTagTable
}

// AS creates new BookmarkTagTable with assigned alias
func (a BookmarkTagTable) AS(alias string) *BookmarkTagTable {
	return newBookmarkTagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BookmarkTagTable with assigned schema name
func (a BookmarkTagTable) FromSchema(schemaName string) *BookmarkTagTable {
	return newBookmarkTagTable(schemaName, a.TableName(), a.Alias())
}

func newBookmarkTagTable(schemaName, tableName, alias string) *BookmarkTagTable {
	return &BookmarkTagTable{
		bookmarkTagTable: newBookmarkTagTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newBookmarkTagTableImpl("", "excluded", ""),
	}
}

func newBookmarkTagTableImpl(schemaName, tableName, alias string) bookmarkTagTable {
	var (
		BookmarkIDColumn = postgres.StringColumn("bookmark_id")
		TagIDColumn      = postgres.StringColumn("tag_id")
		allColumns       = postgres.ColumnList{BookmarkIDColumn, TagIDColumn}
		mutableColumns   = postgres.ColumnList{}
	)

	return bookmarkTagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		BookmarkID: BookmarkIDColumn,
		TagID:      TagIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

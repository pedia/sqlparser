package sqlparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	is := assert.New(t)
	doc, err := Parse("sqlite", "CREATE TABLE `company` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text)")
	is.Nil(err)
	is.Len(doc, 1)
	is.Equal("company", doc[0].CreateTable.Name[0].Identifier.Value)
	is.Equal("id", doc[0].CreateTable.Columns[0].Name.Value)
	is.True(doc[0].CreateTable.Columns[0].PrimaryKey())
	is.True(doc[0].CreateTable.Columns[0].AutoIncrement())
	is.Equal("name", doc[0].CreateTable.Columns[1].Name.Value)
	is.Equal("Integer", doc[0].CreateTable.Columns[0].DataType.Type)
	is.Equal("Text", doc[0].CreateTable.Columns[1].DataType.Type)
}

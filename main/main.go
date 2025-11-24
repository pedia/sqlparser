package main

import (
	"fmt"
	"log"

	"github.com/pedia/sqlparser"
)

func main() {
	doc, err := sqlparser.Parse("sqlite", "CREATE TABLE `company` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text)")

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	if len(doc) > 0 {
		stmt := doc[0].CreateTable

		fmt.Printf("%v\n", stmt.Name[0].Identifier.Value)
		for _, col := range stmt.Columns {
			fmt.Printf("  %s, DataType: %s pk: %v ai: %v\n", col.Name.Value, col.DataType.Type,
				col.PrimaryKey(), col.AutoIncrement())
		}
	}
}

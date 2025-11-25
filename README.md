`sqlparser` wraps rust bindings for sqlparser-rs into a go package.

Since [sqlparser-rs(datafusion-sqlparser-rs)](https://github.com/apache/datafusion-sqlparser-rs) is the most powerful tool contains a lexer and 
parser for SQL that conforms with the ANSI/ISO SQL standard and other dialects,
binding to go is a better choice.

Parsing a SQL query is relatively straight forward:
```go
doc, err := sqlparser.Parse("sqlite", 
    `CREATE TABLE company (id integer PRIMARY KEY AUTOINCREMENT,
    name text)`)

// table name
fmt.Printf("%v\n", stmt.Name[0].Identifier.Value)
// column name, type
for _, col := range stmt.Columns {
    fmt.Printf("  %s, DataType: %s pk: %v ai: %v\n", col.Name.Value, col.DataType.Type,
        col.PrimaryKey(), col.AutoIncrement())
}
```

# TODO:
As a package, `cargo build --release` or `go generate` won't be executed automatically.

Add `replace` in `go.mod` is ugly.
```
replace github.com/pedia/sqlparser => ../sqlparser
```

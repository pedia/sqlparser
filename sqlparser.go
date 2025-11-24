package sqlparser

//go:generate cargo build --release

/*
#cgo LDFLAGS: -Ltarget/release -lsqlparser_go

char* parse(const char* dialect, const char* sql);
void free_rust_string(char* ptr);

#include <stdlib.h> // for free
*/
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"
)

// dialect should be: generic, mysql, postgresql, hive, sqlite, snowflake,
// redshift, mssql, clickhouse, bigquery, ansi, duckdb, databricks
func Parse(dialect, sql string) ([]Statement, error) {
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))
	cdialect := C.CString(dialect)
	defer C.free(unsafe.Pointer(cdialect))

	cr := C.parse(cdialect, csql)
	if cr == nil {
		return nil, errors.New("parse failed")
	}
	defer C.free_rust_string(cr)

	js := C.GoString(cr)

	var doc Document
	err := json.Unmarshal([]byte(js), &doc)
	return doc, err
}

type Document []Statement

type Statement struct {
	CreateTable CreateTable `json:"CreateTable"`
}

type CreateTable struct {
	Name []struct {
		Identifier struct {
			Value      string `json:"value"`
			QuoteStyle string `json:"quote_style"`
			Span       Span   `json:"span"`
		}
	} `json:"name"`
	Columns      []Column `json:"columns"`
	TableOptions string   `json:"table_options"`
	Strict       bool     `json:"strict"`
}

type Span struct {
	Start struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"start"`
	End struct {
		Line   int `json:"line"`
		Column int `json:"end"` // Corrected: original JSON uses "end" key
	} `json:"end"`
}

type ColumnOption struct {
	Name   any       `json:"name"`
	Option OptionDef `json:"option"`
}

type OptionDef map[string]any

func (o OptionDef) PrimaryKey() bool {
	if o["PrimaryKey"] != nil {
		return true
	}
	if u, ok := o["Unique"].(map[string]any); ok {
		return u["is_primary"].(bool)
	}
	return false
}
func (o OptionDef) AutoIncrement() bool {
	if d, ok := o["DialectSpecific"].([]any); ok && len(d) > 0 {
		if w, ok := d[0].(map[string]any)["Word"].(map[string]any); ok {
			return w["keyword"] == "AUTOINCREMENT"
		}
	}
	return false
}

// Custom DataType representation:
type CustomDataType struct {
	Type string // e.g., "Text" or "Integer"
}

// Implement custom UnmarshalJSON for CustomDataType
func (cdt *CustomDataType) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a simple string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		cdt.Type = s
		return nil
	}

	// Try to unmarshal as an object (e.g., {"Integer": null})
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err == nil {
		// Extract the key name (e.g., "Integer")
		for key := range obj {
			cdt.Type = key
			return nil
		}
	}

	return fmt.Errorf("failed to unmarshal DataType: neither string nor object")
}

type Column struct {
	Name struct {
		Value      string `json:"value"`
		QuoteStyle string `json:"quote_style"`
		Span       Span   `json:"span"`
	} `json:"name"`
	DataType CustomDataType `json:"data_type"` // Use the custom type here
	Options  []ColumnOption `json:"options"`
}

func (c Column) PrimaryKey() bool {
	for _, o := range c.Options {
		if o.Option.PrimaryKey() {
			return true
		}
	}
	return false
}

func (c Column) AutoIncrement() bool {
	for _, o := range c.Options {
		if o.Option.AutoIncrement() {
			return true
		}
	}
	return false
}

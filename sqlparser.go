package sqlparser

/*
#cgo LDFLAGS: -Ltarget/release -lsqlparser_go

char* parse(const char* sql);
void free_rust_string(char* ptr);

#include <stdlib.h> // for free
*/
import "C"
import (
	"encoding/json"
	"errors"
	"unsafe"
)

func Parse(sql string) ([]Statement, error) {
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))

	cr := C.parse(csql)
	if cr == nil {
		return nil, errors.New("failed parse")
	}
	defer C.free_rust_string(cr)

	js := C.GoString(cr)

	var statements []Statement
	err := json.Unmarshal([]byte(js), &statements)
	return statements, err
}

type Foo struct{}

type Statement struct {
	OrReplace                  bool               `json:"or_replace"`
	Temporary                  bool               `json:"temporary"`
	External                   bool               `json:"external"`
	Dynamic                    bool               `json:"dynamic"`
	Global                     *Foo               `json:"global"`
	IfNotExist                 bool               `json:"if_not_exists"`
	Transien                   bool               `json:"transient"`
	Volatil                    bool               `json:"volatile"`
	Iceber                     bool               `json:"iceberg"`
	Name                       []map[string]*Name `json:"name"`
	Columns                    []*Name            `json:"columns"`
	Constraints                []Foo              `json:"constraints"`
	TableOptions               string             `json:"table_options"`                   // : "None",
	FileFormat                 *Foo               `json:"file_format"`                     // : null,
	Location                   *Foo               `json:"location"`                        // : null,
	Query                      *Foo               `json:"query"`                           // : null,
	WithoutRowid               bool               `json:"without_rowid"`                   // : false,
	Like                       *Foo               `json:"like"`                            // : null,
	Clone                      *Foo               `json:"clone"`                           // : null,
	Version                    *Foo               `json:"version"`                         // : null,
	Comment                    *Foo               `json:"comment"`                         // : null,
	OnCommit                   *Foo               `json:"on_commit"`                       // : null,
	OnCluster                  *Foo               `json:"on_cluster"`                      // : null,
	PrimaryKey                 *Foo               `json:"primary_key"`                     // : null,
	OrderBy                    *Foo               `json:"order_by"`                        // : null,
	PartitionBy                *Foo               `json:"partition_by"`                    // : null,
	ClusterBy                  *Foo               `json:"cluster_by"`                      // : null,
	ClusteredBy                *Foo               `json:"clustered_by"`                    // : null,
	Inherits                   *Foo               `json:"inherits"`                        // : null,
	Strict                     bool               `json:"strict"`                          // : false,
	CopyGrants                 bool               `json:"copy_grants"`                     // : false,
	EnableSchemaEvolution      *Foo               `json:"enable_schema_evolution"`         // : null,
	ChangeTracking             *Foo               `json:"change_tracking"`                 // : null,
	DataRetentionTimeInDays    *Foo               `json:"data_retention_time_in_days"`     // : null,
	MaxDataExtensionTimeInDays *Foo               `json:"max_data_extension_time_in_days"` // : null,
	DefaultDdlCollation        *Foo               `json:"default_ddl_collation"`           // : null,
	WithAggregationPolicy      *Foo               `json:"with_aggregation_policy"`         // : null,
	WithRowAccessPolicy        *Foo               `json:"with_row_access_policy"`          // : null,
	WithTags                   *Foo               `json:"with_tags"`                       // : null,
	ExternalVolume             *Foo               `json:"external_volume"`                 // : null,
	BaseLocation               *Foo               `json:"base_location"`                   // : null,
	Catalog                    *Foo               `json:"catalog"`                         // : null,
	CatalogSync                *Foo               `json:"catalog_sync"`                    // : null,
	StorageSerializationPolicy *Foo               `json:"storage_serialization_policy"`    // : null,
	TargetLag                  *Foo               `json:"target_lag"`                      // : null,
	Warehouse                  *Foo               `json:"warehouse"`                       // : null,
	RefreshMode                *Foo               `json:"refresh_mode"`                    // : null,
	Initialize                 *Foo               `json:"initialize"`                      // : null,
	RequireUser                bool               `json:"require_user"`                    // : false
}

type Column struct {
	Name     *Name
	DataType any           `json:"data_type"`
	Options  []NamedOption `json:"options"`
}

type Name struct {
	Value      string `json:"value"`
	QuoteStyle string `json:"quote_style"`
	Span       Span   `json:"span"`
}

type NamedOption struct {
	Name   *string          `json:"name"`
	Option []map[string]any `json:"option"`
}

type DialectSpecific struct {
	Word []wrapWord
}

type wrapWord struct {
	Word Word
}

type Word struct {
	Value      string  `json:"value"`
	QuoteStyle *string `json:"quote_style"`
	Keyword    string  `json:"keyword"`
}

type Pos struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Span map[string]Pos

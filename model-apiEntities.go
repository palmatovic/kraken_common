package kraken_common

import "github.com/palmatovic/api_common"

type ApiSchema struct {
	SchemaID     int              `json:"schema_id"`
	SchemaTables []ApiSchemaTable `json:"schema_tables,omitempty"`
	SchemaName   string           `json:"schema_name"`
}

type ApiSchemaTable struct {
	SchemaTableID      int              `json:"schema_table_id"`
	SchemaTableName    string           `json:"schema_table_name"`
	SchemaTableColumns []ApiTableColumn `json:"schema_table_columns"`
}

type ApiTableColumn struct {
	SchemaTableColumnID          int                        `json:"schema_table_column_id,omitempty"`
	SchemaTableColumnName        string                     `json:"schema_table_column_name"`
	SchemaTableColumnDataType    string                     `json:"schema_table_column_data_type"`
	SchemaTableColumnConstraints []ApiTableColumnConstraint `json:"schema_table_column_constraints,omitempty"`
}

type ApiTableColumnConstraint struct {
	SchemaTableColumnConstraintID         int    `json:"schema_table_column_constraint_id"`
	SchemaTableColumnConstraintName       string `json:"schema_table_column_constraint_name"`
	SchemaTableColumnReferencedColumnName string `json:"schema_table_column_referenced_column_name"`
	SchemaTableColumnReferencedTableName  string `json:"schema_table_column_referenced_table_name"`
}

type ApiConfig struct {
	ConfigId   *int        `json:"config_id,omitempty"`
	ConfigType string      `json:"config_type"`
	Host       string      `json:"host"`
	Port       string      `json:"port"`
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Schemas    []ApiSchema `json:"schemas,omitempty"`
}

type ImportConfig struct {
	Data ImportConfigData `json:"data"`
}

type MicroImportConfig struct {
	Status *bool                 `json:"status,omitempty"`
	Data   MicroImportConfigData `json:"data"`
}

type ImportConfigData struct {
	Config ApiConfig `json:"config"`
}

type MicroImportConfigData struct {
	Error     *api_common.Error `json:"error,omitempty"`
	Config    ApiConfig         `json:"config"`
	Operation string            `json:"operation"`
	UserID    string            `json:"user"`
}

type GetConfigData struct {
	Config ApiConfig `json:"config"`
}
type GetConfigsData struct {
	Config *[]ApiConfig `json:"configs"`
	Total  *int64       `json:"total"`
}

type DeleteConfigData struct {
	Message string `json:"message"`
}

type ConfigRequestedData struct {
	Config  ApiConfig `json:"config"`
	Message string    `json:"message"`
}

package models

type InsertDocumentResponse struct {
	Ok  bool   `json:"ok"`
	Id  string `json:"id"`
	Rev string `json:"rev"`
}
type InsertBulkDocumentResponse []InsertDocumentResponse

type FindResponse struct {
	Docs           []any          `json:"docs"`
	Bookmark       string         `json:"bookmark"`
	Warning        string         `json:"warning"`
	ExecutionStats ExecutionStats `json:"execution_stats"`
}
type InsertResponse struct {
	Ok  bool   `json:"ok"`
	Id  string `json:"id"`
	Rev string `json:"rev"`
}
type CreateDBResponse struct {
	Ok bool `json:"ok"`
}
type ExecutionStats struct {
	TotalKeysExamined       int64   `json:"total_keys_examined"`
	TotalDocsExamined       int64   `json:"total_docs_examined"`
	TotalQuorumDocsExamined int64   `json:"total_quorum_docs_examined"`
	ResultsReturned         int64   `json:"results_returned"`
	ExecutionTimeMS         float64 `json:"execution_time_ms"`
}
type UserDBModel struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Type     string   `json:"type"`
	Roles    []string `json:"roles"`
}

type SecurityModel struct {
	Admins  Admins  `json:"admins"`
	Members Members `json:"members"`
}

type Admins struct {
	Names []string      `json:"names"`
	Roles []interface{} `json:"roles"`
}
type Members struct {
	Names []string      `json:"names"`
	Roles []interface{} `json:"roles"`
}

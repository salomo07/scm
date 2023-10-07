package models

type InsertDocumentResponse struct {
	Ok  bool   `json:"ok"`
	Id  string `json:"id"`
	Rev string `json:"rev"`
}

type FindResponse struct {
	Docs           []any          `json:"docs"`
	Bookmark       string         `json:"bookmark"`
	Warning        string         `json:"warning"`
	ExecutionStats ExecutionStats `json:"execution_stats"`
}

type ExecutionStats struct {
	TotalKeysExamined       int64   `json:"total_keys_examined"`
	TotalDocsExamined       int64   `json:"total_docs_examined"`
	TotalQuorumDocsExamined int64   `json:"total_quorum_docs_examined"`
	ResultsReturned         int64   `json:"results_returned"`
	ExecutionTimeMS         float64 `json:"execution_time_ms"`
}

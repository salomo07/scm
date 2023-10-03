package models

type InsertDocumentResponse struct {
	Ok  bool   `json:"ok"`
	Id  string `json:"id"`
	Rev string `json:"rev"`
}

package dtos

type GenerateExcelDto struct {
	Columns   []string   `json:"columns"`
	Rows      [][]string `json:"rows"`
	SheetName string     `json:"sheetName"`
}

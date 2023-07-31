package services

import (
	"fmt"

	"github.com/Felipedelima123/excelor/dtos"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func GenerateExcel(payload dtos.GenerateExcelDto) {
	f := excelize.NewFile()

	fmt.Println(payload.SheetName != "Sheet1")

	if payload.SheetName != "Sheet1" {
		f.NewSheet(payload.SheetName)
		f.DeleteSheet("Sheet1")
	}

	streamWriter, err := f.NewStreamWriter(payload.SheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	header := []interface{}{}
	for _, cell := range payload.Columns {
		header = append(header, cell)
	}
	streamWriter.SetRow("A1", header)

	currentLine := 2

	for _, row := range payload.Rows {
		cells := []interface{}{}
		for _, rowCells := range row {
			cells = append(cells, rowCells)
		}
		streamWriter.SetRow("A"+fmt.Sprint(currentLine), cells)
		currentLine += 1
	}

	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SaveAs("tmp_files/" + uuid.New().String() + ".xlsx"); err != nil {
		fmt.Println(err)
		return
	}

}

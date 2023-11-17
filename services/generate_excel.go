package services

import (
	"fmt"
	"os"

	"github.com/Felipedelima123/excelor/dtos"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func GenerateExcel(payload dtos.GenerateExcelDto) (dtos.ExcelUrlDTO, error) {
	f := excelize.NewFile()

	if payload.SheetName != "Sheet1" {
		f.NewSheet(payload.SheetName)
		f.DeleteSheet("Sheet1")
	}

	streamWriter, err := f.NewStreamWriter(payload.SheetName)
	if err != nil {
		fmt.Println(err)
		return dtos.ExcelUrlDTO{}, err
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
		return dtos.ExcelUrlDTO{}, err
	}

	filename := uuid.New().String() + ".xlsx"

	if err := f.SaveAs("tmp_files/" + filename); err != nil {
		fmt.Println(err)
		return dtos.ExcelUrlDTO{}, err
	}

	bucketName := os.Getenv("BUCKET_NAME")

	UploadToBucket(bucketName, filename)

	fmt.Println("Excel generated successfully!")

	url := GetSignedUrl(bucketName, filename)

	fmt.Println(url)

	return dtos.ExcelUrlDTO{
		Url:      url,
		Filename: filename,
	}, nil

}

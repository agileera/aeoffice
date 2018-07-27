package aetable

import (
	"strings"

	pf "github.com/oltolm/go-pandocfilters"
	"github.com/tealeg/xlsx"
)

type AETable struct {
	Caption     []interface{}
	ColHeader   []interface{}
	ColAlign    []interface{}
	RelColWidth []interface{}
	Rows        [][]interface{}
}

func ParseExcel(config *AETableConfig) (*AETable, error) {
	table := &AETable{}
	table.Caption = []interface{}{pf.Str(config.Title)}

	xlsFile, err := xlsx.OpenFile(config.XLSXFile)
	if err != nil {
		return nil, err
	}

	for _, sheet := range xlsFile.Sheets {
		// fmt.Println(sheet.Name)
		sheetName := strings.TrimSpace(sheet.Name)
		if sheetName == config.Sheet {
			colHeader, colAlign, relColWidth, rows := parseSheet(sheet)
			table.ColHeader = colHeader
			table.ColAlign = colAlign
			table.RelColWidth = relColWidth
			table.Rows = rows
		}
	}

	return table, nil
}

func parseSheet(sheet *xlsx.Sheet) ([]interface{}, []interface{}, []interface{}, [][]interface{}) {
	colHeader := []interface{}{}
	colAlign := []interface{}{}
	relColWidth := []interface{}{}
	rows := [][]interface{}{}
	var jsonrow []interface{}
	for index, row := range sheet.Rows {

		if index != 0 {
			jsonrow = []interface{}{}
			// rows = append(rows, jsonrow)
		}

		for _, cell := range row.Cells {
			text := cell.String()
			if index == 0 {
				plain := []interface{}{}
				plain = append(plain, pf.Str(text))

				colHeader = append(colHeader, []interface{}{pf.Plain(plain)})

				align := map[string]interface{}{"t": "AlignLeft"}
				colAlign = append(colAlign, align)

				relColWidth = append(relColWidth, 0)
				// fmt.Printf("%s\n", text)
			} else {
				// fmt.Printf("%s\n", text)
				plain := []interface{}{}
				plain = append(plain, pf.Str(text))
				jsonrow = append(jsonrow, []interface{}{pf.Plain(plain)})
			}

		}

		if index != 0 {
			rows = append(rows, jsonrow)
		}
	}

	return colHeader, colAlign, relColWidth, rows

}

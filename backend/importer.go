package backend

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
)

type Importer struct {
	ctx context.Context
}

func NewImporter() *Importer {
	return &Importer{}
}

func (i *Importer) Startup(ctx context.Context) {
	i.ctx = ctx
}
func (i *Importer) SelectAndImportExcel() (ExcelData, error) {
	// Tampilkan file picker
	filepath, err := runtime.OpenFileDialog(i.ctx, runtime.OpenDialogOptions{
		Title: "Pilih file Excel",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Excel Files",
				Pattern:     "*.xlsx;*.xls",
			},
		},
	})
	if err != nil || filepath == "" {
		return ExcelData{}, err
	}

	// Proses file Excel (gunakan fungsi yang sama)
	return i.ImportExcel(filepath)
}

type ExcelData struct {
	Filename string     `json:"filename"`
	Header   []string   `json:"header"`
	Details  [][]string `json:"details"`
}

func (i *Importer) ImportExcel(filePath string) (ExcelData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return ExcelData{}, err
	}
	defer f.Close()
	sheetlist := f.GetSheetList()
	rows, err := f.GetRows(sheetlist[0])
	if err != nil {
		return ExcelData{}, err
	}

	var data = ExcelData{
		Header:  []string{},
		Details: [][]string{},
	}
	data.Filename = filePath
	for i, row := range rows {
		if len(row) > 0 {
			if i == 0 {
				// Baris pertama dianggap sebagai header
				data.Header = append(data.Header, "No")
				data.Header = append(data.Header, row[1])
				data.Header = append(data.Header, row[2])
				data.Header = append(data.Header, row[3])
				data.Header = append(data.Header, row[4])
			} else {
				// Baris berikutnya dianggap sebagai detail
				var dtl []string
				dtl = append(dtl, row[0])
				dtl = append(dtl, row[1])
				dtl = append(dtl, row[2])
				dtl = append(dtl, row[3])
				dtl = append(dtl, row[4])
				data.Details = append(data.Details, dtl)
			}
		}
	}

	return data, nil
}

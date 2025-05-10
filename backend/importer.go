package backend

import (
	"context"
	"wails-excel-import/backend/models"

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
	IsSaveDB bool       `json:"isSaveDB"`
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
	runtime.LogDebugf(i.ctx, "Sheet List: %s", rows[0][0])
	var data = ExcelData{
		IsSaveDB: false,
		Header:   []string{},
		Details:  [][]string{},
	}

	if rows[0][0] == "Menu" {
		data, err = MappingListMenuPP(rows, filePath)
		if err != nil {
			return ExcelData{}, err
		}
	} else {
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
	}
	return data, nil
}

func MappingListMenuPP(data [][]string, filename string) (ExcelData, error) {
	var tmp = ExcelData{
		Filename: filename,
		Header:   []string{},
		Details:  [][]string{},
	}
	for i, row := range data {
		if len(row) > 0 {
			if i == 0 {
				// Baris pertama dianggap sebagai header
				tmp.Header = append(tmp.Header, row[0])
			} else {
				// Baris berikutnya dianggap sebagai detail
				var dtl []string
				dtl = append(dtl, row[0])
				tmp.Details = append(tmp.Details, dtl)
				InsertDataToDB(models.Product{
					Name:  row[0],
					Price: 0,
				})
				tmp.IsSaveDB = true
			}
		}
	}
	return tmp, nil
}

func InsertDataToDB(data models.Product) error {
	// Implementasi untuk menyimpan data ke database
	// Misalnya, menggunakan GORM untuk menyimpan data ke tabel yang sesuai
	models.DB.Create(&models.Product{
		Name:  data.Name,
		Price: data.Price,
	})
	return nil
}
func (i *Importer) DeleteAllDataFromDB() (models.Response, error) {
	// Implementasi untuk menghapus data dari database
	// Misalnya, menggunakan GORM untuk menghapus data dari tabel yang sesuai
	res := models.Response{
		Status:  "error",
		Message: "Deleted Failed",
		Data:    nil}
	tx := models.DB.Exec("DELETE FROM products")
	if tx.Error == nil {
		res.Message = "Deleted Success"
		res.Status = "success"
	}
	return res, nil
}

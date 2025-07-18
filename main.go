package main

import (
	"embed"
	"wails-excel-import/backend"
	"wails-excel-import/backend/models"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	importer := backend.NewImporter()
	models.ConnectDatabase()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-excel-import",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			DisableWindowIcon: false,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        importer.Startup,
		Bind: []interface{}{
			app, importer,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

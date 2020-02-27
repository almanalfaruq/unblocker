package main

import (
	"runtime"

	commandrepo "github.com/almanalfaruq/unblocker/repo/command"
	filerepo "github.com/almanalfaruq/unblocker/repo/file"
	iprepo "github.com/almanalfaruq/unblocker/repo/ip"
	usecase "github.com/almanalfaruq/unblocker/usecase/ip"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var ipUsecase *usecase.IP = nil

func writeToHosts(url string) string {
	if ipUsecase == nil {
		return "Usecase is nil"
	}

	system := runtime.GOOS
	err := ipUsecase.WriteToHosts(url, system)
	if err != nil {
		return err.Error()
	}

	return ""
}

func main() {
	ipRepo := iprepo.New()
	fileRepo := filerepo.New()
	commandRepo := commandrepo.New()
	ipUsecase = usecase.New(ipRepo, fileRepo, commandRepo)
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  400,
		Height: 200,
		Title:  "unblocker",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(writeToHosts)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

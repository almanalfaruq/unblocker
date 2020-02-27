package main

import (
	"fmt"
	"runtime"

	commandrepo "github.com/almanalfaruq/unblocker/repo/command"
	filerepo "github.com/almanalfaruq/unblocker/repo/file"
	iprepo "github.com/almanalfaruq/unblocker/repo/ip"
	usecase "github.com/almanalfaruq/unblocker/usecase/ip"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var ipUsecase *usecase.IP = nil

func writeToHosts(url string) error {
	fmt.Println(url)
	if ipUsecase == nil {
		return fmt.Errorf("Usecase is nil")
	}

	system := runtime.GOOS
	return ipUsecase.WriteToHosts(url, system)
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

package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	TotalVisited = 0
	Duplicated   = 0
	TodayVisited = 0
	Blacklisted  = 0
	MostUsed     = "Chrome"
	TodaySession = time.Now().Format("2006-01-02_15-04-05")
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	http.HandleFunc("/receive", func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL.Query().Get("url")

		TodayVisited++
		SaveUrl(url, "TodayVisited")
		SaveUrl(url, "TotalVisited")

		a.AddUrls(url)
	})
	http.ListenAndServe(":8080", nil)
}

func (a *App) AddUrls(url string) {
	runtime.WindowExecJS(a.ctx, fmt.Sprintf("UpdateUrls('%s')", url))
}

func SaveUrl(url, fileName string) error {
	fullFileName := TodaySession + "_" + fileName + ".txt"

	err := ioutil.WriteFile(fullFileName, []byte(url), 0644)
	if err != nil {
		return err
	}

	return nil
}

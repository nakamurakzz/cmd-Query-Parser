package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/atotto/clipboard"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "qp",
		Short: "Visualize query parameters from a URL",
		Run: func(cmd *cobra.Command, args []string) {
			exec()
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exec() {
	fmt.Print("URLを入力してください: ")
	var rawURL string
	fmt.Scanln(&rawURL)

	app := tview.NewApplication()

	infoText := tview.NewTextView()
	infoText.SetText("選択したパラメータがクリップボードにコピーされます")
	infoText.SetBorderPadding(1, 1, 2, 2)

	list := tview.NewList()
	list.SetTitle("Select Query Parameter")
	list.SetBorder(true)

	queries := getQueryParams(rawURL)
	for key, value := range queries {
		list.AddItem(key+": "+value, "", 0, func() {
			clipboard.WriteAll(value)
			app.Stop()
		})
	}

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(infoText, 3, 1, false).
		AddItem(list, 0, 1, true)

	app.SetRoot(flex, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func getQueryParams(rawURL string) map[string]string {
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	qps := map[string]string{}
	for key, values := range u.Query() {
		for _, value := range values {
			qps[key] = value
		}
	}
	return qps
}

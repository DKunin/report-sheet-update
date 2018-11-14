package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
)

type Settings struct {
	 Url string
}

func main() {
	data, err := ioutil.ReadFile("client_secret.json")

	settingsJson, _ := ioutil.ReadFile("./settings.json")
	settings := Settings{}
	json.Unmarshal(settingsJson, &settings)

	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(settings.Url)

	checkError(err)
	sheet, err := spreadsheet.SheetByIndex(0)
	checkError(err)
	var lastRow uint
	for _, row := range sheet.Rows {
		for _, cell := range row {
			lastRow = cell.Row + 1
			fmt.Println(cell.Value)
		}
	}

	// Update cell content
	//sheet.Update(int(lastRow), 0, "hogehoge")
	sheet.Update(int(lastRow), 0, "hogehoge")
	sheet.Update(int(lastRow), 1, "hogehoge 2")
	//spreadsheet.UnmarshalJSON("{some: true}")
	// Make sure call Synchronize to reflect the changes
	err = sheet.Synchronize()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
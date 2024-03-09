package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"net/http"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const appName string = "Skriveleif"
const definitionUrl string = "https://ordbokene.no"
const suggestionsUrl string = "https://ord.uib.no/api/suggest"

func main() {
	a := app.New()
	drv := a.Driver()
	w := drv.(desktop.Driver).CreateSplashWindow()

	i := widget.NewEntry()
	i.SetPlaceHolder("Søk...")

	suggestions := binding.NewStringList()

	l := widget.NewListWithData(suggestions,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	l.Hide()

	l.OnSelected = func(id widget.ListItemID) {
		v, err := suggestions.GetValue(id)
		if err != nil {
			return
		}
		u, _ := url.ParseRequestURI(fmt.Sprintf("%s/bm,nn/%s", definitionUrl, v))
		a.OpenURL(u)
		a.Quit()
	}

	f := canvas.NewText(fmt.Sprintf("%s v0.3.0 Credit: Språkrådet / UiB ordbokene.no", appName), color.White)
	f.TextSize = 12

	i.OnChanged = func(s string) {
		if len(s) < 3 {
			l.Hide()
			w.Resize(fyne.NewSize(0, 0))
			return
		}

		resp, err := http.Get(fmt.Sprintf("%s?include=e&dict=bm,nn&q=%s", suggestionsUrl, s))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error fetching suggestions: %v", err)
			return
		}

		result := new(Suggestions)
		json.NewDecoder(resp.Body).Decode(&result)

		suggs := make([]string, result.Count)
		for i, e := range result.Answer.Exact {
			suggs[i] = e[0].(string)
		}

		suggestions.Set(suggs)
		l.Show()
		w.Resize(fyne.NewSize(0, 200))
	}

	// Move focus to list on enter.
	i.OnSubmitted = func(s string) {
		w.Canvas().Focus(l)
	}

	w.SetContent(container.NewBorder(i, f, nil, nil, l))
	w.Canvas().Focus(i)
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}

type Suggestions struct {
	Count  int `json:"cnt"`
	Answer struct {
		Exact [][]any `json:"exact"`
	} `json:"a"`
}

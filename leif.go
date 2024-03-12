// Leif
// Copyright (C) 2024  Vegard Løkken

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

const appName string = "Leif"
const version string = "v0.1.2"
const credit string = "Credit: Språkrådet / UiB ordbokene.no"

func main() {
	app := app.New()
	drv, ok := app.Driver().(desktop.Driver)
	if !ok {
		log.Fatalf("%s only supports desktop", appName)
	}
	win := drv.CreateSplashWindow()

	list := newSuggestionsList(win)
	// Keep list hidden initially until there are suggestions to show.
	list.Hide()

	// Open browser and quit once a suggestion is selected.
	list.onSelected = func(s suggestion) {
		app.OpenURL(s.getDefinitionUrl())
		app.Quit()
	}

	input := newInput(win)
	input.SetPlaceHolder("Søk...")

	input.OnChanged = func(s string) {
		var result []suggestion

		if len(s) > 1 {
			result = search(s)
			list.Set(result)
		}

		if len(result) > 0 {
			list.Show()
			win.Resize(fyne.NewSize(0, 200))
		} else {
			list.Hide()
			win.Resize(fyne.NewSize(0, 0))
		}
	}

	input.onMoveFocus = func() {
		win.Canvas().Focus(list)
	}

	footer := canvas.NewText(fmt.Sprintf("%s %s %s", appName, version, credit), color.White)
	footer.TextSize = 12

	// This should be a short lived app, quit on lost focus.
	app.Lifecycle().SetOnExitedForeground(func() {
		app.Quit()
	})

	win.SetContent(container.NewBorder(input, footer, nil, nil, list))
	win.Canvas().Focus(input)
	win.SetFixedSize(true)
	win.CenterOnScreen()
	win.ShowAndRun()
}

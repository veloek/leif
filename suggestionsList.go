// Skriveleif
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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type suggestionsList struct {
	widget.List
	window     fyne.Window
	data       binding.UntypedList
	onSelected func(s suggestion)
}

func (s *suggestionsList) TypedKey(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		s.window.Close()
	}
	if key.Name == fyne.KeyReturn {
		// Trigger a selection.
		s.List.TypedKey(&fyne.KeyEvent{Name: fyne.KeySpace})
	}
	s.List.TypedKey(key)
}

func (s *suggestionsList) Set(data []suggestion) {
	s.data.Set(make([]interface{}, 0))

	for _, r := range data {
		s.data.Append(r)
	}
}

func newSuggestionsList(window fyne.Window) *suggestionsList {
	data := binding.NewUntypedList()

	list := &suggestionsList{
		List: widget.List{
			Length: data.Length,
			CreateItem: func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			UpdateItem: func(i widget.ListItemID, o fyne.CanvasObject) {
				item, err := data.GetItem(i)
				if err != nil {
					fyne.LogError(fmt.Sprintf("Error getting data item %d", i), err)
					return
				}
				u, _ := item.(binding.Untyped).Get()
				o.(*widget.Label).SetText(u.(suggestion).String())
			},
		},
		window: window}
	list.ExtendBaseWidget(list)
	data.AddListener(binding.NewDataListener(list.Refresh))
	list.data = data

	list.OnSelected = func(id widget.ListItemID) {
		v, err := data.GetValue(id)
		if err != nil {
			return
		}
		list.onSelected(v.(suggestion))
	}

	return list
}

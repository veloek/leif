// This file is part of Skriveleif.
//
// Skriveleif is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// Skriveleif is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// Skriveleif. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type input struct {
	widget.Entry
	window      fyne.Window
	onMoveFocus func()
}

func (i *input) TypedKey(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		i.window.Close()
	}
	if key.Name == fyne.KeyDown && i.onMoveFocus != nil {
		i.onMoveFocus()
	} else {
		i.Entry.TypedKey(key)
	}
}

func newInput(window fyne.Window) *input {
	e := &input{
		Entry: widget.Entry{
			Wrapping:  fyne.TextTruncate,
			MultiLine: false,
		},
		window: window,
	}
	e.OnSubmitted = func(_ string) {
		if e.onMoveFocus != nil {
			e.onMoveFocus()
		}
	}
	e.ExtendBaseWidget(e)
	return e
}

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
	"log"
	"net/url"
)

type suggestion struct {
	word string
	dict string
}

func (s suggestion) String() string {
	return fmt.Sprintf("%s (%s)", s.word, s.dict)
}

const definitionUrl string = "https://ordbokene.no"

func (s suggestion) getDefinitionUrl() *url.URL {
	u, err := url.ParseRequestURI(fmt.Sprintf("%s/%s/%s", definitionUrl, s.dict, s.word))
	if err != nil {
		log.Fatalf("error building definition URL: %s\n", err)
	}
	return u
}

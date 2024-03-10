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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type apiResponse struct {
	Count  int `json:"cnt"`
	Answer struct {
		Exact [][]any `json:"exact"`
	} `json:"a"`
}

const suggestionsUrl string = "https://ord.uib.no/api/suggest"

func search(s string) []suggestion {
	resp, err := http.Get(fmt.Sprintf("%s?include=e&dict=bm,nn&q=%s", suggestionsUrl, s))
	if err != nil {
		log.Fatalf("error fetching suggestions: %s\n", err)
	}

	result := new(apiResponse)
	json.NewDecoder(resp.Body).Decode(&result)

	suggestions := make([]suggestion, result.Count)
	for i, e := range result.Answer.Exact {
		suggestions[i] = suggestion{word: e[0].(string), dict: parseDictionary(e[1])}
	}

	return suggestions
}

func parseDictionary(d any) string {
	idict := d.([]interface{})
	dict := make([]string, len(idict))
	for i, d := range idict {
		dict[i] = d.(string)
	}
	return strings.Join(dict, ",")
}

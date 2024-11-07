// Copyright 2018 The go-chainstone Authors
// This file is part of the go-chainstone library.
//
// The go-chainstone library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-chainstone library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-chainstone library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://chainstone.org")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "chainstone.org" {
		t.Errorf("expected: %v, got: %v", "chainstone.org", url.Path)
	}

	for _, u := range []string{"chainstone.org", ""} {
		if _, err = parseURL(u); err == nil {
			t.Errorf("input %v, expected err, got: nil", u)
		}
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "chainstone.org"}
	if url.String() != "https://chainstone.org" {
		t.Errorf("expected: %v, got: %v", "https://chainstone.org", url.String())
	}

	url = URL{Scheme: "", Path: "chainstone.org"}
	if url.String() != "chainstone.org" {
		t.Errorf("expected: %v, got: %v", "chainstone.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "chainstone.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://chainstone.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://chainstone.org\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://chainstone.org\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "chainstone.org" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "chainstone.org"}, URL{"https", "chainstone.org"}, 0},
		{URL{"http", "chainstone.org"}, URL{"https", "chainstone.org"}, -1},
		{URL{"https", "chainstone.org/a"}, URL{"https", "chainstone.org"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "chainstone.org"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}

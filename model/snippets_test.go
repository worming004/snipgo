/*
 * Copyright (c) 2020, nwillc@gmail.com
 *
 * Permission to use, copy, modify, and/or distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

const testSnippetsFile = "../test/files/snippets.json"

type SnippetsTestSuite struct {
	suite.Suite
	snippets     Snippets
	badFilename  string
	goodFilename string
}

func (suite *SnippetsTestSuite) TestStringer() {
	snippet := Snippet{
		Category: "Foo",
		Title:    "Bar",
		Body:     "Baz",
	}
	assert.Equal(suite.T(), "Foo: Bar", snippet.String())
}

func (suite *SnippetsTestSuite) SetupTest() {
	suite.snippets = []Snippet{
		{
			Category: "A",
			Title:    "A",
		},
		{
			Category: "A",
			Title:    "B",
		},
	}
	suite.badFilename = "foo"
	suite.goodFilename = testSnippetsFile
}

func (suite *SnippetsTestSuite) TestNonExist() {
	_, ok := ReadSnippets(suite.badFilename)
	assert.NotNil(suite.T(), ok)
}

func (suite *SnippetsTestSuite) TestExist() {
	_, ok := ReadSnippets(suite.goodFilename)
	assert.Nil(suite.T(), ok)
}

func (suite *SnippetsTestSuite) TestWriteFile() {
	tempFile, err := ioutil.TempFile("", "snippets.*.json")
	assert.Nil(suite.T(), err)
	defer os.Remove(tempFile.Name())

	original, ok := ReadSnippets(suite.goodFilename)
	assert.Nil(suite.T(), ok)
	err = original.WriteSnippets(tempFile.Name())
	assert.Nil(suite.T(), err)
	read, ok := ReadSnippets(tempFile.Name())
	assert.Nil(suite.T(), ok)
	assert.Equal(suite.T(), len(original), len(read))
}

func (suite *SnippetsTestSuite) TestMalformedFile() {
	tempFile, err := ioutil.TempFile("", "snippets.*.json")
	assert.Nil(suite.T(), err)
	defer os.Remove(tempFile.Name())
	tempFile.WriteString("not json")

	_, ok := ReadSnippets(tempFile.Name())
	assert.NotNil(suite.T(), ok)
}

func (suite *SnippetsTestSuite) TestLen() {
	assert.Equal(suite.T(), 2, suite.snippets.Len())
}

func (suite *SnippetsTestSuite) TestLess() {
	assert.LessOrEqual(suite.T(), suite.snippets[0].Category, suite.snippets[1].Category)
	assert.LessOrEqual(suite.T(), suite.snippets[0].Title, suite.snippets[1].Title)
	assert.True(suite.T(), suite.snippets.Less(0, 1))
}

func (suite *SnippetsTestSuite) TestSwap() {
	suite.snippets.Swap(0, 1)
	assert.False(suite.T(), suite.snippets.Less(0, 1))
}

func TestSnippetsTestSuite(t *testing.T) {

	suite.Run(t, new(SnippetsTestSuite))
}

func TestSnippetsByCategory(t *testing.T) {
	tests := []struct {
		name       string
		snippets   Snippets
		categories Categories
	}{
		{
			name:       "Empty",
			snippets:   Snippets{},
			categories: Categories{},
		},
		{
			name: "Two Snippets One Category",
			snippets: Snippets{
				{
					Category: "A",
					Title:    "A",
				},
				{
					Category: "A",
					Title:    "B",
				},
			},
			categories: Categories{
				{
					Name:     "A",
					Snippets: []Snippet{{Category: "A", Title: "A"}, {Category: "A", Title: "B"}},
				},
			},
		},
		{
			name: "Two Snippets Two Categories",
			snippets: Snippets{
				{
					Category: "A",
					Title:    "A",
				},
				{
					Category: "B",
					Title:    "B",
				},
			},
			categories: Categories{
				{
					Name:     "A",
					Snippets: []Snippet{{Category: "A", Title: "A"}},
				},
				{
					Name:     "B",
					Snippets: []Snippet{{Category: "B", Title: "B"}},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.snippets.ByCategory()
			assert.Equal(t, len(test.categories), len(result))
			for i, category := range test.categories {
				assert.Equal(t, category.Name, result[i].Name)
				assert.Equal(t, len(category.Snippets), len(result[i].Snippets))
				if len(category.Snippets) == len(result[i].Snippets) {
					for j, snippet := range category.Snippets {
						assert.Equal(t, snippet.Title, result[i].Snippets[j].Title)
					}
				}
			}
		})
	}
}

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
	"sort"
	"strings"
)

// Category is a named set of Snippets
type Category struct {
	Name     string   // category name
	Snippets Snippets // snippets in the category
}

// Categories is a collection of Category
type Categories []Category

// SetCategories is an interface indicating the ability to set Categories
type SetCategories interface {
	SetCategories(categories *Categories)
}

// Categories implements sort.Interface
var _ sort.Interface = (*Categories)(nil)

// ToSnippets will collect all the Snippets in the Categories
func (c Categories) ToSnippets() *Snippets {
	var snippets Snippets
	for _, category := range c {
		snippets = append(snippets, category.Snippets...)
	}
	return &snippets
}

// Len returns the length of the Categories
func (c Categories) Len() int {
	return len(c)
}

// Less compares two Category in the Categories and returns bool true if i is less than j
func (c Categories) Less(i, j int) bool {
	return strings.ToLower(c[i].Name) < strings.ToLower(c[j].Name)
}

// Swap two Category in Categories
func (c Categories) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

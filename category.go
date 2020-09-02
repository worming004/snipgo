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

package main

import (
	"sort"
	"strings"
)

type Category struct {
	Name     string
	Snippets Snippets
}

type Categories []Category

var _ sort.Interface = (*Categories)(nil)

func (c Categories) Len() int {
	return len(c)
}

func (c Categories) Less(i, j int) bool {
	return strings.ToLower(c[i].Name) < strings.ToLower(c[j].Name)
}

func (c Categories) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func SnippetsByCategory(snippets Snippets) Categories {
	catMap := make(map[string]Category)

	for _, snippet := range snippets {
		category, ok := catMap[snippet.Category]
		if !ok {
			category = Category{Name: snippet.Category}
		}
		category.Snippets = append(category.Snippets, snippet)
		catMap[category.Name] = category
	}

	var c []Category
	for _, v := range catMap {
		c = append(c, v)
	}
	sort.Sort(Categories(c))
	return c
}

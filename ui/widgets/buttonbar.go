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

package widgets

import (
	"github.com/rivo/tview"
)

type ButtonBar struct {
	*tview.Grid
	actions []func()
}

func NewButtonBar() *ButtonBar {
	grid := tview.NewGrid().
		SetRows(1).
		SetBorders(false)
	bb := ButtonBar{grid, nil}
	return &bb
}

func (bb *ButtonBar) ItemCount() int {
	return len(bb.actions)
}

func (bb *ButtonBar) AddButton(label string, action func()) *ButtonBar {
	bb.actions = append(bb.actions, action)
	button := tview.NewButton("[ " + label + " ]").SetSelectedFunc(action)
	button.SetLabelColor(tview.Styles.PrimaryTextColor)
	button.SetLabelColorActivated(tview.Styles.PrimaryTextColor)
	button.SetBackgroundColor(tview.Styles.InverseTextColor)
	button.SetBackgroundColorActivated(tview.Styles.InverseTextColor)
	bb.AddItem( button, 0, bb.ItemCount() - 1, 1, 1, 0, 0, false)
	return bb
}

func (bb *ButtonBar) action(i int) {
	bb.actions[i]()
}

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

package ui

import (
	"github.com/nwillc/snipgo/model"
	"github.com/nwillc/snipgo/ui/pages"
	"github.com/nwillc/snipgo/ui/widgets"
	"github.com/rivo/tview"
)

// UI is the tview.Primitive associated with the overall UI.
type UI struct {
	app *tview.Application
	tview.Primitive
	slides []pages.Slide
	pv     *tview.Pages
}

// Implements SetCategories
var _ model.SetCategories = (*UI)(nil)

// NewUI is a factory for UI.
func NewUI() *UI {
	app := tview.NewApplication()
	slides := []pages.Slide{
		pages.NewBrowserPage(),
		pages.NewSnippetPage(),
		pages.NewPreferencesPage(),
		pages.NewAboutPage(),
	}
	pageView := tview.NewPages()

	menu := widgets.NewButtonBar()

	for i, slide := range slides {
		slideIndex := i
		pageView.AddPage(slide.GetName(), slide, true, i == 0)
		menu.AddButton(slide.GetName(), func() {
			pageView.SwitchToPage(slides[slideIndex].GetName())
		})
	}

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(menu, 1, 1, false).
		AddItem(pageView, 0, 1, true)

	ui := UI{
		app,
		layout,
		slides,
		pageView,
	}

	for _, slide := range slides {
		slide.SetCategoryReceiver(func(categories *model.Categories) {
			ui.SetCategories(categories)
		})
	}

	menu.AddButton("Quit", func() {
		app.Stop()
	})
	return &ui
}

// SetCategories sets the model.Categories used on the UI.
func (ui *UI) SetCategories(categories *model.Categories) {
	for _, slide := range ui.slides {
		slide.SetCategories(categories)
	}
	ui.pv.SwitchToPage("Browser")
}

// Run the UI.
func (ui *UI) Run() {
	if err := ui.app.SetRoot(ui, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

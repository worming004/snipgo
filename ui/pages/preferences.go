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

package pages

import (
	"fmt"
	"github.com/nwillc/snipgo/model"
	"github.com/nwillc/snipgo/ui/widgets"
	"github.com/rivo/tview"
)

// PreferencesPage is the Slide to manage model.Preferences.
type PreferencesPage struct {
	tview.Primitive
}

// Implements Slide
var _ Slide = (*PreferencesPage)(nil)

// NewPreferencesPage is a factory for PreferencesPage Slide.
func NewPreferencesPage() *PreferencesPage {
	textView := tview.NewTextView()
	fmt.Fprintln(textView, "Preferences Page")
	fmt.Fprintln(textView, "TODO")
	page := PreferencesPage{widgets.Center(20, 2, textView)}
	return &page
}

// SetCategories sets the model.Categories used on the Slide.
func (p PreferencesPage) SetCategories(categories *model.Categories) {
	// NoOp
}

// GetName returns the name of this Slide.
func (p PreferencesPage) GetName() string {
	return "Preferences"
}

// SetCategoryReceiver inform the Slide where to notify with changes of the model.Categories.
func (p PreferencesPage) SetCategoryReceiver(receiver CategoryReceiver) {
	// NoOp
}

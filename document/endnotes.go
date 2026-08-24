// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// Endnotes contains the endnotes attached to a document. The zero value is
// not usable directly on a Document — use Document.Endnotes()/EnsureEndnotes().
type Endnotes struct {
	d *Document
	x *wml.Endnotes
}

// newEndnotes wraps an existing (possibly nil) *wml.Endnotes.
func newEndnotes(d *Document, x *wml.Endnotes) Endnotes {
	return Endnotes{d: d, x: x}
}

// X returns the inner wrapped XML type, or nil if the document has no
// endnotes part.
func (e Endnotes) X() *wml.Endnotes {
	return e.x
}

// NonEmpty reports whether the document has at least one real endnote
// (excluding the separator/continuationSeparator entries Word always writes).
func (e Endnotes) NonEmpty() bool {
	if e.x == nil {
		return false
	}
	for _, en := range e.x.Endnote {
		if isContentFootnote(en) {
			return true
		}
	}
	return false
}

// AddEndnote appends a new, empty endnote with the next free id and returns
// it. Callers add paragraphs to the returned Footnote via AddParagraph.
func (e Endnotes) AddEndnote() Footnote {
	ex := wml.NewCT_FtnEdn()
	ex.IdAttr = nextFtnEdnID(e.x.Endnote)
	e.x.Endnote = append(e.x.Endnote, ex)
	return Footnote{d: e.d, x: ex}
}

// seedEndnoteSeparators is seedSeparators's endnote twin — CT_Endnotes and
// CT_Footnotes share the CT_FtnEdn element type but are distinct Go structs
// with differently-named slice fields (Footnote vs Endnote).
func seedEndnoteSeparators(x *wml.CT_Endnotes) {
	sep := wml.NewCT_FtnEdn()
	sep.TypeAttr = wml.ST_FtnEdnSeparator
	sep.IdAttr = -1
	cont := wml.NewCT_FtnEdn()
	cont.TypeAttr = wml.ST_FtnEdnContinuationSeparator
	cont.IdAttr = 0
	x.Endnote = append(x.Endnote, sep, cont)
}

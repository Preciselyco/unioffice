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

// Footnotes contains the footnotes attached to a document. The zero value is
// not usable directly on a Document — use Document.Footnotes()/EnsureFootnotes().
type Footnotes struct {
	d *Document
	x *wml.Footnotes
}

// newFootnotes wraps an existing (possibly nil) *wml.Footnotes.
func newFootnotes(d *Document, x *wml.Footnotes) Footnotes {
	return Footnotes{d: d, x: x}
}

// X returns the inner wrapped XML type, or nil if the document has no
// footnotes part.
func (f Footnotes) X() *wml.Footnotes {
	return f.x
}

// NonEmpty reports whether the document has at least one real footnote
// (excluding the separator/continuationSeparator entries Word always writes).
func (f Footnotes) NonEmpty() bool {
	if f.x == nil {
		return false
	}
	for _, fn := range f.x.Footnote {
		if isContentFootnote(fn) {
			return true
		}
	}
	return false
}

// isContentFootnote reports whether fn is a real footnote rather than one of
// the separator marks Word requires in every footnotes.xml.
func isContentFootnote(fn *wml.CT_FtnEdn) bool {
	switch fn.TypeAttr {
	case wml.ST_FtnEdnSeparator, wml.ST_FtnEdnContinuationSeparator:
		return false
	default:
		return true
	}
}

// AddFootnote appends a new, empty footnote with the next free id and returns
// it. Callers add paragraphs to the returned Footnote via AddParagraph.
func (f Footnotes) AddFootnote() Footnote {
	fx := wml.NewCT_FtnEdn()
	fx.IdAttr = nextFtnEdnID(f.x.Footnote)
	f.x.Footnote = append(f.x.Footnote, fx)
	return Footnote{d: f.d, x: fx}
}

// nextFtnEdnID returns the smallest positive id not already used by existing
// — ids 0 and below are reserved for the separator marks (see
// seedSeparators), so real content always starts at 1.
func nextFtnEdnID(existing []*wml.CT_FtnEdn) int64 {
	var max int64
	for _, fn := range existing {
		if fn.IdAttr > max {
			max = fn.IdAttr
		}
	}
	return max + 1
}

// seedSeparators appends the separator and continuationSeparator marks Word
// writes into every footnotes.xml / endnotes.xml part. They use reserved ids
// -1 and 0 so they can never collide with a real footnote's id (which starts
// at 1 — see nextFtnEdnID).
func seedSeparators(x *wml.CT_Footnotes) {
	sep := wml.NewCT_FtnEdn()
	sep.TypeAttr = wml.ST_FtnEdnSeparator
	sep.IdAttr = -1
	cont := wml.NewCT_FtnEdn()
	cont.TypeAttr = wml.ST_FtnEdnContinuationSeparator
	cont.IdAttr = 0
	x.Footnote = append(x.Footnote, sep, cont)
}

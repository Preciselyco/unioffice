// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// TestCT_SdtContentBlock_MarshalXML_interleavesViaContentOrder verifies that
// when ContentOrder is populated and its length matches len(P)+len(Tbl),
// MarshalXML emits w:p and w:tbl interleaved in the recorded order rather
// than all paragraphs followed by all tables.
func TestCT_SdtContentBlock_MarshalXML_interleavesViaContentOrder(t *testing.T) {
	m := wml.NewCT_SdtContentBlock()
	p1, p2 := wml.NewCT_P(), wml.NewCT_P()
	tbl := wml.NewCT_Tbl()
	m.P = []*wml.CT_P{p1, p2}
	m.Tbl = []*wml.CT_Tbl{tbl}
	m.ContentOrder = []wml.CT_SdtContentBlockEltKind{
		wml.CT_SdtContentBlockEltP,
		wml.CT_SdtContentBlockEltTbl,
		wml.CT_SdtContentBlockEltP,
	}

	buf, err := xml.Marshal(m)
	if err != nil {
		t.Fatalf("Marshal failed: %s", err)
	}
	got := string(buf)

	firstP := strings.Index(got, "<w:p>")
	firstTbl := strings.Index(got, "<w:tbl>")
	secondP := strings.Index(got[firstTbl+1:], "<w:p>")
	if firstP == -1 || firstTbl == -1 || secondP == -1 {
		t.Fatalf("expected w:p, w:tbl, w:p in output, got: %s", got)
	}
	if !(firstP < firstTbl) {
		t.Errorf("expected first w:p before w:tbl, got: %s", got)
	}
}

// TestCT_SdtContentBlock_MarshalXML_fallsBackWithoutContentOrder verifies
// that when ContentOrder is absent (or its length doesn't match
// len(P)+len(Tbl)), MarshalXML falls back to the original behavior of
// writing all paragraphs before all tables — so code that appends directly
// to P/Tbl without maintaining ContentOrder keeps working unchanged.
func TestCT_SdtContentBlock_MarshalXML_fallsBackWithoutContentOrder(t *testing.T) {
	m := wml.NewCT_SdtContentBlock()
	m.P = []*wml.CT_P{wml.NewCT_P()}
	m.Tbl = []*wml.CT_Tbl{wml.NewCT_Tbl()}

	buf, err := xml.Marshal(m)
	if err != nil {
		t.Fatalf("Marshal failed: %s", err)
	}
	got := string(buf)

	pIdx := strings.Index(got, "<w:p>")
	tblIdx := strings.Index(got, "<w:tbl>")
	if pIdx == -1 || tblIdx == -1 {
		t.Fatalf("expected both w:p and w:tbl in output, got: %s", got)
	}
	if !(pIdx < tblIdx) {
		t.Errorf("expected fallback order (all P before all Tbl), got: %s", got)
	}
}

// TestCT_SdtContentBlock_UnmarshalXML_recordsContentOrder verifies that
// unmarshaling XML with paragraphs and a table interleaved (p, tbl, p)
// reconstructs P, Tbl, and a ContentOrder that reproduces the original
// document order on a subsequent marshal.
func TestCT_SdtContentBlock_UnmarshalXML_recordsContentOrder(t *testing.T) {
	const src = `<w:sdtContent xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">` +
		`<w:p/><w:tbl/><w:p/>` +
		`</w:sdtContent>`

	m := wml.NewCT_SdtContentBlock()
	if err := xml.Unmarshal([]byte(src), m); err != nil {
		t.Fatalf("Unmarshal failed: %s", err)
	}
	if len(m.P) != 2 {
		t.Fatalf("expected 2 paragraphs, got %d", len(m.P))
	}
	if len(m.Tbl) != 1 {
		t.Fatalf("expected 1 table, got %d", len(m.Tbl))
	}
	want := []wml.CT_SdtContentBlockEltKind{
		wml.CT_SdtContentBlockEltP,
		wml.CT_SdtContentBlockEltTbl,
		wml.CT_SdtContentBlockEltP,
	}
	if len(m.ContentOrder) != len(want) {
		t.Fatalf("expected ContentOrder of length %d, got %d", len(want), len(m.ContentOrder))
	}
	for i, k := range want {
		if m.ContentOrder[i] != k {
			t.Errorf("ContentOrder[%d]: expected %v, got %v", i, k, m.ContentOrder[i])
		}
	}

	// Round-trip: re-marshaling should reproduce document order (p, tbl, p).
	buf, err := xml.Marshal(m)
	if err != nil {
		t.Fatalf("re-Marshal failed: %s", err)
	}
	got := string(buf)
	firstP := strings.Index(got, "<w:p")
	firstTbl := strings.Index(got, "<w:tbl")
	secondP := strings.Index(got[firstTbl+1:], "<w:p")
	if firstP == -1 || firstTbl == -1 || secondP == -1 {
		t.Fatalf("expected w:p, w:tbl, w:p in re-marshaled output, got: %s", got)
	}
	if !(firstP < firstTbl) {
		t.Errorf("expected re-marshaled order to start with w:p before w:tbl, got: %s", got)
	}
}

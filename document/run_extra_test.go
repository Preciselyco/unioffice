// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"testing"

	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

func newRun() Run {
	return Run{nil, wml.NewCT_R()}
}

func TestRunText(t *testing.T) {
	r := newRun()

	// empty run
	if r.Text() != "" {
		t.Errorf("Text() on empty run = %q, want empty", r.Text())
	}

	r.AddText("hello")
	if r.Text() != "hello" {
		t.Errorf("Text() = %q, want hello", r.Text())
	}

	r.AddTab()
	if r.Text() != "hello\t" {
		t.Errorf("Text() with tab = %q, want 'hello\t'", r.Text())
	}
}

func TestRunClearContent(t *testing.T) {
	r := newRun()
	r.AddText("something")
	r.ClearContent()
	if len(r.x.EG_RunInnerContent) != 0 {
		t.Error("ClearContent should empty EG_RunInnerContent")
	}
}

func TestRunAddDeletedText(t *testing.T) {
	r := newRun()
	r.AddDeletedText("removed")
	if len(r.x.EG_RunInnerContent) != 1 {
		t.Fatalf("expected 1 inner content, got %d", len(r.x.EG_RunInnerContent))
	}
	if r.x.EG_RunInnerContent[0].DelText == nil {
		t.Error("DelText should be set")
	}
	if r.x.EG_RunInnerContent[0].DelText.Content != "removed" {
		t.Errorf("DelText.Content = %q, want removed", r.x.EG_RunInnerContent[0].DelText.Content)
	}
}

func TestRunProperties(t *testing.T) {
	r := newRun()
	if r.x.RPr != nil {
		t.Error("RPr should start nil")
	}
	rp := r.Properties()
	if r.x.RPr == nil {
		t.Error("Properties() should initialise RPr")
	}
	_ = rp
}

func TestRunAddBreak(t *testing.T) {
	r := newRun()
	r.AddBreak()
	if len(r.x.EG_RunInnerContent) != 1 {
		t.Fatalf("expected 1 content, got %d", len(r.x.EG_RunInnerContent))
	}
	ic := r.x.EG_RunInnerContent[0]
	if ic.Br == nil {
		t.Error("Br should be set after AddBreak")
	}
	if ic.Br.TypeAttr != wml.ST_BrTypeUnset {
		t.Errorf("AddBreak TypeAttr = %v, want Unset (line break)", ic.Br.TypeAttr)
	}
}

func TestRunAddPageBreak(t *testing.T) {
	r := newRun()
	r.AddPageBreak()
	if len(r.x.EG_RunInnerContent) != 1 {
		t.Fatalf("expected 1 content, got %d", len(r.x.EG_RunInnerContent))
	}
	ic := r.x.EG_RunInnerContent[0]
	if ic.Br == nil {
		t.Error("Br should be set after AddPageBreak")
	}
	if ic.Br.TypeAttr != wml.ST_BrTypePage {
		t.Errorf("AddPageBreak TypeAttr = %v, want Page", ic.Br.TypeAttr)
	}
}

func TestRunDrawingAnchoredEmpty(t *testing.T) {
	r := newRun()
	anchored := r.DrawingAnchored()
	if len(anchored) != 0 {
		t.Errorf("expected empty slice, got %d items", len(anchored))
	}
}

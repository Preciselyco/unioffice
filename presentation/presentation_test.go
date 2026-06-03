// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/Preciselyco/unioffice/presentation"
)

func TestNew(t *testing.T) {
	p := presentation.New()
	if p == nil {
		t.Fatal("New() returned nil")
	}
	if p.X() == nil {
		t.Error("X() must not return nil")
	}
}

func TestNewValidate(t *testing.T) {
	p := presentation.New()
	if err := p.Validate(); err != nil {
		t.Errorf("New() produces invalid presentation: %s", err)
	}
}

func TestNewSaveRoundtrip(t *testing.T) {
	p := presentation.New()
	buf := &bytes.Buffer{}
	if err := p.Save(buf); err != nil {
		t.Fatalf("Save: %s", err)
	}
	if buf.Len() == 0 {
		t.Error("Save produced empty output")
	}
}

func TestAddSlide(t *testing.T) {
	p := presentation.New()
	before := len(p.Slides())
	s := p.AddSlide()
	if s.X() == nil {
		t.Error("AddSlide().X() must not return nil")
	}
	after := len(p.Slides())
	if after != before+1 {
		t.Errorf("Slides count: got %d, want %d", after, before+1)
	}
}

func TestRemoveSlide(t *testing.T) {
	p := presentation.New()
	s1 := p.AddSlide()
	p.AddSlide()
	before := len(p.Slides())

	if err := p.RemoveSlide(s1); err != nil {
		t.Fatalf("RemoveSlide: %s", err)
	}
	after := len(p.Slides())
	if after != before-1 {
		t.Errorf("after RemoveSlide: got %d slides, want %d", after, before-1)
	}
}

func TestSlideMasters(t *testing.T) {
	p := presentation.New()
	masters := p.SlideMasters()
	if len(masters) == 0 {
		t.Error("New presentation should have at least one slide master")
	}
}

func TestSlideLayouts(t *testing.T) {
	p := presentation.New()
	layouts := p.SlideLayouts()
	// A presentation may have no layouts in the minimal new() case; just ensure
	// it returns without panic
	_ = layouts
}

func TestGetLayoutByNameMissing(t *testing.T) {
	p := presentation.New()
	_, err := p.GetLayoutByName("NonExistentLayout")
	if err == nil {
		t.Error("GetLayoutByName with unknown name should return error")
	}
}

func TestSaveToFile(t *testing.T) {
	p := presentation.New()
	p.AddSlide()

	path := filepath.Join(t.TempDir(), "out.pptx")
	if err := p.SaveToFile(path); err != nil {
		t.Fatalf("SaveToFile: %s", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("saved file not found: %s", err)
	}
}

func TestSlideTextBox(t *testing.T) {
	p := presentation.New()
	s := p.AddSlide()
	tb := s.AddTextBox()
	// TextBox has no X(); verify Properties() doesn't panic
	_ = tb.Properties()
}

func TestSlidePlaceHolders(t *testing.T) {
	p := presentation.New()
	s := p.AddSlide()
	// A freshly-added bare slide has no placeholders; just ensure no panic
	phs := s.PlaceHolders()
	_ = phs
}

func TestGetImageByRelIDMissing(t *testing.T) {
	p := presentation.New()
	_, ok := p.GetImageByRelID("nonexistent")
	if ok {
		t.Error("GetImageByRelID with unknown relID should return false")
	}
}

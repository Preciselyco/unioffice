// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"testing"

	"github.com/Preciselyco/unioffice/measurement"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

func TestParagraphStyle(t *testing.T) {
	doc := New()
	p := doc.AddParagraph()

	if p.Style() != "" {
		t.Errorf("Style() on fresh paragraph = %q, want empty", p.Style())
	}

	p.SetStyle("Heading1")
	if p.Style() != "Heading1" {
		t.Errorf("Style() = %q, want Heading1", p.Style())
	}

	p.SetStyle("")
	if p.Style() != "" {
		t.Errorf("SetStyle(\"\") should clear style, got %q", p.Style())
	}
}

func TestParagraphProperties(t *testing.T) {
	doc := New()
	p := doc.AddParagraph()
	pp := p.Properties()
	if pp.x == nil {
		t.Error("Properties() should return a valid ParagraphProperties")
	}
}

func TestParagraphPropertiesX(t *testing.T) {
	doc := New()
	p := doc.AddParagraph()
	pp := p.Properties()
	if pp.X() == nil {
		t.Error("ParagraphProperties.X() must not return nil")
	}
}

func TestParagraphPropertiesStyle(t *testing.T) {
	doc := New()
	p := doc.AddParagraph()
	pp := p.Properties()

	if pp.Style() != "" {
		t.Errorf("Style() = %q, want empty", pp.Style())
	}
	pp.SetStyle("Normal")
	if pp.Style() != "Normal" {
		t.Errorf("Style() = %q, want Normal", pp.Style())
	}
	pp.SetStyle("")
	if pp.Style() != "" {
		t.Errorf("SetStyle(\"\") should clear, got %q", pp.Style())
	}
}

func TestParagraphPropertiesSetAlignment(t *testing.T) {
	doc := New()
	p := doc.AddParagraph()
	pp := p.Properties()

	pp.SetAlignment(wml.ST_JcCenter)
	if pp.x.Jc == nil || pp.x.Jc.ValAttr != wml.ST_JcCenter {
		t.Error("SetAlignment(Center) failed")
	}
	pp.SetAlignment(wml.ST_JcUnset)
	if pp.x.Jc != nil {
		t.Error("SetAlignment(Unset) should clear Jc")
	}
}

func TestParagraphPropertiesKeepWithNext(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.SetKeepWithNext(true)
	if pp.x.KeepNext == nil {
		t.Error("KeepNext should be set")
	}
	pp.SetKeepWithNext(false)
	if pp.x.KeepNext != nil {
		t.Error("KeepNext should be nil after false")
	}
}

func TestParagraphPropertiesKeepOnOnePage(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.SetKeepOnOnePage(true)
	if pp.x.KeepLines == nil {
		t.Error("KeepLines should be set")
	}
	pp.SetKeepOnOnePage(false)
	if pp.x.KeepLines != nil {
		t.Error("KeepLines should be nil after false")
	}
}

func TestParagraphPropertiesSetPageBreakBefore(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.SetPageBreakBefore(true)
	if pp.x.PageBreakBefore == nil {
		t.Error("PageBreakBefore should be set")
	}
	pp.SetPageBreakBefore(false)
	if pp.x.PageBreakBefore != nil {
		t.Error("PageBreakBefore should be nil after false")
	}
}

func TestParagraphPropertiesSetWindowControl(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.SetWindowControl(true)
	if pp.x.WidowControl == nil {
		t.Error("WidowControl should be set")
	}
	pp.SetWindowControl(false)
	if pp.x.WidowControl != nil {
		t.Error("WidowControl should be nil after false")
	}
}

func TestParagraphPropertiesAddTabStop(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.AddTabStop(2*measurement.Inch, wml.ST_TabJcLeft, wml.ST_TabTlcNone)
	if pp.x.Tabs == nil || len(pp.x.Tabs.Tab) != 1 {
		t.Errorf("expected 1 tab stop, got %v", pp.x.Tabs)
	}
}

func TestParagraphPropertiesSetHeadingLevel(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.SetHeadingLevel(2)
	if pp.Style() != "Heading2" {
		t.Errorf("Style() = %q, want Heading2", pp.Style())
	}
	if pp.x.NumPr == nil || pp.x.NumPr.Ilvl == nil || pp.x.NumPr.Ilvl.ValAttr != 2 {
		t.Errorf("NumPr.Ilvl unexpected: %v", pp.x.NumPr)
	}
}

func TestParagraphPropertiesSpacing(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	sp := pp.Spacing()
	if sp.x == nil {
		t.Fatal("Spacing() should return valid ParagraphSpacing")
	}

	sp.SetBefore(12 * measurement.Point)
	if pp.x.Spacing.BeforeAttr == nil {
		t.Error("BeforeAttr should be set")
	}

	sp.SetAfter(6 * measurement.Point)
	if pp.x.Spacing.AfterAttr == nil {
		t.Error("AfterAttr should be set")
	}

	sp.SetLineSpacing(measurement.Distance(240*measurement.Twips), wml.ST_LineSpacingRuleAuto)
	if pp.x.Spacing.LineAttr == nil {
		t.Error("LineAttr should be set")
	}
	sp.SetLineSpacing(0, wml.ST_LineSpacingRuleUnset)
	if pp.x.Spacing.LineAttr != nil {
		t.Error("LineAttr should be nil after Unset")
	}

	sp.SetBeforeAuto(true)
	if pp.x.Spacing.BeforeAutospacingAttr == nil {
		t.Error("BeforeAutospacingAttr should be set")
	}
	sp.SetBeforeAuto(false)
	if pp.x.Spacing.BeforeAutospacingAttr != nil {
		t.Error("BeforeAutospacingAttr should be nil after false")
	}

	sp.SetAfterAuto(true)
	if pp.x.Spacing.AfterAutospacingAttr == nil {
		t.Error("AfterAutospacingAttr should be set")
	}
	sp.SetAfterAuto(false)
	if pp.x.Spacing.AfterAutospacingAttr != nil {
		t.Error("AfterAutospacingAttr should be nil after false")
	}
}

func TestParagraphPropertiesIndent(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	pp.SetFirstLineIndent(0.5 * measurement.Inch)
	if pp.x.Ind == nil || pp.x.Ind.FirstLineAttr == nil {
		t.Error("FirstLineAttr should be set")
	}
	pp.SetFirstLineIndent(measurement.Zero)
	if pp.x.Ind.FirstLineAttr != nil {
		t.Error("FirstLineAttr should be nil after Zero")
	}

	pp.SetStartIndent(1 * measurement.Inch)
	if pp.x.Ind.StartAttr == nil {
		t.Error("StartAttr should be set")
	}
	pp.SetStartIndent(measurement.Zero)
	if pp.x.Ind.StartAttr != nil {
		t.Error("StartAttr should be nil after Zero")
	}

	pp.SetEndIndent(1 * measurement.Inch)
	if pp.x.Ind.EndAttr == nil {
		t.Error("EndAttr should be set")
	}
	pp.SetEndIndent(measurement.Zero)
	if pp.x.Ind.EndAttr != nil {
		t.Error("EndAttr should be nil after Zero")
	}

	pp.SetHangingIndent(0.25 * measurement.Inch)
	if pp.x.Ind.HangingAttr == nil {
		t.Error("HangingAttr should be set")
	}
	pp.SetHangingIndent(measurement.Zero)
	if pp.x.Ind.HangingAttr != nil {
		t.Error("HangingAttr should be nil after Zero")
	}
}

func TestParagraphPropertiesAddSection(t *testing.T) {
	doc := New()
	pp := doc.AddParagraph().Properties()

	s := pp.AddSection(wml.ST_SectionMarkNextPage)
	if pp.x.SectPr == nil {
		t.Error("SectPr should be set")
	}
	if s.x == nil {
		t.Error("Section.x should not be nil")
	}
	if pp.x.SectPr.Type == nil || pp.x.SectPr.Type.ValAttr != wml.ST_SectionMarkNextPage {
		t.Error("SectPr.Type.ValAttr should be NextPage")
	}

	pp2 := doc.AddParagraph().Properties()
	pp2.AddSection(wml.ST_SectionMarkUnset)
	if pp2.x.SectPr.Type != nil {
		t.Error("SectPr.Type should be nil when Unset")
	}
}

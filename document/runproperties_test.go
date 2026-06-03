// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"testing"

	"github.com/Preciselyco/unioffice/color"
	"github.com/Preciselyco/unioffice/measurement"
	"github.com/Preciselyco/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

func newRP() RunProperties {
	return RunProperties{wml.NewCT_RPr()}
}

func TestRunPropertiesX(t *testing.T) {
	rp := newRP()
	if rp.X() == nil {
		t.Error("X() must not return nil")
	}
}

func TestRunPropertiesSetStyle(t *testing.T) {
	rp := newRP()
	rp.SetStyle("MyStyle")
	if rp.x.RStyle == nil || rp.x.RStyle.ValAttr != "MyStyle" {
		t.Errorf("SetStyle: got %v", rp.x.RStyle)
	}
	rp.SetStyle("")
	if rp.x.RStyle != nil {
		t.Error("SetStyle(\"\") should clear RStyle")
	}
}

func TestRunPropertiesSetFontFamily(t *testing.T) {
	rp := newRP()
	rp.SetFontFamily("Arial")
	if rp.x.RFonts == nil {
		t.Fatal("RFonts should not be nil after SetFontFamily")
	}
	if *rp.x.RFonts.AsciiAttr != "Arial" {
		t.Errorf("AsciiAttr = %q, want Arial", *rp.x.RFonts.AsciiAttr)
	}
	if *rp.x.RFonts.HAnsiAttr != "Arial" {
		t.Errorf("HAnsiAttr = %q, want Arial", *rp.x.RFonts.HAnsiAttr)
	}
	if *rp.x.RFonts.EastAsiaAttr != "Arial" {
		t.Errorf("EastAsiaAttr = %q, want Arial", *rp.x.RFonts.EastAsiaAttr)
	}
}

func TestRunPropertiesFontsLazyInit(t *testing.T) {
	rp := newRP()
	f := rp.Fonts()
	if f.x == nil {
		t.Error("Fonts() should initialise RFonts")
	}
}

func TestRunPropertiesItalic(t *testing.T) {
	rp := newRP()
	if rp.IsItalic() {
		t.Error("expected IsItalic=false initially")
	}
	if rp.ItalicValue() != OnOffValueUnset {
		t.Errorf("ItalicValue() = %v, want Unset", rp.ItalicValue())
	}

	rp.SetItalic(true)
	if !rp.IsItalic() {
		t.Error("expected IsItalic=true after SetItalic(true)")
	}
	rp.SetItalic(false)
	if rp.IsItalic() {
		t.Error("expected IsItalic=false after SetItalic(false)")
	}
	if rp.x.I != nil {
		t.Error("I should be nil after SetItalic(false)")
	}
}

func TestRunPropertiesSetBold(t *testing.T) {
	rp := newRP()
	rp.SetBold(true)
	if !rp.IsBold() {
		t.Error("expected IsBold=true after SetBold(true)")
	}
	if rp.x.BCs == nil {
		t.Error("BCs should be set alongside B")
	}
	rp.SetBold(false)
	if rp.IsBold() {
		t.Error("expected IsBold=false after SetBold(false)")
	}
	if rp.x.B != nil || rp.x.BCs != nil {
		t.Error("B and BCs should be nil after SetBold(false)")
	}
}

func TestRunPropertiesSetAllCaps(t *testing.T) {
	rp := newRP()
	rp.SetAllCaps(true)
	if rp.x.Caps == nil {
		t.Error("Caps should be set")
	}
	rp.SetAllCaps(false)
	if rp.x.Caps != nil {
		t.Error("Caps should be nil after SetAllCaps(false)")
	}
}

func TestRunPropertiesSetSmallCaps(t *testing.T) {
	rp := newRP()
	rp.SetSmallCaps(true)
	if rp.x.SmallCaps == nil {
		t.Error("SmallCaps should be set")
	}
	rp.SetSmallCaps(false)
	if rp.x.SmallCaps != nil {
		t.Error("SmallCaps should be nil after false")
	}
}

func TestRunPropertiesSetStrikeThrough(t *testing.T) {
	rp := newRP()
	rp.SetStrikeThrough(true)
	if rp.x.Strike == nil {
		t.Error("Strike should be set")
	}
	rp.SetStrikeThrough(false)
	if rp.x.Strike != nil {
		t.Error("Strike should be nil after false")
	}
}

func TestRunPropertiesSetDoubleStrikeThrough(t *testing.T) {
	rp := newRP()
	rp.SetDoubleStrikeThrough(true)
	if rp.x.Dstrike == nil {
		t.Error("Dstrike should be set")
	}
	rp.SetDoubleStrikeThrough(false)
	if rp.x.Dstrike != nil {
		t.Error("Dstrike should be nil after false")
	}
}

func TestRunPropertiesSetOutline(t *testing.T) {
	rp := newRP()
	rp.SetOutline(true)
	if rp.x.Outline == nil {
		t.Error("Outline should be set")
	}
	rp.SetOutline(false)
	if rp.x.Outline != nil {
		t.Error("Outline should be nil after false")
	}
}

func TestRunPropertiesSetShadow(t *testing.T) {
	rp := newRP()
	rp.SetShadow(true)
	if rp.x.Shadow == nil {
		t.Error("Shadow should be set")
	}
	rp.SetShadow(false)
	if rp.x.Shadow != nil {
		t.Error("Shadow should be nil after false")
	}
}

func TestRunPropertiesSetEmboss(t *testing.T) {
	rp := newRP()
	rp.SetEmboss(true)
	if rp.x.Emboss == nil {
		t.Error("Emboss should be set")
	}
	rp.SetEmboss(false)
	if rp.x.Emboss != nil {
		t.Error("Emboss should be nil after false")
	}
}

func TestRunPropertiesSetImprint(t *testing.T) {
	rp := newRP()
	rp.SetImprint(true)
	if rp.x.Imprint == nil {
		t.Error("Imprint should be set")
	}
	rp.SetImprint(false)
	if rp.x.Imprint != nil {
		t.Error("Imprint should be nil after false")
	}
}

func TestRunPropertiesColorRoundtrip(t *testing.T) {
	rp := newRP()
	rp.SetColor(color.RGB(255, 0, 0))
	if rp.x.Color == nil {
		t.Fatal("Color should not be nil after SetColor")
	}
	if *rp.x.Color.ValAttr.ST_HexColorRGB != "ff0000" {
		t.Errorf("color hex = %q, want ff0000", *rp.x.Color.ValAttr.ST_HexColorRGB)
	}
	rp.ClearColor()
	if rp.x.Color != nil {
		t.Error("Color should be nil after ClearColor")
	}
}

func TestRunPropertiesSetHighlight(t *testing.T) {
	rp := newRP()
	rp.SetHighlight(wml.ST_HighlightColorYellow)
	if rp.x.Highlight == nil {
		t.Fatal("Highlight should not be nil")
	}
	if rp.x.Highlight.ValAttr != wml.ST_HighlightColorYellow {
		t.Errorf("Highlight.ValAttr = %v, want Yellow", rp.x.Highlight.ValAttr)
	}
}

func TestRunPropertiesSetEffect(t *testing.T) {
	rp := newRP()
	rp.SetEffect(wml.ST_TextEffectShimmer)
	if rp.x.Effect == nil {
		t.Error("Effect should not be nil")
	}
	rp.SetEffect(wml.ST_TextEffectUnset)
	if rp.x.Effect != nil {
		t.Error("Effect should be nil after Unset")
	}
}

func TestRunPropertiesSetVerticalAlignment(t *testing.T) {
	rp := newRP()
	rp.SetVerticalAlignment(sharedTypes.ST_VerticalAlignRunSuperscript)
	if rp.x.VertAlign == nil {
		t.Fatal("VertAlign should not be nil")
	}
	if rp.x.VertAlign.ValAttr != sharedTypes.ST_VerticalAlignRunSuperscript {
		t.Errorf("VertAlign = %v, want Superscript", rp.x.VertAlign.ValAttr)
	}
	rp.SetVerticalAlignment(sharedTypes.ST_VerticalAlignRunUnset)
	if rp.x.VertAlign != nil {
		t.Error("VertAlign should be nil after Unset")
	}
}

func TestRunPropertiesSetUnderline(t *testing.T) {
	rp := newRP()
	rp.SetUnderline(wml.ST_UnderlineSingle, color.RGB(0, 0, 255))
	if rp.x.U == nil {
		t.Fatal("U should not be nil after SetUnderline")
	}
	if rp.x.U.ValAttr != wml.ST_UnderlineSingle {
		t.Errorf("U.ValAttr = %v, want Single", rp.x.U.ValAttr)
	}
	if rp.x.U.ColorAttr == nil || *rp.x.U.ColorAttr.ST_HexColorRGB != "0000ff" {
		t.Errorf("underline color = %v, want 0000ff", rp.x.U.ColorAttr)
	}
	rp.SetUnderline(wml.ST_UnderlineUnset, color.Auto)
	if rp.x.U != nil {
		t.Error("U should be nil after Unset")
	}
}

func TestRunPropertiesColor(t *testing.T) {
	rp := newRP()
	c := rp.Color()
	if rp.x.Color == nil {
		t.Error("Color() should initialise Color")
	}
	_ = c
}

func TestRunPropertiesSetSize(t *testing.T) {
	rp := newRP()
	rp.SetSize(12 * measurement.Point)
	if rp.x.Sz == nil {
		t.Fatal("Sz should not be nil after SetSize")
	}
	wantHalfPoints := uint64(24)
	if *rp.x.Sz.ValAttr.ST_UnsignedDecimalNumber != wantHalfPoints {
		t.Errorf("Sz = %d, want %d", *rp.x.Sz.ValAttr.ST_UnsignedDecimalNumber, wantHalfPoints)
	}
}

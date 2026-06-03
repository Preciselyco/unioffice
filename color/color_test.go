// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package color_test

import (
	"testing"

	"github.com/Preciselyco/unioffice/color"
)

func TestRGB(t *testing.T) {
	c := color.RGB(255, 128, 0)
	if c.IsAuto() {
		t.Errorf("RGB() should not be auto")
	}
	got := *c.AsRGBString()
	if got != "ff8000" {
		t.Errorf("AsRGBString() = %q, want %q", got, "ff8000")
	}
}

func TestRGBA(t *testing.T) {
	c := color.RGBA(255, 128, 0, 64)
	if c.IsAuto() {
		t.Errorf("RGBA() should not be auto")
	}
	// AsRGBAString is alpha-first: a, r, g, b
	got := *c.AsRGBAString()
	if got != "40ff8000" {
		t.Errorf("AsRGBAString() = %q, want %q", got, "40ff8000")
	}
}

func TestAsRGBStringZero(t *testing.T) {
	c := color.RGB(0, 0, 0)
	got := *c.AsRGBString()
	if got != "000000" {
		t.Errorf("AsRGBString() = %q, want %q", got, "000000")
	}
}

func TestAsRGBStringMax(t *testing.T) {
	c := color.RGB(255, 255, 255)
	got := *c.AsRGBString()
	if got != "ffffff" {
		t.Errorf("AsRGBString() = %q, want %q", got, "ffffff")
	}
}

func TestAuto(t *testing.T) {
	c := color.Auto
	if !c.IsAuto() {
		t.Errorf("Auto.IsAuto() = false, want true")
	}
}

func TestFromHex(t *testing.T) {
	tests := []struct {
		input string
		want  string
		auto  bool
	}{
		{"ff8000", "ff8000", false},
		{"#ff8000", "ff8000", false},
		{"000000", "000000", false},
		{"ffffff", "ffffff", false},
		{"", "", true},   // empty → Auto
		{"gg0000", "", true}, // invalid → Auto
	}
	for _, tc := range tests {
		c := color.FromHex(tc.input)
		if c.IsAuto() != tc.auto {
			t.Errorf("FromHex(%q).IsAuto() = %v, want %v", tc.input, c.IsAuto(), tc.auto)
		}
		if !tc.auto {
			got := *c.AsRGBString()
			if got != tc.want {
				t.Errorf("FromHex(%q).AsRGBString() = %q, want %q", tc.input, got, tc.want)
			}
		}
	}
}

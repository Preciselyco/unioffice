// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/Preciselyco/unioffice/schema/soo/dml"
)

func TestCT_GvmlGraphicFrameNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGraphicFrameNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGraphicFrameNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGraphicFrameNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlGraphicFrameNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGraphicFrameNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGraphicFrameNonVisual()
	xml.Unmarshal(buf, v2)
}

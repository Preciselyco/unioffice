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

func TestCT_TextShapeAutofitConstructor(t *testing.T) {
	v := dml.NewCT_TextShapeAutofit()
	if v == nil {
		t.Errorf("dml.NewCT_TextShapeAutofit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextShapeAutofit should validate: %s", err)
	}
}

func TestCT_TextShapeAutofitMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextShapeAutofit()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextShapeAutofit()
	xml.Unmarshal(buf, v2)
}

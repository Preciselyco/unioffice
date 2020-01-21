// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/Preciselyco/unioffice/schema/soo/sml"
)

func TestCT_GroupLevelsConstructor(t *testing.T) {
	v := sml.NewCT_GroupLevels()
	if v == nil {
		t.Errorf("sml.NewCT_GroupLevels must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GroupLevels should validate: %s", err)
	}
}

func TestCT_GroupLevelsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GroupLevels()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GroupLevels()
	xml.Unmarshal(buf, v2)
}

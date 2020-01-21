// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/Preciselyco/unioffice/schema/soo/ofc/math"
)

func TestCT_MPrConstructor(t *testing.T) {
	v := math.NewCT_MPr()
	if v == nil {
		t.Errorf("math.NewCT_MPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MPr should validate: %s", err)
	}
}

func TestCT_MPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MPr()
	xml.Unmarshal(buf, v2)
}

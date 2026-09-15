// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

// Extra (xsd:any) content comes before w:movie/w:control in the schema, so it
// must be written first.
func TestCT_PictureExtraWrittenBeforeControl(t *testing.T) {
	const in = `<w:pict xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" ` +
		`xmlns:v="urn:schemas-microsoft-com:vml" ` +
		`xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">` +
		`<v:rect fillcolor="red"/><w:control r:id="rId5"/></w:pict>`

	v := wml.NewCT_Picture()
	if err := xml.Unmarshal([]byte(in), v); err != nil {
		t.Fatalf("Unmarshal: %s", err)
	}
	if len(v.Extra) != 1 {
		t.Fatalf("want 1 Extra element, got %d", len(v.Extra))
	}
	if v.Control == nil {
		t.Fatal("w:control not decoded")
	}

	// Declare w and r on the start element the way the enclosing part root
	// would, so the output can be decoded again on its own.
	var buf bytes.Buffer
	start := xml.StartElement{
		Name: xml.Name{Local: "w:pict"},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"},
			{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"},
		},
	}
	if err := xml.NewEncoder(&buf).EncodeElement(v, start); err != nil {
		t.Fatalf("Marshal: %s", err)
	}
	out := buf.Bytes()
	s := string(out)
	rect, control := strings.Index(s, ":rect"), strings.Index(s, "w:control")
	if rect < 0 || control < 0 || rect > control {
		t.Errorf("want v:rect before w:control, got %s", s)
	}

	v2 := wml.NewCT_Picture()
	if err := xml.Unmarshal(out, v2); err != nil {
		t.Fatalf("re-Unmarshal: %s\n%s", err, s)
	}
	if len(v2.Extra) != 1 || v2.Control == nil {
		t.Errorf("content lost on second decode: %s", s)
	}
}

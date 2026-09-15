// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"image"
	"image/png"
	"io"
	"strings"
	"testing"

	"github.com/Preciselyco/unioffice/document"
)

const (
	pictWNS   = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	pictVNS   = "urn:schemas-microsoft-com:vml"
	pictONS   = "urn:schemas-microsoft-com:office:office"
	pictRelNS = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"

	pictRootNS = `xmlns:w="` + pictWNS + `" xmlns:v="` + pictVNS + `" xmlns:o="` + pictONS + `" xmlns:r="` + pictRelNS + `"`

	// pictWatermark is what Word writes for Design > Watermark > Picture.
	pictWatermark = `<w:pict>` +
		`<v:shapetype id="_x0000_t75" coordsize="21600,21600" o:spt="75" o:preferrelative="t" path="m@4@5l@4@11@9@11@9@5xe" filled="f" stroked="f">` +
		`<v:stroke joinstyle="miter"/><o:lock v:ext="edit" aspectratio="t"/>` +
		`</v:shapetype>` +
		`<v:shape id="WordPictureWatermark1" o:spid="_x0000_s1025" type="#_x0000_t75" o:allowincell="f" ` +
		`style="position:absolute;margin-left:0;margin-top:0;width:595.2pt;height:841.9pt;z-index:-251658239">` +
		`<v:imagedata r:id="rId1" o:title="letterhead"/>` +
		`</v:shape>` +
		`</w:pict>`
)

// pictVMLDocx builds a DOCX with a VML rect in the body and a picture
// watermark in the default header.
func pictVMLDocx(t *testing.T) []byte {
	t.Helper()
	var img bytes.Buffer
	if err := png.Encode(&img, image.NewRGBA(image.Rect(0, 0, 1, 1))); err != nil {
		t.Fatal(err)
	}
	parts := map[string]string{
		"[Content_Types].xml": `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Default Extension="xml" ContentType="application/xml"/>
  <Default Extension="png" ContentType="image/png"/>
  <Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
  <Override PartName="/word/header1.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml"/>
</Types>`,
		"_rels/.rels": `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/>
</Relationships>`,
		"word/_rels/document.xml.rels": `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/header" Target="header1.xml"/>
</Relationships>`,
		"word/document.xml": `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document ` + pictRootNS + `><w:body>` +
			`<w:p><w:r><w:pict><v:rect style="width:100pt;height:50pt" fillcolor="red"/></w:pict></w:r></w:p>` +
			`<w:sectPr><w:headerReference w:type="default" r:id="rId2"/></w:sectPr>` +
			`</w:body></w:document>`,
		"word/_rels/header1.xml.rels": `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="media/image1.png"/>
</Relationships>`,
		"word/header1.xml": `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:hdr ` + pictRootNS + `><w:p><w:r>` + pictWatermark + `</w:r></w:p></w:hdr>`,
		"word/media/image1.png": img.String(),
	}

	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	for name, content := range parts {
		f, err := w.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := f.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func pictRoundTrip(t *testing.T, in []byte) []byte {
	t.Helper()
	doc, err := document.Read(bytes.NewReader(in), int64(len(in)))
	if err != nil {
		t.Fatalf("Read: %s", err)
	}
	var out bytes.Buffer
	if err := doc.Save(&out); err != nil {
		t.Fatalf("Save: %s", err)
	}
	return out.Bytes()
}

// pictPart returns the content of the first part whose name matches.
func pictPart(t *testing.T, docx []byte, match func(string) bool) (string, string) {
	t.Helper()
	r, err := zip.NewReader(bytes.NewReader(docx), int64(len(docx)))
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range r.File {
		if !match(f.Name) {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			t.Fatal(err)
		}
		var b bytes.Buffer
		_, err = b.ReadFrom(rc)
		if cerr := rc.Close(); err == nil {
			err = cerr
		}
		if err != nil {
			t.Fatal(err)
		}
		return f.Name, b.String()
	}
	t.Fatalf("no matching part in package")
	return "", ""
}

// pictElements decodes part, failing on malformed XML, and returns the
// start elements in namespace space found inside a w:pict.
func pictElements(t *testing.T, part, space string) map[string][]xml.StartElement {
	t.Helper()
	found := map[string][]xml.StartElement{}
	dec := xml.NewDecoder(strings.NewReader(part))
	depth := 0
	for {
		tok, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				return found
			}
			t.Fatalf("part is not well-formed XML: %s", err)
		}
		switch el := tok.(type) {
		case xml.StartElement:
			if el.Name.Space == pictWNS && el.Name.Local == "pict" {
				depth++
			} else if depth > 0 && el.Name.Space == space {
				found[el.Name.Local] = append(found[el.Name.Local], el.Copy())
			}
		case xml.EndElement:
			if el.Name.Space == pictWNS && el.Name.Local == "pict" {
				depth--
			}
		}
	}
}

func pictAttr(el xml.StartElement, space, local string) string {
	for _, a := range el.Attr {
		if a.Name.Space == space && a.Name.Local == local {
			return a.Value
		}
	}
	return ""
}

func TestPictVMLSurvivesRoundTripInBody(t *testing.T) {
	out := pictVMLDocx(t)
	for pass := 1; pass <= 2; pass++ {
		out = pictRoundTrip(t, out)
		_, body := pictPart(t, out, func(n string) bool { return n == "word/document.xml" })
		rects := pictElements(t, body, pictVNS)["rect"]
		if len(rects) != 1 {
			t.Fatalf("pass %d: want 1 v:rect inside w:pict, got %d\n%s", pass, len(rects), body)
		}
		if got := pictAttr(rects[0], "", "fillcolor"); got != "red" {
			t.Errorf("pass %d: v:rect fillcolor = %q, want %q", pass, got, "red")
		}
	}
}

func TestPictVMLWatermarkSurvivesRoundTripInHeader(t *testing.T) {
	out := pictVMLDocx(t)
	for pass := 1; pass <= 2; pass++ {
		out = pictRoundTrip(t, out)
		name, hdr := pictPart(t, out, func(n string) bool {
			return strings.HasPrefix(n, "word/header") && strings.HasSuffix(n, ".xml")
		})

		v := pictElements(t, hdr, pictVNS)
		if len(v["shapetype"]) != 1 || len(v["stroke"]) != 1 {
			t.Errorf("pass %d: v:shapetype or its children dropped\n%s", pass, hdr)
		}
		if len(pictElements(t, hdr, pictONS)["lock"]) != 1 {
			t.Errorf("pass %d: o:lock dropped\n%s", pass, hdr)
		}
		shapes := v["shape"]
		if len(shapes) != 1 {
			t.Fatalf("pass %d: want 1 v:shape inside w:pict, got %d\n%s", pass, len(shapes), hdr)
		}
		if got := pictAttr(shapes[0], "", "id"); got != "WordPictureWatermark1" {
			t.Errorf("pass %d: v:shape id = %q", pass, got)
		}
		if got := pictAttr(shapes[0], pictONS, "spid"); got != "_x0000_s1025" {
			t.Errorf("pass %d: v:shape o:spid = %q", pass, got)
		}
		imgs := v["imagedata"]
		if len(imgs) != 1 {
			t.Fatalf("pass %d: want 1 v:imagedata, got %d\n%s", pass, len(imgs), hdr)
		}
		rid := pictAttr(imgs[0], pictRelNS, "id")
		if rid == "" {
			t.Fatalf("pass %d: v:imagedata lost its r:id\n%s", pass, hdr)
		}

		relsName := "word/_rels/" + strings.TrimPrefix(name, "word/") + ".rels"
		_, rels := pictPart(t, out, func(n string) bool { return n == relsName })
		if !strings.Contains(rels, `Id="`+rid+`"`) || !strings.Contains(rels, "media/image1.png") {
			t.Errorf("pass %d: %s does not resolve %s to media/image1.png\n%s", pass, relsName, rid, rels)
		}
	}
}

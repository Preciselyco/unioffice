// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Preciselyco/unioffice"
	"github.com/Preciselyco/unioffice/common"
	"github.com/Preciselyco/unioffice/document"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
	"github.com/Preciselyco/unioffice/testhelper"
)

func TestSimpleDoc(t *testing.T) {
	doc := document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("foo")
	// _ = doc.SaveToFile("testdata/simple-1.docx") // Uncomment to update file
	got := bytes.Buffer{}
	if err := doc.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	doc.Save(&got)
	testhelper.CompareGoldenZip(t, "simple-1.docx", got.Bytes())
}

func TestOpen(t *testing.T) {
	wb, err := document.Open("testdata/simple-1.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareZip(t, "simple-1.docx", got.Bytes(), true)
}

func TestOpenStrict(t *testing.T) {
	strict, err := document.Open("testdata/strict.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	gotStrict := bytes.Buffer{}
	if err := strict.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	strict.Save(&gotStrict)
	os.WriteFile("testdata/non-strict.docx", gotStrict.Bytes(), 0644)

	// run test assuming that the doc is a valid non-strict doc
	nonStrict, err := document.Open("testdata/non-strict.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	gotNonStrict := bytes.Buffer{}
	if err := nonStrict.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	nonStrict.Save(&gotNonStrict)
	testhelper.CompareZip(t, "non-strict.docx", gotNonStrict.Bytes(), true)

	os.Remove("testdata/non-strict.docx")
}

func TestOpenHeaderFooter(t *testing.T) {
	wb, err := document.Open("testdata/header-footer-multiple.docx")
	if err != nil {
		t.Errorf("error opening document: %s", err)
	}

	// _ = wb.SaveToFile("testdata/header-footer-multiple.docx") // Uncomment to update file
	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareGoldenZip(t, "header-footer-multiple.docx", got.Bytes())
}

func TestAddParagraph(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	doc.AddParagraph()
	if len(doc.Paragraphs()) != 1 {
		t.Errorf("expected 1 paragraphs, got %d", len(doc.Paragraphs()))
	}
	doc.AddParagraph()
	if len(doc.Paragraphs()) != 2 {
		t.Errorf("expected 2 paragraphs, got %d", len(doc.Paragraphs()))
	}
}

func TestOpenWord2016(t *testing.T) {
	doc, err := document.Open("../testdata/Office2016/Word-Windows.docx")
	if err != nil {
		t.Errorf("error opening Windows Word 2016 document: %s", err)
	}
	got := bytes.Buffer{}
	if err := doc.Save(&got); err != nil {
		t.Errorf("error saving W216 file: %s", err)
	}
	testhelper.CompareGoldenZipFilesOnly(t, "../../testdata/Office2016/Word-Windows.docx", got.Bytes())
}

func TestInsertParagraph(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	p := doc.AddParagraph()
	before := doc.InsertParagraphBefore(p)
	after := doc.InsertParagraphAfter(p)
	if len(doc.Paragraphs()) != 3 {
		t.Errorf("expected 3 paragraphs, got %d", len(doc.Paragraphs()))
	}
	if doc.Paragraphs()[0].X() != before.X() {
		t.Error("InsertParagraphBefore failed")
	}
	if doc.Paragraphs()[2].X() != after.X() {
		t.Error("InsertParagraphAfter failed")
	}
}

func TestInsertTable(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	p1 := doc.AddParagraph()
	p2 := doc.AddParagraph()
	beforeP1 := doc.InsertTableBefore(p1)
	afterP1 := doc.InsertTableAfter(p1)
	beforeP2 := doc.InsertTableBefore(p2)
	afterP2 := doc.InsertTableAfter(p2)
	if len(doc.Tables()) != 4 {
		t.Errorf("expected 4 tables, got %d", len(doc.Tables()))
	}
	if doc.Tables()[0].X() != beforeP1.X() {
		t.Error("InsertTableBefore 1st paragraph failed")
	}
	if doc.Tables()[1].X() != afterP1.X() {
		t.Error("InsertTableAfter 1st paragraph failed")
	}
	if doc.Tables()[2].X() != beforeP2.X() {
		t.Error("InsertTableBefore 2nd paragraph failed")
	}
	if doc.Tables()[3].X() != afterP2.X() {
		t.Error("InsertTableAfter 2nd paragraph failed")
	}
}

func TestInsertRun(t *testing.T) {
	doc := document.New()
	if len(doc.Paragraphs()) != 0 {
		t.Errorf("expected 0 paragraphs, got %d", len(doc.Paragraphs()))
	}
	p := doc.AddParagraph()
	middle := p.AddRun()
	before := p.InsertRunBefore(middle)
	after := p.InsertRunAfter(middle)
	middle.AddText("middle")
	before.AddText("before")
	after.AddText("after")
	if len(p.Runs()) != 3 {
		t.Errorf("expected 3 runs, got %d", len(p.Runs()))
	}
	if p.Runs()[0].X() != before.X() {
		t.Error("InsertParagraphBefore failed")
	}
	if p.Runs()[2].X() != after.X() {
		t.Error("InsertParagraphAfter failed")
	}

	p.RemoveRun(after)

	if len(p.Runs()) != 2 {
		t.Errorf("expected 2 runs, got %d", len(p.Runs()))
	}
	if p.Runs()[0].X() != before.X() {
		t.Error("InsertParagraphBefore failed")
	}
	p.RemoveRun(before)

	if len(p.Runs()) != 1 {
		t.Errorf("expected 1 runs, got %d", len(p.Runs()))
	}

	if p.Runs()[0].X() != middle.X() {
		t.Errorf("remove failed")
	}
}

func TestInsertBookmarks(t *testing.T) {
	doc := document.New()
	if len(doc.Bookmarks()) != 0 {
		t.Errorf("expected 0 bookmarks, got %d", len(doc.Bookmarks()))
	}

	p := doc.AddParagraph()
	p.AddBookmark(1, "bookmark1")
	p.AddBookmark(2, "bookmark2")

	if len(doc.Bookmarks()) != 2 {
		t.Errorf("expected 2 bookmarks, got %d", len(doc.Bookmarks()))
	}
}

func TestDuplicateBookmarks(t *testing.T) {
	doc := document.New()
	if len(doc.Bookmarks()) != 0 {
		t.Errorf("expected 0 bookmarks, got %d", len(doc.Bookmarks()))
	}

	p := doc.AddParagraph()
	p.AddBookmark(1, "bookmark1")
	p.AddBookmark(1, "bookmark1")

	if len(doc.Bookmarks()) != 2 {
		t.Errorf("expected 2 bookmarks, got %d", len(doc.Bookmarks()))
	}

	if err := doc.Validate(); err == nil {
		t.Errorf("expected error due to duplicate bookmark names")
	}
}

func TestHeaderAndFooterImages(t *testing.T) {
	doc := document.New()
	img1, err := common.ImageFromFile("testdata/gopher.png")
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}
	img2, err := common.ImageFromFile("testdata/gophercolor.png")
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}
	png3x3 := []byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
		0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x03,
		0x08, 0x02, 0x00, 0x00, 0x00, 0xd9, 0x4a, 0x22,
		0xe8, 0x00, 0x00, 0x00, 0x1e, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xd7, 0x63, 0xf8, 0xc5, 0x1e, 0xf8,
		0x9d, 0xfd, 0xd7, 0x34, 0xf6, 0x5f, 0x0c, 0x10,
		0x8a, 0x9d, 0xf7, 0x17, 0x03, 0x84, 0x62, 0xf7,
		0xf9, 0x05, 0x00, 0xd2, 0x6f, 0x0d, 0x71, 0x26,
		0x33, 0x2f, 0xe1, 0x00, 0x00, 0x00, 0x00, 0x49,
		0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	}
	img3, err := common.ImageFromBytes(png3x3)
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}

	dir1, err := doc.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	if dir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", dir1.RelID())
	}

	dir2, err := doc.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	if dir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", dir2.RelID())
	}

	dir3, err := doc.AddImage(img3)
	if err != nil {
		t.Fatalf("unable to add image to doc: %s", err)
	}
	if dir3.RelID() != "rId3" {
		t.Errorf("expected rId3 != %s", dir3.RelID())
	}

	hdr := doc.AddHeader()
	ftr := doc.AddFooter()

	hir1, err := hdr.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to header: %s", err)
	}
	if hir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", hir1.RelID())
	}

	hir2, err := hdr.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to header: %s", err)
	}
	if hir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", hir2.RelID())
	}

	hir3, err := hdr.AddImage(img3)
	if err != nil {
		t.Fatalf("unable to add image to header: %s", err)
	}
	if hir3.RelID() != "rId3" {
		t.Errorf("expected rId3 != %s", hir3.RelID())
	}

	fir1, err := ftr.AddImage(img1)
	if err != nil {
		t.Fatalf("unable to add image to footer: %s", err)
	}
	if fir1.RelID() != "rId1" {
		t.Errorf("expected rId1 != %s", fir1.RelID())
	}

	fir2, err := ftr.AddImage(img2)
	if err != nil {
		t.Fatalf("unable to add image to footer: %s", err)
	}
	if fir2.RelID() != "rId2" {
		t.Errorf("expected rId2 != %s", fir2.RelID())
	}

	fir3, err := ftr.AddImage(img3)
	if err != nil {
		t.Fatalf("unable to add image to footer: %s", err)
	}
	if fir3.RelID() != "rId3" {
		t.Errorf("expected rId3 != %s", fir3.RelID())
	}
}

func TestIssue198(t *testing.T) {
	// this tests the image fixes performed as part of issue 198
	// where we were breaking jpg images
	fn := "issue198.docx"
	doc, err := document.Open("testdata/" + fn)
	if err != nil {
		t.Errorf("error reading %s: %s", fn, err)
		return
	}
	// _ = doc.SaveToFile("testdata/" + fn + ".golden") // Uncomment to update file
	got := bytes.Buffer{}
	doc.Save(&got)
	testhelper.CompareGoldenZip(t, fn+".golden", got.Bytes())
}

func TestGetTables(t *testing.T) {
	doc := document.New()
	table := doc.AddTable()
	tables := doc.Tables()

	if len(tables) != 1 {
		t.Errorf("expected 1 table, got %d", len(tables))
		return
	}

	if table != tables[0] {
		t.Error("retrieved table != added table")
		return
	}

	tbl := document.New().AddTable().X()

	tc := table.AddRow().AddCell().X()
	elts := wml.NewEG_BlockLevelElts()
	tc.EG_BlockLevelElts = append(tc.EG_BlockLevelElts, elts)
	c := wml.NewEG_ContentBlockContent()
	elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, c)
	c.Tbl = append(c.Tbl, tbl)

	tables = doc.Tables()
	if len(tables) < 2 {
		t.Errorf("nested table not enumerated. found %d, expected 2", len(tables))
	}
}

// recodeDocx parses a docx byte slice, validates it and saves to a new byte slice.
func recodeDocx(t *testing.T, docx []byte) []byte {
	doc, err := document.ReadFromBytes(docx)
	if err != nil {
		t.Errorf("failed to read document: %s", err)
	}
	if err := doc.Validate(); err != nil {
		t.Errorf("failed to validate document: %s", err)
	}
	buf := bytes.Buffer{}
	doc.Save(&buf)
	return buf.Bytes()
}

func TestRunTrackChange(t *testing.T) {
	doc := document.New()

	para := doc.AddParagraph()
	para.AddRun().AddText("Here is ")
	rtc := para.AddRunTrackChange(document.RunTrackChangeDelete).SetAuthor("Gullman Rollger").SetDate(time.Date(2020, time.November, 9, 18, 19, 0, 0, time.UTC)).SetID(1)
	rtc.AddRun().AddDeletedText("some")
	rtc = para.AddRunTrackChange(document.RunTrackChangeInsert).SetAuthor("Gullman Rollger").SetDate(time.Date(2020, time.November, 9, 18, 19, 0, 0, time.UTC)).SetID(2)
	rtc.AddRun().AddText("a little")
	para.AddRun().AddText(" text.")

	if err := doc.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	// _ = doc.SaveToFile("testdata/runtrackchange.docx") // Uncomment to update file
	got := bytes.Buffer{}
	doc.Save(&got)
	testhelper.CompareGoldenZip(t, "runtrackchange.docx", got.Bytes())
	recoded := recodeDocx(t, got.Bytes())
	testhelper.CompareGoldenZip(t, "runtrackchange.docx", recoded)
}

func TestComments(t *testing.T) {
	doc := document.New()

	para := doc.AddParagraph()
	para.AddRun().AddText("Lorem ipsum ")
	para.AddCommentRangeStart().SetID(0)
	para.AddCommentRangeStart().SetID(1)
	para.AddCommentRangeStart().SetID(2)
	para.AddRun().AddText("dolor")
	para.AddCommentRangeEnd().SetID(0)
	para.AddRun().AddCommentReference().SetID(0)
	para.AddCommentRangeEnd().SetID(1)
	para.AddRun().AddCommentReference().SetID(1)
	para.AddCommentRangeEnd().SetID(2)
	para.AddRun().AddCommentReference().SetID(2)
	para.AddRun().AddText(" sit amet")

	com := doc.Comments.AddComment().SetID(0).SetAuthor("Gullman Rollger").SetInitials("GR").SetDate(time.Date(2020, time.November, 9, 18, 19, 0, 0, time.UTC))
	para = com.AddParagraph()
	para.AddRun().AddAnnotationReference()
	para.AddRun().AddText("First paragraph of Gullman's comment")
	para = com.AddParagraph()
	para.SetParagraphID("11111111")
	para.X().RsidPAttr = unioffice.String("ABABABAB")        // Test that this is kept separated from ParaIdAttr
	para.X().RsidRDefaultAttr = unioffice.String("CDCDCDCD") // Test that this is kept separated from ParaIdAttr
	para.AddRun().AddText("Second paragraph of Gullman's comment")

	com = doc.Comments.AddComment().SetID(1).SetAuthor("Jar-Jar Krutiainen").SetInitials("JK").SetDate(time.Date(2020, time.November, 11, 14, 57, 0, 0, time.UTC))
	para = com.AddParagraph()
	para.AddRun().AddAnnotationReference()
	para.AddRun().AddText("First paragraph of Jar-Jar's comment")
	para = com.AddParagraph()
	para.SetParagraphID("22222222")
	para.AddRun().AddText("Second paragraph of Jar-Jar's comment")

	com = doc.Comments.AddComment().SetID(2).SetAuthor("Karls von Randolf").SetInitials("KR").SetDate(time.Date(2020, time.November, 13, 22, 21, 0, 0, time.UTC))
	para = com.AddParagraph()
	para.AddRun().AddAnnotationReference()
	para.AddRun().AddText("First paragraph of Karls' comment")
	para = com.AddParagraph()
	para.SetParagraphID("33333333")
	para.AddRun().AddText("Second paragraph of Karls' comment")

	// Call SetDone() to test that the "paraIdParent" attribute is kept separate from "done"
	doc.CommentsExtended.AddCommentExtended().SetParagraphID("11111111").SetDone(false)
	doc.CommentsExtended.AddCommentExtended().SetParagraphID("22222222").SetParentParagraphID("11111111").SetDone(false)
	doc.CommentsExtended.AddCommentExtended().SetParagraphID("33333333").SetParentParagraphID("11111111").SetDone(false)

	if err := doc.Validate(); err != nil {
		t.Errorf("created an invalid document: %s", err)
	}
	// _ = doc.SaveToFile("testdata/comments.docx") // Uncomment to update file
	got := bytes.Buffer{}
	doc.Save(&got)
	testhelper.CompareGoldenZip(t, "comments.docx", got.Bytes())
	recoded := recodeDocx(t, got.Bytes())
	testhelper.CompareGoldenZip(t, "comments.docx", recoded)
}

// TestCloseOnlyRemovesTempDir verifies that Close() does not delete directories
// whose path merely starts with os.TempDir() but are not under it (e.g.
// "/tmpfoo/..." on Linux where os.TempDir() == "/tmp").
func TestCloseOnlyRemovesTempDir(t *testing.T) {
	// Create a directory that shares the temp-dir prefix but is a sibling, not
	// a child.  On Linux this is e.g. /tmp + "foo" = /tmpfoo.
	// We use a real subdirectory under the real temp dir so the test works on
	// all platforms (Windows, macOS) without needing root or special paths.
	// The important thing is to verify the HasPrefix+separator guard: we set
	// TmpPath to a path that does NOT have os.TempDir()+separator as a prefix
	// and confirm it is left alone.
	outsideDir := filepath.Join(os.TempDir()+"suffix_test", "gooxml-close-test")
	if err := os.MkdirAll(outsideDir, 0o700); err != nil {
		t.Skipf("cannot create test directory %s: %v", outsideDir, err)
	}
	defer os.RemoveAll(filepath.Dir(outsideDir))

	doc := document.New()
	doc.TmpPath = outsideDir

	if err := doc.Close(); err != nil {
		t.Fatalf("Close() returned unexpected error: %v", err)
	}

	if _, err := os.Stat(outsideDir); os.IsNotExist(err) {
		t.Errorf("Close() deleted %q which is not under os.TempDir()", outsideDir)
	}
}

// TestReadCleansUpTmpDirOnError verifies that Read() removes its temporary
// directory when parsing fails, so callers that receive (nil, err) do not leak
// the directory.
func TestReadCleansUpTmpDirOnError(t *testing.T) {
	// Snapshot gooxml-docx* entries in the system temp dir before the call.
	before := tmpDocxDirs(t)

	_, err := document.ReadFromBytes([]byte("this is not a valid zip"))
	if err == nil {
		t.Fatal("expected ReadFromBytes to fail on invalid input")
	}

	after := tmpDocxDirs(t)

	// Any gooxml-docx* directory that appeared during the call should be gone.
	for dir := range after {
		if !before[dir] {
			t.Errorf("temp directory %q was not cleaned up after Read error", dir)
		}
	}
}

// tmpDocxDirs returns the set of gooxml-docx* directory names present in
// os.TempDir().
func tmpDocxDirs(t *testing.T) map[string]bool {
	t.Helper()
	entries, err := os.ReadDir(os.TempDir())
	if err != nil {
		t.Fatalf("cannot read temp dir: %v", err)
	}
	dirs := map[string]bool{}
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "gooxml-docx") {
			dirs[e.Name()] = true
		}
	}
	return dirs
}

// TestSetFontTableRoundTrip verifies that FontTable/SetFontTable preserve the
// *wml.Fonts pointer and that EnsureFontTableRel adds the expected relationship.
func TestSetFontTableRoundTrip(t *testing.T) {
	src := document.New()
	// FontTable is nil on a freshly constructed document.
	if src.FontTable() != nil {
		t.Fatal("expected nil FontTable on new document")
	}

	fonts := wml.NewFonts()
	src.SetFontTable(fonts)
	if src.FontTable() != fonts {
		t.Error("SetFontTable/FontTable round-trip failed: pointer mismatch")
	}

	// Transfer to a second document and ensure the relationship is wired up.
	dst := document.New()
	dst.SetFontTable(src.FontTable())
	dst.EnsureFontTableRel()

	// Verify the relationship exists by saving and checking the zip contents.
	var buf bytes.Buffer
	if err := dst.Save(&buf); err != nil {
		t.Fatalf("Save failed: %v", err)
	}
	contents := buf.String()
	if !strings.Contains(contents, "fontTable") {
		t.Error("saved document does not reference fontTable")
	}
}

// TestSetFontTableDocstringWarning is a compile-time documentation test: it
// calls SetFontTable with a value from another document to confirm the API
// exists and accepts the right type, and documents via fmt.Sprintf that the
// caller is responsible for copying ExtraFiles.
func TestSetFontTableAPIExists(t *testing.T) {
	src := document.New()
	src.SetFontTable(wml.NewFonts())
	dst := document.New()
	dst.SetFontTable(src.FontTable())
	// Embedded font binaries are in src.ExtraFiles and must be copied separately.
	_ = fmt.Sprintf("ExtraFiles to copy: %d", len(src.ExtraFiles))
}

// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet_test

import (
	"testing"

	"github.com/Preciselyco/unioffice/spreadsheet"
)

func newWorkbookSheet() (*spreadsheet.Workbook, spreadsheet.Sheet) {
	wb := spreadsheet.New()
	return wb, wb.AddSheet()
}

func TestCellIsEmpty(t *testing.T) {
	_, sh := newWorkbookSheet()
	row := sh.AddRow()
	c := row.AddCell()

	if !c.IsEmpty() {
		t.Error("fresh cell should be empty")
	}

	c.SetNumber(42)
	if c.IsEmpty() {
		t.Error("cell with number should not be empty")
	}
}

func TestCellIsNumber(t *testing.T) {
	_, sh := newWorkbookSheet()
	row := sh.AddRow()

	cNum := row.AddCell()
	cNum.SetNumber(3.14)
	if !cNum.IsNumber() {
		t.Error("number cell should IsNumber()==true")
	}

	cStr := row.AddCell()
	cStr.SetString("hello")
	if cStr.IsNumber() {
		t.Error("string cell should not IsNumber()")
	}

	cBool := row.AddCell()
	cBool.SetBool(true)
	if cBool.IsNumber() {
		t.Error("bool cell should not IsNumber()")
	}
}

func TestCellIsBool(t *testing.T) {
	_, sh := newWorkbookSheet()
	row := sh.AddRow()

	c := row.AddCell()
	if c.IsBool() {
		t.Error("fresh cell should not IsBool()")
	}
	c.SetBool(true)
	if !c.IsBool() {
		t.Error("bool cell should IsBool()==true")
	}
}

func TestCellGetFormula(t *testing.T) {
	_, sh := newWorkbookSheet()
	row := sh.AddRow()
	c := row.AddCell()

	if c.GetFormula() != "" {
		t.Errorf("GetFormula() on plain cell = %q, want empty", c.GetFormula())
	}

	c.SetFormulaRaw("SUM(A1:A3)")
	if c.GetFormula() != "SUM(A1:A3)" {
		t.Errorf("GetFormula() = %q, want SUM(A1:A3)", c.GetFormula())
	}
}

func TestCellGetCachedFormulaResult(t *testing.T) {
	_, sh := newWorkbookSheet()
	row := sh.AddRow()
	c := row.AddCell()

	if c.GetCachedFormulaResult() != "" {
		t.Errorf("GetCachedFormulaResult() on empty cell = %q, want empty", c.GetCachedFormulaResult())
	}

	c.SetCachedFormulaResult("99")
	if c.GetCachedFormulaResult() != "99" {
		t.Errorf("GetCachedFormulaResult() = %q, want 99", c.GetCachedFormulaResult())
	}
}

func TestCellColumn(t *testing.T) {
	_, sh := newWorkbookSheet()
	row := sh.Row(1)
	c := row.Cell("B")

	col, err := c.Column()
	if err != nil {
		t.Fatalf("Column() error: %s", err)
	}
	if col != "B" {
		t.Errorf("Column() = %q, want B", col)
	}
}

func TestSheetIsValid(t *testing.T) {
	wb := spreadsheet.New()
	sh := wb.AddSheet()
	if !sh.IsValid() {
		t.Error("newly added sheet should be valid")
	}

	var zero spreadsheet.Sheet
	if zero.IsValid() {
		t.Error("zero-value Sheet should not be valid")
	}
}

func TestSheetColumn(t *testing.T) {
	_, sh := newWorkbookSheet()
	col := sh.Column(1)
	// Column returns a wrapper; just verify it doesn't panic and has an inner value
	if col.X() == nil {
		t.Error("Column(1).X() should not be nil")
	}
}

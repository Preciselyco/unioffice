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

func TestWorkbookSetActiveSheet(t *testing.T) {
	wb := spreadsheet.New()
	sh := wb.AddSheet()
	sh.SetName("First")
	sh2 := wb.AddSheet()
	sh2.SetName("Second")

	// Should not panic; we can't easily read ActiveTabAttr through the public
	// API but we verify the call doesn't error.
	wb.SetActiveSheet(sh2)
	wb.SetActiveSheet(sh)
}

func TestWorkbookSetActiveSheetIndex(t *testing.T) {
	wb := spreadsheet.New()
	wb.AddSheet()
	wb.AddSheet()
	wb.SetActiveSheetIndex(1)
	// Verify bookviews was initialised
	if wb.X().BookViews == nil || len(wb.X().BookViews.WorkbookView) == 0 {
		t.Error("SetActiveSheetIndex should initialise BookViews")
	}
	if wb.X().BookViews.WorkbookView[0].ActiveTabAttr == nil {
		t.Error("ActiveTabAttr should be set")
	}
}

func TestWorkbookTablesEmpty(t *testing.T) {
	wb := spreadsheet.New()
	tables := wb.Tables()
	if tables != nil {
		t.Error("new workbook should have nil tables")
	}
}

func TestWorkbookClearCachedFormulaResults(t *testing.T) {
	wb := spreadsheet.New()
	sh := wb.AddSheet()
	row := sh.AddRow()
	c := row.AddCell()
	c.SetFormulaRaw("1+1")
	c.SetCachedFormulaResult("2")

	// Should not panic and should clear cached results
	wb.ClearCachedFormulaResults()
	if c.GetCachedFormulaResult() != "" {
		t.Errorf("expected empty cached result, got %q", c.GetCachedFormulaResult())
	}
}

func TestWorkbookRecalculateFormulas(t *testing.T) {
	wb := spreadsheet.New()
	sh := wb.AddSheet()
	row := sh.AddRow()
	c1 := row.AddCell()
	c1.SetNumber(5)

	row2 := sh.AddRow()
	c2 := row2.AddCell()
	c2.SetFormulaRaw("A1*2")

	// Should not panic; just a smoke test
	wb.RecalculateFormulas()
}

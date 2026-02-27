// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/Preciselyco/unioffice/common"
)

func TestCustomPropertiesNew(t *testing.T) {
	cp := common.NewCustomProperties()
	if cp.X() == nil {
		t.Fatal("expected non-nil inner type")
	}
	if cp.NonEmpty() {
		t.Error("expected empty custom properties")
	}
}

func TestCustomPropertiesSetString(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsString("Author", "John Doe")

	prop, ok := cp.GetPropertyByName("Author")
	if !ok {
		t.Fatal("expected to find property 'Author'")
	}
	if prop.Name() != "Author" {
		t.Errorf("expected name 'Author', got '%s'", prop.Name())
	}
	if v, ok := prop.Value().(string); !ok || v != "John Doe" {
		t.Errorf("expected value 'John Doe', got '%v'", prop.Value())
	}
}

func TestCustomPropertiesSetInt(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsInt("Version", 42)

	prop, ok := cp.GetPropertyByName("Version")
	if !ok {
		t.Fatal("expected to find property 'Version'")
	}
	if v, ok := prop.Value().(int32); !ok || v != 42 {
		t.Errorf("expected value 42, got '%v'", prop.Value())
	}
}

func TestCustomPropertiesSetBool(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsBool("Approved", true)

	prop, ok := cp.GetPropertyByName("Approved")
	if !ok {
		t.Fatal("expected to find property 'Approved'")
	}
	if v, ok := prop.Value().(bool); !ok || !v {
		t.Errorf("expected value true, got '%v'", prop.Value())
	}
}

func TestCustomPropertiesSetFloat64(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsFloat64("Score", 3.14)

	prop, ok := cp.GetPropertyByName("Score")
	if !ok {
		t.Fatal("expected to find property 'Score'")
	}
	if v, ok := prop.Value().(float64); !ok || v != 3.14 {
		t.Errorf("expected value 3.14, got '%v'", prop.Value())
	}
}

func TestCustomPropertiesSetDate(t *testing.T) {
	cp := common.NewCustomProperties()
	dt := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	cp.SetPropertyAsDate("ReviewDate", dt)

	prop, ok := cp.GetPropertyByName("ReviewDate")
	if !ok {
		t.Fatal("expected to find property 'ReviewDate'")
	}
	if v, ok := prop.Value().(time.Time); !ok || !v.Equal(dt) {
		t.Errorf("expected value %v, got '%v'", dt, prop.Value())
	}
}

func TestCustomPropertiesUpdateExisting(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsString("Status", "Draft")
	cp.SetPropertyAsString("Status", "Final")

	props := cp.Properties()
	count := 0
	for _, p := range props {
		if p.Name() == "Status" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("expected exactly 1 'Status' property, got %d", count)
	}

	prop, ok := cp.GetPropertyByName("Status")
	if !ok {
		t.Fatal("expected to find property 'Status'")
	}
	if v, ok := prop.Value().(string); !ok || v != "Final" {
		t.Errorf("expected value 'Final', got '%v'", prop.Value())
	}
}

func TestCustomPropertiesChangeType(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsString("Mixed", "text")
	cp.SetPropertyAsInt("Mixed", 99)

	prop, ok := cp.GetPropertyByName("Mixed")
	if !ok {
		t.Fatal("expected to find property 'Mixed'")
	}
	if v, ok := prop.Value().(int32); !ok || v != 99 {
		t.Errorf("expected int32 value 99, got '%v'", prop.Value())
	}
}

func TestCustomPropertiesRemoveByName(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsString("ToRemove", "value")
	cp.SetPropertyAsString("ToKeep", "value")

	cp.RemoveByName("ToRemove")

	if _, ok := cp.GetPropertyByName("ToRemove"); ok {
		t.Error("expected 'ToRemove' to be removed")
	}
	if _, ok := cp.GetPropertyByName("ToKeep"); !ok {
		t.Error("expected 'ToKeep' to still exist")
	}
}

func TestCustomPropertiesPIDAssignment(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsString("First", "a")
	cp.SetPropertyAsString("Second", "b")
	cp.SetPropertyAsString("Third", "c")

	props := cp.Properties()
	if len(props) != 3 {
		t.Fatalf("expected 3 properties, got %d", len(props))
	}

	// PIDs should be unique and >= 2
	pids := make(map[int32]bool)
	for _, p := range props {
		pid := p.X().PidAttr
		if pid < 2 {
			t.Errorf("expected PID >= 2, got %d", pid)
		}
		if pids[pid] {
			t.Errorf("duplicate PID %d", pid)
		}
		pids[pid] = true
	}
}

func TestCustomPropertiesNonEmpty(t *testing.T) {
	cp := common.NewCustomProperties()
	if cp.NonEmpty() {
		t.Error("expected NonEmpty to return false for empty properties")
	}
	cp.SetPropertyAsString("Key", "Value")
	if !cp.NonEmpty() {
		t.Error("expected NonEmpty to return true after adding a property")
	}
}

func TestCustomPropertiesXMLRoundTrip(t *testing.T) {
	cp := common.NewCustomProperties()
	cp.SetPropertyAsString("StringProp", "hello world")
	cp.SetPropertyAsInt("IntProp", 123)
	cp.SetPropertyAsBool("BoolProp", true)
	cp.SetPropertyAsFloat64("FloatProp", 2.718)
	dt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	cp.SetPropertyAsDate("DateProp", dt)

	// Marshal
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	if err := enc.Encode(cp.X()); err != nil {
		t.Fatalf("error marshaling custom properties: %s", err)
	}
	enc.Flush()

	// Unmarshal
	cp2 := common.NewCustomProperties()
	dec := xml.NewDecoder(&buf)
	if err := dec.Decode(cp2.X()); err != nil {
		t.Fatalf("error unmarshaling custom properties: %s", err)
	}

	// Verify
	tests := []struct {
		name string
		want interface{}
	}{
		{"StringProp", "hello world"},
		{"IntProp", int32(123)},
		{"BoolProp", true},
		{"FloatProp", 2.718},
	}
	for _, tc := range tests {
		prop, ok := cp2.GetPropertyByName(tc.name)
		if !ok {
			t.Errorf("expected to find property '%s' after round-trip", tc.name)
			continue
		}
		if prop.Value() != tc.want {
			t.Errorf("property '%s': expected %v (%T), got %v (%T)", tc.name, tc.want, tc.want, prop.Value(), prop.Value())
		}
	}

	// Date needs special comparison
	dateProp, ok := cp2.GetPropertyByName("DateProp")
	if !ok {
		t.Error("expected to find property 'DateProp' after round-trip")
	} else {
		if v, ok := dateProp.Value().(time.Time); !ok || !v.Equal(dt) {
			t.Errorf("property 'DateProp': expected %v, got %v", dt, dateProp.Value())
		}
	}
}

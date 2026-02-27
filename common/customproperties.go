// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"time"

	"github.com/Preciselyco/unioffice/schema/soo/ofc/custom_properties"
)

// CustomProperties contains user-defined custom properties of the document.
type CustomProperties struct {
	x *custom_properties.Properties
}

// NewCustomProperties constructs a new CustomProperties.
func NewCustomProperties() CustomProperties {
	return CustomProperties{x: custom_properties.NewProperties()}
}

// X returns the inner wrapped XML type.
func (c CustomProperties) X() *custom_properties.Properties {
	return c.x
}

// SetPropertyAsString sets or adds a custom property with a string value.
func (c CustomProperties) SetPropertyAsString(name, value string) {
	prop := c.getOrCreate(name)
	c.clearValue(prop)
	prop.Lpwstr = &value
}

// SetPropertyAsInt sets or adds a custom property with an int32 value.
func (c CustomProperties) SetPropertyAsInt(name string, value int32) {
	prop := c.getOrCreate(name)
	c.clearValue(prop)
	prop.I4 = &value
}

// SetPropertyAsBool sets or adds a custom property with a bool value.
func (c CustomProperties) SetPropertyAsBool(name string, value bool) {
	prop := c.getOrCreate(name)
	c.clearValue(prop)
	prop.Bool = &value
}

// SetPropertyAsFloat64 sets or adds a custom property with a float64 value.
func (c CustomProperties) SetPropertyAsFloat64(name string, value float64) {
	prop := c.getOrCreate(name)
	c.clearValue(prop)
	prop.R8 = &value
}

// SetPropertyAsDate sets or adds a custom property with a date/time value.
func (c CustomProperties) SetPropertyAsDate(name string, value time.Time) {
	prop := c.getOrCreate(name)
	c.clearValue(prop)
	prop.Filetime = &value
}

// GetPropertyByName returns the custom property with the given name, if it exists.
func (c CustomProperties) GetPropertyByName(name string) (CustomProperty, bool) {
	for _, prop := range c.x.Property {
		if prop.NameAttr != nil && *prop.NameAttr == name {
			return CustomProperty{x: prop}, true
		}
	}
	return CustomProperty{}, false
}

// Properties returns all custom properties.
func (c CustomProperties) Properties() []CustomProperty {
	ret := make([]CustomProperty, 0, len(c.x.Property))
	for _, prop := range c.x.Property {
		ret = append(ret, CustomProperty{x: prop})
	}
	return ret
}

// RemoveByName removes a custom property by name.
func (c CustomProperties) RemoveByName(name string) {
	for i, prop := range c.x.Property {
		if prop.NameAttr != nil && *prop.NameAttr == name {
			copy(c.x.Property[i:], c.x.Property[i+1:])
			c.x.Property = c.x.Property[:len(c.x.Property)-1]
			return
		}
	}
}

// NonEmpty returns true if there are any custom properties defined.
func (c CustomProperties) NonEmpty() bool {
	return len(c.x.Property) > 0
}

func (c CustomProperties) nextPID() int32 {
	maxPID := int32(1)
	for _, prop := range c.x.Property {
		if prop.PidAttr > maxPID {
			maxPID = prop.PidAttr
		}
	}
	return maxPID + 1
}

func (c CustomProperties) getOrCreate(name string) *custom_properties.CT_Property {
	for _, prop := range c.x.Property {
		if prop.NameAttr != nil && *prop.NameAttr == name {
			return prop
		}
	}
	prop := custom_properties.NewCT_Property()
	prop.NameAttr = &name
	prop.PidAttr = c.nextPID()
	c.x.Property = append(c.x.Property, prop)
	return prop
}

func (c CustomProperties) clearValue(prop *custom_properties.CT_Property) {
	prop.Lpwstr = nil
	prop.Lpstr = nil
	prop.I4 = nil
	prop.R8 = nil
	prop.Bool = nil
	prop.Filetime = nil
}

// CustomProperty is a single custom property with a name and typed value.
type CustomProperty struct {
	x *custom_properties.CT_Property
}

// X returns the inner wrapped XML type.
func (cp CustomProperty) X() *custom_properties.CT_Property {
	return cp.x
}

// Name returns the name of the custom property.
func (cp CustomProperty) Name() string {
	if cp.x.NameAttr != nil {
		return *cp.x.NameAttr
	}
	return ""
}

// Value returns the value of the custom property as an interface{}.
func (cp CustomProperty) Value() interface{} {
	if cp.x.Lpwstr != nil {
		return *cp.x.Lpwstr
	}
	if cp.x.Lpstr != nil {
		return *cp.x.Lpstr
	}
	if cp.x.I4 != nil {
		return *cp.x.I4
	}
	if cp.x.R8 != nil {
		return *cp.x.R8
	}
	if cp.x.Bool != nil {
		return *cp.x.Bool
	}
	if cp.x.Filetime != nil {
		return *cp.x.Filetime
	}
	return nil
}

package custom_properties

import "github.com/Preciselyco/unioffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Property", NewCT_Property)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Properties", NewCT_Properties)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "Properties", NewProperties)
}

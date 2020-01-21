// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"

	"github.com/Preciselyco/unioffice"
)

type CT_SlideTransitionChoice struct {
	Blinds    *CT_OrientationTransition
	Checker   *CT_OrientationTransition
	Circle    *CT_Empty
	Dissolve  *CT_Empty
	Comb      *CT_OrientationTransition
	Cover     *CT_EightDirectionTransition
	Cut       *CT_OptionalBlackTransition
	Diamond   *CT_Empty
	Fade      *CT_OptionalBlackTransition
	Newsflash *CT_Empty
	Plus      *CT_Empty
	Pull      *CT_EightDirectionTransition
	Push      *CT_SideDirectionTransition
	Random    *CT_Empty
	RandomBar *CT_OrientationTransition
	Split     *CT_SplitTransition
	Strips    *CT_CornerDirectionTransition
	Wedge     *CT_Empty
	Wheel     *CT_WheelTransition
	Wipe      *CT_SideDirectionTransition
	Zoom      *CT_InOutTransition
}

func NewCT_SlideTransitionChoice() *CT_SlideTransitionChoice {
	ret := &CT_SlideTransitionChoice{}
	return ret
}

func (m *CT_SlideTransitionChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Blinds != nil {
		seblinds := xml.StartElement{Name: xml.Name{Local: "p:blinds"}}
		e.EncodeElement(m.Blinds, seblinds)
	}
	if m.Checker != nil {
		sechecker := xml.StartElement{Name: xml.Name{Local: "p:checker"}}
		e.EncodeElement(m.Checker, sechecker)
	}
	if m.Circle != nil {
		secircle := xml.StartElement{Name: xml.Name{Local: "p:circle"}}
		e.EncodeElement(m.Circle, secircle)
	}
	if m.Dissolve != nil {
		sedissolve := xml.StartElement{Name: xml.Name{Local: "p:dissolve"}}
		e.EncodeElement(m.Dissolve, sedissolve)
	}
	if m.Comb != nil {
		secomb := xml.StartElement{Name: xml.Name{Local: "p:comb"}}
		e.EncodeElement(m.Comb, secomb)
	}
	if m.Cover != nil {
		secover := xml.StartElement{Name: xml.Name{Local: "p:cover"}}
		e.EncodeElement(m.Cover, secover)
	}
	if m.Cut != nil {
		secut := xml.StartElement{Name: xml.Name{Local: "p:cut"}}
		e.EncodeElement(m.Cut, secut)
	}
	if m.Diamond != nil {
		sediamond := xml.StartElement{Name: xml.Name{Local: "p:diamond"}}
		e.EncodeElement(m.Diamond, sediamond)
	}
	if m.Fade != nil {
		sefade := xml.StartElement{Name: xml.Name{Local: "p:fade"}}
		e.EncodeElement(m.Fade, sefade)
	}
	if m.Newsflash != nil {
		senewsflash := xml.StartElement{Name: xml.Name{Local: "p:newsflash"}}
		e.EncodeElement(m.Newsflash, senewsflash)
	}
	if m.Plus != nil {
		seplus := xml.StartElement{Name: xml.Name{Local: "p:plus"}}
		e.EncodeElement(m.Plus, seplus)
	}
	if m.Pull != nil {
		sepull := xml.StartElement{Name: xml.Name{Local: "p:pull"}}
		e.EncodeElement(m.Pull, sepull)
	}
	if m.Push != nil {
		sepush := xml.StartElement{Name: xml.Name{Local: "p:push"}}
		e.EncodeElement(m.Push, sepush)
	}
	if m.Random != nil {
		serandom := xml.StartElement{Name: xml.Name{Local: "p:random"}}
		e.EncodeElement(m.Random, serandom)
	}
	if m.RandomBar != nil {
		serandomBar := xml.StartElement{Name: xml.Name{Local: "p:randomBar"}}
		e.EncodeElement(m.RandomBar, serandomBar)
	}
	if m.Split != nil {
		sesplit := xml.StartElement{Name: xml.Name{Local: "p:split"}}
		e.EncodeElement(m.Split, sesplit)
	}
	if m.Strips != nil {
		sestrips := xml.StartElement{Name: xml.Name{Local: "p:strips"}}
		e.EncodeElement(m.Strips, sestrips)
	}
	if m.Wedge != nil {
		sewedge := xml.StartElement{Name: xml.Name{Local: "p:wedge"}}
		e.EncodeElement(m.Wedge, sewedge)
	}
	if m.Wheel != nil {
		sewheel := xml.StartElement{Name: xml.Name{Local: "p:wheel"}}
		e.EncodeElement(m.Wheel, sewheel)
	}
	if m.Wipe != nil {
		sewipe := xml.StartElement{Name: xml.Name{Local: "p:wipe"}}
		e.EncodeElement(m.Wipe, sewipe)
	}
	if m.Zoom != nil {
		sezoom := xml.StartElement{Name: xml.Name{Local: "p:zoom"}}
		e.EncodeElement(m.Zoom, sezoom)
	}
	return nil
}

func (m *CT_SlideTransitionChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideTransitionChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "blinds"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "blinds"}:
				m.Blinds = NewCT_OrientationTransition()
				if err := d.DecodeElement(m.Blinds, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "checker"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "checker"}:
				m.Checker = NewCT_OrientationTransition()
				if err := d.DecodeElement(m.Checker, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "circle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "circle"}:
				m.Circle = NewCT_Empty()
				if err := d.DecodeElement(m.Circle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "dissolve"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "dissolve"}:
				m.Dissolve = NewCT_Empty()
				if err := d.DecodeElement(m.Dissolve, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "comb"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "comb"}:
				m.Comb = NewCT_OrientationTransition()
				if err := d.DecodeElement(m.Comb, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cover"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cover"}:
				m.Cover = NewCT_EightDirectionTransition()
				if err := d.DecodeElement(m.Cover, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cut"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cut"}:
				m.Cut = NewCT_OptionalBlackTransition()
				if err := d.DecodeElement(m.Cut, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "diamond"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "diamond"}:
				m.Diamond = NewCT_Empty()
				if err := d.DecodeElement(m.Diamond, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "fade"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "fade"}:
				m.Fade = NewCT_OptionalBlackTransition()
				if err := d.DecodeElement(m.Fade, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "newsflash"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "newsflash"}:
				m.Newsflash = NewCT_Empty()
				if err := d.DecodeElement(m.Newsflash, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "plus"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "plus"}:
				m.Plus = NewCT_Empty()
				if err := d.DecodeElement(m.Plus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pull"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pull"}:
				m.Pull = NewCT_EightDirectionTransition()
				if err := d.DecodeElement(m.Pull, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "push"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "push"}:
				m.Push = NewCT_SideDirectionTransition()
				if err := d.DecodeElement(m.Push, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "random"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "random"}:
				m.Random = NewCT_Empty()
				if err := d.DecodeElement(m.Random, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "randomBar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "randomBar"}:
				m.RandomBar = NewCT_OrientationTransition()
				if err := d.DecodeElement(m.RandomBar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "split"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "split"}:
				m.Split = NewCT_SplitTransition()
				if err := d.DecodeElement(m.Split, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "strips"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "strips"}:
				m.Strips = NewCT_CornerDirectionTransition()
				if err := d.DecodeElement(m.Strips, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "wedge"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "wedge"}:
				m.Wedge = NewCT_Empty()
				if err := d.DecodeElement(m.Wedge, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "wheel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "wheel"}:
				m.Wheel = NewCT_WheelTransition()
				if err := d.DecodeElement(m.Wheel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "wipe"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "wipe"}:
				m.Wipe = NewCT_SideDirectionTransition()
				if err := d.DecodeElement(m.Wipe, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "zoom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "zoom"}:
				m.Zoom = NewCT_InOutTransition()
				if err := d.DecodeElement(m.Zoom, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SlideTransitionChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideTransitionChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideTransitionChoice and its children
func (m *CT_SlideTransitionChoice) Validate() error {
	return m.ValidateWithPath("CT_SlideTransitionChoice")
}

// ValidateWithPath validates the CT_SlideTransitionChoice and its children, prefixing error messages with path
func (m *CT_SlideTransitionChoice) ValidateWithPath(path string) error {
	if m.Blinds != nil {
		if err := m.Blinds.ValidateWithPath(path + "/Blinds"); err != nil {
			return err
		}
	}
	if m.Checker != nil {
		if err := m.Checker.ValidateWithPath(path + "/Checker"); err != nil {
			return err
		}
	}
	if m.Circle != nil {
		if err := m.Circle.ValidateWithPath(path + "/Circle"); err != nil {
			return err
		}
	}
	if m.Dissolve != nil {
		if err := m.Dissolve.ValidateWithPath(path + "/Dissolve"); err != nil {
			return err
		}
	}
	if m.Comb != nil {
		if err := m.Comb.ValidateWithPath(path + "/Comb"); err != nil {
			return err
		}
	}
	if m.Cover != nil {
		if err := m.Cover.ValidateWithPath(path + "/Cover"); err != nil {
			return err
		}
	}
	if m.Cut != nil {
		if err := m.Cut.ValidateWithPath(path + "/Cut"); err != nil {
			return err
		}
	}
	if m.Diamond != nil {
		if err := m.Diamond.ValidateWithPath(path + "/Diamond"); err != nil {
			return err
		}
	}
	if m.Fade != nil {
		if err := m.Fade.ValidateWithPath(path + "/Fade"); err != nil {
			return err
		}
	}
	if m.Newsflash != nil {
		if err := m.Newsflash.ValidateWithPath(path + "/Newsflash"); err != nil {
			return err
		}
	}
	if m.Plus != nil {
		if err := m.Plus.ValidateWithPath(path + "/Plus"); err != nil {
			return err
		}
	}
	if m.Pull != nil {
		if err := m.Pull.ValidateWithPath(path + "/Pull"); err != nil {
			return err
		}
	}
	if m.Push != nil {
		if err := m.Push.ValidateWithPath(path + "/Push"); err != nil {
			return err
		}
	}
	if m.Random != nil {
		if err := m.Random.ValidateWithPath(path + "/Random"); err != nil {
			return err
		}
	}
	if m.RandomBar != nil {
		if err := m.RandomBar.ValidateWithPath(path + "/RandomBar"); err != nil {
			return err
		}
	}
	if m.Split != nil {
		if err := m.Split.ValidateWithPath(path + "/Split"); err != nil {
			return err
		}
	}
	if m.Strips != nil {
		if err := m.Strips.ValidateWithPath(path + "/Strips"); err != nil {
			return err
		}
	}
	if m.Wedge != nil {
		if err := m.Wedge.ValidateWithPath(path + "/Wedge"); err != nil {
			return err
		}
	}
	if m.Wheel != nil {
		if err := m.Wheel.ValidateWithPath(path + "/Wheel"); err != nil {
			return err
		}
	}
	if m.Wipe != nil {
		if err := m.Wipe.ValidateWithPath(path + "/Wipe"); err != nil {
			return err
		}
	}
	if m.Zoom != nil {
		if err := m.Zoom.ValidateWithPath(path + "/Zoom"); err != nil {
			return err
		}
	}
	return nil
}

// docx-tool is a command-line utility for debugging docx files and unioffice.

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-xmlfmt/xmlfmt" // Regex-based XML formatter library, not an XML parser

	"github.com/Preciselyco/unioffice/document"
)

type DumpXmlCmd struct {
	DocxFile string `arg required help:"Input file"`
	XmlFile  string `arg optional help:"File to extract"`
}

func (cmd *DumpXmlCmd) Run() error {
	zipReader, err := zip.OpenReader(cmd.DocxFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Sort list of files to make it diffable
	sort.Slice(zipReader.File, func(i, j int) bool { return zipReader.File[i].Name < zipReader.File[j].Name })

	for _, file := range zipReader.File {
		if cmd.XmlFile == "" {
			// Default mode; list all files
			fmt.Println("--------------------------------------------------------------------------------")
			fmt.Println("File: " + file.Name)
			fmt.Println("--------------------------------------------------------------------------------")
			if strings.HasSuffix(file.Name, ".xml") || strings.HasSuffix(file.Name, ".rels") {
				if err := cmd.DumpFile(file); err != nil {
					return err
				}
			} else {
				fmt.Println("Not an XML file")
			}
		} else if file.Name == cmd.XmlFile {
			// Filter on filename if provided
			if err := cmd.DumpFile(file); err != nil {
				return err
			}
		}
	}
	return nil
}

func (cmd *DumpXmlCmd) DumpFile(file *zip.File) error {
	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	fmt.Println(strings.Trim(xmlfmt.FormatXML(string(data), "", "  "), "\r\n"))
	return nil
}

type SpewCmd struct {
	NoNil    bool   `help:"Filter out nil values"`
	DocxFile string `arg required help:"Input file"`
}

func (cmd *SpewCmd) Run() error {
	docx, err := document.Open(cmd.DocxFile)
	if err != nil {
		return err
	}
	if err := docx.Validate(); err != nil {
		return err
	}
	str := spew.Sdump(docx)
	if cmd.NoNil {
		// Filter out lines with nil pointers for shorter output
		str = regexp.MustCompile(`(?m)^.*<nil>.*$\n`).ReplaceAllString(str, "")
	}
	fmt.Print(str)
	return nil
}

type RecodeCmd struct {
	InFile  string `arg required help:"Input file"`
	OutFile string `arg required help:"Output file"`
}

func (cmd *RecodeCmd) Run() error {
	docx, err := document.Open(cmd.InFile)
	if err != nil {
		return err
	}
	if err := docx.Validate(); err != nil {
		return err
	}
	if err := docx.SaveToFile(cmd.OutFile); err != nil {
		return err
	}
	return nil
}

func main() {
	cli := struct {
		DumpXml DumpXmlCmd `cmd help:"Pretty-print XML file(s) from DOCX file"`
		Spew    SpewCmd    `cmd help:"Parse DOCX file and dump data structure"`
		Recode  RecodeCmd  `cmd help:"Parse DOCX file and save to new file"`
	}{}
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}

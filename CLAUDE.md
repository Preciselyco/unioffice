# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`unioffice` is a pure-Go library for creating, reading, and editing ECMA-376 Office Open
XML files: Word (`.docx`), Excel (`.xlsx`), and PowerPoint (`.pptx`). This repo is
**Preciselyco's fork** of `unidoc/unioffice` (module path
`github.com/Preciselyco/unioffice`), maintained with targeted compatibility and
correctness fixes rather than upstream feature work.

## Build / test / lint

Standard Go modules (`go 1.13`). There is no Makefile.

```bash
go build ./...                              # build everything
go test ./...                               # run all tests
go vet ./...                                # vet
go test -run TestName ./document/           # single test in a package
go test -v ./spreadsheet/...                # one subtree, verbose
```

Note: `.travis.yml`, `build-examples.sh`, and `test-coverage.sh` still reference the old
`github.com/unidoc/unioffice` GOPATH layout and are **not** an accurate guide for this
fork — use the module commands above. The `_examples/` tree is not part of the module
build; each example is its own `main.go`.

## Architecture

The library has two layers:

1. **Raw schema types (`schema/`)** — generated Go structs covering the full ECMA-376
   standard, mirroring the XML 1:1 (`CT_*`, `ST_*`, `EG_*` types). Organized by namespace,
   e.g. `schema/soo/wml` (WordprocessingML), `schema/soo/sml` (SpreadsheetML),
   `schema/soo/pml` (PresentationML), `schema/soo/dml` (DrawingML). This is **generated
   code (~3500 files) — do not hand-edit it.**

2. **Wrapper packages** — ergonomic APIs over the raw types:
   - `document/` — Word documents
   - `spreadsheet/` — Excel workbooks (incl. `spreadsheet/formula` evaluation,
     `spreadsheet/format`, `spreadsheet/reference`)
   - `presentation/` — PowerPoint presentations

Every wrapper exposes an `X()` method returning its underlying raw `schema` struct, so
any part of the spec not covered by the friendly API can be manipulated directly. Wrapper
types are **value types with non-pointer methods**, except the three base documents —
`document.Document`, `spreadsheet.Workbook`, `presentation.Presentation` — which are
pointers and own the document lifecycle.

### Cross-cutting packages

- `common/` — shared OOXML plumbing used by all three formats: `docbase.go`,
  relationships, content types, core/app/custom properties, theme, and image handling.
- `zippkg/` — (de)serialization of the OOXML package (a zip of XML parts); marshaling,
  decode mapping, self-closing-tag writer.
- `chart/`, `drawing/`, `vmldrawing/`, `color/`, `measurement/`, `algo/` — helper packages.
- Root package `unioffice` — content-type/relationship constants (`schemas.go`),
  `xsd:any` machinery (`RegisterConstructor`/`CreateElement` in `creator.go`),
  licensing (`license.go`), `optional.go` pointer helpers.

### Document lifecycle

`New`, `Open`, `OpenTemplate`, `Read`, `ReadFromBytes` construct/load; `Save`,
`SaveToFile` write; `Close` releases temp resources. The same New/Open/Save pattern holds
for `spreadsheet` and `presentation`. See `document/document.go` for the canonical
implementation.

## Conventions

- Fixes to raw-XML serialization behavior belong in the `schema/` `MarshalXML`/
  `UnmarshalXML` methods, but treat the directory as generated — changes there must be
  applied systematically across all affected types (see recent namespace-prefix fix
  commits for the pattern).
- Tests live next to source as `*_test.go` and use `stretchr/testify`. Fixtures are in
  per-package `testdata/` dirs and the top-level `testdata/`; `testhelper/` has
  comparison helpers for document equality checks.

## CLI tools

`cmd/` contains standalone utilities: `docx-tool`, `docx2md`, `csv2xlsx`, `xlsx2csv`,
`catdoc`, `unprotect-xlsx`, `test-open-write`. Build/run one with
`go run ./cmd/<name>`.

## Licensing

Dual-licensed AGPLv3 / commercial. License enforcement is in `license.go`; an open-source
license is installed by default.

# docx-tool

Command-line tool for reverse-engineering docx files and testing the unioffice library.

### List the XML contents of a docx file

```
docx-tool dump-xml file.docx
```

### Dump the unioffice data structure of a docx file

```
docx-tool spew [--no-nil] file.docx
```

The `--no-nil` option filters out nil struct members and makes the output considerably shorter.

### Read a docx file with unioffice and write it to file again

```
docx-tool recode infile.docx outfile.docx
```

This is useful for testing that unioffice does not change the data.

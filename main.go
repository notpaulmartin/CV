package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/notpaulmartin/mdParser"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/bep/godartsass/v2"
)

const (
	mdFname   = "cv.md"
	sassFname = "styling.scss"
	htmlFname = "cv_out.html"
	pdfFname  = "cv_pmartin.pdf"
)

func main() {
	// Read input Markdown string from file
	mdBytes, err := ioutil.ReadFile(mdFname)
	if err != nil {
		log.Panic("Could not read input file: " + mdFname)
	}

	// convert md to html
	htmlChan := make(chan string)
	go mdToHtml(mdBytes, htmlChan) // NOTE: myMdToHtml uses my custom markdown parser

	// convert scss to css
	cssChan := make(chan string)
	go parseSass(sassFname, cssChan)

	// join html with css
	html := "<!DOCTYPE html>\n <html>" +
		"<head>" +
		"<meta charset=\"utf-8\" />" +
		"<style>" +
		<-cssChan +
		"</style>" +
		"</head>" +
		"<body>" +
		<-htmlChan +
		"</body>" +
		"</html>"

	saveHtml(html, "./"+htmlFname)

	// Convert to PDF
	pdfg := htmlToPdf(html)

	// Write PDF to file
	err = pdfg.WriteFile("./" + pdfFname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote ./" + pdfFname)
}

func mdToHtml(mdBytes []byte, c chan string) {
	var buf bytes.Buffer
	parser := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.TaskList),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)

	if err := parser.Convert(mdBytes, &buf); err != nil {
		panic(err)
	}

	c <- buf.String()
}

func myMdToHtml(mdBytes []byte, c chan string) {
	c <- mdParser.MdToHtml(string(mdBytes))
}

func parseSass(fname string, c chan string) {
    // Read input Sass string from file
    scssBytes, err := ioutil.ReadFile(fname)
    if err != nil {
        log.Panic("Could not read input file: " + fname) // Also fixed variable name here
    }

    transpiler, err := godartsass.Start(godartsass.Options{})
    if err != nil {
        log.Panic("Failed to start Sass transpiler: ", err)
    }

    args := godartsass.Args{
        OutputStyle: godartsass.OutputStyleExpanded,
        Source:      string(scssBytes),
    }

    result, err := transpiler.Execute(args)
    if err != nil {
        log.Panic("Failed to execute Sass transpiler: ", err)
    }

    c <- result.CSS
}

func saveHtml(html string, fname string) {
	// Open file
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Write to file
	_, err = f.WriteString(html)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote " + fname)
}

func htmlToPdf(html string) *wkhtmltopdf.PDFGenerator {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// Add to document
	page := wkhtmltopdf.NewPageReader(strings.NewReader(string(html)))
	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	return pdfg
}

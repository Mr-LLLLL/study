package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"pdf/tmpl"
	"sync"
	"time"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	printByStream()
}

func printhtml() {
	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	htmlfile, err := os.Open("order_print.html")
	if err != nil {
		log.Fatal(err)
	}
	defer htmlfile.Close()

	pdfg.OutputFile = "test.pdf"

	pdfg.AddPage(pdf.NewPageReader(htmlfile))
	pdfg.AddPage(pdf.NewPageReader(htmlfile))
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}

func printByStream() {
	model := tmpl.TemplateModel{
		PrintDate:  time.Now().String(),
		OrderTitle: "hello",
		OrderInfo: tmpl.OrderInfo{
			ShowOrderInfo: true,
		},
	}
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "wworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworldorlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworld")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "wworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworldorlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworld")
	model.OrderInfo.TemplateModel.Append("hello : ", "wworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworldorlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworld")
	model.OrderInfo.TemplateModel.Append("hello : ", "wworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworldorlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworld")
	model.OrderInfo.TemplateModel.Append("hello : ", "wworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworldorlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworlddddddddworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworlworldddddddddddddddddddworld")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")
	model.OrderInfo.TemplateModel.Append("hello : ", "world")

	t, err := template.New("letter").Parse(tmpl.Html)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, model)
	if err != nil {
		log.Fatal(err)
	}

	page := pdf.NewPageReader(buf)
	// page.DisableSmartShrinking.Set(true)

	pdfg := pdf.NewPDFPreparer()
	pdfg.AddPage(page)
	jb, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	pdfgFromJSON, err := pdf.NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	fileName := "test.pdf"

	err = pdfgFromJSON.WriteFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
}

func print() {
	buf := new(bytes.Buffer)
	html, err := ioutil.ReadFile("test.template")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("letter").Parse(string(html))
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	ch := make(chan *tmpl.TemplateModel)
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// _ = t.Execute(buf, model)
			// ch <- &model
		}()
	}
	go func() {
		for c := range ch {
			t.Execute(buf, c)
		}
	}()
	wg.Wait()
	close(ch)

	pdfg := pdf.NewPDFPreparer()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "--auto-size", "")

	err = pdfg.CreateContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// htmlfile2, err := ioutil.ReadFile("order_print.html")
	// if err != nil {
	//     log.Fatal(err)
	// }

	// buf.Write(htmlfile1)
	// buf.Write(htmlfile2)

	pdfg.AddPage(pdf.NewPageReader(buf))
	// pdfg.Dpi.Set(600)
	// pdfg.AddPage(pdf.NewPageReader(bytes.NewReader(htmlfile)))

	// The contents of htmlsimple.html are saved as base64 string in the JSON file
	jb, err := pdfg.ToJSON()
	// if err != nil {
	//     log.Fatal(err)
	// }

	// Server code
	// pdfgFromJSON, err := pdf.NewPDFGeneratorFromJSON(buf)
	pdfgFromJSON, err := pdf.NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	pdfgFromJSON.WriteFile("test.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

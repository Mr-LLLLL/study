package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const temp1 = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(temp1))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>" // 不受信任的纯文本
	data.B = "<b>Hello!</b>" // 受信任的HTML
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

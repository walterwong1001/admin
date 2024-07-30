package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Field struct {
	Name string
	Type string
}

type TemplateData struct {
	UpperName string
	LowerName string
	Fields    []Field
}

func main() {
	data := TemplateData{
		UpperName: "Role",
		LowerName: "role",
		Fields: []Field{
			{Name: "ID", Type: "uint64"},
			{Name: "Name", Type: "string"},
			{Name: "Description", Type: "string"},
			{Name: "CreateTime", Type: "int64"},
		},
	}

	targetFileName := data.LowerName + ".go"
	templates := []struct {
		Filename  string
		Template  string
		TargetDir string
	}{
		{targetFileName, "model_template.go.tmpl", "../models"},
		{targetFileName, "repository_template.go.tmpl", "../repositories"},
		{targetFileName, "service_template.go.tmpl", "../services"},
		{targetFileName, "handler_template.go.tmpl", "../handlers"},
	}

	for _, tmpl := range templates {
		t, err := template.New(tmpl.Template).Funcs(template.FuncMap{
			"ToLower": strings.ToLower,
		}).ParseFiles(tmpl.Template)
		if err != nil {
			panic(err)
		}

		// 创建目标目录（如果不存在）
		if err := os.MkdirAll(tmpl.TargetDir, os.ModePerm); err != nil {
			panic(err)
		}
		targetFilePath := filepath.Join(tmpl.TargetDir, tmpl.Filename)
		f, err := os.Create(targetFilePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = t.Execute(f, data)
		if err != nil {
			panic(err)
		}
	}
	println("Code generated successfully!")
}

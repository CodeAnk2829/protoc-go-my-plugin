package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue;
			}
			generatedFile(gen, f)
		}
		return nil
	})
}

func generatedFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_my-plugin.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by my-plugin. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.P(`import "github.com/common-nighthawk/go-figure"`)
	g.P()
	
	for _, msg := range file.Messages {
		g.P("func (req *", msg.GoIdent, ") ShowMessages() string {")
		g.P("var result string")
		g.P(`fontStyle := "doom"`)
		
		for _, field := range msg.Fields {
			fieldName := field.GoName
			fieldType := field.Desc.Kind()
			
			switch fieldType {
			case protoreflect.StringKind:
				g.P("result += req.", fieldName)
			default:
				g.P("// other fields are not added yet")
			}
		}
		
		g.P("fig := figure.NewFigure(result, fontStyle, true)")
		g.P("return fig.String()")
		g.P("}")
	}
	
	return g
}
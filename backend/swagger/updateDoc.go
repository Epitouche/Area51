package swagger

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"
)

type SwaggerUpdateDoc interface {
	UpdateDocTemplate(filepath string) (string, error)
	RemoveSchemesLine(rawValue string) string
	UpdateDocTemplateWithJSON(filepath, tmpFilepath string) error
}

func UpdateDocTemplate(filepath string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filepath, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok.String() != "const" {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Names) == 0 {
				continue
			}

			if valueSpec.Names[0].Name == "docTemplate" {
				rawValue := valueSpec.Values[0].(*ast.BasicLit).Value
				rawValue = strings.Trim(rawValue, "`")
				rawValue = RemoveSchemesLine(rawValue)
				os.WriteFile("tmp.json", []byte(rawValue), 0644)
				return rawValue, nil
			}
		}
	}

	fmt.Println("docTemplate constant not found.")
	return "", nil
}

func RemoveSchemesLine(rawValue string) string {
	re := regexp.MustCompile(`(?m)^\s*"schemes":.*\n`)
	updatedValue := re.ReplaceAllString(rawValue, "")
	return updatedValue
}

func UpdateDocTemplateWithJSON(filepath, tmpFilepath string) error {
	tmpContent, err := os.ReadFile(tmpFilepath)
	if err != nil {
		return fmt.Errorf("error reading tmp.json: %w", err)
	}

	prefixedContent := fmt.Sprintf(`{
		"schemes": {{ marshal .Schemes }},
	%s`, tmpContent[1:])

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse Go file: %w", err)
	}

	found := false
	ast.Inspect(node, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			return true
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Names) == 0 {
				continue
			}

			if valueSpec.Names[0].Name == "docTemplate" {
				rawString := fmt.Sprintf("`%s`", prefixedContent)
				valueSpec.Values[0] = &ast.BasicLit{
					Kind: token.STRING,
					Value: rawString,
				}
				found = true
				return false
			}
		}
		return true
	})

	if !found {
		return fmt.Errorf("docTemplate constant not found in file: %s", filepath)
	}

	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, node); err != nil {
		return fmt.Errorf("error printing updated Go file: %w", err)
	}

	err = os.WriteFile(filepath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing updated Go file: %w", err)
	}

	fmt.Printf("Successfully updated docTemplate in file: %s\n", filepath)
	return nil
}
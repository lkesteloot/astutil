// Copyright 2013 Lawrence Kesteloot

package astutil

import (
	"go/ast"
	"go/token"
	"strings"
)

func pathMatchesPkgName(path, pkgName string) bool {
	// The path will be quoted, so remove that first.
	if len(path) < 2 {
		return false
	}
	path = path[1:len(path)-1]

	// The path will include several sections. Remove all but the last.
	i := strings.LastIndex(path, "/")
	if i >= 0 {
		path = path[i+1:]
	}

	return path == pkgName
}

// Add an import for pkgName to f if it's used in the tree and not already imported.
func AddImport(f *ast.File, pkgName string) {
	// XXX We don't handle pkgName like "go/ast".

	// See if the package is already imported and/or used.
	packageImported := false
	packageUsed := false

	ast.Inspect(f, func(n ast.Node) bool {
		switch e := n.(type) {
		case *ast.SelectorExpr:
			i, ok := e.X.(*ast.Ident)
			if ok && i.Name == pkgName {
				packageUsed = true
			}
		case *ast.ImportSpec:
			if e.Path.Kind == token.STRING && pathMatchesPkgName(e.Path.Value, pkgName) {
				packageImported = true
			}
		}

		// Recurse.
		return true
	})

	if packageUsed && !packageImported {
		importSpec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind: token.STRING,
				Value: "\"" + pkgName + "\"",
			},
		}
		// This could be more elegant. We could look for an existing import list
		// and add it, sorted.
		f.Decls = append([]ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					importSpec,
				},
			},
		}, f.Decls...)
		f.Imports = append(f.Imports, importSpec)
	}
}

// Copyright 2013 Lawrence Kesteloot

package astutil

import (
	"go/ast"
)

// Inspector that sets positions of n to 0.
func stripOutPosFunc(n ast.Node) bool {
	switch e := n.(type) {
	case *ast.Comment:
		e.Slash = 0
	case *ast.FieldList:
		e.Opening = 0
		e.Closing = 0
	case *ast.BadExpr:
		e.From = 0
		e.To = 0
	case *ast.Ident:
		e.NamePos = 0
	case *ast.Ellipsis:
		e.Ellipsis = 0
	case *ast.BasicLit:
		e.ValuePos = 0
	case *ast.CompositeLit:
		e.Lbrace = 0
		e.Rbrace = 0
	case *ast.ParenExpr:
		e.Lparen = 0
		e.Rparen = 0
	case *ast.IndexExpr:
		e.Lbrack = 0
		e.Rbrack = 0
	case *ast.SliceExpr:
		e.Lbrack = 0
		e.Rbrack = 0
	case *ast.CallExpr:
		e.Lparen = 0
		e.Ellipsis = 0
		e.Rparen = 0
	case *ast.StarExpr:
		e.Star = 0
	case *ast.UnaryExpr:
		e.OpPos = 0
	case *ast.BinaryExpr:
		e.OpPos = 0
	case *ast.KeyValueExpr:
		e.Colon = 0
	case *ast.ArrayType:
		e.Lbrack = 0
	case *ast.StructType:
		e.Struct = 0
	case *ast.FuncType:
		e.Func = 0
	case *ast.InterfaceType:
		e.Interface = 0
	case *ast.MapType:
		e.Map = 0
	case *ast.ChanType:
		e.Begin = 0
	case *ast.BadStmt:
		e.From = 0
		e.To = 0
	case *ast.EmptyStmt:
		e.Semicolon = 0
	case *ast.LabeledStmt:
		e.Colon = 0
	case *ast.SendStmt:
		e.Arrow = 0
	case *ast.IncDecStmt:
		e.TokPos = 0
	case *ast.AssignStmt:
		e.TokPos = 0
	case *ast.GoStmt:
		e.Go = 0
	case *ast.DeferStmt:
		e.Defer = 0
	case *ast.ReturnStmt:
		e.Return = 0
	case *ast.BranchStmt:
		e.TokPos = 0
	case *ast.BlockStmt:
		e.Lbrace = 0
		e.Rbrace = 0
	case *ast.IfStmt:
		e.If = 0
	case *ast.CaseClause:
		e.Case = 0
		e.Colon = 0
	case *ast.SwitchStmt:
		e.Switch = 0
	case *ast.TypeSwitchStmt:
		e.Switch = 0
	case *ast.CommClause:
		e.Case = 0
		e.Colon = 0
	case *ast.SelectStmt:
		e.Select = 0
	case *ast.ForStmt:
		e.For = 0
	case *ast.RangeStmt:
		e.For = 0
		e.TokPos = 0
	case *ast.ImportSpec:
		e.EndPos = 0
	case *ast.BadDecl:
		e.From = 0
		e.To = 0
	case *ast.GenDecl:
		e.TokPos = 0
		// We can't set these to 0 because that would cause lists of declarations
		// (like multiple imports) to be reduced to a single import. I don't know why
		// they look at the parens position for this and not whether len(Specs) > 1,
		// but we set them to 1 (if they're already non-zero) so that the position
		// doesn't affect anything when printing.
		if e.Lparen > 0 {
			e.Lparen = 1
			e.Rparen = 1
		}
	case *ast.File:
		e.Package = 0
	}

	return true
}

// Sets all positions in the node tree to 0 (undefined).
func StripOutPos(n ast.Node) {
	ast.Inspect(n, stripOutPosFunc)
}

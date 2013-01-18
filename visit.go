// Copyright 2013 Lawrence Kesteloot

// Modification of the original walk.go file from the standard library,
// modified to pass around pointers to nodes.

package astutil

import (
	"fmt"
	"go/ast"
)

type Visitor interface {
	ProcessNode(node ast.Node)
	ProcessIdent(ident **ast.Ident)
	ProcessExpr(expr *ast.Expr)
	ProcessStmt(stmt *ast.Stmt)
	ProcessDecl(decl *ast.Decl)
}

func walkIdentList(v Visitor, list []*ast.Ident) {
	for i, _ := range list {
		visitIdent(v, &list[i])
	}
}

func walkExprList(v Visitor, list []ast.Expr) {
	for i, _ := range list {
		VisitExpr(v, &list[i])
	}
}

func walkStmtList(v Visitor, list []ast.Stmt) {
	for i, _ := range list {
		VisitStmt(v, &list[i])
	}
}

func walkDeclList(v Visitor, list []ast.Decl) {
	for i, _ := range list {
		VisitDecl(v, &list[i])
	}
}

func visitBlockStmt(v Visitor, stmt *ast.BlockStmt) {
	v.ProcessNode(stmt)

	walkStmtList(v, stmt.List)
}

func visitArrayType(v Visitor, a *ast.ArrayType) {
	v.ProcessNode(a)

	if a.Len != nil {
		VisitExpr(v, &a.Len)
	}
	VisitExpr(v, &a.Elt)
}

func visitIdent(v Visitor, ident **ast.Ident) {
	v.ProcessIdent(ident)
}

func VisitExpr(v Visitor, expr *ast.Expr) {
	v.ProcessExpr(expr)

	switch n := (*expr).(type) {
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		if n.Elt != nil {
			VisitExpr(v, &n.Elt)
		}

	case *ast.FuncLit:
		VisitNode(v, n.Type)
		visitBlockStmt(v, n.Body)

	case *ast.CompositeLit:
		if n.Type != nil {
			VisitExpr(v, &n.Type)
		}
		walkExprList(v, n.Elts)

	case *ast.ParenExpr:
		VisitExpr(v, &n.X)

	case *ast.SelectorExpr:
		VisitExpr(v, &n.X)
		visitIdent(v, &n.Sel)

	case *ast.IndexExpr:
		VisitExpr(v, &n.X)
		VisitExpr(v, &n.Index)

	case *ast.SliceExpr:
		VisitExpr(v, &n.X)
		if n.Low != nil {
			VisitExpr(v, &n.Low)
		}
		if n.High != nil {
			VisitExpr(v, &n.High)
		}

	case *ast.TypeAssertExpr:
		VisitExpr(v, &n.X)
		if n.Type != nil {
			VisitExpr(v, &n.Type)
		}

	case *ast.CallExpr:
		VisitExpr(v, &n.Fun)
		walkExprList(v, n.Args)

	case *ast.StarExpr:
		VisitExpr(v, &n.X)

	case *ast.UnaryExpr:
		VisitExpr(v, &n.X)

	case *ast.BinaryExpr:
		VisitExpr(v, &n.X)
		VisitExpr(v, &n.Y)

	case *ast.KeyValueExpr:
		VisitExpr(v, &n.Key)
		VisitExpr(v, &n.Value)

	// Types
	case *ast.ArrayType:
		visitArrayType(v, n)

	default:
		fmt.Printf("ast.VisitExpr: unexpected node type %T", n)
		panic("ast.VisitExpr")
	}
}

func VisitStmt(v Visitor, stmt *ast.Stmt) {
	v.ProcessStmt(stmt)

	switch n := (*stmt).(type) {
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		VisitDecl(v, &n.Decl)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		visitIdent(v, &n.Label)
		VisitStmt(v, &n.Stmt)

	case *ast.ExprStmt:
		VisitExpr(v, &n.X)

	case *ast.SendStmt:
		VisitExpr(v, &n.Chan)
		VisitExpr(v, &n.Value)

	case *ast.IncDecStmt:
		VisitExpr(v, &n.X)

	case *ast.AssignStmt:
		walkExprList(v, n.Lhs)
		walkExprList(v, n.Rhs)

	case *ast.GoStmt:
		VisitNode(v, n.Call)

	case *ast.DeferStmt:
		VisitNode(v, n.Call)

	case *ast.ReturnStmt:
		walkExprList(v, n.Results)

	case *ast.BranchStmt:
		if n.Label != nil {
			VisitNode(v, n.Label)
		}

	case *ast.BlockStmt:
		visitBlockStmt(v, n)

	case *ast.IfStmt:
		if n.Init != nil {
			VisitStmt(v, &n.Init)
		}
		VisitExpr(v, &n.Cond)
		visitBlockStmt(v, n.Body)
		if n.Else != nil {
			VisitStmt(v, &n.Else)
		}

	case *ast.CaseClause:
		walkExprList(v, n.List)
		walkStmtList(v, n.Body)

	case *ast.SwitchStmt:
		if n.Init != nil {
			VisitStmt(v, &n.Init)
		}
		if n.Tag != nil {
			VisitExpr(v, &n.Tag)
		}
		visitBlockStmt(v, n.Body)

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			VisitStmt(v, &n.Init)
		}
		VisitStmt(v, &n.Assign)
		visitBlockStmt(v, n.Body)

	case *ast.CommClause:
		if n.Comm != nil {
			VisitStmt(v, &n.Comm)
		}
		walkStmtList(v, n.Body)

	case *ast.SelectStmt:
		visitBlockStmt(v, n.Body)

	case *ast.ForStmt:
		if n.Init != nil {
			VisitStmt(v, &n.Init)
		}
		if n.Cond != nil {
			VisitExpr(v, &n.Cond)
		}
		if n.Post != nil {
			VisitStmt(v, &n.Post)
		}
		visitBlockStmt(v, n.Body)

	case *ast.RangeStmt:
		VisitExpr(v, &n.Key)
		if n.Value != nil {
			VisitExpr(v, &n.Value)
		}
		VisitExpr(v, &n.X)
		visitBlockStmt(v, n.Body)

	default:
		fmt.Printf("ast.VisitStmt: unexpected node type %T", n)
		panic("ast.VisitStmt")
	}
}

func VisitDecl(v Visitor, decl *ast.Decl) {
	v.ProcessDecl(decl)

	switch n := (*decl).(type) {
	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		for _, s := range n.Specs {
			VisitNode(v, s)
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		if n.Recv != nil {
			VisitNode(v, n.Recv)
		}
		visitIdent(v, &n.Name)
		VisitNode(v, n.Type)
		if n.Body != nil {
			visitBlockStmt(v, n.Body)
		}

	default:
		fmt.Printf("ast.VisitDecl: unexpected node type %T", n)
		panic("ast.VisitDecl")
	}
}

func VisitNode(v Visitor, node ast.Node) {
	v.ProcessNode(node)

	switch n := node.(type) {
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	// Comments and fields
	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		for _, c := range n.List {
			VisitNode(v, c)
		}

	case *ast.Field:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		walkIdentList(v, n.Names)
		VisitExpr(v, &n.Type)
		if n.Tag != nil {
			VisitNode(v, n.Tag)
		}
		if n.Comment != nil {
			VisitNode(v, n.Comment)
		}

	case *ast.FieldList:
		for _, f := range n.List {
			VisitNode(v, f)
		}

	// Types
	case *ast.ArrayType:
		visitArrayType(v, n)

	case *ast.StructType:
		VisitNode(v, n.Fields)

	case *ast.FuncType:
		VisitNode(v, n.Params)
		if n.Results != nil {
			VisitNode(v, n.Results)
		}

	case *ast.InterfaceType:
		VisitNode(v, n.Methods)

	case *ast.MapType:
		VisitNode(v, n.Key)
		VisitNode(v, n.Value)

	case *ast.ChanType:
		VisitExpr(v, &n.Value)

	case *ast.ImportSpec:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		if n.Name != nil {
			visitIdent(v, &n.Name)
		}
		VisitNode(v, n.Path)
		if n.Comment != nil {
			VisitNode(v, n.Comment)
		}

	case *ast.ValueSpec:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		walkIdentList(v, n.Names)
		if n.Type != nil {
			VisitExpr(v, &n.Type)
		}
		walkExprList(v, n.Values)
		if n.Comment != nil {
			VisitNode(v, n.Comment)
		}

	case *ast.TypeSpec:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		visitIdent(v, &n.Name)
		VisitExpr(v, &n.Type)
		if n.Comment != nil {
			VisitNode(v, n.Comment)
		}


	// Files and packages
	case *ast.File:
		if n.Doc != nil {
			VisitNode(v, n.Doc)
		}
		visitIdent(v, &n.Name)
		walkDeclList(v, n.Decls)
		for _, g := range n.Comments {
			VisitNode(v, g)
		}
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		for _, f := range n.Files {
			VisitNode(v, f)
		}

	default:
		fmt.Printf("ast.VisitNode: unexpected node type %T", n)
		panic("ast.VisitNode")
	}
}

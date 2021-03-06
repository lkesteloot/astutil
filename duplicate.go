// Copyright 2013 Lawrence Kesteloot

package astutil

import (
	"fmt"
	"go/ast"
)

// Makes a deep copy of the syntax tree rooted at the Node.
func DuplicateNode(x ast.Node) ast.Node {
	if x == nil {
		return nil
	}

	switch y := x.(type) {
	case ast.Expr:
		return DuplicateExpr(y)
	case ast.Stmt:
		return DuplicateStmt(y)
	case ast.Decl:
		return DuplicateDecl(y)
	case ast.Spec:
		return DuplicateSpec(y)
	case *ast.File:
		return DuplicateFile(y)
	case *ast.Package:
		return DuplicatePackage(y)
	}

	panic(fmt.Sprintf("Unknown node %T", x))
}

// Makes a deep copy of the syntax tree rooted at the Expr.
func DuplicateExpr(x ast.Expr) ast.Expr {
	if x == nil {
		return nil
	}

	switch y := x.(type) {
	case *ast.BadExpr:
		return DuplicateBadExpr(y)
	case *ast.Ident:
		return DuplicateIdent(y)
	case *ast.Ellipsis:
		return DuplicateEllipsis(y)
	case *ast.BasicLit:
		return DuplicateBasicLit(y)
	case *ast.FuncLit:
		return DuplicateFuncLit(y)
	case *ast.CompositeLit:
		return DuplicateCompositeLit(y)
	case *ast.ParenExpr:
		return DuplicateParenExpr(y)
	case *ast.SelectorExpr:
		return DuplicateSelectorExpr(y)
	case *ast.IndexExpr:
		return DuplicateIndexExpr(y)
	case *ast.SliceExpr:
		return DuplicateSliceExpr(y)
	case *ast.TypeAssertExpr:
		return DuplicateTypeAssertExpr(y)
	case *ast.CallExpr:
		return DuplicateCallExpr(y)
	case *ast.StarExpr:
		return DuplicateStarExpr(y)
	case *ast.UnaryExpr:
		return DuplicateUnaryExpr(y)
	case *ast.BinaryExpr:
		return DuplicateBinaryExpr(y)
	case *ast.KeyValueExpr:
		return DuplicateKeyValueExpr(y)
	case *ast.ArrayType:
		return DuplicateArrayType(y)
	case *ast.StructType:
		return DuplicateStructType(y)
	case *ast.FuncType:
		return DuplicateFuncType(y)
	case *ast.InterfaceType:
		return DuplicateInterfaceType(y)
	case *ast.MapType:
		return DuplicateMapType(y)
	case *ast.ChanType:
		return DuplicateChanType(y)
	}

	panic(fmt.Sprintf("Unknown expr %T", x))
}

// Makes a deep copy of the slice of syntax trees rooted at the Exprs.
func DuplicateExprSlice(x []ast.Expr) []ast.Expr {
	if x == nil {
		return nil
	}

	y := []ast.Expr{}

	for _, z := range x {
		y = append(y, DuplicateExpr(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the Stmt.
func DuplicateStmt(x ast.Stmt) ast.Stmt {
	if x == nil {
		return nil
	}

	switch y := x.(type) {
	case *ast.BadStmt:
		return DuplicateBadStmt(y)
	case *ast.DeclStmt:
		return DuplicateDeclStmt(y)
	case *ast.EmptyStmt:
		return DuplicateEmptyStmt(y)
	case *ast.LabeledStmt:
		return DuplicateLabeledStmt(y)
	case *ast.ExprStmt:
		return DuplicateExprStmt(y)
	case *ast.SendStmt:
		return DuplicateSendStmt(y)
	case *ast.IncDecStmt:
		return DuplicateIncDecStmt(y)
	case *ast.AssignStmt:
		return DuplicateAssignStmt(y)
	case *ast.GoStmt:
		return DuplicateGoStmt(y)
	case *ast.DeferStmt:
		return DuplicateDeferStmt(y)
	case *ast.ReturnStmt:
		return DuplicateReturnStmt(y)
	case *ast.BranchStmt:
		return DuplicateBranchStmt(y)
	case *ast.BlockStmt:
		return DuplicateBlockStmt(y)
	case *ast.IfStmt:
		return DuplicateIfStmt(y)
	case *ast.CaseClause:
		return DuplicateCaseClause(y)
	case *ast.SwitchStmt:
		return DuplicateSwitchStmt(y)
	case *ast.TypeSwitchStmt:
		return DuplicateTypeSwitchStmt(y)
	case *ast.CommClause:
		return DuplicateCommClause(y)
	case *ast.SelectStmt:
		return DuplicateSelectStmt(y)
	case *ast.ForStmt:
		return DuplicateForStmt(y)
	case *ast.RangeStmt:
		return DuplicateRangeStmt(y)
	}

	panic(fmt.Sprintf("Unknown stmt %T", x))
}

// Makes a deep copy of the slice of syntax trees rooted at the Stmts.
func DuplicateStmtSlice(x []ast.Stmt) []ast.Stmt {
	if x == nil {
		return nil
	}

	y := []ast.Stmt{}

	for _, z := range x {
		y = append(y, DuplicateStmt(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the Decl.
func DuplicateDecl(x ast.Decl) ast.Decl {
	if x == nil {
		return nil
	}

	switch y := x.(type) {
	case *ast.BadDecl:
		return DuplicateBadDecl(y)
	case *ast.GenDecl:
		return DuplicateGenDecl(y)
	case *ast.FuncDecl:
		return DuplicateFuncDecl(y)
	}

	panic(fmt.Sprintf("Unknown decl %T", x))
}

// Makes a deep copy of the slice of syntax trees rooted at the Decls.
func DuplicateDeclSlice(x []ast.Decl) []ast.Decl {
	if x == nil {
		return nil
	}

	y := []ast.Decl{}

	for _, z := range x {
		y = append(y, DuplicateDecl(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the Spec.
func DuplicateSpec(x ast.Spec) ast.Spec {
	if x == nil {
		return nil
	}

	switch y := x.(type) {
	case *ast.ImportSpec:
		return DuplicateImportSpec(y)
	case *ast.ValueSpec:
		return DuplicateValueSpec(y)
	case *ast.TypeSpec:
		return DuplicateTypeSpec(y)
	}

	panic(fmt.Sprintf("Unknown spec %T", x))
}

// Makes a deep copy of the slice of syntax trees rooted at the Specs.
func DuplicateSpecSlice(x []ast.Spec) []ast.Spec {
	if x == nil {
		return nil
	}

	y := []ast.Spec{}

	for _, z := range x {
		y = append(y, DuplicateSpec(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the Comment.
func DuplicateComment(x *ast.Comment) *ast.Comment {
	if x == nil {
		return nil
	}

	return &ast.Comment{
		Slash: x.Slash,
		Text: x.Text,
	}
}

// Makes a deep copy of the slice of syntax trees rooted at the Comments.
func DuplicateCommentSlice(x []*ast.Comment) []*ast.Comment {
	if x == nil {
		return nil
	}

	y := []*ast.Comment{}

	for _, z := range x {
		y = append(y, DuplicateComment(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the CommentGroup.
func DuplicateCommentGroup(x *ast.CommentGroup) *ast.CommentGroup {
	if x == nil {
		return nil
	}

	return &ast.CommentGroup{
		List: DuplicateCommentSlice(x.List),
	}
}

// Makes a deep copy of the syntax tree rooted at the Field.
func DuplicateField(x *ast.Field) *ast.Field {
	if x == nil {
		return nil
	}

	return &ast.Field {
		Doc: DuplicateCommentGroup(x.Doc),
		Names: DuplicateIdentSlice(x.Names),
		Type: DuplicateExpr(x.Type),
		Tag: DuplicateBasicLit(x.Tag),
		Comment: DuplicateCommentGroup(x.Comment),
	}
}

// Makes a deep copy of the slice of syntax trees rooted at the Fields.
func DuplicateFieldSlice(x []*ast.Field) []*ast.Field {
	if x == nil {
		return nil
	}

	y := []*ast.Field{}

	for _, z := range x {
		y = append(y, DuplicateField(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the FieldList.
func DuplicateFieldList(x *ast.FieldList) *ast.FieldList {
	if x == nil {
		return nil
	}

	return &ast.FieldList{
		Opening: x.Opening,
		List: DuplicateFieldSlice(x.List),
		Closing: x.Closing,
	}
}

// Makes a deep copy of the syntax tree rooted at the BadExpr.
func DuplicateBadExpr(x *ast.BadExpr) *ast.BadExpr {
	if x == nil {
		return nil
	}

	return &ast.BadExpr{
		From: x.From,
		To: x.To,
	}
}

// Makes a deep copy of the syntax tree rooted at the Ident. The Object
// pointed to by the Obj field is not duplicated. Its fields point to the tree
// being copied.
func DuplicateIdent(x *ast.Ident) *ast.Ident {
	if x == nil {
		return nil
	}

	return &ast.Ident {
		NamePos: x.NamePos,
		Name: x.Name,
		Obj: x.Obj,
	}
}

// Makes a deep copy of the slice of syntax trees rooted at the Idents.
func DuplicateIdentSlice(x []*ast.Ident) []*ast.Ident {
	if x == nil {
		return nil
	}

	y := []*ast.Ident{}

	for _, z := range x {
		y = append(y, DuplicateIdent(z))
	}

	return y
}

// Makes a deep copy of the syntax tree rooted at the Ellipsis.
func DuplicateEllipsis(x *ast.Ellipsis) *ast.Ellipsis {
	if x == nil {
		return nil
	}

	return &ast.Ellipsis{
		Ellipsis: x.Ellipsis,
		Elt: DuplicateExpr(x.Elt),
	}
}

// Makes a deep copy of the syntax tree rooted at the BasicLit.
func DuplicateBasicLit(x *ast.BasicLit) *ast.BasicLit {
	if x == nil {
		return nil
	}

	return &ast.BasicLit{
		ValuePos: x.ValuePos,
		Kind: x.Kind,
		Value: x.Value,
	}
}

// Makes a deep copy of the syntax tree rooted at the FuncLit.
func DuplicateFuncLit(x *ast.FuncLit) *ast.FuncLit {
	if x == nil {
		return nil
	}

	return &ast.FuncLit{
		Type: DuplicateFuncType(x.Type),
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the CompositeLit.
func DuplicateCompositeLit(x *ast.CompositeLit) *ast.CompositeLit {
	if x == nil {
		return nil
	}

	return &ast.CompositeLit{
		Type: DuplicateExpr(x.Type),
		Lbrace: x.Lbrace,
		Elts: DuplicateExprSlice(x.Elts),
		Rbrace: x.Rbrace,
	}
}

// Makes a deep copy of the syntax tree rooted at the ParenExpr.
func DuplicateParenExpr(x *ast.ParenExpr) *ast.ParenExpr {
	if x == nil {
		return nil
	}

	return &ast.ParenExpr{
		Lparen: x.Lparen,
		X: DuplicateExpr(x.X),
		Rparen: x.Rparen,
	}
}

// Makes a deep copy of the syntax tree rooted at the SelectorExpr.
func DuplicateSelectorExpr(x *ast.SelectorExpr) *ast.SelectorExpr {
	if x == nil {
		return nil
	}

	return &ast.SelectorExpr{
		X: DuplicateExpr(x.X),
		Sel: DuplicateIdent(x.Sel),
	}
}

// Makes a deep copy of the syntax tree rooted at the IndexExpr.
func DuplicateIndexExpr(x *ast.IndexExpr) *ast.IndexExpr {
	if x == nil {
		return nil
	}

	return &ast.IndexExpr{
		X: DuplicateExpr(x.X),
		Lbrack: x.Lbrack,
		Index: DuplicateExpr(x.Index),
		Rbrack: x.Rbrack,
	}
}

// Makes a deep copy of the syntax tree rooted at the SliceExpr.
func DuplicateSliceExpr(x *ast.SliceExpr) *ast.SliceExpr {
	if x == nil {
		return nil
	}

	return &ast.SliceExpr{
		X: DuplicateExpr(x.X),
		Lbrack: x.Lbrack,
		Low: DuplicateExpr(x.Low),
		High: DuplicateExpr(x.High),
		Rbrack: x.Rbrack,
	}
}

// Makes a deep copy of the syntax tree rooted at the TypeAssertExpr.
func DuplicateTypeAssertExpr(x *ast.TypeAssertExpr) *ast.TypeAssertExpr {
	if x == nil {
		return nil
	}

	return &ast.TypeAssertExpr{
		X: DuplicateExpr(x.X),
		Type: DuplicateExpr(x.Type),
	}
}

// Makes a deep copy of the syntax tree rooted at the CallExpr.
func DuplicateCallExpr(x *ast.CallExpr) *ast.CallExpr {
	if x == nil {
		return nil
	}

	return &ast.CallExpr{
		Fun: DuplicateExpr(x.Fun),
		Lparen: x.Lparen,
		Args: DuplicateExprSlice(x.Args),
		Ellipsis: x.Ellipsis,
		Rparen: x.Rparen,
	}
}

// Makes a deep copy of the syntax tree rooted at the StarExpr.
func DuplicateStarExpr(x *ast.StarExpr) *ast.StarExpr {
	if x == nil {
		return nil
	}

	return &ast.StarExpr{
		Star: x.Star,
		X: DuplicateExpr(x.X),
	}
}

// Makes a deep copy of the syntax tree rooted at the UnaryExpr.
func DuplicateUnaryExpr(x *ast.UnaryExpr) *ast.UnaryExpr {
	if x == nil {
		return nil
	}

	return &ast.UnaryExpr{
		OpPos: x.OpPos,
		Op: x.Op,
		X: DuplicateExpr(x.X),
	}
}

// Makes a deep copy of the syntax tree rooted at the BinaryExpr.
func DuplicateBinaryExpr(x *ast.BinaryExpr) *ast.BinaryExpr {
	if x == nil {
		return nil
	}

	return &ast.BinaryExpr{
		X: DuplicateExpr(x.X),
		OpPos: x.OpPos,
		Op: x.Op,
		Y: DuplicateExpr(x.Y),
	}
}

// Makes a deep copy of the syntax tree rooted at the KeyValueExpr.
func DuplicateKeyValueExpr(x *ast.KeyValueExpr) *ast.KeyValueExpr {
	if x == nil {
		return nil
	}

	return &ast.KeyValueExpr{
		Key: DuplicateExpr(x.Key),
		Colon: x.Colon,
		Value: DuplicateExpr(x.Value),
	}
}

// Makes a deep copy of the syntax tree rooted at the ArrayType.
func DuplicateArrayType(x *ast.ArrayType) *ast.ArrayType {
	if x == nil {
		return nil
	}

	return &ast.ArrayType{
		Lbrack: x.Lbrack,
		Len: DuplicateExpr(x.Len),
		Elt: DuplicateExpr(x.Elt),
	}
}

// Makes a deep copy of the syntax tree rooted at the StructType.
func DuplicateStructType(x *ast.StructType) *ast.StructType {
	if x == nil {
		return nil
	}

	return &ast.StructType{
		Struct: x.Struct,
		Fields: DuplicateFieldList(x.Fields),
		Incomplete: x.Incomplete,
	}
}

// Makes a deep copy of the syntax tree rooted at the FuncType.
func DuplicateFuncType(x *ast.FuncType) *ast.FuncType {
	if x == nil {
		return nil
	}

	return &ast.FuncType{
		Func: x.Func,
		Params: DuplicateFieldList(x.Params),
		Results: DuplicateFieldList(x.Results),
	}
}

// Makes a deep copy of the syntax tree rooted at the InterfaceType.
func DuplicateInterfaceType(x *ast.InterfaceType) *ast.InterfaceType {
	if x == nil {
		return nil
	}

	return &ast.InterfaceType{
		Interface: x.Interface,
		Methods: DuplicateFieldList(x.Methods),
		Incomplete: x.Incomplete,
	}
}

// Makes a deep copy of the syntax tree rooted at the MapType.
func DuplicateMapType(x *ast.MapType) *ast.MapType {
	if x == nil {
		return nil
	}

	return &ast.MapType{
		Map: x.Map,
		Key: DuplicateExpr(x.Key),
		Value: DuplicateExpr(x.Value),
	}
}

// Makes a deep copy of the syntax tree rooted at the ChanType.
func DuplicateChanType(x *ast.ChanType) *ast.ChanType {
	if x == nil {
		return nil
	}

	return &ast.ChanType{
		Begin: x.Begin,
		Dir: x.Dir,
		Value: DuplicateExpr(x.Value),
	}
}

// Makes a deep copy of the syntax tree rooted at the BadStmt.
func DuplicateBadStmt(x *ast.BadStmt) *ast.BadStmt {
	if x == nil {
		return nil
	}

	return &ast.BadStmt{
		From: x.From,
		To: x.To,
	}
}

// Makes a deep copy of the syntax tree rooted at the DeclStmt.
func DuplicateDeclStmt(x *ast.DeclStmt) *ast.DeclStmt {
	if x == nil {
		return nil
	}

	return &ast.DeclStmt{
		Decl: DuplicateDecl(x.Decl),
	}
}

// Makes a deep copy of the syntax tree rooted at the EmptyStmt.
func DuplicateEmptyStmt(x *ast.EmptyStmt) *ast.EmptyStmt {
	if x == nil {
		return nil
	}

	return &ast.EmptyStmt{
		Semicolon: x.Semicolon,
	}
}

// Makes a deep copy of the syntax tree rooted at the LabeledStmt.
func DuplicateLabeledStmt(x *ast.LabeledStmt) *ast.LabeledStmt {
	if x == nil {
		return nil
	}

	return &ast.LabeledStmt{
		Label: DuplicateIdent(x.Label),
		Colon: x.Colon,
		Stmt: DuplicateStmt(x.Stmt),
	}
}

// Makes a deep copy of the syntax tree rooted at the ExprStmt.
func DuplicateExprStmt(x *ast.ExprStmt) *ast.ExprStmt {
	if x == nil {
		return nil
	}

	return &ast.ExprStmt{
		X: DuplicateExpr(x.X),
	}
}

// Makes a deep copy of the syntax tree rooted at the SendStmt.
func DuplicateSendStmt(x *ast.SendStmt) *ast.SendStmt {
	if x == nil {
		return nil
	}

	return &ast.SendStmt{
		Chan: DuplicateExpr(x.Chan),
		Arrow: x.Arrow,
		Value: DuplicateExpr(x.Value),
	}
}

// Makes a deep copy of the syntax tree rooted at the IncDecStmt.
func DuplicateIncDecStmt(x *ast.IncDecStmt) *ast.IncDecStmt {
	if x == nil {
		return nil
	}

	return &ast.IncDecStmt{
		X: DuplicateExpr(x.X),
		TokPos: x.TokPos,
		Tok: x.Tok,
	}
}

// Makes a deep copy of the syntax tree rooted at the AssignStmt.
func DuplicateAssignStmt(x *ast.AssignStmt) *ast.AssignStmt {
	if x == nil {
		return nil
	}

	return &ast.AssignStmt{
		Lhs: DuplicateExprSlice(x.Lhs),
		TokPos: x.TokPos,
		Tok: x.Tok,
		Rhs: DuplicateExprSlice(x.Rhs),
	}
}

// Makes a deep copy of the syntax tree rooted at the GoStmt.
func DuplicateGoStmt(x *ast.GoStmt) *ast.GoStmt {
	if x == nil {
		return nil
	}

	return &ast.GoStmt{
		Go: x.Go,
		Call: DuplicateCallExpr(x.Call),
	}
}

// Makes a deep copy of the syntax tree rooted at the DeferStmt.
func DuplicateDeferStmt(x *ast.DeferStmt) *ast.DeferStmt {
	if x == nil {
		return nil
	}

	return &ast.DeferStmt{
		Defer: x.Defer,
		Call: DuplicateCallExpr(x.Call),
	}
}

// Makes a deep copy of the syntax tree rooted at the ReturnStmt.
func DuplicateReturnStmt(x *ast.ReturnStmt) *ast.ReturnStmt {
	if x == nil {
		return nil
	}

	return &ast.ReturnStmt{
		Return: x.Return,
		Results: DuplicateExprSlice(x.Results),
	}
}

// Makes a deep copy of the syntax tree rooted at the BranchStmt.
func DuplicateBranchStmt(x *ast.BranchStmt) *ast.BranchStmt {
	if x == nil {
		return nil
	}

	return &ast.BranchStmt{
		TokPos: x.TokPos,
		Tok: x.Tok,
		Label: DuplicateIdent(x.Label),
	}
}

// Makes a deep copy of the syntax tree rooted at the BlockStmt.
func DuplicateBlockStmt(x *ast.BlockStmt) *ast.BlockStmt {
	if x == nil {
		return nil
	}

	return &ast.BlockStmt{
		Lbrace: x.Lbrace,
		List: DuplicateStmtSlice(x.List),
		Rbrace: x.Rbrace,
	}
}

// Makes a deep copy of the syntax tree rooted at the IfStmt.
func DuplicateIfStmt(x *ast.IfStmt) *ast.IfStmt {
	if x == nil {
		return nil
	}

	return &ast.IfStmt{
		If: x.If,
		Init: DuplicateStmt(x.Init),
		Cond: DuplicateExpr(x.Cond),
		Body: DuplicateBlockStmt(x.Body),
		Else: DuplicateStmt(x.Else),
	}
}

// Makes a deep copy of the syntax tree rooted at the CaseClause.
func DuplicateCaseClause(x *ast.CaseClause) *ast.CaseClause {
	if x == nil {
		return nil
	}

	return &ast.CaseClause{
		Case: x.Case,
		List: DuplicateExprSlice(x.List),
		Colon: x.Colon,
		Body: DuplicateStmtSlice(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the SwitchStmt.
func DuplicateSwitchStmt(x *ast.SwitchStmt) *ast.SwitchStmt {
	if x == nil {
		return nil
	}

	return &ast.SwitchStmt{
		Switch: x.Switch,
		Init: DuplicateStmt(x.Init),
		Tag: DuplicateExpr(x.Tag),
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the TypeSwitchStmt.
func DuplicateTypeSwitchStmt(x *ast.TypeSwitchStmt) *ast.TypeSwitchStmt {
	if x == nil {
		return nil
	}

	return &ast.TypeSwitchStmt{
		Switch: x.Switch,
		Init: DuplicateStmt(x.Init),
		Assign: DuplicateStmt(x.Assign),
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the CommClause.
func DuplicateCommClause(x *ast.CommClause) *ast.CommClause {
	if x == nil {
		return nil
	}

	return &ast.CommClause{
		Case: x.Case,
		Comm: DuplicateStmt(x.Comm),
		Colon: x.Colon,
		Body: DuplicateStmtSlice(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the SelectStmt.
func DuplicateSelectStmt(x *ast.SelectStmt) *ast.SelectStmt {
	if x == nil {
		return nil
	}

	return &ast.SelectStmt{
		Select: x.Select,
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the ForStmt.
func DuplicateForStmt(x *ast.ForStmt) *ast.ForStmt {
	if x == nil {
		return nil
	}

	return &ast.ForStmt{
		For: x.For,
		Init: DuplicateStmt(x.Init),
		Cond: DuplicateExpr(x.Cond),
		Post: DuplicateStmt(x.Post),
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the RangeStmt.
func DuplicateRangeStmt(x *ast.RangeStmt) *ast.RangeStmt {
	if x == nil {
		return nil
	}

	return &ast.RangeStmt{
		For: x.For,
		Key: DuplicateExpr(x.Key),
		Value: DuplicateExpr(x.Value),
		TokPos: x.TokPos,
		Tok: x.Tok,
		X: DuplicateExpr(x.X),
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the ImportSpec.
func DuplicateImportSpec(x *ast.ImportSpec) *ast.ImportSpec {
	if x == nil {
		return nil
	}

	return &ast.ImportSpec{
		Doc: DuplicateCommentGroup(x.Doc),
		Name: DuplicateIdent(x.Name),
		Path: DuplicateBasicLit(x.Path),
		Comment: DuplicateCommentGroup(x.Comment),
		EndPos: x.EndPos,
	}
}

// Makes a deep copy of the syntax tree rooted at the ValueSpec.
func DuplicateValueSpec(x *ast.ValueSpec) *ast.ValueSpec {
	if x == nil {
		return nil
	}

	return &ast.ValueSpec{
		Doc: DuplicateCommentGroup(x.Doc),
		Names: DuplicateIdentSlice(x.Names),
		Type: DuplicateExpr(x.Type),
		Values: DuplicateExprSlice(x.Values),
		Comment: DuplicateCommentGroup(x.Comment),
	}
}

// Makes a deep copy of the syntax tree rooted at the TypeSpec.
func DuplicateTypeSpec(x *ast.TypeSpec) *ast.TypeSpec {
	if x == nil {
		return nil
	}

	return &ast.TypeSpec{
		Doc: DuplicateCommentGroup(x.Doc),
		Name: DuplicateIdent(x.Name),
		Type: DuplicateExpr(x.Type),
		Comment: DuplicateCommentGroup(x.Comment),

	}
}

// Makes a deep copy of the syntax tree rooted at the BadDecl.
func DuplicateBadDecl(x *ast.BadDecl) *ast.BadDecl {
	if x == nil {
		return nil
	}

	return &ast.BadDecl{
		From: x.From,
		To: x.To,
	}
}

// Makes a deep copy of the syntax tree rooted at the GenDecl.
func DuplicateGenDecl(x *ast.GenDecl) *ast.GenDecl {
	if x == nil {
		return nil
	}

	return &ast.GenDecl{
		Doc: DuplicateCommentGroup(x.Doc),
		TokPos: x.TokPos,
		Tok: x.Tok,
		Lparen: x.Lparen,
		Specs: DuplicateSpecSlice(x.Specs),
		Rparen: x.Rparen,
	}
}

// Makes a deep copy of the syntax tree rooted at the FuncDecl.
func DuplicateFuncDecl(x *ast.FuncDecl) *ast.FuncDecl {
	if x == nil {
		return nil
	}

	return &ast.FuncDecl{
		Doc: DuplicateCommentGroup(x.Doc),
		Recv: DuplicateFieldList(x.Recv),
		Name: DuplicateIdent(x.Name),
		Type: DuplicateFuncType(x.Type),
		Body: DuplicateBlockStmt(x.Body),
	}
}

// Makes a deep copy of the syntax tree rooted at the File. The Scope field is
// not duplicated; it points to the old tree's object. The Imports, Unresolved,
// and Comments fields are nil.
func DuplicateFile(x *ast.File) *ast.File {
	if x == nil {
		return nil
	}

	// XXX Some of these should be references to items within the new tree.
	// Perhaps I should set up the links from scratch.
	return &ast.File{
		Doc: DuplicateCommentGroup(x.Doc),
		Package: x.Package,
		Name: DuplicateIdent(x.Name),
		Decls: DuplicateDeclSlice(x.Decls),
		Scope: x.Scope,
		Imports: nil,
		Unresolved: nil,
		Comments: nil,
	}
}

// Makes a deep copy of the syntax tree rooted at the Package. The Scope field
// is not duplicated; it points to the old tree's object. The Imports and Files
// fields are nil.
func DuplicatePackage(x *ast.Package) *ast.Package {
	if x == nil {
		return nil
	}

	// XXX Should make internal references.
	return &ast.Package{
		Name: x.Name,
		Scope: x.Scope,
		Imports: nil,
		Files: nil,
	}
}

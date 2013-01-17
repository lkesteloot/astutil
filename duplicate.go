// Copyright 2013 Lawrence Kesteloot

package astutil

import (
	"go/ast"
)

func DuplicateNode(x ast.Node) ast.Node {
	if x == nil {
		return nil
	}

	// XXX Switch on all nodes (interfaces?) and call specific function.
	return nil
}

func DuplicateExpr(x ast.Expr) ast.Expr {
	if x == nil {
		return nil
	}

	// XXX Switch on all nodes (interfaces?) and call specific function.
	return nil
}

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

func DuplicateStmt(x ast.Stmt) ast.Stmt {
	if x == nil {
		return nil
	}

	// XXX Switch on all nodes (interfaces?) and call specific function.
	return nil
}

func DuplicateDecl(x ast.Decl) ast.Decl {
	if x == nil {
		return nil
	}

	// XXX Switch on all nodes (interfaces?) and call specific function.
	return nil
}

func DuplicateComment(x *ast.Comment) *ast.Comment {
	if x == nil {
		return nil
	}

	return &ast.Comment{
		Slash: x.Slash,
		Text: x.Text,
	}
}

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

func DuplicateCommentGroup(x *ast.CommentGroup) *ast.CommentGroup {
	if x == nil {
		return nil
	}

	return &ast.CommentGroup{
		List: DuplicateCommentSlice(x.List),
	}
}

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

func DuplicateBadExpr(x *ast.BadExpr) *ast.BadExpr {
	if x == nil {
		return nil
	}

	return &ast.BadExpr{
		From: x.From,
		To: x.To,
	}
}

// The Object pointed to by the Obj field is not duplicated. Its fields point to the
// tree being copied.
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

func DuplicateEllipsis(x *ast.Ellipsis) *ast.Ellipsis {
	if x == nil {
		return nil
	}

	return &ast.Ellipsis{
		Ellipsis: x.Ellipsis,
		Elt: DuplicateExpr(x.Elt),
	}
}

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

func DuplicateFuncLit(x *ast.FuncLit) *ast.FuncLit {
	if x == nil {
		return nil
	}

	return &ast.FuncLit{
		Type: DuplicateFuncType(x.Type),
		Body: DuplicateBlockStmt(x.Body),
	}
}

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

func DuplicateSelectorExpr(x *ast.SelectorExpr) *ast.SelectorExpr {
	if x == nil {
		return nil
	}

	return &ast.SelectorExpr{
		X: DuplicateExpr(x.X),
		Sel: DuplicateIdent(x.Sel),
	}
}

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

func DuplicateTypeAssertExpr(x *ast.TypeAssertExpr) *ast.TypeAssertExpr {
	if x == nil {
		return nil
	}

	return &ast.TypeAssertExpr{
		X: DuplicateExpr(x.X),
		Type: DuplicateExpr(x.Type),
	}
}

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

func DuplicateStarExpr(x *ast.StarExpr) *ast.StarExpr {
	if x == nil {
		return nil
	}

	return &ast.StarExpr{
		Star: x.Star,
		X: DuplicateExpr(x.X),
	}
}

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

/*
	// A BadStmt node is a placeholder for statements containing
	// syntax errors for which no correct statement nodes can be
	// created.
	//
	BadStmt struct {
		From, To token.Pos // position range of bad statement
	}

	// A DeclStmt node represents a declaration in a statement list.
	DeclStmt struct {
		Decl Decl
	}

	// An EmptyStmt node represents an empty statement.
	// The "position" of the empty statement is the position
	// of the immediately preceding semicolon.
	//
	EmptyStmt struct {
		Semicolon token.Pos // position of preceding ";"
	}

	// A LabeledStmt node represents a labeled statement.
	LabeledStmt struct {
		Label *Ident
		Colon token.Pos // position of ":"
		Stmt  Stmt
	}

	// An ExprStmt node represents a (stand-alone) expression
	// in a statement list.
	//
	ExprStmt struct {
		X Expr // expression
	}

	// A SendStmt node represents a send statement.
	SendStmt struct {
		Chan  Expr
		Arrow token.Pos // position of "<-"
		Value Expr
	}

	// An IncDecStmt node represents an increment or decrement statement.
	IncDecStmt struct {
		X      Expr
		TokPos token.Pos   // position of Tok
		Tok    token.Token // INC or DEC
	}

	// An AssignStmt node represents an assignment or
	// a short variable declaration.
	//
	AssignStmt struct {
		Lhs    []Expr
		TokPos token.Pos   // position of Tok
		Tok    token.Token // assignment token, DEFINE
		Rhs    []Expr
	}

	// A GoStmt node represents a go statement.
	GoStmt struct {
		Go   token.Pos // position of "go" keyword
		Call *CallExpr
	}

	// A DeferStmt node represents a defer statement.
	DeferStmt struct {
		Defer token.Pos // position of "defer" keyword
		Call  *CallExpr
	}

	// A ReturnStmt node represents a return statement.
	ReturnStmt struct {
		Return  token.Pos // position of "return" keyword
		Results []Expr    // result expressions; or nil
	}

	// A BranchStmt node represents a break, continue, goto,
	// or fallthrough statement.
	//
	BranchStmt struct {
		TokPos token.Pos   // position of Tok
		Tok    token.Token // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
		Label  *Ident      // label name; or nil
	}

	// A BlockStmt node represents a braced statement list.
	BlockStmt struct {
		Lbrace token.Pos // position of "{"
		List   []Stmt
		Rbrace token.Pos // position of "}"
	}

	// An IfStmt node represents an if statement.
	IfStmt struct {
		If   token.Pos // position of "if" keyword
		Init Stmt      // initialization statement; or nil
		Cond Expr      // condition
		Body *BlockStmt
		Else Stmt // else branch; or nil
	}

	// A CaseClause represents a case of an expression or type switch statement.
	CaseClause struct {
		Case  token.Pos // position of "case" or "default" keyword
		List  []Expr    // list of expressions or types; nil means default case
		Colon token.Pos // position of ":"
		Body  []Stmt    // statement list; or nil
	}

	// A SwitchStmt node represents an expression switch statement.
	SwitchStmt struct {
		Switch token.Pos  // position of "switch" keyword
		Init   Stmt       // initialization statement; or nil
		Tag    Expr       // tag expression; or nil
		Body   *BlockStmt // CaseClauses only
	}

	// An TypeSwitchStmt node represents a type switch statement.
	TypeSwitchStmt struct {
		Switch token.Pos  // position of "switch" keyword
		Init   Stmt       // initialization statement; or nil
		Assign Stmt       // x := y.(type) or y.(type)
		Body   *BlockStmt // CaseClauses only
	}

	// A CommClause node represents a case of a select statement.
	CommClause struct {
		Case  token.Pos // position of "case" or "default" keyword
		Comm  Stmt      // send or receive statement; nil means default case
		Colon token.Pos // position of ":"
		Body  []Stmt    // statement list; or nil
	}

	// An SelectStmt node represents a select statement.
	SelectStmt struct {
		Select token.Pos  // position of "select" keyword
		Body   *BlockStmt // CommClauses only
	}

	// A ForStmt represents a for statement.
	ForStmt struct {
		For  token.Pos // position of "for" keyword
		Init Stmt      // initialization statement; or nil
		Cond Expr      // condition; or nil
		Post Stmt      // post iteration statement; or nil
		Body *BlockStmt
	}

	// A RangeStmt represents a for statement with a range clause.
	RangeStmt struct {
		For        token.Pos   // position of "for" keyword
		Key, Value Expr        // Value may be nil
		TokPos     token.Pos   // position of Tok
		Tok        token.Token // ASSIGN, DEFINE
		X          Expr        // value to range over
		Body       *BlockStmt
	}

	// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.
	Spec interface {
		Node
		specNode()
	}

	// An ImportSpec node represents a single package import.
	ImportSpec struct {
		Doc     *CommentGroup // associated documentation; or nil
		Name    *Ident        // local package name (including "."); or nil
		Path    *BasicLit     // import path
		Comment *CommentGroup // line comments; or nil
		EndPos  token.Pos     // end of spec (overrides Path.Pos if nonzero)
	}

	// A ValueSpec node represents a constant or variable declaration
	// (ConstSpec or VarSpec production).
	//
	ValueSpec struct {
		Doc     *CommentGroup // associated documentation; or nil
		Names   []*Ident      // value names (len(Names) > 0)
		Type    Expr          // value type; or nil
		Values  []Expr        // initial values; or nil
		Comment *CommentGroup // line comments; or nil
	}

	// A TypeSpec node represents a type declaration (TypeSpec production).
	TypeSpec struct {
		Doc     *CommentGroup // associated documentation; or nil
		Name    *Ident        // type name
		Type    Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
		Comment *CommentGroup // line comments; or nil
	}

	// A BadDecl node is a placeholder for declarations containing
	// syntax errors for which no correct declaration nodes can be
	// created.
	//
	BadDecl struct {
		From, To token.Pos // position range of bad declaration
	}

	// A GenDecl node (generic declaration node) represents an import,
	// constant, type or variable declaration. A valid Lparen position
	// (Lparen.Line > 0) indicates a parenthesized declaration.
	//
	// Relationship between Tok value and Specs element type:
	//
	//	token.IMPORT  *ImportSpec
	//	token.CONST   *ValueSpec
	//	token.TYPE    *TypeSpec
	//	token.VAR     *ValueSpec
	//
	GenDecl struct {
		Doc    *CommentGroup // associated documentation; or nil
		TokPos token.Pos     // position of Tok
		Tok    token.Token   // IMPORT, CONST, TYPE, VAR
		Lparen token.Pos     // position of '(', if any
		Specs  []Spec
		Rparen token.Pos // position of ')', if any
	}

	// A FuncDecl node represents a function declaration.
	FuncDecl struct {
		Doc  *CommentGroup // associated documentation; or nil
		Recv *FieldList    // receiver (methods); or nil (functions)
		Name *Ident        // function/method name
		Type *FuncType     // position of Func keyword, parameters and results
		Body *BlockStmt    // function body; or nil (forward declaration)
	}
)

type File struct {
	Doc        *CommentGroup   // associated documentation; or nil
	Package    token.Pos       // position of "package" keyword
	Name       *Ident          // package name
	Decls      []Decl          // top-level declarations; or nil
	Scope      *Scope          // package scope (this file only)
	Imports    []*ImportSpec   // imports in this file
	Unresolved []*Ident        // unresolved identifiers in this file
	Comments   []*CommentGroup // list of all comments in the source file
}

type Package struct {
	Name    string             // package name
	Scope   *Scope             // package scope across all files
	Imports map[string]*Object // map of package id -> package object
	Files   map[string]*File   // Go source files by filename
}

*/

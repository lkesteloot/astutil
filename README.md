Utilities for manipulating Go abstract syntax tree (ast) nodes.

Visiting
--------

The standard "go/ast" library provides two methods for inspecting the
tree: Inspect() and Walk(). These let you modify the contents of each node,
but not let you replace the node wholesale. Replacing the node is useful
if you've found an expression of one type (say an identifier) and want to replace
it with another. This can be done using Inspect() and Walk(), but they require
you to test the children of every node type. This forces you to create huge switch
statements and know about the contents of every struct.

This library's VisitNode() function passes you a pointer to the node, which lets
you replace it with a different node. The Visitor interface has a few functions
to let you replace expressions and statement.

The code below converts calls of the form:

    "%03d: %02X"(a, b)

to:
    
    fmt.Sprintf("%03d: %02X", a, b)

This is essentially "calling" the format string.

    type stringify bool

    func (p stringify) ProcessNode(node ast.Node) {)
    func (p stringify) ProcessExpr(expr *ast.Expr) {
        switch e := (*expr).(type) {
        case *ast.CallExpr: 
            b, ok := e.Fun.(*ast.BasicLit)
            if ok && b.Kind == token.STRING {
                *expr = &ast.CallExpr{
                    Fun: &ast.SelectorExpr{
                        X: &ast.Ident{
                            Name: "fmt",
                        },
                        Sel: &ast.Ident {
                            Name: "Sprintf",
                        },
                    },
                    Args: append([]ast.Expr{e.Fun}, e.Args...),
                }
            }
        }
    }
    func (p stringify) ProcessStmt(stmt *ast.Stmt) {}

    var p stringify
    astutil.VisitNode(p, f)

Duplicating
-----------

This function will make a deep copy of a tree:

    astutil.DuplicateNode(n)

One caveat is that some pointers in the File and Package objects point to
other nodes within the tree. These are not copied, nor do they point to
internal node -- they're simply nil. Also Object pointers are not copied;
they point to the old objects. Each function's documentation lists the
fields not deep copied.

Zeroing position
----------------

If you parse a Go file and heavily modify it, you'll get odd newlines when
printing it back out. That's because the AST keeps positioning information
and tries to insert newlines to keep the old formatting. It's sometimes better
to zero out the position altogether:

    astutil.StripOutPos(n)

Imports
-------

If you've modified the AST to add function calls into other modules, you may need
to add imports for those modules. Use this functions to do that:

    astutil.AddImport(f, "fmt")

This is, in general, a tricky thing to do, since imports can have paths, be imported
as aliases, and can be shadowed by local variables. This function only handles the
simplest of cases.

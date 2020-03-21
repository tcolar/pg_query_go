package deparse

// Context holds the context during deparsing
type Context interface {
	Kind() ContextKind
}

type SimpleContext struct {
	kind ContextKind
}

func NewContext(kind ContextKind) Context {
	return SimpleContext{
		kind: kind,
	}
}

func NewEmptyContext() Context {
	return SimpleContext{
		kind: EmptyContext,
	}
}

func PassOrEmptyContext(ctx Context) Context {
	if ctx != nil {
		return ctx
	}
	return NewEmptyContext()
}

func (ctx SimpleContext) Kind() ContextKind {
	return ctx.kind
}

type ContextKind int

const (
	EmptyContext ContextKind = iota
	SelectContext
	UpdateContext
	AconstContext
	FuncCall
	ColumnOperator
	TypeName
	DefNameAs
)

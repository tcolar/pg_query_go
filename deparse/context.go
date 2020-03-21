package deparse

// Context holds the context during deparsing
type Context interface {
	Kind() ContextKind
}

type SimpleContext struct {
	kind ContextKind
}

func NewSimpleContext(kind ContextKind) Context {
	return SimpleContext{
		kind: kind,
	}
}

func (ctx SimpleContext) Kind() ContextKind {
	return ctx.kind
}

type ContextKind int

const (
	SelectContext ContextKind = iota
	UpdateContext
)

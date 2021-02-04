package parser

import (
	"fmt"
	"sk-go/ast"
	"sk-go/types"
)

type Result struct {
	result ast.Node
	error  *types.Error
	index  int
}

func newResult() *Result {
	return &Result{}
}

func newFailure(description string) *Result {
	return &Result{
		error: types.NewError("ParsingError", description),
	}
}

func newFailureSpf(format string, a ...interface{}) *Result {
	err := fmt.Sprintf(format, a...)
	return newFailure(err)
}

func (r *Result) Register(other *Result) ast.Node {
	r.index += other.index
	if !other.IsSuccess() {
		r.error = other.error
	}
	return other.result
}

func (r *Result) Advance() {
	r.index++
}

func (r *Result) Success(result ast.Node) *Result {
	r.result = result
	return r
}

func (r *Result) Failure(err *types.Error) *Result {
	if r.error == nil || r.index == 0 {
		r.error = err
	}
	return r
}

func (r *Result) IsSuccess() bool {
	return r.error == nil
}

package parser

import (
	"sk-go/ast"
	"sk-go/token"
	"sk-go/types"
)

const EOF = token.EOF

type Parser struct {
	// The list of tokens from the Lexer
	tokens []*token.Token

	// The last token
	last *token.Token

	// The current token
	token *token.Token

	// The next token
	next *token.Token

	// The index in the tokens where the parser is located
	index int
}

func New(tokens []*token.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) advance() {
	p.index++
	if p.index == len(p.tokens) {
		return
	}
	p.last = p.token
	p.token = p.next
	p.next = p.tokens[p.index]
}

func (p *Parser) PrimaryExpr() *Result {
	if p.token.Type == EOF {
		return newFailure("Unexpected end of source")
	}

	var node ast.Node
	result := newResult()

	switch p.token.Type {
	case token.KEYWORD:
		switch p.token.Value {
		case "fn":
			res := p.FuncDecl()
			node = result.Register(res)
			if !result.IsSuccess() {
				return result
			}
		case "let", "const":
			res := p.VarDecl()
			node = result.Register(res)
			if !result.IsSuccess() {
				return result
			}
		case "return":
			p.advance()
			res := p.Expr()
			node = result.Register(res)
			if !result.IsSuccess() {
				return result
			}
		}
		fallthrough
		// ^ on non match, approach outer default case (read expr)
	case token.LBRACE:
		res := p.Scope()
		node = result.Register(res)
		if !result.IsSuccess() {
			return result
		}
	default:
		res := p.Expr()
		node = result.Register(res)
		if !result.IsSuccess() {
			return result
		}
	}

	if p.token.Type != token.SEMICOLON {
		return newFailureSpf("Expected ';' but instead found '%v'", p.token.String())
	}
	p.advance()

	return result.Success(node)
}

func (p *Parser) Expr() *Result {
	panic("unimplemented")
}

func (p *Parser) Scope() *Result {
	if p.token.Type == token.LBRACE {
		return p.Scope()
	} else {
		result := newResult()
		expression := p.Expr()
		expr := result.Register(expression)
		if !expression.IsSuccess() {
			return expression
		}
		node := &ast.BlockNode{
			TokenPtr:   p.token,
			Statements: []ast.Node{expr},
		}
		return result.Success(node)
	}
}

func (p *Parser) FuncDecl() *Result {
	result := newResult()

	if p.token.Type == token.KEYWORD && p.token.Value == "fn" {
		err := types.NewError("ParsingError", "Expected 'fn' but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	p.advance()
	result.Advance() // todo ??

	if p.token.Type != token.IDENTIFIER {
		err := types.NewError("ParsingError", "Expected identifier but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	name := p.token
	p.advance()
	result.Advance()

	params := make([]*ast.FunctionParamNode, 0)
	if p.token.Type == token.LPAREN {
		for p.token.Type != token.RPAREN {
			p.advance()
			result.Advance()
			if p.token.Type == token.RPAREN {
				break
			}

			paramResult := p.FuncParam()
			param := result.Register(paramResult)
			if !result.IsSuccess() {
				return result
			}

			params = append(params, param.(*ast.FunctionParamNode))
		}

		if p.token.Type != token.RPAREN {
			err := types.NewError("ParsingError", "Expected ')' but found '"+p.token.Value+"'")
			return result.Failure(err)
		}

		p.advance()
		result.Advance()
	}

	if p.token.Type == token.SEMICOLON {
		p.advance()
		result.Advance()

		fn := &ast.FunctionNode{
			Name:       name,
			Parameters: params,
			Body:       &ast.BlockNode{Statements: make([]ast.Node, 0)},
		}
		return result.Success(fn)
	}

	scopeResult := p.StrictScope()
	scope := result.Register(scopeResult).(*ast.BlockNode)
	if !result.IsSuccess() {
		return result
	}

	fn := &ast.FunctionNode{
		Name:       name,
		Parameters: params,
		Body:       scope,
	}
	return result.Success(fn)
}

func (p *Parser) StrictScope() *Result {
	result := newResult()
	start := p.token

	if p.token.Type != token.LBRACE {
		err := types.NewError("ParsingError", "Expected '{' but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	p.advance()
	result.Advance()

	expressions := make([]ast.Node, 0)
	for p.token.Type != token.RBRACE {
		expression := p.PrimaryExpr()
		expr := result.Register(expression)
		if !result.IsSuccess() {
			return result
		}

		expressions = append(expressions, expr)
	}

	if p.token.Type != token.RBRACE {
		err := types.NewError("ParsingError", "Expected '}' but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	p.advance()
	result.Advance()

	node := &ast.BlockNode{
		TokenPtr:   start,
		Statements: expressions,
	}
	return result.Success(node)
}

func (p *Parser) FuncParam() *Result {
	result := newResult()

	spread := p.token.Type == token.THREE_DOT
	if spread {
		p.advance()
		result.Advance()
	}

	if p.token.Type != token.IDENTIFIER {
		err := types.NewError("ParsingError", "Expected identifier but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	name := p.token
	p.advance()
	result.Advance()

	// todo implement type specifier

	var defaultValue ast.Node
	if p.token.Type == token.EQUALS {
		p.advance()
		result.Advance()

		defaultResult := p.Expr()
		value := result.Register(defaultResult)
		if !result.IsSuccess() {
			return result
		}

		defaultValue = value
	}

	param := &ast.FunctionParamNode{
		Name:    name,
		// todo implement type
		Default: defaultValue,
		Spread:  spread,
	}
	return result.Success(param)
}

func (p *Parser) VarDecl() *Result {
	panic("unimplemented")
}

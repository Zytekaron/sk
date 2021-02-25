package parser

import (
	"math/big"
	"sk-go/ast"
	"sk-go/token"
	"sk-go/types"
	"strconv"
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

func (p *Parser) init() {
	p.index = -1
	p.advance()
	p.advance()
}

func (p *Parser) advance() {
	p.index++
	p.last = p.token
	p.token = p.next
	if p.index >= len(p.tokens) {
		p.next = token.NewType(EOF)
		return
	}
	p.next = p.tokens[p.index]
}

func (p *Parser) Parse() *Result {
	p.init()
	if p.token.Type == EOF {
		return newResult().Success(nil)
	}
	var result *Result
	for {
		result = p.PrimaryExpr()
		if !result.IsSuccess() {
			return result
		}
		if p.token.Type == EOF {
			break
		}
	}
	return result
}

func (p *Parser) PrimaryExpr() *Result {
	if p.token.Type == EOF {
		return newFailure("Unexpected end of source")
	}

	var node ast.Node
	result := newResult()

	outer:
	switch p.token.Type {
	case token.KEYWORD:
		switch p.token.Value {
		case "fn":
			res := p.FuncDecl()
			node = result.Register(res)
			if !result.IsSuccess() {
				return result
			}
			break outer
		case "let", "const":
			res := p.VarDecl()
			node = result.Register(res)
			if !result.IsSuccess() {
				return result
			}
			break outer
		case "return":
			p.advance()
			res := p.Expr()
			node = result.Register(res)
			if !result.IsSuccess() {
				return result
			}
			break outer
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
	result := newResult()

	if p.token.Type == EOF {
		err := types.NewError("ParsingError", "Unexpected end of input: found '"+p.token.Value+"'")
		return newResult().Failure(err)
	}

	nodeResult := p.binOp(p.Comparison, p.Comparison, token.SimpleComparison)
	node := result.Register(nodeResult)
	if !result.IsSuccess() {
		return result
	}

	return result.Success(node)
}

func (p *Parser) Comparison() *Result {
	result := newResult()

	if p.token.Type == token.NOT {
		tok := p.token
		p.advance()
		result.Advance()

		exprResult := p.Comparison()
		compExpr := result.Register(exprResult)
		if !result.IsSuccess() {
			return result
		}

		return result.Success(&ast.PrefixOperationNode{
			Token:    tok,
			Operator: token.NOT,
			Right:    compExpr,
		})
	}

	return p.binOp(p.Math, p.Math, token.ComplexComparison)
}

func (p *Parser) Math() *Result {
	return p.binOp(p.Term, p.Term, token.SimpleMath)
}

func (p *Parser) Term() *Result {
	return p.binOp(p.Factor, p.Factor, token.ComplexMath)
}

func (p *Parser) Factor() *Result {
	result := newResult()

	tok := p.token

	if containsType(token.Prefix, tok.Type) {
		p.advance()
		result.Advance()

		factorResult := p.Factor()
		factor := result.Register(factorResult)

		if !result.IsSuccess() {
			return result
		}

		return result.Success(&ast.PrefixOperationNode{
			Token:    tok,
			Operator: tok.Value,
			Right:    factor,
		})
	}

	return p.Power()
}

func (p *Parser) Power() *Result {
	return p.binOp(p.Atom, p.Factor, []token.Type{token.POW})
}

func (p *Parser) Atom() *Result {
	result := newResult()
	tok := p.token

	switch tok.Type {
	case token.INT:
		p.advance()
		result.Advance()
		num, _ := strconv.ParseInt(tok.Value, 10, 32)
		return result.Success(&ast.IntNode{
			Token: tok,
			Value: int32(num),
		})
	case token.LONG:
		p.advance()
		result.Advance()
		num, _ := strconv.ParseInt(tok.Value, 10, 64)
		return result.Success(&ast.LongNode{
			Token: tok,
			Value: num,
		})
	case token.BIGINT:
		p.advance()
		result.Advance()
		var val big.Int
		val.SetString(tok.Value, 10)
		return result.Success(&ast.BigIntNode{
			Token: tok,
			Value: val,
		})
	case token.FLOAT:
		p.advance()
		result.Advance()
		num, _ := strconv.ParseFloat(tok.Value, 32)
		return result.Success(&ast.FloatNode{
			Token: tok,
			Value: float32(num),
		})
	case token.DOUBLE:
		p.advance()
		result.Advance()
		num, _ := strconv.ParseFloat(tok.Value, 64)
		return result.Success(&ast.DoubleNode{
			Token: tok,
			Value: num,
		})
	case token.BIGFLOAT:
		p.advance()
		result.Advance()
		var val big.Float
		val.SetString(tok.Value)
		return result.Success(&ast.BigFloatNode{
			Token: tok,
			Value: val,
		})
	case token.BOOL:
		p.advance()
		result.Advance()
		return result.Success(&ast.BoolNode{
			Token: tok,
			Value: tok.Value == "true",
		})
	case token.STRING:
		p.advance()
		result.Advance()
		node := &ast.StringNode{
			Token: tok,
			Value: tok.Value,
		}
		return result.Success(node)
	case token.LPAREN:
		p.advance()
		result.Advance()

		expr := p.Expr()
		node := result.Register(expr)
		if !result.IsSuccess() {
			return result
		}

		if p.token.Type == token.RPAREN {
			p.advance()
			result.Advance()
			return result.Success(node)
		} else {
			err := types.NewError("ParsingError", "Expected ')' but instead found '"+p.token.Value+"'")
			return result.Failure(err)
		}
	case token.LBRACKET:
		arrayResult := p.ArrayLiteral()
		array := result.Register(arrayResult)
		if !result.IsSuccess() {
			return result
		}
		return result.Success(array)
	// todo check on this, not sure if it will allow multi-accessing via [x][y][z]
	case token.IDENTIFIER:
		identifier := p.token
		p.advance()
		result.Advance()

		if containsType(token.Assignment, p.token.Type) {
			op := p.token.Value
			p.advance()
			result.Advance()

			exprResult := p.Expr()
			expr := result.Register(exprResult)
			if !result.IsSuccess() {
				return result
			}

			return result.Success(&ast.VarAssignNode{
				Name:     identifier,
				Operator: op,
				Value:    expr,
			})
		}

		if p.token.Type == token.LBRACKET {
			elemResult := p.Expr()
			elem := result.Register(elemResult)
			if !result.IsSuccess() {
				return result
			}

			if p.token.Type != token.RBRACKET {
				err := types.NewError("ParsingError", "Expected ']' but found '"+p.token.Value+"'")
				return result.Failure(err)
			}
			p.advance()
			result.Advance()

			return result.Success(&ast.AccessNode{
				Token: identifier,
				Value: elem, // fixme eh?
			})
		}

		if p.token.Type == token.LPAREN {
			args := make([]ast.Node, 0)

			for p.token.Type != token.RPAREN {
				p.advance()
				result.Advance()

				if p.token.Type == token.RPAREN {
					break
				}

				paramResult := p.Expr()
				param := result.Register(paramResult)
				if !result.IsSuccess() {
					return result
				}
				args = append(args, param)
			}

			if p.token.Type != token.RPAREN {
				err := types.NewError("ParsingError", "Expected ')' but instead found '"+p.token.Value+"'")
				return result.Failure(err)
			}
			p.advance()
			result.Advance()

			return result.Success(&ast.FunctionCallNode{
				Token:     identifier,
				Function:  identifier,
				Arguments: args,
			})
		}

		return result.Success(&ast.VarAccessNode{
			Token: identifier,
			Value: identifier.Value,
		})
	}

	err := types.NewError("ParsingError", "Did not pass: atom '"+p.token.Value+"'")
	return result.Failure(err)
}

func (p *Parser) Scope() *Result {
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

	return result.Success(&ast.BlockNode{
		Token:      start,
		Statements: expressions,
	})
}

func (p *Parser) FuncDecl() *Result {
	result := newResult()

	if p.token.Type != token.KEYWORD || p.token.Value != "fn" {
		err := types.NewError("ParsingError", "Expected 'fn' but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	p.advance()
	result.Advance()

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

		return result.Success(&ast.FunctionNode{
			Name:       name,
			Parameters: params,
			Body:       &ast.BlockNode{Statements: make([]ast.Node, 0)},
		})
	}

	scopeResult := p.Scope()
	scopeNode := result.Register(scopeResult)
	if !result.IsSuccess() {
		return result
	}
	scope := scopeNode.(*ast.BlockNode)

	return result.Success(&ast.FunctionNode{
		Name:       name,
		Parameters: params,
		Body:       scope,
	})
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
	if p.token.Type == token.ASSIGN {
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
		Name: name,
		Default: defaultValue,
		Spread:  spread,
	}
	return result.Success(param)
}

func (p *Parser) VarDecl() *Result {
	result := newResult()

	p.advance()
	result.Advance()

	if p.token.Type == token.IDENTIFIER {
		err := types.NewError("ParsingError", "Expected identifier but found '"+p.token.Value+"'")
		return result.Failure(err)
	}

	name := p.token

	p.advance()
	result.Advance()

	var expr ast.Node
	if containsType(token.Assignment, p.token.Type) {
		p.advance()
		result.Advance()

		expression := p.Expr()
		expr = result.Register(expression)
		if !result.IsSuccess() {
			return result
		}
	}

	// todo check up on this!
	node := &ast.VarAssignNode{
		Name:     name,
		Operator: p.token.Value,
		Value:    expr,
	}
	return result.Success(node)
}

func (p *Parser) ArrayLiteral() *Result {
	result := newResult()

	if p.token.Type != token.LBRACKET {
		err := types.NewError("ParsingError", "Expected '[' but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	start := p.token

	nodes := make([]ast.Node, 0)
	for p.token.Type != token.RBRACKET {
		p.advance()
		result.Advance()

		if p.token.Type == token.RBRACKET {
			break
		}

		nodeResult := p.Expr()
		node := result.Register(nodeResult)
		if !result.IsSuccess() {
			return result
		}
		nodes = append(nodes, node)
	}

	if p.token.Type != token.RBRACKET {
		err := types.NewError("ParsingError", "Expected ']' but found '"+p.token.Value+"'")
		return result.Failure(err)
	}
	//end := p.token
	p.advance()
	result.Advance()

	return result.Success(&ast.ArrayNode{
		Token:    start,
		Elements: nodes,
	})
}

func (p *Parser) ObjectLiteral() *Result {
	panic("unimplemented")
}

func (p *Parser) binOp(leftFn, rightFn func() *Result, types []token.Type) *Result {
	result := newResult()

	leftResult := leftFn()
	left := result.Register(leftResult)
	if !result.IsSuccess() {
		return result
	}

	for containsType(types, p.token.Type) {
		op := p.token

		p.advance()
		result.Advance()

		rightResult := rightFn()
		right := result.Register(rightResult)
		if !result.IsSuccess() {
			return result
		}

		left = &ast.BinaryOperationNode{
			Token:    op, // todo fix, earlier token
			Left:     left,
			Operator: op.Value,
			Right:    right,
		}
	}

	return result.Success(left)
}

func containsType(types []token.Type, t token.Type) bool {
	for _, e := range types {
		if e == t {
			return true
		}
	}
	return false
}

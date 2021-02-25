package main

import (
	"fmt"
	"reflect"
	"sk-go/lexer"
	"sk-go/parser"
	"sk-go/token"
	"sk-go/types"
)

type Value interface {
	Format() string
	Type() string
}

type Object struct {}
func (o *Object) Format() string {
	fmt.Println("Format() called on Object")
	return "{?}"
}
func (o *Object) Type() string {
	fmt.Println("Type() called on Object")
	return "object"
}

type String struct {
	*Object
	Value string
}
func (s *String) Format() string {
	fmt.Println("Format() called on String")
	return s.Value
}
func (s *String) Type() string {
	fmt.Println("Type() called on String")
	return "string"
}

func sliceContains(slice interface{}, value interface{}) bool {
	val := reflect.ValueOf(slice)
	for i := 0; i < val.Len(); i++ {
		if val.Index(i).Interface() == value {
			return true
		}
	}
	return false
}

func main() {
	//var lex = lexer.New("0 1 2 'hi \" there' `what's up` % x true false if else switch case string int")
	lex := lexer.New("fn main(a, b = 0) { []; 123; x[0]; };")

	tokens := make([]*token.Token, 0)
	for {
		tok := lex.Next()
		if tok == nil {
			break
		}
		tokens = append(tokens, tok)
	}
	fmt.Println(tokens)

	pars := parser.New(tokens)
	res := pars.Parse()

	fmt.Println("Success:", res.IsSuccess())
	fmt.Println("Error:", res.Error())
	fmt.Println("Result:", res.Result())
	//if res.Result() != nil {
	//	fmt.Println("Result:", res.Result().Format(0, 4, []ast.Node{}))
	//} else {
	//	fmt.Println("Result:", res.Result())
	//}

	ch := iter(lex)
	for t := range ch {
		fmt.Printf("%s '%s'\n", t.Type, t.Value)
	}
}

func iter(lex *lexer.Lexer) chan *token.Token {
	ch := make(chan *token.Token)
	go func() {
		for {
			t := lex.Next()
			if t == nil {
				close(ch)
				break
			}
			ch <- t
		}
	}()
	return ch
}

func test() {
	obj := types.NewObject()
	obj.Set("x", types.NewInt(123))
	obj.Set("x", types.NewFloat(3.141592))
	obj.Set("y", types.NewString("lil peep"))
	obj.Set("err", types.NewError("oh", "fuck"))

	bob := types.NewObject()
	bob.Set("name", types.NewString("Bob"))
	bob.Set("age", types.NewInt(69))
	bob.Set("birthday", types.NewString("15/26/701"))

	obj.Set("bob", bob)

	r1 := obj.Get("x")
	r2 := obj.Get("y")
	fmt.Println(r1, r2)
	fmt.Println(obj.Format(0, 4, make([]types.Value, 0)))
}

func stringContains(str string, r rune) bool {
	for _, c := range []rune(str) {
		if c == r {
			return true
		}
	}
	return false
}

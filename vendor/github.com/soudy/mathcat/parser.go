// Copyright 2016 Steven Oud. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.

package mathcat

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Parser holds the lexed tokens, token position, declared variables and stacks
// used throughout the parsing of an expression.
//
// By default, variables always contains the constants defined below. These can
// however be overwritten.
type Parser struct {
	Tokens    Tokens
	Variables map[string]float64

	pos int
	tok *Token

	operands, operators, arity stack
}

var (
	ErrDivionByZero         = errors.New("Divison by zero")
	ErrUnmatchedParentheses = errors.New("Unmatched parentheses")
	ErrMisplacedComma       = errors.New("Misplaced ‘,’")
	ErrAssignToLiteral      = errors.New("Can't assign to literal")
)

// New initializes a new Parser instance, useful when you want to run multiple
// expression and/or use variables.
func New() *Parser {

	return &Parser{
		pos: 0,
		Variables: map[string]float64{
			"pi":  math.Pi,
			"tau": math.Pi * 2,
			"phi": math.Phi,
			"e":   math.E,
		},
	}
}

// Eval evaluates an expression and returns its result and any errors found.
//
// Example:
//     res, err := mathcat.Eval("2 * 2 * 2") // 8
func Eval(expr string) (float64, error) {
	tokens, err := Lex(expr)

	// If a lexer error occurred don't parse
	if err != nil {
		return -1, err
	}

	p := New()
	p.Tokens = tokens

	return p.parse()
}

// Run executes an expression on an existing parser instance. Useful for
// variable assignment.
//
// Example:
//     p.Run("a = 555")
//     p.Run("a += 45")
//     res, err := p.Run("a + a") // 1200
func (p *Parser) Run(expr string) (float64, error) {
	tokens, err := Lex(expr)

	if err != nil {
		return -1, err
	}

	p.reset()
	p.Tokens = tokens

	return p.parse()
}

// Exec executes an expression with a given map of variables.
//
// Example:
//     res, err := mathcat.Exec("a + b * b", map[string]float64{
//         "a": 1,
//         "b": 3,
//     }) // 10
func Exec(expr string, vars map[string]float64) (float64, error) {
	tokens, err := Lex(expr)

	if err != nil {
		return -1, err
	}

	p := New()
	p.Tokens = tokens

	for name, val := range vars {
		if !IsValidIdent(name) {
			return -1, fmt.Errorf("Invalid variable name: ‘%s’", name)
		}
		p.Variables[name] = val
	}

	return p.parse()
}

// GetVar gets an existing variable.
//
// Example:
//     p.Run("酷 = -33")
//     if val, err := p.GetVar("酷"); !err {
//         fmt.Printf("%f\n", val) // -33
//     }
func (p *Parser) GetVar(index string) (float64, error) {
	if val, ok := p.Variables[index]; ok {
		return val, nil
	}

	return -1, fmt.Errorf("Undefined variable ‘%s’", index)
}

func (p *Parser) parse() (float64, error) {
	// Initializing current token value
	p.tok = p.Tokens[0]

	for !p.eat().Is(EOL) {
		switch {
		case p.tok.IsLiteral():
			if p.peek().Is(LPAREN) {
				// It's a function call, push to operators stack instead
				p.operators.Push(p.tok)

				// Check ahead if the function call has any argument at all, so
				// we can do accurate tracking of arity
				if p.peekN(2).Is(RPAREN) {
					p.arity.Push(0)
				} else {
					p.arity.Push(1)
				}
				break
			}
			p.operands.Push(p.tok)
		case p.tok.Is(LPAREN):
			p.operators.Push(p.tok)
		case p.tok.Is(COMMA):
			for {
				if p.operators.Empty() {
					return -1, ErrMisplacedComma
				}

				if p.operators.Top().(*Token).Is(LPAREN) {
					break
				}

				val, err := p.evaluate(p.operators.Pop().(*Token))
				if err != nil {
					return -1, err
				}

				p.operands.Push(val)
			}
			p.arity.Push(p.arity.Pop().(int) + 1)
		case p.tok.IsOperator():
			if err := p.handleOperator(); err != nil {
				return -1, err
			}
		case p.tok.Is(RPAREN):
			for {
				if p.operators.Empty() {
					return -1, ErrUnmatchedParentheses
				}

				top := p.operators.Pop().(*Token)
				if top.Is(LPAREN) {
					break
				}

				val, err := p.evaluate(top)
				if err != nil {
					return -1, err
				}

				p.operands.Push(val)
			}
		}
	}

	// Evaluate remaining operators
	for !p.operators.Empty() {
		top := p.operators.Pop().(*Token)

		if top.Is(LPAREN) {
			return -1, ErrUnmatchedParentheses
		}

		val, err := p.evaluate(top)
		if err != nil {
			return -1, err
		}

		p.operands.Push(val)
	}

	// If there are no operands, the expression is useless and doesn't do
	// anything, for example `()`
	if p.operands.Empty() {
		return 0, nil
	}

	// Single operand left means the expression was evaluated successful
	if len(p.operands) == 1 {
		return p.lookup(p.operands[0])
	}

	// Leftover token on operand stack indicates invalid syntax
	return -1, fmt.Errorf("Unexpected ‘%s’", p.operands.Top())
}

func (p *Parser) handleOperator() error {
	var o1, o2 *operator

	o1 = operators[p.tok.Type]

	if !p.operators.Empty() {
		if p.operators.Top().(*Token).Is(IDENT) {
			// Special case, if the token on top of the operators stack is
			// a function call, always take precedence above an operator.
			function := p.operators.Pop().(*Token)
			val, err := p.evaluateFunc(function)
			if err != nil {
				return err
			}

			p.operands.Push(val)

			if p.operators.Empty() {
				p.operators.Push(p.tok)
				return nil
			}
		}

		// If the item on top of the operators stack isn't an
		// operator, it's a function call. We then don't need to check
		// precedence.
		if !p.operators.Top().(*Token).IsOperator() {
			p.operators.Push(p.tok)
			return nil
		}

		o2 = operators[p.operators.Top().(*Token).Type]

		if o2.hasHigherPrecThan(o1) {
			operator := p.operators.Pop().(*Token)
			val, err := p.evaluateOp(operator)
			if err != nil {
				return err
			}
			p.operands.Push(val)
		}
	}
	p.operators.Push(p.tok)

	return nil
}

// evaluate gets called when an operator or function call has to be evaluated
// for a result. In case of a function, evaluateFunc is called and in case of
// an operator evaluateOp is called.
func (p *Parser) evaluate(tok *Token) (float64, error) {
	if tok.IsOperator() {
		return p.evaluateOp(tok)
	}

	return p.evaluateFunc(tok)
}

func (p *Parser) evaluateFunc(tok *Token) (float64, error) {
	var (
		function *function
		ok       bool
		i        int
	)

	if function, ok = funcs[tok.Value]; !ok {
		return -1, fmt.Errorf("Undefined function ‘%s’", tok)
	}

	if arity := p.arity.Pop().(int); arity != function.arity {
		return -1, fmt.Errorf("Invalid argument count for ‘%s’ (expected %d, got %d)", tok, function.arity, arity)
	}

	// Start popping off arguments for the function call
	args := make([]float64, function.arity)
	for i = function.arity - 1; i >= 0; i-- {
		if p.operands.Empty() {
			return -1, ErrMisplacedComma
		}

		arg, err := p.lookup(p.operands.Pop())
		if err != nil {
			return -1, err
		}

		args[i] = arg
	}

	return function.fn(args), nil
}

func (p *Parser) evaluateOp(operator *Token) (float64, error) {
	var (
		result      float64
		left, right float64
		err         error
		lhsToken    interface{}
	)

	if p.operands.Empty() {
		return -1, fmt.Errorf("Unexpected ‘%s’", operator)
	}

	if right, err = p.lookup(p.operands.Pop()); err != nil {
		return -1, err
	}

	// Unary operators have no left hand side
	if op := operators[operator.Type]; !op.unary {
		if p.operands.Empty() {
			return -1, fmt.Errorf("Unexpected ‘%s’", operator)
		}
		// Save the token in case of a assignment variable is used and we need
		// to save the result in a variable
		lhsToken = p.operands.Pop()

		// Don't lookup the left hand side if = is used so we can do initial
		// assignment
		if !operator.Is(EQ) {
			left, err = p.lookup(lhsToken)
			if err != nil {
				return -1, err
			}
		}
	}

	result, err = execute(operator, left, right)
	if err != nil {
		return -1, err
	}

	if operator.IsAssignment() {
		// Save result in variable
		if val, ok := lhsToken.(*Token); !(ok && val.Is(IDENT)) {
			return -1, ErrAssignToLiteral
		}
		p.Variables[lhsToken.(*Token).Value] = result
	}

	return result, nil
}

func execute(operator *Token, lhs, rhs float64) (float64, error) {
	var result float64

	// Both lhs and rhs have to be whole numbers for bitwise operations
	if operator.IsBitwise() && !(IsWholeNumber(lhs) && IsWholeNumber(rhs)) {
		return -1, fmt.Errorf("Unsupported type (float) for ‘%s’", operator)
	}

	switch operator.Type {
	case ADD, ADD_EQ:
		result = lhs + rhs
	case SUB, SUB_EQ:
		result = lhs - rhs
	case UNARY_MIN:
		result = -rhs
	case DIV, DIV_EQ:
		if rhs == 0 {
			return -1, ErrDivionByZero
		}
		result = lhs / rhs
	case MUL, MUL_EQ:
		result = lhs * rhs
	case POW, POW_EQ:
		result = math.Pow(lhs, rhs)
	case REM, REM_EQ:
		if rhs == 0 {
			return -1, ErrDivionByZero
		}
		result = math.Mod(lhs, rhs)
	case AND, AND_EQ:
		result = float64(int64(lhs) & int64(rhs))
	case OR, OR_EQ:
		result = float64(int64(lhs) | int64(rhs))
	case XOR, XOR_EQ:
		result = float64(int64(lhs) ^ int64(rhs))
	case LSH, LSH_EQ:
		result = float64(uint64(lhs) << uint64(rhs))
	case RSH, RSH_EQ:
		result = float64(uint64(lhs) >> uint64(rhs))
	case NOT:
		result = float64(^int64(rhs))
	case EQ:
		result = rhs
	case EQ_EQ:
		result = bool2float(lhs == rhs)
	case BANG_EQ:
		result = bool2float(lhs != rhs)
	case GT:
		result = bool2float(lhs > rhs)
	case GT_EQ:
		result = bool2float(lhs >= rhs)
	case LT:
		result = bool2float(lhs < rhs)
	case LT_EQ:
		result = bool2float(lhs <= rhs)
	default:
		return -1, fmt.Errorf("Invalid operator ‘%s’", operator)
	}

	return result, nil
}

// Look up a literal. If it's an identifier, check the parser's variables map,
// otherwise convert the tokenized string to a float64.
func (p *Parser) lookup(val interface{}) (float64, error) {
	// val can be a token or a float64, if it's a float64 it has been already
	// evaluated and we don't need to do anything
	if v, ok := val.(float64); ok {
		return v, nil
	}

	var (
		tmp uint64
		res float64
		err error
	)

	bases := [...]int{
		HEX:    16,
		BINARY: 2,
		OCTAL:  8,
	}

	tok := val.(*Token)
	switch tok.Type {
	case NUMBER:
		res, err = strconv.ParseFloat(tok.Value, 64)
	case HEX, BINARY, OCTAL:
		// Remove 0x part of literal and convert to uint first
		tmp, err = strconv.ParseUint(tok.Value[2:], bases[tok.Type], 64)

		// Then convert to float
		res = float64(tmp)
	case IDENT:
		res, err = p.GetVar(tok.Value)
		if err != nil {
			return -1, err
		}
		return res, nil
	default:
		return -1, fmt.Errorf("Invalid lookup type ‘%s’", tok)
	}

	if err != nil {
		if numError, ok := err.(*strconv.NumError); ok && numError.Err == strconv.ErrRange {
			return -1, fmt.Errorf("Error parsing ‘%s’: %s", tok.Value, strconv.ErrRange)
		}
		return -1, fmt.Errorf("Error parsing ‘%s’: invalid %s", tok.Value, tok)
	}

	return res, nil
}

func (p *Parser) reset() {
	p.Tokens = nil
	p.pos = 0

	p.operators = nil
	p.operands = nil
	p.arity = nil
}

func (p *Parser) peek() *Token {
	return p.Tokens[p.pos]
}

func (p *Parser) peekN(n int) *Token {
	return p.Tokens[p.pos-1+n]
}

func (p *Parser) eat() *Token {
	p.tok = p.peek()
	p.pos++
	return p.tok
}

// IsWholeNumber checks if a float is a whole number
func IsWholeNumber(n float64) bool {
	return float64(int64(n)) == n
}

func bool2float(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

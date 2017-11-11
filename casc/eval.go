package main

import (
	"strconv"
	"strings"
)

func evalExpr(expr string) float64 {
	toks := lexExpr(expr)
	return evalToks(toks, 0)
}

func evalToks(toks *tokens, depth int) float64 {
	var f, next float64
	var op typ = typPlus
	neg := false
	var err error
	resolve := func() {
		if neg {
			next = -next
			neg = false
		}
		switch op {
		case typNeg:
			f -= next
		case typPlus:
			f += next
		case typDiv:
			f /= next
		case typMul:
			f *= next
		default:
			panic(op)
		}
		next = 0
		op = 0
	}
	for toks.next() {
		t := toks.get()
		switch t.typ {
		case typNum:
			next, err = strconv.ParseFloat(t.val, 64)
			if err != nil {
				panic(err)
			}
			resolve()
			continue
		case typNeg:
			if op != 0 {
				neg = !neg
				continue
			}
		case typClose:
			if depth == 0 {
				continue
			}
			return f
		case typOpen:
			next = evalToks(toks, depth+1)
			resolve()
			continue
		}
		if isOp[t.typ] {
			// Sometimes an op will start the string, ignore it.
			if op != 0 {
				continue
			}
			op = t.typ
			continue
		}
		panic(t.typ)
	}
	return f
}

type tokens struct {
	i    int
	toks []token
}

func (t *tokens) next() bool {
	t.i++
	return t.i <= len(t.toks)
}

func (t *tokens) get() token {
	return t.toks[t.i-1]
}

func lexExpr(expr string) *tokens {
	expr = strings.Replace(expr, " ", "", -1)
	var toks tokens
	for i := 0; i < len(expr); i++ {
		next := expr[i]
		t := typMap[next]
		if t != typNum {
			toks.toks = append(toks.toks, token{
				typ: t,
				val: string(next),
			})
			continue
		}
		start := i
		for i++; i < len(expr); i++ {
			if typMap[expr[i]] != typNum {
				break
			}
		}
		cur := expr[start:i]
		toks.toks = append(toks.toks, token{
			typ: typNum,
			val: cur,
		})
		i--
	}
	return &toks
}

type token struct {
	typ typ
	val string
}

type typ int

const (
	typNum typ = iota
	typNeg
	typPlus
	typDiv
	typMul
	typOpen
	typClose
)

var typMap = map[byte]typ{
	'-': typNeg,
	'+': typPlus,
	'/': typDiv,
	'*': typMul,
	'(': typOpen,
	')': typClose,
}

var isOp = map[typ]bool{
	typNeg:  true,
	typPlus: true,
	typDiv:  true,
	typMul:  true,
}

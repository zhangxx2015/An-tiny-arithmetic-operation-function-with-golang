package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Eval(express string) float64 {
	express = strings.Trim(express, " ")
	for {
		lb := strings.IndexAny(express, "(")
		if lb == -1 {
			break
		}
		rb := strings.IndexAny(express[lb+1:], ")")
		if -1 == rb {
			panic("Symbol is not match!")
		}
		rb = lb + rb + 1
		express = fmt.Sprintf("%s %f %s", express[0:lb], Eval(express[lb+1:rb]), express[rb+1:])
	}
	ret, err := strconv.ParseFloat(express, 64)
	if nil == err {
		return ret
	}
	op := strings.IndexAny(express, "+")
	if -1 != op {
		return Eval(express[0:op]) + Eval(express[1+op:])
	}
	op = strings.IndexAny(express, "-")
	if -1 != op {
		return Eval(express[0:op]) - Eval(express[1+op:])
	}
	op = strings.IndexAny(express, "*")
	if -1 != op {
		return Eval(express[0:op]) * Eval(express[1+op:])
	}
	op = strings.IndexAny(express, "/")
	if -1 != op {
		return Eval(express[0:op]) / Eval(express[1+op:])
	}
	panic(fmt.Sprintf("Unknow error![%s]", express))
}

func main() {
	fmt.Println(Eval("92 + 5 + 5 * 27 - (92 - 12) / 4 + 26"))
}

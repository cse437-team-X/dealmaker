package util

import (
	"fmt"
	"testing"
)

func TestCopyFields(t *testing.T) {
	type t1 struct {
		AA string
		BB int
		CC string
	}

	type t2 struct {
		AA string
		BB int
		CC string
		DD float64
	}

	a := t1{
		AA: "t1",
		BB: 2,
		CC: "t1c",
	}

	b:= t2{}

	CopyFieldsByName(a,&b)
	fmt.Println(b)
}
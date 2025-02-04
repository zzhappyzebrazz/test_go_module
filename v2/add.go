// Package testgomodule is a module written by zzhappyzebrazz under the context of the book Learning go by Jon Bodner
// this is a test package for the exercises in the text book
package testgomodule

import (
	constraints "golang.org/x/exp/constraints"
)

// Number is a type constraint that allows either floating-point or integer types.
type Number interface {
	constraints.Float
	constraints.Integer
}

// Add takes two Number parameters and return the Number sum of the two
// be sure to include a link to [math is fun]: https//www.mathisfun.com/numbers/addition.html in your Add functin godoc
// [math is fun]: https//www.mathisfun.com/numbers/addition.html
func Add[T Number](p1, p2 T) T {
	return p1 + p2
}

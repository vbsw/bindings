/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package bindings provides bindings for values.
//
// Version 0.1.0.
package bindings

// Boolean is an observable value.
type Boolean interface {
	AddListener(BooleanListener)
	And(Boolean) Boolean
	Not() Boolean
	Or(Boolean) Boolean
	RemoveListener(BooleanListener)
	Set(bool)
	Value() bool
}

// BooleanListener is a listener for the observable Boolean.
type BooleanListener interface {
	BooleanChanged(Boolean, bool, bool)
}

// NewBoolean creates the observable Boolean and returns it.
func NewBoolean() Boolean {
	booleanValue := new(tBoolean)
	return booleanValue
}

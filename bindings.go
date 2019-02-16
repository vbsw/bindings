/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package bindings provides bindings for values.
//
// Version 0.2.0.
package bindings

// Boolean is an observable value.
type Boolean interface {
	AddListener(BooleanListener)
	And(Boolean) Boolean
	EqualTo(Boolean) Boolean
	Not() Boolean
	NotEqualTo(Boolean) Boolean
	Or(Boolean) Boolean
	RemoveListener(BooleanListener)
	Set(bool)
	Value() bool
}

// Float64 is an observable value.
type Float64 interface {
	AddListener(Float64Listener)
	Divide(Float64) Float64
	EqualTo(Float64) Boolean
	GreaterThan(Float64) Boolean
	GreaterThanOrEqualTo(Float64) Boolean
	LessThan(Float64) Boolean
	LessThanOrEqualTo(Float64) Boolean
	Minus(Float64) Float64
	Multiply(Float64) Float64
	NotEqualTo(Float64) Boolean
	Plus(Float64) Float64
	RemoveListener(Float64Listener)
	Set(float64)
	Value() float64
}

// BooleanListener is a listener for the observable Boolean. Function BooleanChanged is called
// only when observable value has changed, i.e. new value is not equal to old value.
type BooleanListener interface {
	BooleanChanged(Boolean, bool, bool)
}

// Float64Listener is a listener for the observable Float64. Function Float64Changed is called
// only when observable value has changed, i.e. new value is not equal to old value.
type Float64Listener interface {
	Float64Changed(Float64, float64, float64)
}

// NewBoolean creates the observable Boolean and returns it.
func NewBoolean() Boolean {
	booleanValue := new(tBoolean)
	return booleanValue
}

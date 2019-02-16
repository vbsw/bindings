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

type CycleChecker interface {
	CheckCycle(CycleChecker)
}

type BooleanListener interface {
	BooleanChanged(bool, bool)
}

type Boolean interface {
	AddListener(BooleanListener)
	Set(bool)
	Value() bool
}

type tBooleanValue struct {
	listeners []BooleanListener
	value     bool
}

func (booleanValue *tBooleanValue) AddListener(listener BooleanListener) {
	if !containsBooleanListener(booleanValue.listeners, listener) {
		booleanValue.listeners = append(booleanValue.listeners, listener)
	}
}

func (booleanValue *tBooleanValue) Set(b bool) {
	if booleanValue.value != b {
		for _, listener := range booleanValue.listeners {
			listener.BooleanChanged(booleanValue.value, b)
		}
		booleanValue.value = b
	}
}

func (booleanValue *tBooleanValue) CheckCycle(checker CycleChecker) {
}

func (booleanValue *tBooleanValue) Value() bool {
	return booleanValue.value
}

func NewBoolean() Boolean {
	booleanValue := new(tBooleanValue)
	return booleanValue
}

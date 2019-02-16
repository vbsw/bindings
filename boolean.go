/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

type tBoolean struct {
	listeners []BooleanListener
	value     bool
}

type tBooleanAnd struct {
	tBoolean
	parentA Boolean
	parentB Boolean
}

type tBooleanNot struct {
	tBoolean
}

type tBooleanOr struct {
	tBoolean
	parentA Boolean
	parentB Boolean
}

func (booleanValue *tBoolean) AddListener(listener BooleanListener) {
	if !containsBooleanListener(booleanValue.listeners, listener) {
		booleanValue.listeners = append(booleanValue.listeners, listener)
	}
}

func (booleanValueA *tBoolean) And(booleanValueB Boolean) Boolean {
	booleanValueAnd := new(tBooleanAnd)
	booleanValueAnd.parentA = booleanValueA
	booleanValueAnd.parentB = booleanValueB
	booleanValueA.AddListener(booleanValueAnd)
	booleanValueB.AddListener(booleanValueAnd)
	return booleanValueAnd
}

func (booleanValue *tBoolean) Not() Boolean {
	booleanValueNot := new(tBooleanNot)
	booleanValue.AddListener(booleanValueNot)
	return booleanValueNot
}

func (booleanValueA *tBoolean) Or(booleanValueB Boolean) Boolean {
	booleanValueOr := new(tBooleanOr)
	booleanValueOr.parentA = booleanValueA
	booleanValueOr.parentB = booleanValueB
	booleanValueA.AddListener(booleanValueOr)
	booleanValueB.AddListener(booleanValueOr)
	return booleanValueOr
}

func (booleanValue *tBoolean) RemoveListener(listener BooleanListener) {
	i := indexBooleanListener(booleanValue.listeners, listener)
	if i >= 0 {
		booleanValue.listeners = removeBooleanListener(booleanValue.listeners, i)
	}
}

func (booleanValue *tBoolean) Set(newValue bool) {
	if booleanValue.value != newValue {
		booleanValue.updateValue(newValue)
	}
}

func (booleanValue *tBoolean) updateValue(newValue bool) {
	oldValue := booleanValue.value
	observable := Boolean(booleanValue)
	booleanValue.value = newValue
	for _, listener := range booleanValue.listeners {
		listener.BooleanChanged(observable, oldValue, newValue)
	}
}

func (booleanValue *tBoolean) Value() bool {
	return booleanValue.value
}

func (booleanValue *tBooleanAnd) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if booleanValue.parentA == observable {
		booleanValue.updateValue(booleanValue.parentB.Value() && newValue)
	} else {
		booleanValue.updateValue(booleanValue.parentA.Value() && newValue)
	}
}

func (booleanValue *tBooleanNot) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	booleanValue.updateValue(!newValue)
}

func (booleanValue *tBooleanOr) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if booleanValue.parentA == observable {
		booleanValue.updateValue(booleanValue.parentB.Value() || newValue)
	} else {
		booleanValue.updateValue(booleanValue.parentA.Value() || newValue)
	}
}

func NewBoolean() Boolean {
	booleanValue := new(tBoolean)
	return booleanValue
}

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

type tBooleanNot struct {
	tBoolean
}

func (booleanValue *tBoolean) AddListener(listener BooleanListener) {
	if !containsBooleanListener(booleanValue.listeners, listener) {
		booleanValue.listeners = append(booleanValue.listeners, listener)
	}
}

func (booleanValue *tBoolean) CheckCycle(checker CycleChecker) {
}

func (booleanValue *tBoolean) Not() Boolean {
	booleanValueNot := new(tBooleanNot)
	booleanValue.AddListener(booleanValueNot)
	return booleanValueNot
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

func (booleanValue *tBooleanNot) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	booleanValue.updateValue(!newValue)
}

func NewBoolean() Boolean {
	booleanValue := new(tBoolean)
	return booleanValue
}

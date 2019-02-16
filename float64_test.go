/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "testing"

type testFloat64Listener struct {
	called   bool
	oldValue float64
	newValue float64
}

func (float64Listener *testFloat64Listener) Float64Changed(observable Float64, oldValue, newValue float64) {
	float64Listener.called = true
	float64Listener.oldValue = oldValue
	float64Listener.newValue = newValue
}

func TestFloat64ValueAddListener(t *testing.T) {
	float64Value := new(tFloat64)
	float64Listener := new(testFloat64Listener)

	float64Value.AddListener(float64Listener)
	if len(float64Value.listeners) != 1 {
		t.Error(len(float64Value.listeners))
	}

	float64Value.AddListener(float64Listener)
	if len(float64Value.listeners) != 1 {
		t.Error(len(float64Value.listeners))
	}
}

func TestFloat64ValueSet(t *testing.T) {
	float64Value := new(tFloat64)
	initialValue := float64Value.value

	if float64Value.Value() != initialValue {
		t.Error(float64Value.Value())
	}

	float64Value.Set(10)
	if float64Value.Value() != 10 {
		t.Error(float64Value.Value())
	}
}

func TestFloat64ValueFloat64Changed(t *testing.T) {
	float64Value := new(tFloat64)
	float64Listener := new(testFloat64Listener)
	initialValue := float64Value.value

	float64Value.AddListener(float64Listener)
	float64Value.Set(initialValue)
	if float64Listener.called != false {
		t.Error(float64Listener.called)
	}

	float64Value.Set(10)
	if float64Listener.called != true {
		t.Error(float64Listener.called)
	}
	if float64Listener.oldValue != initialValue {
		t.Error(float64Listener.oldValue)
	}
	if float64Listener.newValue != 10 {
		t.Error(float64Listener.newValue)
	}

	float64Value.Set(50)
	if float64Listener.oldValue != 10 {
		t.Error(float64Listener.oldValue)
	}
	if float64Listener.newValue != 50 {
		t.Error(float64Listener.newValue)
	}
}

func TestFloat64ValueEqualTo(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	booleanValueEqual := float64ValueA.EqualTo(float64ValueB)
	booleanListener := new(testBooleanListener)
	initialValueA := float64ValueA.value
	initialValueB := float64ValueB.value

	booleanValueEqual.AddListener(booleanListener)
	float64ValueA.Set(initialValueA)
	float64ValueB.Set(initialValueB)
	if booleanValueEqual.Value() != false {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	float64ValueA.Set(10)
	if booleanValueEqual.Value() != false {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}

	float64ValueB.Set(10)
	if booleanValueEqual.Value() != true {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}
}

func TestFloat64ValueNotEqualTo(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	booleanValueNotEqual := float64ValueA.NotEqualTo(float64ValueB)
	booleanListener := new(testBooleanListener)
	initialValueA := float64ValueA.value
	initialValueB := float64ValueB.value

	booleanValueNotEqual.AddListener(booleanListener)
	float64ValueA.Set(initialValueA)
	float64ValueB.Set(initialValueB)
	if booleanValueNotEqual.Value() != false {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	float64ValueA.Set(10)
	if booleanValueNotEqual.Value() != true {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.called != true {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	float64ValueB.Set(10)
	if booleanValueNotEqual.Value() != false {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.oldValue != true {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}
}

/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

// containsBooleanListener returns true, if list contains key.
func containsBooleanListener(list []BooleanListener, key BooleanListener) bool {
	for _, currentKey := range list {
		if currentKey == key {
			return true
		}
	}
	return false
}

// indexBooleanListener returns index of key in list. If list does not contain key, returns -1.
func indexBooleanListener(list []BooleanListener, key BooleanListener) int {
	for i, currentKey := range list {
		if currentKey == key {
			return i
		}
	}
	return -1
}

// removeBooleanListener removes an entry from list.
func removeBooleanListener(list []BooleanListener, index int) []BooleanListener {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

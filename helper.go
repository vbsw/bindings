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

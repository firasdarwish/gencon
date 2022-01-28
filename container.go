/*
 * Copyright (c) 2020 Firas M. Darwish <firas@dev.sy> ( https://firas.dev.sy )
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package gencon

import (
	"fmt"
	"reflect"
)

var services []service[any]

//Add adds a new interface/implementation pair to the container
func Add[T comparable](interf, impl T) {
	s := service[any]{
		inter: interf,
		impl:  impl,
	}
	services = append(services, s)
}

//Get retrieves an implementation of T from the container
// panics if no implementation was found
func Get[T comparable]() T {
	for _, s := range services {
		s, ok := s.impl.(T) // checks whether service implements T
		if ok {
			return s
		}
	}

	var requestedTypeInterface T
	panic[string](fmt.Sprintf("NO IMPLEMENTATION FOR %s WAS FOUND", reflect.TypeOf(&requestedTypeInterface).String()))
}

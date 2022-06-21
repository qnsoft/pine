// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

type AbstractCache interface {
	Get(string) ([]byte, error)
	GetWithUnmarshal(string, interface{}) error

	Set(string, []byte, ...int) error
	SetWithMarshal(string, interface{}, ...int) error

	Delete(string) error
	Exists(string) bool
}

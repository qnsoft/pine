// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sessions

import (
	"fmt"
)

type Session struct {
	id    string
	store AbstractSessionStore
}

type entry struct {
	Val   string
	Flush bool
}

func newSession(id string, store AbstractSessionStore) (*Session, error) {
	sess := &Session{
		store: store,
		id:    id,
	}
	return sess, nil
}

func (sess *Session) Set(key string, val string) {
	sess.store.Save(sess.makeKey(key), &entry{Val: val})
}

func (sess *Session) Get(key string) string {
	var val entry
	if err := sess.store.Get(sess.makeKey(key), &val); err == nil {
		if val.Flush {
			sess.Remove(key)
		}
	}
	return val.Val
}

func (sess *Session) AddFlush(key string, val string) {
	sess.store.Save(sess.makeKey(key), &entry{Val: val, Flush: true})
}

func (sess *Session) Remove(key string) {
	sess.store.Delete(sess.makeKey(key))
}

func (sess *Session) makeKey(key string) string {
	return fmt.Sprintf("%s_%s", sess.id, key)
}

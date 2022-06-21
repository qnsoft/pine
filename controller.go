// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pine

import (
	"reflect"

	"github.com/qnsoft/logger"
	"github.com/qnsoft/pine/sessions"
)

const ControllerSuffix = "Controller"

type Controller struct {
	context *Context
}

var _ IController = (*Controller)(nil)

type IController interface {
	Ctx() *Context

	Render() *Render

	Logger() logger.AbstractLogger
	Session() sessions.AbstractSession
	Cookie() *sessions.Cookie
}

var reflectingNeedIgnoreMethods = map[string]struct{}{}

func init() {
	rt := reflect.TypeOf(&Controller{})
	for i := 0; i < rt.NumMethod(); i++ {
		reflectingNeedIgnoreMethods[rt.Method(i).Name] = struct{}{}
	}
}

func (c *Controller) Ctx() *Context {
	return c.context
}

func (c *Controller) Cookie() *sessions.Cookie {
	return c.context.cookie
}

func (c *Controller) Render() *Render {
	return c.context.Render()
}

func (c *Controller) Param() params {
	return c.context.params
}

func (c *Controller) View(name string) {
	c.Render().HTML(name)
}

func (c *Controller) Logger() logger.AbstractLogger {
	return c.context.Logger()
}

func (c *Controller) Session() sessions.AbstractSession {
	return c.context.Session()
}

func (c *Controller) ViewData(key string, val interface{}) {
	c.Render().ViewData(key, val)
}

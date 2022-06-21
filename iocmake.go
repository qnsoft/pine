// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pine

import (
	"github.com/qnsoft/logger"
	"github.com/qnsoft/pine/di"
)

func Make(service interface{}, params ...interface{}) interface{} {
	return di.MustGet(service, params...)
}

func Logger() logger.AbstractLogger {
	return Make(di.ServicePineLogger).(logger.AbstractLogger)
}

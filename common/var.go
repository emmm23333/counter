package common

import (
	"io"

	"go.uber.org/zap"
)

// Log object
var Log *zap.SugaredLogger

// log output target object
var Outer io.Writer

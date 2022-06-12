package core

import (
	"github.com/depocket/lab/base"
	"github.com/depocket/support/dynamic"
)

var (
	ServiceRegisters = dynamic.NewScheme[base.Adapter]()
)

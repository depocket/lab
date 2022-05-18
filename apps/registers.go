package apps

import (
	"github.com/depocket/lab/apps/lido"
	"github.com/depocket/lab/core"
)

var (
	ServiceRegisters = core.NewScheme()
)

func init() {
	ServiceRegisters.RegisterService("ethereum_lido", lido.EthereumService{})
}

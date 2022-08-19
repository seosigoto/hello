package keeper

import (
	"github.com/seosigoto/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}

// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package consts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
)

// token1rvzhmceq997zntgvravfagsks6w0ryud3rylh4cdvayry0dl97nsjzf3yp
//./build/token-cli chain import-anr
// 2LdF8NEFmzDNf9eP31zyVcxEEjPUn1xvGmXew1uFXbf7oM3rTn
// 2usbMDkzdaCMEtRZAGeDiWD72rZQSKiPWR8wxPmQUCGQFC8q56

const (
	// TODO: choose a human-readable part for your hyperchain
	HRP = "SMASHALE"
	// TODO: choose a name for your hyperchain
	Name = "ShashankMashale"
	// TODO: choose a token symbol
	Symbol = "SM"
)

var ID ids.ID

func init() {
	b := make([]byte, consts.IDLen)
	copy(b, []byte(Name))
	vmID, err := ids.ToID(b)
	if err != nil {
		panic(err)
	}
	ID = vmID
}

// Instantiate registry here so it can be imported by any package. We set these
// values in [controller/registry].
var (
	ActionRegistry *codec.TypeParser[chain.Action, *warp.Message, bool]
	AuthRegistry   *codec.TypeParser[chain.Auth, *warp.Message, bool]
)

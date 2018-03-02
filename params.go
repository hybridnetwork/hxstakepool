// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"github.com/hybridnetwork/hxd/chaincfg"
	"github.com/hybridnetwork/hxd/wire"
)

// activeNetParams is a pointer to the parameters specific to the
// currently active HX network.
var activeNetParams = &mainNetParams

// params is used to group parameters for various networks such as the main
// network and test networks.
type params struct {
	*chaincfg.Params
	StakepooldRPCServerPort string
	WalletRPCServerPort     string
}

// mainNetParams contains parameters specific to the main network
// (wire.MainNet).  NOTE: The RPC port is intentionally different than the
// reference implementation because hxd does not handle wallet requests.  The
// separate wallet process listens on the well-known port and forwards requests
// it does not handle on to hxd.  This approach allows the wallet process
// to emulate the full reference implementation RPC API.
var mainNetParams = params{
	Params:                  &chaincfg.MainNetParams,
	StakepooldRPCServerPort: "9898",
	WalletRPCServerPort:     "9898",
}

// testNet2Params contains parameters specific to the test network (version 0)
// (wire.TestNet).  NOTE: The RPC port is intentionally different than the
// reference implementation - see the mainNetParams comment for details.

var testNet2Params = params{
	Params:                  &chaincfg.TestNet2Params,
	StakepooldRPCServerPort: "12013",
	WalletRPCServerPort:     "12010",
}

// simNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var simNetParams = params{
	Params:                  &chaincfg.SimNetParams,
	StakepooldRPCServerPort: "29898",
	WalletRPCServerPort:     "29898",
}

// netName returns the name used when referring to a decred network.  At the
// time of writing, hxd currently places blocks for testnet version 0 in the
// data and log directory "testnet", which does not match the Name field of the
// chaincfg parameters.  This function can be used to override this directory name
// as "testnet" when the passed active network matches wire.TestNet.
//
// A proper upgrade to move the data and log directories for this network to
// "testnet" is planned for the future, at which point this function can be
// removed and the network parameter's name used instead.
func netName(chainParams *params) string {
	switch chainParams.Net {
	case wire.TestNet2:
		return "testnet2"
	default:
		return chainParams.Name
	}
}

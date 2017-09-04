// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the mian WhaleCoin network.
var MainnetBootnodes = []string{

	// WhaleCoinOrg  Go Bootnodes
  "enode://173d88a03853011bd7505b205f40cc144646bbf843f93682457447c4e19407d9f2389a8657ce5f4eceede4076d4bf329a3630b554a6f353ae8d397b150b8c5e6@34.231.64.228:30371",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// WhaleCoin test network.

var TestnetBootnodes = []string{
	// "enode://4f0f14c68ae3dc1e9e883ff6f6746871f56446356569e927dfbd039f5a09f5dc76fe1b1c84e239a618f338c3013c9db0400d80a32c216df22664aacdfc6a0666@34.231.48.74:30371",
	// "enode://e6a4a7b8e6f0adf290af9c7a979c3bd332bf638e57a4b3d2a641cbafd4717f8e442e65ffac711e6ae5e84eae3161dc824fd407685ab21aa2a7957eb1f98800c8@34.192.101.147:30373",
}

var DiscoveryV5Bootnodes = []string{}
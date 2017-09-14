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
  "enode://173d88a03853011bd7505b205f40cc144646bbf843f93682457447c4e19407d9f2389a8657ce5f4eceede4076d4bf329a3630b554a6f353ae8d397b150b8c5e6@34.231.64.228:30373",
  "enode://197598fb2872bf03aef8eceba4c39b04bea9638ffc62209cc92bcb5e8545bd4bf842665b29a14b92c2cf58b500f9dcb31d4e312d4a2570118f17fb4302717f58@34.197.224.180:30373",
  "enode://c584fc5843efad1e2339977807ba7c0c77b8dcf2b4a954a50dd36920bb4c464ff91cfbe016537d87a79aacbb12cd74ebc594091dbdd92007590334c0f3e88d25@213.32.53.181:30371",
  "enode://a2deaff121392e2286c59e52325b8b682fd7b09991df9fd91146a20d1b52d1eaaf6c9bbdaffded2083177a62fb026c2f9f71348018c279cc7be75f02a05296e1@138.68.29.211:30371",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// WhaleCoin test network.

var TestnetBootnodes = []string{

	"enode://b6708caf671522f9c376b430db37fae0e336badc25446e8055ab46d3497c91e765349ba26c2811deaa1b7f3ce6cb1bdda8c72831acd9b493c85bbbc3f6bef08d@34.231.48.74:30373",
	// "enode://271d3c7b3422e5447df4d6b8478b6c913c41bbd9490e34b92cdf7910b958ec938001ba2d60955f53812f66206fbedd0b9daea1c44621ddfb7df25e24dfaba51c@34.192.101.147:30373",
}

var DiscoveryV5Bootnodes = []string{}

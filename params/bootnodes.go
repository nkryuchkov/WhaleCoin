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
  "enode://ff82427382f1c4b0c6b0113e742cb0b06264385ff2c074e65780a667325358fd327f9f7b3fd65e74e89a2c8081448533a80324d52da25c781a224d749d629441@34.235.197.115:30371",
  "enode://690db490f15a366b15126f67221bbcb0d4a6e6089e96db13787e805d3d273f3ffeccabc8bd149e6bbf6a4ab4744ed51ab7723c129007348dd7e0ff9e91df064e@213.32.53.181:30371",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// WhaleCoin test network.

var TestnetBootnodes = []string{

//	"enode://b6708caf671522f9c376b430db37fae0e336badc25446e8055ab46d3497c91e765349ba26c2811deaa1b7f3ce6cb1bdda8c72831acd9b493c85bbbc3f6bef08d@34.231.48.74:30373",
	// "enode://271d3c7b3422e5447df4d6b8478b6c913c41bbd9490e34b92cdf7910b958ec938001ba2d60955f53812f66206fbedd0b9daea1c44621ddfb7df25e24dfaba51c@34.192.101.147:30373",
}

var DiscoveryV5Bootnodes = []string{}

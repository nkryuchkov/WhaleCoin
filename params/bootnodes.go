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
// the main Ethereum network.
var MainnetBootnodes = []string{

	// Ethereum Foundation Go Bootnodes
  "enode://244220001945e09367fb9863a3701c17083bf01b016e9ec9fd8940bf6b25153962c1f4b9ca67549796f142f9c059de1cbec44dbf0e35524f1ebf4ff5eecab5d3@34.231.64.228:30371",

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
/*
{
  enode: "enode://cd15a411c96b2dceb43a7365247bfb6d0017b3b04f45498804e1b8ea3c27c8b1329cf73a5a87b0fad71bf6cc84ddf59b19213980b71c81181d3400259e372ef5@[::]:30373?discport=0",
  id: "cd15a411c96b2dceb43a7365247bfb6d0017b3b04f45498804e1b8ea3c27c8b1329cf73a5a87b0fad71bf6cc84ddf59b19213980b71c81181d3400259e372ef5",
  ip: "::",
  listenAddr: "[::]:30373",
  name: "gwhale/v1.6.7-stable-26c696fc/linux-amd64/go1.8.3",
  ports: {
    discovery: 0,
    listener: 30373
  },
  protocols: {
    eth: {
      difficulty: 711726653,
      genesis: "0xb3756bec461cefd484dea6fbb115fc2db0cd1c06fae53c39cfc0bff9f3f40bf7",
      head: "0x1ecb7960ebb203981ba38331078e14e2a642626b9e0b81622be918f363328086",
      network: 1211
    }
  }
}
*/
var TestnetBootnodes = []string{
	"enode://7ce558ae86b58988123782fcda7e99deaac9603d3c2db320ff4f1ae6c2a68782ad2c9eb232444db407cf84dbdbd557dfc0bbe4c5ccaaa24f50ed43cbf042a6ca@34.231.48.74:30371",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	// "enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30373", // IE
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	// "enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30373?discport=30304", // IE
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	// "enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	// "enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	// "enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
}

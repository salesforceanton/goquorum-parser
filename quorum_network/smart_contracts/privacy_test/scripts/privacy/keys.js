
// WARNING: the keys here are demo purposes ONLY. Please use a tool like EthSigner for production, rather than hard coding private keys

module.exports = {
  tessera: {
    member1: {
      publicKey: "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI=",
    },
    member2: {
      publicKey: "xOt7tRhG5lDBljQfAW7DC0ZDAJlBYbcFZVx4LJEo7xg=",
    },
    member3: {
      publicKey: "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao=",
    },
  },
  quorum: {
    validator1: {
      name: "validator1",
      url: "http://127.0.0.1:21001",
      nodekey: "91737ad95f833033f6362ff3cdda9a2926d8002a12116d836819c4c0952ea34a",
      accountAddress: "0x7b0c8c3def4d34aedd54412d373acecfb96326b2",
      accountPrivateKey: "0x87aed0d209febbfca2469ee9af458c0fed776172648adf507a5ea35e6f81012a"
    },
    validator2: {
      name: "validator2",
      url: "http://127.0.0.1:21002",
      nodekey: "08e2067d87eadc13c43fd1173fc035c1eee1fb530b82e3762306867d84b9d179",
      accountAddress: "0x30a21ccca26f11f7f34a303342c68d79eb8a0945",
      accountPrivateKey: "0xe6e3a3f4d85b420a581218f31e70cf189bb39ee4456b7c4b42c60cc9651c23e1"
    },
    validator3: {
      name: "validator3",
      url: "http://127.0.0.1:21003",
      nodekey: "96cbe05cc632bd514e583f551a871704f3953b0311a9ced3a9db8dce56f0ef2b",
      accountAddress: "0x97c3f95a6883fa7a314bea7e54c74f9fa2f81e1a",
      accountPrivateKey: "0x4fb219a85669dc4dcd12def796574e7abcf2a9bf4acc0dc09e6dd0add99cfb90"
    },
    bootnode1: {
      name: "bootnode1",
      url: "http://127.0.0.1:21004",
      nodekey: "2c92b6f8a89a79dfe8e692fcc998ab91f38260a8164b9bd2975794413a496678",
      accountAddress: "0x51c9ade0a2a39e680b48f53931f61217ee952866",
      accountPrivateKey: "0xe0f71a0ec76978818df0f2f74baecef304ed51999a2bdadc2f73eb24b448e79b"
    },   
    bootnode2: {
      name: "bootnode2",
      url: "http://127.0.0.1:21005",
      nodekey: "ee5032363aca6b0de0df3270e3b20f0ccc8f5233e7cdd3535f180eb7b14b1908",
      accountAddress: "0x45fa074ed10da818b86491a73917fd29e76825ff",
      accountPrivateKey: "0x7cad97241b8a022d25a4e4a42e6f3c7112913dccd4ce1c9be1f8a62742c23754"
    },       
    rpcnode: {
      name: "rpcnode",
      url: "http://127.0.0.1:8545",
      wsUrl: "ws://127.0.0.1:8546",
      nodekey: "99ef3c369c50f39f396e8499d1979898347b76dfb02ffdc0ce28c4c01cb528a3",
      accountAddress: "0xdfdc34b8add053ac7f13ef05aa36566d2c884501",
      accountPrivateKey: "0x2f6bdb1173d3eee59eba25f9f6b542422c0bba6b60c12220b82a6d87390ac2eb"
    },       
    member1: {
      name: "member1",
      url: "http://127.0.0.1:20000",
      wsUrl: "ws://127.0.0.1:20001",
      privateUrl: "http://127.0.0.1:9081",
      nodekey: "fafcdc9cca0a707a4b45efb5c9d734ae0336408a03d5a57d62941dad475598d3",
      accountAddress: "0x111b0e42105a3412eee350a06a46edf986ce9172",
      accountPrivateKey: "0x865c9c89000a23d4ff92a9f776284ff7184799169c327303de8efc438ea05171"
    },
    member2: {
      name: "member2",
      url: "http://127.0.0.1:20002",
      wsUrl: "ws://127.0.0.1:20003",
      privateUrl: "http://127.0.0.1:9082",
      nodekey: "076aa334c12fdf76b8edd2ef741405111fead0a4e2334eb8add918e777145f00",
      accountAddress: "0x9eb168ff6ef4369fff996506a1e584e8c0e84a8c",
      accountPrivateKey: "0x4071c7719b3a4a06f0c7b85208d439d3a4ebdc2bace0f01cfea0225246fc6e94"
    },
    member3: {
      name: "member3",
      url: "http://127.0.0.1:20004",
      wsUrl: "ws://127.0.0.1:20005",
      privateUrl: "http://localhost:9083",
      nodekey: "abd79d275195292529053f778aedd87a1c261cf1a2912551fb8baafc76409124",
      accountAddress: "0x5fd43bdb977cfeadb8b3e3a7a205b1bffa954ab1",
      accountPrivateKey: "0x0934357397d81ab0616fa8da09148d4dc6f3e79e2d62f72b87b676fccee5f338"
    }
  },
};

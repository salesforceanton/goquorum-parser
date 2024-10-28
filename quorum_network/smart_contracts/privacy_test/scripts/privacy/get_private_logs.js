const Web3 = require("web3");
const Web3Quorum = require("web3js-quorum");

const { tessera, quorum } = require("../../../keys.js");
const client = quorum.member1;

const web3 = new Web3(client.url);
const web3quorum = new Web3Quorum(
  web3,
  { privateUrl: client.privateUrl },
  true
);
const contractAddress = "0xebaf93b936e40dcd4f5fc78c15d01e91d09e085d";

function run() {
  const filter = {
    address: contractAddress,
  };

  return web3quorum.priv.getLogs(1, filter).then((logs) => {
    console.log("Received logs\n", logs);
    return logs;
  });
}

run();
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
const contractAddress = "0xe54d2778f6fecc525f255f1c786422d115809822";

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
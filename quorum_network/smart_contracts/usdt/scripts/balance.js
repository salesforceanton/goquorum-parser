const Web3 = require("web3");

// validator1 details
const { quorum, accounts } = require("../../keys.js");
const host = quorum.validator1.url;

let web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider(host));

const balanceOfABI = [
  {
    constant: true,
    inputs: [
      {
        name: "_owner",
        type: "address",
      },
    ],
    name: "balanceOf",
    outputs: [
      {
        name: "balance",
        type: "uint256",
      },
    ],
    payable: false,
    stateMutability: "view",
    type: "function",
  },
];

// USDT token contract
const tokenContract = "0x064f860b6683223b03b38252853d5d2c210cce19";
// A USDT token holder
const tokenHolder = "0x064f860b6683223b03b38252853d5d2c210cce19";
const contract = new web3.eth.Contract(balanceOfABI, tokenContract);

async function getTokenBalance() {
  const result = await contract.methods.balanceOf(tokenHolder).call();
  console.log(result);
}

getTokenBalance();

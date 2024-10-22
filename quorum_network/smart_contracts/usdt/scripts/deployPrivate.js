const path = require("path");
const fs = require("fs-extra");
const Web3 = require("web3");
const Web3Quorum = require("web3js-quorum");

// WARNING: the keys here are demo purposes ONLY. Please use a tool like EthSigner for production, rather than hard coding private keys
const { tessera, quorum } = require("../../keys.js");
const chainId = 1338;
const privacyFlag = 0;
const contractName = "TetherToken"

const contractAbi = JSON.parse(
  fs.readFileSync(
      path.resolve(
        __dirname,
        "../",
        "contracts/output",
        `${contractName}.abi`
      )
    )
  );
const contractBytecode = 
  fs.readFileSync(
      path.resolve(
        __dirname,
        "../",
        "contracts/output",
        `${contractName}.bin`
      )
    )
    .toString();

async function createContract(
  client,
  fromPrivateKey,
  fromPublicKey,
  toPublicKey
) {
  const web3 = new Web3(client.url);
  const web3quorum = new Web3Quorum(
    web3,
    { privateUrl: client.privateUrl },
    true
  );

  // unlock the account so you can sign the tx; uses web3.eth.accounts.decrypt(keystoreJsonV3, password);
  const accountKeyPath = path.resolve(
    __dirname,
    "../../../",
    "config/nodes",
    client.name,
    "accountKeystore"
  );
  const accountKey = JSON.parse(fs.readFileSync(accountKeyPath));
  const signingAccount = web3.eth.accounts.decrypt(accountKey, "");

  // get the nonce for the accountAddress
  const accountAddress = client.accountAddress;
  const txCount = await web3.eth.getTransactionCount(`${accountAddress}`);

  // initialize the default constructor with a value `47 = 0x2F`; this value is appended to the bytecode
  const initialSupply = 100000000000000
  const name = 'TetherToken'
  const symbol = 'USDT'
  const decimals = 6

  const contractConstructorInit = web3.eth.abi.encodeParameters(
    ["uint256", "string", "string", "uint256"],
    [initialSupply, name, symbol, decimals]
  );    

  const txOptions = {
    chainId,
    nonce: txCount,
    gasPrice: 0, //ETH per unit of gas
    gasLimit: 0x24a22, //max number of gas units the tx is allowed to use
    value: 0,
    data: "0x" + contractBytecode + contractConstructorInit,
    from: signingAccount,
    isPrivate: true,
    privateKey: fromPrivateKey,
    privateFrom: fromPublicKey,
    privateFor: [toPublicKey],
    privacyFlag,
  };
  console.log("Creating contract...");

  // Generate and send the Raw transaction to the Besu node using the eea_sendRawTransaction JSON-RPC call
  const txHash = await web3quorum.priv.generateAndSendRawTransaction(txOptions);
  console.log("Getting contractAddress from txHash: ", txHash);
  return txHash;
}

async function main() {
  const res = await createContract(
    quorum.member1,
    quorum.member1.accountPrivateKey,
    tessera.member1.publicKey,
    tessera.member3.publicKey
  )
      
  console.log(`Contract Address: ${res.contractAddress}`);
}

if (require.main === module) {
  main();
}

module.exports = exports = main;

// 0xe54d2778f6fecc525f255f1c786422d115809822

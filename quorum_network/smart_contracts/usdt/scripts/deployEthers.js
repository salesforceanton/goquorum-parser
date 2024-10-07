const ethers = require("ethers");
const fs = require("fs-extra");
const path = require("path");

// member1 details
const { quorum } = require("../../keys.js");
const host = quorum.validator1.url;

const provider = new ethers.providers.JsonRpcProvider(host);

const privateKey = quorum.validator1.accountPrivateKey;
const wallet = new ethers.Wallet(privateKey, provider);

async function contractDeployer(contractName, ...argumentss) {
  const abi = JSON.parse(
    fs.readFileSync(
      path.resolve(
        __dirname,
        "../",
        "contracts/output",
        `${contractName}.abi`
      )
    )
  );
  const bytecode = fs
    .readFileSync(
      path.resolve(
        __dirname,
        "../",
        "contracts/output",
        `${contractName}.bin`
      )
    )
    .toString();
  const factory = new ethers.ContractFactory(abi, bytecode, wallet);
  const contract = await factory.deploy(...argumentss);
  await contract.deployTransaction.wait();
  console.log(`${contractName} Contract Address: ${contract.address}`);
  return contract.address;
}

function createContractsConfigFile(
    usdtAddress,
  ) {
    let cfg = {
      usdtAddress: usdtAddress,
    };
    fs.writeJsonSync("contracts-config.json", cfg, { spaces: 2 });
  }

(async function main() {
  const initialSupply = 100000000000000
  const name = 'TetherToken'
  const symbol = 'USDT'
  const decimals = 6

  const contractInit = [initialSupply, name, symbol, decimals]
  
  const usdtAddress = await contractDeployer(
    "TetherToken",
    ...contractInit,
  );

  // tx hash: 
  // eth.getTransactionReceipt("0xdce98d88e0eca4faf17a9c3b5508a591cd957d1a2aeee4910282f8814be32ad8")
  
  createContractsConfigFile(usdtAddress)
})();



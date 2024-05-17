// demonstrates how to get a transaction status
const { providers } = require("near-api-js");
// import providers from "near-api-js"
//network config (replace testnet with mainnet or betanet)
const provider = new providers.JsonRpcProvider(
    "https://archival-rpc.testnet.near.org"
);

const TX_HASH = "4DVSGALhRJatgvWoNbhxap9H5y65L1fcuemvubEEGrhi";
// account ID associated with the transaction
const ACCOUNT_ID = "ypenghui7.testnet";

getState(TX_HASH, ACCOUNT_ID);

async function getState(txHash, accountId) {
    const result = await provider.txStatus(txHash, accountId);
    console.log("Result: ", Buffer.from(JSON.stringify(result.transaction.actions[0].FunctionCall.args),'base64').toString());
}
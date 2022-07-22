# chainlink_contract

Notice: This is only the onchain part repo, offchain(node) part is as same as official deployment

https://docs.chain.link/docs/running-a-chainlink-node/



# Deployment

The contract to deploy chainlink with our own token

![image|684x500](./img/files.png?raw=true)



main contracts are `MLINK.sol`, `caller.sol` and `Oracle.sol`, other contracats are dependencies(but you still need them since they will be used in `Oracle.sol`).

`MLINK.sol` is the ERC20 token to replace the LINK payment.

`Oracle.sol` is the oracle contract.

 `caller.sol` is a caller contract to call oracle.

you can mimic `function requestEthereumPrice(address _oracle, string memory _jobId)`in `caller.sol` in your contract to interact with oracle. 



# Appendix

## chainlink structure

![image|684x500](./img/chainlink.drawio.png?raw=true)




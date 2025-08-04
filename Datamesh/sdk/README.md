# DataMesh SDK

## Purpose
The DataMesh SDK provides tools, libraries, and code generation utilities to launch modular appchains or sovereign rollups using DataMesh infrastructure.

## Technology
- **Go**: CLI and codegen tools
- **Protobuf**: Inter-module messages and APIs

## Responsibilities
- CLI to scaffold new appchains/rollups
- Codegen for proto definitions and client libraries
- Developer utilities for interacting with DataMesh modules

---

## Advanced Modules & Integration Notes

### Consensus (Tendermint/CometBFT ABCI)
- See `consensus/abci.go` for ABCI app stubs
- [Tendermint Docs](https://docs.tendermint.com/)

### Data Availability
- **Erasure Coding:** `data_availability/erasure.go` (uses [klauspost/reedsolomon](https://github.com/klauspost/reedsolomon))
- **KZG Commitments:** `data_availability/kzg.go` (see [protolambda/go-kzg](https://github.com/protolambda/go-kzg))
- **DAS:** `data_availability/das.go` (sampling API stub)

### Execution
- **EVM:** `execution/evm.go` (see [evmos/ethermint](https://github.com/evmos/ethermint), [go-ethereum](https://github.com/ethereum/go-ethereum))
- **WASM:** `execution/wasm.go` (see [wasmerio/wasmer-go](https://github.com/wasmerio/wasmer-go), CosmWasm)
- **zkVM:** `execution/zkvm.go` (see [zkevm-circuits](https://github.com/privacy-scaling-explorations/zkevm-circuits))

### Fraud Proofs & Shared Security
- **Fraud Proofs:** `consensus/fraudproof.go` (API stubs)
- **Staking/Slashing:** `consensus/shared_security.go` (API stubs)

### Interop (IBC-like)
- **Interop Layer:** `interop/interop.go` (cross-chain message stubs)

---

## Library Recommendations
- **Cosmos SDK:** [cosmos/cosmos-sdk](https://github.com/cosmos/cosmos-sdk)
- **IPFS:** [ipfs/go-ipfs](https://github.com/ipfs/go-ipfs)
- **Rollkit:** [rollkit/rollkit](https://github.com/rollkit/rollkit)

---

## Next Steps
- Implement real integrations for each stubbed module
- Expand SDK to Rust, Solidity, Web3.js (see below)

---

## Rust & Solidity SDKs (Planned)
- Use gRPC/protobuf codegen for Rust ([tonic](https://github.com/hyperium/tonic))
- Use OpenAPI or protobuf for Solidity/Web3.js bindings
- See `proto/` for message definitions

---

For architecture and inter-module communication, see the main [ARCHITECTURE.md](../docs/ARCHITECTURE.md). 

The Data Availability module now has real business logic:

- **Blobs are stored in-memory** and indexed by their SHA-256 hash.
- **PostBlob** returns the real blob ID (hash).
- **SampleBlob** returns the actual stored data for a given blob ID.

**You can now:**
1. Post a blob:
   ```sh
   go run sdk/main.go -cmd da-postblob -addr :50052 -data "hello world"
   ```
   (This will return a blob ID, e.g., `09ca7e4eaa6e8ae9c7d261167129184883644d0745d6c2d3...`)

2. Sample the blob:
   ```sh
   go run sdk/main.go -cmd da-postblob -addr :50052 -data "hello world"
   # Copy the returned blob ID, then:
   go run sdk/main.go -cmd da-postblob -addr :50052 -data "hello world"
   # (For sampling, you would use a new CLI command to call SampleBlob by ID; let me know if you want this added!)
   ```

---

**Next, I can:**
- Add a CLI command for sampling blobs by ID.
- Upgrade the Consensus module to track real block height and status.
- Add tests or further business logic.

Let me know your preference, or just say "proceed" to continue with the next logical step! 

You now have a CLI command for sampling blobs by ID:

- **Post a blob:**
  ```sh
  go run sdk/main.go -cmd da-postblob -addr :50052 -data "hello world"
  ```
  (Copy the returned blob ID.)

- **Sample the blob:**
  ```sh
  go run sdk/main.go -cmd da-sampleblob -addr :50052 -blobid <BLOB_ID>
  ```

---

**Next, I will upgrade the Consensus module to track real block height and status.**  
After that, I can add tests or further business logic as you prefer.

Would you like to proceed with the Consensus module upgrade, or focus on something else? 
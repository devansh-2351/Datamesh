# DataMesh Architecture

## Overview
DataMesh is a modular Layer 1 blockchain that separates Consensus, Data Availability, and Execution into distinct, pluggable modules. This enables scalable, flexible, and secure infrastructure for data-intensive decentralized applications.

## Core Modules

### 1. Consensus Layer
- **Technology:** Tendermint/CometBFT
- **Responsibilities:**
  - Block production and finality
  - Validator set management
  - Staking and slashing
  - Shared security (optional)

### 2. Data Availability Layer
- **Technology:** Erasure coding (Reed-Solomon), KZG commitments
- **Responsibilities:**
  - Storing and distributing transaction/data blobs
  - Providing proofs of data availability
  - Data Availability Sampling (DAS) for light clients
  - Blobstream for rollups

### 3. Execution Layer(s)
- **Technology:** EVM (Ethermint), WASM (CosmWasm), MoveVM, zkVM (future)
- **Responsibilities:**
  - State transition execution
  - Pluggable VM support
  - Fraud proof integration for rollups

### 4. SDK
- **Technology:** Go or Rust
- **Responsibilities:**
  - CLI and codegen for appchains/rollups
  - Protobuf definitions for inter-module communication
  - Developer utilities

## Inter-Module Communication
- **Protobuf** is used for defining messages and APIs between modules.
- **gRPC** or ABCI++ (Tendermint) for module interaction.

## Security Features
- Data Availability Sampling (DAS)
- Fraud Proofs
- Shared Security (staking layer)

## Extensibility
- New execution environments can be added as modules.
- Rollups can plug into DA and consensus layers via SDK.

## Diagram
```
flowchart TD
  Consensus <--> DataAvailability
  DataAvailability <--> Execution
  SDK --> Consensus
  SDK --> DataAvailability
  SDK --> Execution
```

---

For more details, see each module's README. 
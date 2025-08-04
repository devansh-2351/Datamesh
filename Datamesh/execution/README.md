# Execution Module

## Purpose
The Execution module is responsible for processing transactions, executing state transitions, and supporting multiple virtual machines (VMs) for diverse application needs.

## Technology
- **EVM (Ethermint)**: Ethereum-compatible execution environment.
- **WASM (CosmWasm)**: WebAssembly-based smart contracts.
- **MoveVM, zkVM**: Future/pluggable VMs for advanced use cases.

## Responsibilities
- State transition execution
- Pluggable VM support
- Fraud proof integration for rollups

## Initial Implementation Plan
1. Integrate Ethermint (EVM on Tendermint) for initial execution.
2. Stub out WASM/zkVM/MoveVM for future expansion.
3. Define interfaces for pluggable execution environments.

---

For architecture and inter-module communication, see the main [ARCHITECTURE.md](../docs/ARCHITECTURE.md). 
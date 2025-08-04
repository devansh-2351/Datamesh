# Proto Definitions

## Purpose
This directory contains Protobuf definitions for messages and APIs used in inter-module communication across DataMesh components.

## Responsibilities
- Define messages for consensus, data availability, and execution modules
- Enable gRPC/ABCI++ communication between modules

## Initial Plan
1. Define base messages for block, transaction, and blob data
2. Specify APIs for validator, DA, and execution interactions
3. Generate code for Go/Rust SDKs

## Compiling Protobufs
To generate Go code from .proto files, use:

```
protoc --go_out=. --go-grpc_out=. *.proto
```

To generate Rust code (using tonic):
```
protoc --tonic_out=. *.proto
```

To generate Solidity/Web3.js bindings, use OpenAPI or protobuf-to-solidity tools.

---

For architecture and inter-module communication, see the main [ARCHITECTURE.md](../docs/ARCHITECTURE.md). 
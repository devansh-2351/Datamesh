package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/datamesh-labs/datamesh/consensus"
	"github.com/datamesh-labs/datamesh/data_availability"
	"github.com/datamesh-labs/datamesh/execution"
)

func main() {
	cmd := flag.String("cmd", "", "SDK command: scaffold, run, status, run-consensus, run-da, run-execution, consensus-status, da-postblob, da-sampleblob, execution-execute, consensus-add-validator, consensus-remove-validator, consensus-list-validators")
	addr := flag.String("addr", ":50051", "gRPC server address")
	data := flag.String("data", "", "Data for DA or execution commands (as string)")
	blobID := flag.String("blobid", "", "Blob ID for DA sampleblob command")
	id := flag.String("id", "", "Validator ID for consensus commands")
	pubkey := flag.String("pubkey", "", "Validator pubkey for consensus-add-validator")
	flag.Parse()

	switch *cmd {
	case "scaffold":
		fmt.Println("[SDK] Scaffolding new appchain/rollup (stub)")
	case "run":
		fmt.Println("[SDK] Running DataMesh node (stub)")
	case "status":
		fmt.Println("[SDK] Node status: (stub)")
	case "run-consensus":
		log.Fatal(consensus.RunConsensusGRPCServer(*addr))
	case "run-da":
		log.Fatal(data_availability.RunDAGRPCServer(*addr))
	case "run-execution":
		log.Fatal(execution.RunExecutionGRPCServer(*addr))
	case "consensus-status":
		callConsensusStatus(*addr)
	case "da-postblob":
		callDAPostBlob(*addr, []byte(*data))
	case "da-sampleblob":
		callDASampleBlob(*addr, *blobID)
	case "execution-execute":
		callExecutionExecuteTx(*addr, []byte(*data))
	case "consensus-add-validator":
		callConsensusAddValidator(*addr, *id, *pubkey)
	case "consensus-remove-validator":
		callConsensusRemoveValidator(*addr, *id)
	case "consensus-list-validators":
		callConsensusListValidators(*addr)
	default:
		fmt.Println("[SDK] Unknown command. Use -cmd scaffold|run|status|run-consensus|run-da|run-execution|consensus-status|da-postblob|da-sampleblob|execution-execute")
		os.Exit(1)
	}
} 
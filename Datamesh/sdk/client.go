package main

import (
	"context"
	"fmt"
	"log"
	"time"

	datameshpb "github.com/datamesh-labs/datamesh/proto"
	"google.golang.org/grpc"
)

func callConsensusStatus(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewConsensusServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.GetStatus(ctx, &datameshpb.StatusRequest{})
	if err != nil {
		log.Fatalf("Consensus GetStatus error: %v", err)
	}
	fmt.Printf("Consensus status: %s\n", resp.Status)
}

func callDAPostBlob(addr string, data []byte) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewDAServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.PostBlob(ctx, &datameshpb.PostBlobRequest{Data: data})
	if err != nil {
		log.Fatalf("DA PostBlob error: %v", err)
	}
	fmt.Printf("DA PostBlob ID: %s\n", resp.Id)
}

func callExecutionExecuteTx(addr string, tx []byte) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewExecutionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.ExecuteTx(ctx, &datameshpb.ExecuteTxRequest{Tx: tx})
	if err != nil {
		log.Fatalf("Execution ExecuteTx error: %v", err)
	}
	fmt.Printf("Execution result: %s\n", string(resp.Result))
}

func callDASampleBlob(addr string, id string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewDAServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.SampleBlob(ctx, &datameshpb.SampleBlobRequest{Id: id})
	if err != nil {
		log.Fatalf("DA SampleBlob error: %v", err)
	}
	fmt.Printf("DA SampleBlob Data: %s\n", string(resp.Data))
}

func callConsensusAddValidator(addr, id, pubkey string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewConsensusServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.AddValidator(ctx, &datameshpb.AddValidatorRequest{
		Validator: &datameshpb.Validator{Id: id, Pubkey: pubkey},
	})
	if err != nil {
		log.Fatalf("AddValidator error: %v", err)
	}
	fmt.Printf("AddValidator success: %v\n", resp.Success)
}

func callConsensusRemoveValidator(addr, id string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewConsensusServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.RemoveValidator(ctx, &datameshpb.RemoveValidatorRequest{Id: id})
	if err != nil {
		log.Fatalf("RemoveValidator error: %v", err)
	}
	fmt.Printf("RemoveValidator success: %v\n", resp.Success)
}

func callConsensusListValidators(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := datameshpb.NewConsensusServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.ListValidators(ctx, &datameshpb.ListValidatorsRequest{})
	if err != nil {
		log.Fatalf("ListValidators error: %v", err)
	}
	fmt.Println("Validators:")
	for _, v := range resp.Validators {
		fmt.Printf("- ID: %s, Pubkey: %s\n", v.Id, v.Pubkey)
	}
} 
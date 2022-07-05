package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"peg-go-client"
	operationv1 "peg-go-client/gen/operation/v1"
	"peg-go-client/gen/operation/v1/operationv1connect"
	typesv1 "peg-go-client/gen/types/v1"
)

func main() {
	// grpc用の環境変数を読み込みます
	log.Printf("Load env file for grpc: %s\n", ".env.local")
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("failed to load env file: %v", err)
	}

	// アカウントの環境変数を読み込みます
	envFile := flag.String("env", ".env.local", "path of env file, default is .env.local")
	flag.Parse()
	log.Printf("Load env file for auth: %s\n", *envFile)
	if err := godotenv.Load(*envFile); err != nil {
		log.Fatalf("failed to load env file: %v", err)
	}

	// firebaseからトークンを取得します
	idToken, uid, err := peg_go_client.GetAuthToken(os.Getenv("PEG_EMAIL"), os.Getenv("PEG_PASSWORD"))
	if err != nil {
		log.Fatalf("failed to get idToken: %v", err)
	}
	log.Printf("welcome: %v\n", uid)
	log.Printf("idToken: %v\n", "***")

	// grpc clientを準備
	addr := fmt.Sprintf("http://%s:%s", os.Getenv("OPERATION_SERVICE_HOST"), os.Getenv("OPERATION_SERVICE_PORT"))
	client := operationv1connect.NewOperationServiceClient(
		http.DefaultClient,
		addr,
	)

	// ストリーミング
	req := connect.NewRequest(&operationv1.FetchOperationsRequest{OnsetOperationId: 0})
	req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", idToken))
	stream, err := client.FetchOperations(
		context.Background(),
		req,
	)
	if err != nil {
		log.Fatalf("failed to fetch ops: %v", err)
	}
	for stream.Receive() {
		fmt.Println()
		op := stream.Msg().Operation
		log.Printf("ReceiveOp: %v\n", op)
		if op.Type == typesv1.OperationType_OPERATION_TYPE_SEND_MESSAGE || op.Type == typesv1.OperationType_OPERATION_TYPE_SEND_MESSAGE_RECV {
			log.Printf("Msg: %v\n", stream.Msg().Message)
		}
		fmt.Println()
	}
	if err := stream.Err(); err != nil {
		log.Fatalf("stream error: %v", err)
	}

}

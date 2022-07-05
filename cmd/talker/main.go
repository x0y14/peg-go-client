package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	peg_go_client "peg-go-client"
	talkv1 "peg-go-client/gen/talk/v1"
	"peg-go-client/gen/talk/v1/talkv1connect"
	typesv1 "peg-go-client/gen/types/v1"
	"strings"
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
	addr := fmt.Sprintf("http://%s:%s", os.Getenv("TALK_SERVICE_HOST"), os.Getenv("TALK_SERVICE_PORT"))
	client := talkv1connect.NewTalkServiceClient(
		http.DefaultClient,
		addr,
	)

	var receiver string

	for {
		fmt.Printf("receiver: %v\n", receiver)
		fmt.Printf("you(%v)>", uid)
		scanner := bufio.NewScanner(os.Stdin)
	inputLoop:
		for {
			scanner.Scan()
			in := scanner.Text()
			switch {
			case in == "exit", in == "quit":
				log.Printf("終了します\n")
				os.Exit(0)
			case in == "":
				log.Printf("文字を入力してください\n")
				break inputLoop
			case strings.Contains(in, "set:"):
				//
				receiver = strings.Replace(in, "set:", "", 1)
				log.Printf("receiverを%vに設定しました\n", receiver)
				break inputLoop
			default:
				if receiver == "" {
					log.Println("set:$receiverIdで送信先を設定してください")
					break inputLoop
				}
				req := connect.NewRequest(&talkv1.SendMessageRequest{Message: &typesv1.Message{
					Id:          "",
					From:        "",
					To:          receiver,
					ContentType: typesv1.MessageContentType_MESSAGE_CONTENT_TYPE_TEXT,
					Text:        in,
					Metadata:    "{}",
				}})
				req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", idToken))
				_, err := client.SendMessage(
					context.Background(),
					req,
				)
				if err != nil {
					log.Printf("failed to send message: %v", err)
				}
				break inputLoop
			}
		}
		fmt.Println()
	}

}

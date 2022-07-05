# peg-go-client
guiを伴わなず、情報取得と送信を分離した動作確認用クライアント

.env.localをプロジェクトルートに作成し、以下の4点を記載してください。  
```dotenv
OPERATION_SERVICE_HOST=<host># e.g. 127.0.0.1
OPERATION_SERVICE_PORT=<port># e.g. 8080
TALK_SERVICE_HOST=<host># e.g. 127.0.0.1
TALK_SERVICE_PORT=<port># e.g. 8081
```

同様に.env.*.localを作成しアカウントの情報を記載してください。  
(.env.localに記載しても問題はない。)  
クライアントのリクエスト検証に使用します。
- PEG_EMAIL  
アカウント作成時に設定したメールアドレス
- PEG_PASSWORD  
アカウント作成時に設定した生のパスワード

```dotenv
# e.g. filename: .env.account1.local 
PEG_EMAIL="test@test.test"
PEG_PASSWORD="test-test-test"
```

---
以下を実行することにより`gen/`にbufにより生成されたprotocol bufファイル群が生成されます。
```shell
buf generate
```

buf.yaml, buf.gen.yamlの書き方は以下をご覧ください.
- [buf.gen.yaml](https://docs.buf.build/configuration/v1/buf-gen-yaml#plugins)
- [13 Use managed mode](https://docs.buf.build/tour/use-managed-mode)

---
以下で、operationを取得し続けるプロセスを起動できます。  
(バックグラウンドで走るプロセスを想定)
```shell
# PEG_EMAIL, PEG_PASSWORDを記載したenvファイルを指定してください。
go run cmd/fetcher/main.go 
# or
go run cmd/fetcher/main.go --env=${envファイル名}
# ${envファイル名}は実際のファイル名と置き換えてください。
```

---
以下でメッセージ送信を可能とするプロセスを起動できます  
(トーク画面のメッセージフォーム、送信ボタンを想定)
```shell
# PEG_EMAIL, PEG_PASSWORDを記載したenvファイルを指定してください。
go run cmd/talker/main.go
# or
go run cmd/talker/main.go --env=${envファイル名}
# ${envファイル名}は実際のファイル名と置き換えてください。
```
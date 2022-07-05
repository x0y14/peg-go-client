# peg-go-client
guiを伴わなず、情報取得と送信を分離した動作確認用クライアント

.env.localをプロジェクトルートに作成し、以下の二点を記載してください。  
複数のアカウントを使用する場合は、.env.account1.local, .env.account2.localのようにファイルを分けてください。   
クライアントのリクエスト検証に使用します。
- PEG_EMAIL  
アカウント作成時に設定したメールアドレス
- PEG_PASSWORD  
アカウント作成時に設定した生のパスワード

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
# envが指定されていなかった場合は.env.localが読み込まれます。
go run cmd/fetcher/main.go 
# or
go run cmd/fetcher/main.go --env=${envファイル名}
# ${envファイル名}は実際のファイル名と置き換えてください。
```

---
以下でメッセージ送信を可能とするプロセスを起動できます  
(トーク画面のメッセージフォーム、送信ボタンを想定)
```shell
# envが指定されていなかった場合は.env.localが読み込まれます。
go run cmd/talker/main.go
# or
go run cmd/talker/main.go --env=${envファイル名}
# ${envファイル名}は実際のファイル名と置き換えてください。
```
# peg-go-client
guiを伴わなず、情報取得と送信を分離した動作確認用クライアント

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
(mobile clientなどではバックグラウンドで動かす物を想定)
```shell
go run cmd/fetcher/main.go
```

---
以下でメッセージ送信を可能とするプロセスを起動できます
```shell
go run cmd/talker/main.go
```
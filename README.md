# simple-http-server
言語学習用のシンプルなhttpサーバの実装です。[todokr/simple-http-server](https://github.com/todokr/simple-http-server)をもとにしています。

## 仕様
仕様は[todokr/simple-http-server](https://github.com/todokr/simple-http-server)と同じです。

- localhost:8080で待ち受け、HTTPリクエストを受けとり、HTTPレスポンスを返す
- 対応するHTTPリクエストメソッドはGETのみ
- リソースのMIMEは外部ファイルで設定できる
- リクエストをブロックしない（マルチスレッド方式）
- Keep-Aliveはしない
- HTTP Cacheはしない

## 実装で行うこと
- 
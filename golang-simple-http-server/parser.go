package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

const PATTERN = `(?P<method>[A-Z]+) (?P<path>[^ ]+) (?P<httpVersion>[^ ]+)`

func Parse(conn net.Conn) (request Request, err error) {
	reader := bufio.NewReader(conn)

	// リクエストを読み込む
	// 実際は二番目の返り値でisPrefixが返ってくる。これがtrueだと行がデフォルトバッファより長いので、さらに読み込まないといけない。
	// ただ今回は一行目が短いことがわかっているので無視している。
	reqLine, _, err := reader.ReadLine()
	if err != nil {
		fmt.Printf("Failed to read request: %s", err)
		return Request{}, err
	}

	// リクエストをパースする
	// TODO: 不便なのでbufioを噛ませる
	re, err := regexp.Compile(PATTERN)
	if (err != nil) {
		fmt.Printf("Failed to parse regex pattern: %s", err)
		os.Exit(1)
	}

	// 正規表現にグループを使ったときは、FindStringSubmatchを使うことで、グループごとのmatchを配列で取れる
	match := re.FindStringSubmatch(string(reqLine))

	return Request{
		method: string(match[1]),
		path: string(match[2]),
		httpVersion: string(match[3]),
	}, nil
}

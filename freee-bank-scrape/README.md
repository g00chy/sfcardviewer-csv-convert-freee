# Freee対応のバンク一覧抽出

GOでスクレピングの練習として本プロジェクトを作成  
[https://secure.freee.co.jp/walletables/sync_bank_list#credit_cards](対象)となるページから取得したhtmlファイルから各データを抽出し、CSVにて出力を行う。

## 動作

上記URLのHTMLをindex.htmlとしてダウンロード  

```shell
go mod vendor
go run main.go
```

## 参考サイト

[Go言語でのファイル読み取り \- Qiita](https://qiita.com/Kashiwara/items/9a8365ea800e6f39713f)  
[【Go】Webスクレイピングのやり方 \- Qiita](https://qiita.com/kou_pg_0131/items/dab4bcbb1df1271a17b6)  


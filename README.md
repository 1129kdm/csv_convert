# csv_convert
csv convert by golang.

## Quick start
### initialize
- 変換するCSVファイルと同じ階層にgoファイルを設置してください

```
cd /path/to/directory # anywhere you want
git clone git@github.dip-net.co.jp:s-kodama/csv_convert.git # clone app
```

### build

```
go build convert.go
```

### execute
- 変換するCSVファイル名を引数にして実行してください

```
./convert [args 1] [args 2] ... [args n]
```

- example

```
./convert test.csv test2.csv test3.csv
```

- go run でも実行可能

```
go run convert.go [args 1] ... [args n]
```


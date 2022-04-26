# go_hash_vs_crypt
MD5とAES128 CBC pkcs7 Decryptのベンチ検証

正確には以下のような追加条件あり
- MD5
    - 文字列からJSON構造をパース
    - jSON構造からMD5を作成
- AES128
    - byteで入力
    - Decrypt後の文字列をJSONとしてパースする

サーバーサイドのエンドポイントに、簡単な検証機能として採用する場合、どちらが最適な手法か？を検証してみた。

## 比較結果

| 項目          | 比率 hash : crypt |
|---------------|-------------------|
| ns/op         | 1 : 1.1           |
| mem Byte/op   | 1 : 2.5           |
| mem allocs/op | 1 : 1.6           |

```
$ go test -bench . -benchmem -count 10
goos: linux
goarch: amd64
pkg: example.com/hash_vs_crypt
cpu: 11th Gen Intel(R) Core(TM) i7-11700K @ 3.60GHz
BenchmarkChecksum-16              554444              2065 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              568454              2067 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              558429              2059 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              535000              2050 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              560889              2054 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              571147              2080 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              542085              2052 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              566052              2057 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              531561              2055 ns/op             568 B/op          9 allocs/op
BenchmarkChecksum-16              562804              2054 ns/op             568 B/op          9 allocs/op
BenchmarkDecrypt-16               521656              2236 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               512527              2245 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               528735              2280 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               525717              2250 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               517525              2222 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               523347              2235 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               512638              2232 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               524133              2218 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               529066              2221 ns/op            2488 B/op         15 allocs/op
BenchmarkDecrypt-16               503889              2227 ns/op            2488 B/op         15 allocs/op
PASS
ok      example.com/hash_vs_crypt       23.564s
```

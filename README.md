# yaml2json-deno

DenoからGo製Wasmを呼び出し、標準入力に入ったyamlファイルをjsonファイルにして標準出力に出します。
パフォーマンスはdenoのstdに含まれるYAMLモジュールに比べて10倍ほど悪いです。

$ cat hogehoge.yaml | deno run -A main.js > hogehoge.json

40万行ほどのyamlを処理するためにWSL上の手元のRYZEN7 2700マシンでは13秒ほどかかりました。

ライセンス
主張しません
wasm_exec.jsのBSDライセンスに準拠します
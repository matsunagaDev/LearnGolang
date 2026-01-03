## godoc コマンドが使えない

ターミナルで`godoc`コマンドを使用したが、使えなかった。
正しくは以下の記法で実行する。

```diff
- godoc fmt Println
+ go doc fmt Println
```

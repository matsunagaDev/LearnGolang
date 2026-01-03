## コンパイル

```zsh
go run main.go
```

## godoc コマンドが使えない

ターミナルで`godoc`コマンドを使用したが、使えなかった。
正しくは以下の記法で実行する。

```diff
- godoc fmt Println
+ go doc fmt Println
```

## var と ショートの違い

var 関数は関数の外に出しても実行できる

```go
package main

import (
	"fmt"
)

var (
		i int
		f64 float64
		 s string
		 t, f bool
)

func foo(){
	xi := 1
	xi = 2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T", xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
```

## const

- 型を定義しない

## 配列とスライスの違い

- 配列は拡張できない
- スライスは拡張できる(Java の list と同じ考え方)

```go
// var b [2]int = [2]int{100, 200}
	// b.append(b, 300)
	// fmt.Println(b)

	// slice
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
```

## デバッガを使えるのは 1.24 が最新

```
brew install go@1.24
```

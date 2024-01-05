package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	// 使用这种方式其实是在遍历底层的byte数组，所以当我们将byte转化成rune的时候，只有ascii code能被成功转换
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}

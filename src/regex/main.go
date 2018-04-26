// regex project main.go
package main

import (
	"fmt"
	"regexp"
)

const text = `My email is talenthb8761@msn.com
lfsajlkjdfs@jkasfdjk.com
jasdflj@jasjf.com
asljkfdlj@adfs.com.cn`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}

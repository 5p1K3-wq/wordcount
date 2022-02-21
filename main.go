package main

import (
	"fmt"
	"github.com/5p1K3-wq/wordcount/word"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args[1:], " ")
	/*
		s, err := readString()
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	count := word.Count(s)
	fmt.Println(count)
}

/*
func readString() (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
*/

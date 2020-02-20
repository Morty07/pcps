package pcpsd

import (
	"fmt"
	"testing"
)

func TestReadFileByLine(t *testing.T) {
	result := ParseResult{}
	s, err := ReadFileByLine(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(result)
}

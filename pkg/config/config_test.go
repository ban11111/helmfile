package config

import (
	"fmt"
	"testing"
)

func Test_parseCommaSeparatedString(t *testing.T) {
	test := `a=1,b=2\,3,c=9\,10\,11`
	result := parseCommaSeparatedString(test)
	fmt.Println(result)
}

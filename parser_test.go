package parser

import (
	"fmt"
	"testing"
)

func TestCases(t *testing.T) {

	fmt.Println("(){}[]", ContainerParser("(){}[]"))
	fmt.Println("((){}[])", ContainerParser("((){}[])"))
	fmt.Println("(])", ContainerParser("(])"))
	fmt.Println("([)]", ContainerParser("([)]"))
	fmt.Println("{[}", ContainerParser("{[}"))
	fmt.Println("{", ContainerParser("{"))
	fmt.Println("{})", ContainerParser("{})"))
	fmt.Println("{", ContainerParser("{"))
	fmt.Println("]", ContainerParser("]"))
	fmt.Println("([{}])", ContainerParser("([{}])"))

}

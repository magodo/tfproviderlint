package a

import "fmt"

func foo() {
	// lintignore:a
	fmt.Println("Hello") // want `a ignored`
}

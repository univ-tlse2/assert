package assert

import "fmt"

func Assert(test bool, mess string) {
	defer noMessages()
	if !test {
		panic(mess)
	}
}

func noMessages() {
	if r := recover(); r != nil {
		fmt.Println("Assert failed :", r)
	}
}

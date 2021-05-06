package assert

import "fmt"

// Assert vérifie que le test réussit ou panique et affiche le message dans le cas contraire
func Assert(test bool, messOk, messNok string) {
	defer noMessages()
	if !test {
		panic(messNok)
	} else {
		fmt.Println(messOk)
	}
}

func noMessages() {
	if r := recover(); r != nil {
		fmt.Println("Assert failed :", r)
	}
}

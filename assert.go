package assert

import "fmt"

// Assert vérifie que le test réussit ou panique et affiche le message dans le cas contraire
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

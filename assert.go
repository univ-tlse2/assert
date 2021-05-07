package assert

import "fmt"

// Assert vérifie que le test réussit ou panique et affiche le message dans le cas contraire
func Assert(test bool, messOk, messNok string) {
	defer noMessages()
	if !test {
		panic(messNok)
	} else {
		if messOk != "" {
			fmt.Println("Succès :", messOk)
		}
	}
}

func noMessages() {
	if r := recover(); r != nil {
		if r != "" {
			fmt.Println("Échec :", r)
		}
	}
}

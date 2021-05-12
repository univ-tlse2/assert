// Package assert ajoute une fonction Assert pour tester une condition
package assert

import "fmt"

// Assert vérifie que le test réussit ou panique et affiche le message dans le cas contraire
func Assert(test bool, messOk, messNok string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Échec :", r)
		}
	}()

	if !test {
		panic(messNok)
	} else {
		if messOk != "" {
			fmt.Println("Succès :", messOk)
		}
	}
}

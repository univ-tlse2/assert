package main

import (
	"github.com/univ-tlse2/assert"
	"strconv"
	"strings"
)

func main() {
	val := 20 * 2
	assert.Assert(val == 42, "val n'a pas la bonne valeur : "+strconv.Itoa(val))

	mess := "coucou"
	assert.Assert(strings.HasPrefix(mess, "cou"), "mess n'a pas la bonne valeur")
}

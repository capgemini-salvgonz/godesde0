package ejerinterfaces

import (
	"fmt"

	"github.com/chava.gnolasco/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando\n", hu.Gender())
}

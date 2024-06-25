package models

type Mujer struct {
	Hombre
}

func (h *Mujer) Respirar() {
	h.respirando = true
}

func (h *Mujer) Pensar() {
	h.pensando = true
}

func (h *Mujer) Comiendo() {
	h.comiendo = true
}

func (h *Mujer) Gender() string {
	return "Mujer"
}

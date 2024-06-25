package models

type Hombre struct {
	edad       int
	altura     float64
	peso       float64
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}

func (h *Hombre) Pensar() {
	h.pensando = true
}

func (h *Hombre) Comer() {
	h.comiendo = true
}

func (h *Hombre) Gender() string {
	return "Hombre"
}

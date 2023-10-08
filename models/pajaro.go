package models

type Pajaro struct {
	PosicionX int
	PosicionY int
}

func NuevoPajaro() *Pajaro {
	return &Pajaro{
		PosicionX: 50,
		PosicionY: 50,
	}
}

func (p *Pajaro) MoverDerecha() {
	p.PosicionX += 80
}

func (p *Pajaro) MoverIzquierda() {
	p.PosicionX -= 80
}

func (p *Pajaro) MoverArriba() {
	p.PosicionY -= 60
}

func (p *Pajaro) MoverAbajo() {
	p.PosicionY += 60
}
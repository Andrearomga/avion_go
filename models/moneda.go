package models

import (
	
	"math/rand" 
	"time"
)

type Moneda struct {
	PosicionX int
	PosicionY int
	Color     string
}

// constructor que se utiliza para crear una nueva instancia de la estructura Moneda
func NuevaMoneda(x, y int, color string) *Moneda {
	return &Moneda{PosicionX: x, PosicionY: y, Color: color}
}
//Este es un método de la estructura Moneda. 
func (m *Moneda) IniciarHilo() {
	go func() {
		for {
			// Hace que la moneda desaparezca durante 1 segundo.
			m.PosicionX = -1000 
			time.Sleep(time.Second)

			// Coloca la moneda en una posición aleatoria y la muestra durante 3 segundos.
			m.PosicionX = rand.Intn(600)
			m.PosicionY = rand.Intn(400)
			time.Sleep(time.Second * 3)
		}
	}()
}

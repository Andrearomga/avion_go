package models

import (
	
	"time"
)

type Nube struct {
	PosicionX int
	PosicionY int
}

//esta función es un constructor que crea y devuelve una nueva instancia de Nube.
func NuevaNube(x, y int) *Nube {
	return &Nube{PosicionX: x, PosicionY: y}
}


func (n *Nube) IniciarHilo() {
	go func() {
		for {
			// Mueve la nube de izquierda a derecha.
			for i := 0; i < 200; i++ {
				n.PosicionX++
				time.Sleep(time.Millisecond * 80)  //simula el movimiento lento
			}

			// Mueve la nube de derecha a izquierda.
			for i := 0; i < 200; i++ {
				n.PosicionX--
				time.Sleep(time.Millisecond * 80) 
			}
		}
	}()
}

package models

import (
	"time"
	"math/rand" 
)

type Obstaculo struct {
	PosicionX int
	PosicionY int
	Color     string
}

func NuevoObstaculo(x, y int, color string) *Obstaculo {
	return &Obstaculo{PosicionX: x, PosicionY: y, Color: color}
}

func (o *Obstaculo) IniciarHilo() {
	go func() {
		for {
			// Mueve el obstáculo de izquierda a derecha.
			for i := 0; i < 500; i++ {
				o.PosicionX++
				time.Sleep(time.Millisecond * 10) 
			}

			// Hace que el obstáculo desaparezca durante 1 segundo.
			o.PosicionX = -100
			time.Sleep(time.Second)

			// Coloca el obstáculo en una posición aleatoria y lo muestra durante 3 segundos.
			o.PosicionX = rand.Intn(600)
			o.PosicionY = rand.Intn(400)
			time.Sleep(time.Second * 3)
		}
	}()
}
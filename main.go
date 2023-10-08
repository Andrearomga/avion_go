package main

import (
	"fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2"
	"jueguito/views" 
)

func main() {
    aplicacion := app.New()
    window := aplicacion.NewWindow("Juego del Avioncito")
    window.Resize(fyne.NewSize(950, 300))
    views.IniciarVentana(window)
    window.ShowAndRun()
}

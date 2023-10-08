package views

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "math/rand"
    "time"
    "jueguito/models"    
    
)

func crearNubes() ([]*models.Nube, []*canvas.Image) {
    var nubes []*models.Nube
    var nubesImagenes []*canvas.Image
    posicionesNubes := []struct{x, y int}{
        {100, 50},
        {200, 70},
        {300, 30},
        {400, 50},
        {500, 70},
        {600, 90},
       
    }

    for _, pos := range posicionesNubes {
        nube := models.NuevaNube(pos.x, pos.y)
        nube.IniciarHilo()
        nubes = append(nubes, nube)
        nubeImagen := canvas.NewImageFromFile("assets/nube.png")
        nubeImagen.FillMode = canvas.ImageFillOriginal
        nubeImagen.Resize(fyne.NewSize(100, 100))
        nubesImagenes = append(nubesImagenes, nubeImagen)
    }

    return nubes, nubesImagenes
}



func IniciarVentana(window fyne.Window) {
    nubes, nubesImagenes := crearNubes()

    pajaro := models.NuevoPajaro()
pajaroImagen := canvas.NewImageFromFile("assets/Avion3.png")
pajaroImagen.FillMode = canvas.ImageFillOriginal
pajaroImagen.Resize(fyne.NewSize(60, 60))


window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
    switch e.Name {
    case fyne.KeyUp:
        pajaro.MoverArriba()
    case fyne.KeyDown:
        pajaro.MoverAbajo()
    case fyne.KeyLeft:
        pajaro.MoverIzquierda()
    case fyne.KeyRight:
        pajaro.MoverDerecha()
    }
  
})
 

    var monedas []*models.Moneda
    var monedasImagenes []*canvas.Image
    for i := 0; i < 25; i++ {
        x := rand.Intn(600)
        y := rand.Intn(400)
        color := "gold" 
        moneda := models.NuevaMoneda(x, y, color)
        moneda.IniciarHilo()
        monedas = append(monedas, moneda)
        monedaImagen := canvas.NewImageFromFile("assets/moneda.png")
        monedaImagen.FillMode = canvas.ImageFillOriginal
        monedaImagen.Resize(fyne.NewSize(30, 30))
        monedasImagenes = append(monedasImagenes, monedaImagen)
    }

    var obstaculos []*models.Obstaculo
    var obstaculosImagenes []*canvas.Image
    for i := 0; i < 10; i++ {
        x := rand.Intn(300)
        y := rand.Intn(200)
        color := "red" 
        obstaculo := models.NuevoObstaculo(x, y, color)
        obstaculo.IniciarHilo()
        obstaculos = append(obstaculos, obstaculo)
        obstaculoImagen := canvas.NewImageFromFile("assets/obstaculo.png")
        obstaculoImagen.FillMode = canvas.ImageFillOriginal
        obstaculoImagen.Resize(fyne.NewSize(70, 70))
        obstaculosImagenes = append(obstaculosImagenes, obstaculoImagen)
    }



    fondo := canvas.NewImageFromFile("assets/fondo1.png")
    fondo.FillMode = canvas.ImageFillOriginal
    fondo.Resize(fyne.NewSize(950,350))


content := fyne.NewContainerWithoutLayout(fondo, pajaroImagen)

    
    for _, monedaImagen := range monedasImagenes {
        content.Add(monedaImagen)
    }
    for _, obstaculoImagen := range obstaculosImagenes {
        content.Add(obstaculoImagen)
    }
    for _, nubeImagen := range nubesImagenes {
        content.Add(nubeImagen)
    }
   
    
    go func() {
        for {
           
            pajaroImagen.Move(fyne.NewPos(float32(pajaro.PosicionX), float32(pajaro.PosicionY)))
            for i, moneda := range monedas {
                // Actualiza la posición de la imagen de la moneda.
                monedasImagenes[i].Move(fyne.NewPos(float32(moneda.PosicionX), float32(moneda.PosicionY)))
            }
            for i, obstaculo := range obstaculos {
                // Actualiza la posición de la imagen del obstáculo.
                obstaculosImagenes[i].Move(fyne.NewPos(float32(obstaculo.PosicionX), float32(obstaculo.PosicionY)))
            }
            for i, nube := range nubes {
                nubesImagenes[i].Move(fyne.NewPos(float32(nube.PosicionX), float32(nube.PosicionY)))
            }
    
            time.Sleep(time.Millisecond * 30) // espera solo 100 milisegundos
           
            content.Refresh()
        }
    }()
    
    
    window.SetContent(content)
}
package main

import (
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"simulador-estacionamiento/models"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simulador de Estacionamiento con Animación")

	var wg sync.WaitGroup

	estacionamiento := models.NuevoEstacionamiento(20)

	title := canvas.NewText("Simulador de Estacionamiento", color.NRGBA{R: 34, G: 34, B: 34, A: 255})
	title.TextSize = 28
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	label := canvas.NewText("Espacios disponibles: 20", color.NRGBA{R: 34, G: 34, B: 34, A: 255})
	label.TextSize = 18
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Italic: true}

	respawnBackground := canvas.NewRectangle(color.NRGBA{R: 173, G: 216, B: 230, A: 255}) // Azul claro
	respawnBackground.SetMinSize(fyne.NewSize(300, 400))
	respawnLabel := canvas.NewText("Área de Respawn", color.NRGBA{R: 34, G: 34, B: 34, A: 255})
	respawnLabel.TextSize = 16
	respawnLabel.Alignment = fyne.TextAlignCenter

	respawnContainer := container.NewVBox(
		container.NewCenter(respawnLabel),
		respawnBackground,
	)

	parkingBackground := canvas.NewRectangle(color.NRGBA{R: 220, G: 220, B: 220, A: 255}) // Gris claro
	parkingBackground.SetMinSize(fyne.NewSize(500, 400))
	vehiculosContainer := container.NewGridWithColumns(5)
	parkingContainer := container.NewMax(parkingBackground, vehiculosContainer)

	
	splitContainer := container.NewHSplit(respawnContainer, parkingContainer)
	splitContainer.SetOffset(0.3)

	labelUpdater := func(text string) {
		label.Text = text
		label.Color = color.RGBA{R: 255, G: 69, B: 0, A: 255}
		label.Refresh()

		time.AfterFunc(500*time.Millisecond, func() {
			label.Color = color.NRGBA{R: 34, G: 34, B: 34, A: 255}
			label.Refresh()
		})
	}

	startButton := widget.NewButtonWithIcon("Iniciar Simulación", theme.MediaPlayIcon(), func() {
		go estacionamiento.IniciarSimulacion(labelUpdater, vehiculosContainer, &wg)
	})
	startButton.Importance = widget.HighImportance

	headerContainer := container.NewVBox(
		container.NewCenter(title),
		container.NewCenter(label),
		container.NewCenter(startButton),
	)


	myWindow.SetContent(container.NewBorder(headerContainer, nil, nil, nil, splitContainer))
	myWindow.ShowAndRun()
}

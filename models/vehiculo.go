package models

import (
	"fmt"
	"image/color"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Vehiculo struct {
	ID                 int
	Estacionamiento    *Estacionamiento
	VehiculosContainer *fyne.Container
	UpdateLabel        func(string)
	Wg                 *sync.WaitGroup
}

func NewVehiculo(id int, estacionamiento *Estacionamiento, vehiculosContainer *fyne.Container, updateLabel func(string), wg *sync.WaitGroup) *Vehiculo {
	return &Vehiculo{
		ID:                 id,
		Estacionamiento:    estacionamiento,
		VehiculosContainer: vehiculosContainer,
		UpdateLabel:        updateLabel,
		Wg:                 wg,
	}
}

func (v *Vehiculo) MoverVehiculo() {
	defer v.Wg.Done()

	select {
	case v.Estacionamiento.puerta <- struct{}{}:

		select {
		case v.Estacionamiento.espacios <- struct{}{}:
			v.Estacionamiento.mu.Lock()
			v.Estacionamiento.EspaciosDisponibles--
			v.Estacionamiento.mu.Unlock()
			v.UpdateLabel(fmt.Sprintf("🅿️: %d", v.Estacionamiento.EspaciosDisponibles))
			v.Entrar()
			<-v.Estacionamiento.puerta 

			time.Sleep(time.Duration(rand.Intn(5)+14) * time.Second)

			<-v.Estacionamiento.espacios 
			v.Estacionamiento.liberarEspacio() 
			v.Salir()

		default:
			<-v.Estacionamiento.puerta 
			fmt.Printf("Vehículo %d: No hay espacio disponible, esperando...\n", v.ID)
		}
	default:
		fmt.Printf("Vehículo %d: La puerta está ocupada, esperando...\n", v.ID)
	}
}

func (v *Vehiculo) Entrar() {
	image := canvas.NewImageFromFile("C:/Users/eduar/Downloads/simulador-20241104T145239Z-001/simulador/assets/carro.png")
	image.SetMinSize(fyne.NewSize(100, 50))
	numeroTexto := canvas.NewText(fmt.Sprintf("%d", v.ID), color.Black)
	numeroTexto.TextSize = 14

	vehiculoContainer := container.NewVBox(numeroTexto, image)
	v.VehiculosContainer.Add(vehiculoContainer)

	v.iniciarAnimacion(vehiculoContainer, -200, 300, float32(50+v.ID*10))
}

func (v *Vehiculo) Salir() {
	if len(v.VehiculosContainer.Objects) > 0 {
		vehiculoContainer, ok := v.VehiculosContainer.Objects[0].(*fyne.Container)
		if !ok {
			fmt.Println("Error: No se pudo convertir vehiculoContainer a *fyne.Container")
			return
		}
		v.iniciarAnimacion(vehiculoContainer, vehiculoContainer.Position().X, 800, float32(50+v.ID*10))

		time.AfterFunc(time.Second*2, func() {
			if len(v.VehiculosContainer.Objects) > 0 {
				v.VehiculosContainer.Objects = v.VehiculosContainer.Objects[1:]
				v.VehiculosContainer.Refresh()
			}
		})
	}
}


func (v *Vehiculo) iniciarAnimacion(vehiculoContainer *fyne.Container, startX, endX, y float32) {
	animation := canvas.NewPositionAnimation(
		fyne.NewPos(startX, y),
		fyne.NewPos(endX, y),
		time.Second*2,
		v.actualizarPosicion(vehiculoContainer),
	)
	animation.Start()
}

func (v *Vehiculo) actualizarPosicion(vehiculoContainer *fyne.Container) func(pos fyne.Position) {
	return func(pos fyne.Position) {
		vehiculoContainer.Move(pos)
		vehiculoContainer.Refresh()
	}
}

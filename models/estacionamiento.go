package models

import (
	"fmt"
	"sync"
	"time"
	"fyne.io/fyne/v2"
	"math/rand"
)

type Estacionamiento struct {
	EspaciosDisponibles int
	observadores        []func()
	mu                  sync.Mutex
	puerta              chan struct{}
	espacios            chan struct{}
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
	return &Estacionamiento{
		EspaciosDisponibles: capacidad,
		puerta:              make(chan struct{}, 1),
		espacios:            make(chan struct{}, capacidad),
	}
}

func (e *Estacionamiento) AgregarObservador(f func()) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.observadores = append(e.observadores, f)
}

func (e *Estacionamiento) NotificarObservadores() {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, observador := range e.observadores {
		observador()
	}
}

func (e *Estacionamiento) liberarEspacio() {
	e.mu.Lock()
	e.EspaciosDisponibles++
	e.mu.Unlock()
	e.NotificarObservadores()
}

func (e *Estacionamiento) IniciarSimulacion(updateLabel func(string), vehiculosContainer *fyne.Container, wg *sync.WaitGroup) {
	e.AgregarObservador(func() {
		fmt.Println("Observador notificado: Se ha liberado un espacio.")
	})

	for i := 0; i < 100; i++ {
		wg.Add(1)
		vehiculo := NewVehiculo(i, e, vehiculosContainer, updateLabel, wg)
		go vehiculo.MoverVehiculo()
		time.Sleep(time.Duration(rand.Intn(1000)+900) * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("Simulación finalizada")
}

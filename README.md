# Simulador de Estacionamiento aplicando conocimientos en Fyne

## Introducción
Hola mi nombre es Eduardo Garcia y quería compartir mi experiencia al desarrollar un proyecto en Go usando la biblioteca Fyne. Este proyecto fue un simulador de estacionamiento con animaciones, una práctica que me permitió aprender bastante sobre concurrencia. Espero que este reporte les sea útil si están explorando Fyne o si simplemente buscan inspiración para sus propios proyectos.

## ¿Por qué Fyne?
Decidí usar Fyne porque me interesaba mucho la idea de crear interfaces gráficas en Go. Fyne es una biblioteca que simplifica el trabajo con interfaces, haciéndolo accesible para aquellos que, como yo, aún estamos en el proceso de aprender a manejar elementos gráficos. Con Fyne, puedes crear ventanas, botones, etiquetas y mucho más, y lo mejor es que está diseñado específicamente para Go, lo que facilita bastante el desarrollo.

## Mi Proyecto: Simulador de Estacionamiento
Este simulador de estacionamiento tiene un enfoque en la concurrencia, ya que simula la entrada y salida de vehículos en un estacionamiento con una capacidad limitada. Quería que los vehículos se movieran en la pantalla de forma realista, y para eso, usé las funciones de animación de Fyne. Además, utilicé WaitGroups y canales para controlar la concurrencia y los recursos compartidos, como el acceso a la "puerta" del estacionamiento.

## Aprendizajes Clave
Aquí algunos de los aprendizajes que obtuve al trabajar en este proyecto:

### Concurrencia en Go
Go tiene herramientas muy útiles como los WaitGroups y los canales. Aprendí a controlar cuándo un vehículo puede entrar o salir del estacionamiento y cómo comunicar entre rutinas usando canales. Esto fue esencial para que el simulador funcione correctamente sin errores de sincronización.

### Animaciones con Fyne
Fyne tiene una función de animación que permite que los objetos en la pantalla se muevan suavemente de un punto a otro. Esto fue clave para darle vida al simulador, y realmente aprendí mucho ajustando la velocidad y el trayecto de cada vehículo.

### Patrón de Observador
Implementé el patrón de observador para notificar cuando se liberaba un espacio en el estacionamiento. Esto fue una manera interesante de hacer que el sistema responda a los cambios en tiempo real y mostrar información actualizada al usuario.

### Separación de Responsabilidades
Para hacer el código más limpio y manejable, dividí el proyecto en tres partes principales:

1. **La interfaz (ui)**: Aquí se encuentra la lógica que gestiona la interfaz gráfica.
2. **El modelo del estacionamiento (models/estacionamiento.go)**: Gestiona la lógica de los espacios disponibles y los observadores.
3. **El modelo de vehículo (models/vehiculo.go)**: Controla el comportamiento de cada vehículo individual, incluyendo la animación de entrada y salida.

## Conclusiones 
Este proyecto me ayudo a comprender cómo estructurar un programa con múltiples componentes que deben trabajar juntos de manera sincronizada. Fyne se ha convertido en una herramienta muy valiosa para mí, y estoy seguro de que muchos de ustedes también podrán sacarle provecho en sus propios proyectos.

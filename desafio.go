package main

import "fmt"

const ebulicaoK float64 = 373

func main ()  {


	tempK := ebulicaoK
	tempC :=  tempK - 273

	fmt.Printf("A tempetarua de ebulição em °K é: %g (%T) e A tempetarua de ebulição em °C é: %g (%T) ",tempK, tempK, tempC, tempC)
}
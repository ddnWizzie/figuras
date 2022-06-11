package figuras

import (
	"math"
)

type Circulo struct {
	Radio float32
}

func (cir *Circulo) calcArea() float32 {
	area := math.Pi * (cir.Radio)
	return area
}

func (cir *Circulo) calcPerimetro() float32 {
	perimetro := 2 * math.Pi * cir.Radio
	return perimetro
}

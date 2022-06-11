package figuras

import (
	"fmt"
)

type CalculoEspacio interface {
	calcArea() float32
	calcPerimetro() float32
}

func CalcEspa(calc CalculoEspacio) {
	fmt.Println("El area es: ", calc.calcArea())
	fmt.Println("El perimetro es: ", calc.calcPerimetro())
}

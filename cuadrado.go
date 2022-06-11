package figuras

type Cuadrado struct {
	Lado float32
}

func (cua *Cuadrado) calcArea() float32 {
	area := cua.Lado * cua.Lado
	return area
}

func (cua *Cuadrado) calcPerimetro() float32 {
	perimetro := (4 * cua.Lado)
	return perimetro
}

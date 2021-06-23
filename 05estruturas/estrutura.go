package estruturas

import "math"

type Retangulo struct {
	Largura, Altura float64
}

type Triangulo struct {
	Base, Altura float64
}

type Circulo struct {
	Raio float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura)/2
}



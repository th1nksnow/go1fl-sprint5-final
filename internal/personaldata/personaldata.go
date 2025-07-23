package personaldata

import "fmt"

type Personal struct {
	Name           string
	Weight, Height float64
}

func (p Personal) Print() {

	outputFormat := "Имя: %d\nВес: %.2f\nРост: %.2f\n"
	fmt.Printf(outputFormat, p.Name, p.Weight, p.Height)
}

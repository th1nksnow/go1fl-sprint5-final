package personaldata

import "fmt"

type Personal struct {
	Name           string
	Weight, Height float64
}

func (p Personal) Print() {

	outputFormat := "Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n"
	fmt.Printf(outputFormat, p.Name, p.Weight, p.Height)
}

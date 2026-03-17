package model

import "fmt"

type Currence float64

func (c Currence) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", c)), nil
}

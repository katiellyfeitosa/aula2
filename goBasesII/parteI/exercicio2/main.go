package main
import (
"fmt"
"errors"
)

func CalcMedia(notas []float32) (float32, error) {
var soma float32

for _, v := range notas {
	if v < 0 {
		return 0, errors.New("erro: número negativo")
	}
	soma += v

}

result := soma/float32(len(notas))
return result, nil
}

func main() {
notas := []float32{2.0, 3.0, 4.0, -6.0}

	result, err := CalcMedia(notas)
	if err != nil {
 fmt.Println(err)
} else {
	fmt.Println(result)
	}
}
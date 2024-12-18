package libs

import (
	"math"
	"strconv"

	"github.com/bonavadeur/miporin/pkg/miporin"
)

var (
	CLIENTSET = miporin.GetClientSet()
)

func init() {

}

func AddMatrix(MatA [][]int32, MatB [][]int32) [][]int32 {
	n := len(MatA)
	MatC := make([][]int32, n)
	for i := range MatC {
		MatC[i] = make([]int32, n)
		for j := range MatC[i] {
			MatC[i][j] = MatA[i][j] + MatB[i][j]
		}

	}
	return MatC
}

func String2RoundedInt(s string) int32 {
	floatValue, _ := strconv.ParseFloat(s, 32)
	if math.IsNaN(floatValue) {
		floatValue = 0.0
	}
	intValue := int32(math.Round(floatValue))
	return intValue
}

func Average(slice []int32) int32 {
	sum := int32(0)
	for _, value := range slice {
		sum += value
	}
	ret := math.Round(float64(sum) / float64(len(slice)))
	if math.IsNaN(ret) {
		return int32(0.0)
	} else {
		return int32(ret)
	}
}

package fmt

import "strconv"

func ToCurrency(amout float64) string{
	return "$" + strconv.FormatFloat(amout,'f',2, 64)
}

package kata18_str_int

import "fmt"

func SumMix(arr []any) int {
	//var mixedArray = []interface{}{"1", 2 , 3, "4", 5}

	sum := 0
	for _, value := range arr {
		switch v := value.(type) {
		case string:
			// Convert string to int
			var num int
			fmt.Sscanf(v, "%d", &num)
			sum += num
		case int:
			sum += v
		default:
		}
	}
	return sum

}

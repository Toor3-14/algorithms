package search

import "gihub.com/Toor3-14/algorithms/abc"

type gen interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Binary[T gen](array []T, num T) int {

	result := -1 // as default - not found
	start := 0
	end := len(array) - 1

	for start <= end {
		mid := (start + end) / 2
		if array[mid] == num {
			result = mid // found
			break
		} else if array[mid] < num {
			start = mid + 1
		} else if array[mid] > num {
			end = mid - 1
		}
	}
	
	return result
}


func BinaryCustom(array []string, str string) int {

	result := -1 // as default - not found
	start := 0
	end := len(array) - 1

	for start <= end {
		mid := (start + end) / 2
		if array[mid] == str {
			result = mid // found
			break
		}else {
			var min int = 0
			if len(array[mid]) > len(str) {
				min = len(str)
			}else {
				min = len(array[mid])
			}
			for j := 0; j < min; j++ {
				if abc.GetPosition(array[mid][j:j+1],abc.ABC) < abc.GetPosition(str[j:j+1],abc.ABC) {
					start = mid + 1
					break
				} else if abc.GetPosition(array[mid][j:j+1],abc.ABC) > abc.GetPosition(str[j:j+1],abc.ABC) {
					end = mid - 1
					break
				}else {
					continue
				}
			}
		}
		
	}
	return result
}
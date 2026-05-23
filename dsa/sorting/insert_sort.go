package sorting

// def insertion_sort(arr: List[int]) -> List[int]:
//     for i in range(1, len(arr)):
//         key = arr[i]
//         j = i - 1  # starts on 0
//         while j >= 0 and key < arr[j]:
//             arr[j + 1] = arr[j]
//             j -= 1
//         arr[j + 1] = key
//     return arr

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}

package searching

// returning *int so we can return nil, i could return -1 or something but i prefer returning nil if i found nothing
func BinarySearch(arr []int, val int) *int {
	// calc start, end, mid
	start := 0
	mid := len(arr) / 2 // -> lowering down
	end := len(arr) - 1

	// infinite loop
	for {
		if arr[mid] == val {
			return &arr[mid]
		} else if arr[mid] < val {
			// if the val is on the right portion of the slice -> update the start pointer
			if start >= end {
				return nil
			}
			start = mid + 1
		} else {
			// if the val is on the left portion of the slice -> update the end pointer
			if end <= start {
				return nil
			}
			end = mid - 1
		}
		mid = (start + end) / 2
	}
}

// val = 9
// 0 1 2 3 4 5 6 7
// 1 2 3 4 5 6 7 8
// s = 0
// m = 4
// e = 7

// 5 6 7
// 6 7 8
// s = 5
// m = 12 / 2 = 6
// e = 7

// 7
// 8
// s = 7
// m = 7
// e = 7

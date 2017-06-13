package bubble_sort

import(
  "github.com/sasZilla/algoritms-on-golang/utils"
)

// BubbleSort algoritm has complexity O(n*n)
func BubbleSort(arr []int) {
  arrLen := len(arr)
  for i := 0; i < arrLen - 1; i++ {
    for j := i + 1; j < arrLen; j++ {
      if arr[i] > arr[j] {
        utils.Swap(arr, i, j)
      }
    }
  }
}

package bubble_sort

import(
  "github.com/sasZilla/algoritms-on-golang/utils"
)

func Sort(arr []int, direction string) {
  arr_len := len(arr)
  for i := 0; i < arr_len - 1; i++ {
    for j := i + 1; j < arr_len; j++ {
      if direction == "asc" && arr[i] > arr[j] {
         utils.Swap(arr, i, j)
      }
      if direction == "desc" && arr[i] < arr[j] {
        utils.Swap(arr, i, j)
      }
    }
  }
}

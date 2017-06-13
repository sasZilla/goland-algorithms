package utils

func Swap(arr []int, left, right int) {
  tmp := arr[left]
  arr[left] = arr[right]
  arr[right] = tmp
}

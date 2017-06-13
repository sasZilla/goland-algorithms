package bubble_sort_test

import(
  "testing"
  "github.com/sasZilla/algoritms-on-golang/sort/bubble"
)

var tests = []int{
  1,3,2,56,3,2,45,3,7,
}

var expected = []int{
  1,2,2,3,3,3,7,45,56,
}

func TestBubbleSort(t *testing.T) {
  bubble_sort.BubbleSort(tests)

  for i,v := range expected {
    if tests[i] != v {
      t.Error("fails with sorting", tests[i], v)
    }
  }
}
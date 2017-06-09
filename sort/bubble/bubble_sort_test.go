package bubble_sort_test

import(
  "testing"
  "github.com/sasZilla/algoritms-on-golang/sort/bubble"
)

var tests = []int{
  1,3,2,56,3,2,45,3,7,
}

var asc_expected = []int{
  1,2,2,3,3,3,7,45,56,
}

var desc_expected = []int{
  56, 45, 7, 3, 3, 3, 2, 2, 1,
}

func TestBubbleSort(t *testing.T) {
  bubble_sort.Sort(tests, "asc")

  for i,v := range asc_expected {
    if tests[i] != v {
      t.Error("fails with sorting", tests[i], v)
    }
  }

  bubble_sort.Sort(tests, "desc")
  for i,v := range desc_expected {
    if tests[i] != v {
      t.Error("fails with sorting", tests[i], v)
    }
  }

}
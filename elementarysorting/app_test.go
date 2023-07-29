package elementarysorting

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testData []int

func TestMain(t *testing.M) {
	testData = createRandomNumberArray(10000)
	t.Run()
}

func Test_insertionSort(t *testing.T) {
	type in struct {
		a []int
	}
	cases := map[string]struct {
		in   in
		want []int
	}{
		"配列内の数字がユニーク": {
			in: in{
				[]int{5, 1, 3, 6, 8, 2, 4, 9, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		"配列内に同じ値を含む": {
			in: in{
				[]int{5, 2, 3, 6, 8, 2, 3, 9, 7},
			},
			want: []int{2, 2, 3, 3, 5, 6, 7, 8, 9},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			insertionSort(tt.in.a)
			if diff := cmp.Diff(tt.in.a, tt.want); diff != "" {
				t.Errorf("insertionSort() (-got +want):\n%s", diff)
			}
		})
	}
}

func Benchmark_insertionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		insertionSort(testData)
	}
}

func createRandomNumberArray(n int) []int {
	a := make([]int, n)
	i := 0
	for i < n {
		a[i] = rand.Intn(n)
		i++
	}
	return a
}

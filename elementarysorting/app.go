package elementarysorting

// 挿入ソート
// ・先頭の要素をソート済みとする
// ・未ソートの部分がなくなるまで、以下の処理を繰り返す
// 　・未ソートの先頭から要素を1つ取り出しvに記録する
// 　・ソート済みの部分において、vより大きい要素を後方へ1つずつ移動する
// 　・最後に空いた位置に「取り出した要素v」を挿入する
func insertionSort(a []int) {
	for c := 1; c < len(a); c++ {
		evalTarget := a[c]
		i := c - 1
		for i >= 0 && a[i] > evalTarget {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = evalTarget
	}
}

// バブルソート
// 順番が逆になっている隣接要素がなくなるまで、次の処理を繰り返す
// ・配列の末尾から隣接する要素を比べていき、大小関係が逆ならば交換する
func bubbleSort(a []int) {
	done := true
	for done {
		done = false
		for i := len(a) - 1; i >= 1; i-- {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				done = true
			}
		}
	}
}

// 選択ソート
// ・見ソートの部分から最小の要素の位置を特定する
// ・最小の要素の位置にある要素と未ソートの部分の先頭要素を交換する
func selectionSort(a []int) {
	for c := 0; c < len(a); c++ {
		minValueIndex := c
		for i := c; i < len(a); i++ {
			if a[i] < a[minValueIndex] {
				minValueIndex = i
			}
		}
		a[c], a[minValueIndex] = a[minValueIndex], a[c]
	}
}

func stableSort(a []int) {

}

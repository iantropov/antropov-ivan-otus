package bucket

type BucketNode struct {
	val  int
	next *BucketNode
}

func Sort(a []int) []int {
	max := a[0]
	for _, el := range a {
		if el > max {
			max = el
		}
	}

	buckets := make([]*BucketNode, len(a))

	for i := 0; i < len(a); i++ {
		elBucketIdx := a[i] / (max + 1) * len(a)

		node := buckets[elBucketIdx]
		if node == nil {
			node = new(BucketNode)
			node.val = a[i]
			buckets[elBucketIdx] = node
			continue
		}

		next := node
		prev := (*BucketNode)(nil)
		for ; next != nil; next = next.next {
			if a[i] < next.val {
				break
			}
			prev = next
		}

		newNode := new(BucketNode)
		newNode.val = a[i]
		newNode.next = next
		if prev == nil {
			buckets[elBucketIdx] = newNode
		} else {
			prev.next = newNode
		}
	}

	idx := 0
	for _, bucket := range buckets {
		for node := bucket; node != nil; node = node.next {
			a[idx] = node.val
			idx++
		}
	}

	return a
}

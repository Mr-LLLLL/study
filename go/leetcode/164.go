package main

func Code_164(nums []int) int {
	return base_sort(nums)
}

func base_sort(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := 0; i < n; i++ {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return ans
}

func bucket_sort(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal := min(nums...)
	maxVal := max(nums...)
	d := (maxVal - minVal) / (n - 1)
	bucketSize := (maxVal-minVal)/d + 1

	buckets := make([]pair, bucketSize)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}
	for _, v := range nums {
		bid := (v - minVal) / d
		if buckets[bid].v1 == -1 {
			buckets[bid].v1 = v
			buckets[bid].v2 = v
		} else {
			buckets[bid].v1 = min(buckets[bid].v1, v)
			buckets[bid].v2 = max(buckets[bid].v2, v)
		}
	}

	prev := -1
	ans := 0
	for i, b := range buckets {
		if b.v1 == -1 {
			continue
		}
		if prev != -1 {
			ans = max(ans, b.v1-buckets[prev].v2)
		}
		prev = i
	}
	return ans
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

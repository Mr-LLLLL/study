pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let mut m = m as usize;
    let mut n = n as usize;
    let mut right = nums1.len();

    while n > 0 {
        right -= 1;

        if m == 0 {
            n -= 1;
            nums1[right] = nums2[n];
        } else {
            nums1[right] = if nums1[m - 1] > nums2[n - 1] {
                m -= 1;
                nums1[m]
            } else {
                n -= 1;
                nums2[n]
            }
        }
    }
}

#[test]
fn merge_test() {
    let tests = vec![(
        (vec![1, 2, 3, 0, 0, 0], 3, vec![2, 5, 6], 3),
        vec![1, 2, 2, 3, 5, 6],
    )];

    for mut v in tests {
        merge(&mut v.0 .0, v.0 .1, &mut v.0 .2, v.0 .3);
        for i in 0..v.0 .0.len() {
            if v.0 .0[i] != v.1[i] {
                assert!(false)
            }
        }
    }
}

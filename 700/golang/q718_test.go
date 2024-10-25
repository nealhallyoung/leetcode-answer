package main

import "testing"

func TestFindLength(t *testing.T) {
	// ecp=3
	// nums1 := []int{1, 2, 3, 2, 1}
	// nums2 := []int{3, 2, 1, 4, 7}
	// exp=4
	// nums1 := []int{1, 0, 0, 0, 0}
	// nums2 := []int{0, 0, 0, 0, 1}
	nums1 := []int{1, 1, 2, 1, 3, 1}
	nums2 := []int{2, 1}
	expected := 2
	solution := findLength(nums1, nums2)
	if solution != expected {
		t.Fatalf("error expected %d, but got %d", expected, solution)
	}
}

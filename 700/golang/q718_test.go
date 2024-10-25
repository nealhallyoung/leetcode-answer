package main

import "testing"

func TestFindLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	// nums1 := []int{1, 0, 0, 0, 0}
	// nums2 := []int{0, 0, 0, 0, 1}
	expected := 3
	solution := findLength(nums1, nums2)
	if solution != expected {
		t.Fatalf("error expected %d, but got %d", expected, solution)
	}
}

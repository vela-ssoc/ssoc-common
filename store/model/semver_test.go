package model

import (
	"slices"
	"testing"
)

func TestSemver(t *testing.T) {
	semvers := []string{
		"5.3.4-alpha",
		"5.3.4-beta",
		"5.3.4",
		"5.3.4-abcdef",
		"5.3.4-abcdeg",
		"0.0.1",
		"0.0.1+sha1",
		"99.99.99",
		"99.99.100",
		"99.99.101",
	}

	nums := make([]uint64, 0, len(semvers))
	for _, version := range semvers {
		num := Semver(version).Uint64()
		nums = append(nums, num)
	}

	slices.Sort(nums)
	for _, num := range nums {
		str := SemverFromUint64(num)
		t.Log(str)
	}
}

func TestSemver_Compare(t *testing.T) {
	arr := []Semver{
		"10.3.2",
		"9.3.4",
		"5.3.4-alpha",
		"5.3.4-beta",
		"5.3.4",
		"5.3.4-abcdef",
		"5.3.4-abcdeg",
		"0.0.1",
		"0.0.1+sha1",
		"99.99.99",
		"99.99.100",
		"99.99.101",
	}
	slices.Sort(arr)
	t.Log(arr) // 直接排序是字符串字典序，错误。

	slices.SortFunc(arr, func(a, b Semver) int { return a.Compare(b) })
	t.Log(arr)
}

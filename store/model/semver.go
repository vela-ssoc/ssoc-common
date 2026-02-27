package model

import (
	"encoding/binary"
	"strconv"
	"strings"
)

// Semver https://semver.org/lang/zh-CN/
type Semver string

func (s Semver) Compare(u Semver) int {
	a, b := s.Uint64(), u.Uint64()
	if a < b {
		return -1
	} else if a > b {
		return 1
	}

	return 0
}

// Uint64 返回版本号的数字表示形式，方便比较大小。
//
// FIXME 这种方式仅适用于特定环境下的版本号，不具备普适性。
//
//	版本号 X.Y.Z 中 X Y Z 均要大于等于 0 小于等于 255 才可用。
//	prerelease 至多保留 5 个字符。
//
// 内存布局：
// MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]
// +-+-+-+-+-+-+-+-+
// | | | |         |
// +-+-+-+-+-+-+-+-+
//
//	0 1 2 3 4 5 6 7
//	M M P PRERELEASE
//	A I A
//	J N T
//	O O C
//	R R H
func (s Semver) Uint64() uint64 {
	// 版本号 MAJOR.MINOR.PATCH 位按照点分十进制依次比较。
	//
	// - 号后面是 PRERELEASE，PRERELEASE 按照字典序大小比较，
	// 所有带 PRERELEASE 的版本号，都要小于主版本号，例如：
	// 0.0.1 > 0.0.1-alpha
	// 0.0.1 > 0.0.1-balabala
	//
	// + 号后面的是 BUILD，不影响版本优先级比较。例如：
	// 0.0.1 == 0.0.1+balabala 大小是一样的。

	str := string(s)
	str, _, _ = strings.Cut(str, "+") // 移除 BUILD 部分
	semver, prerelease, _ := strings.Cut(str, "-")
	bucket := make([]byte, 8) // uint64
	for i, v := range strings.SplitN(semver, ".", 3) {
		n, _ := strconv.Atoi(v) // 假定是符合预期的
		bucket[i] = byte(n)
	}

	idx := 3 // PRERELEASE 剩余位置
	if prerelease == "" {
		// 无 PRERELEASE 说明是主版本号，填充 0xff 保证结果最大。
		for ; idx < len(bucket); idx++ {
			bucket[idx] = 0xff
		}
		return binary.BigEndian.Uint64(bucket)
	}

	for _, b := range []byte(prerelease) {
		if idx >= len(bucket) { // 满了
			break
		}
		bucket[idx] = b
		idx++
	}

	return binary.BigEndian.Uint64(bucket)
}

func SemverFromUint64(n uint64) Semver {
	slot := make([]byte, 8)
	binary.BigEndian.PutUint64(slot, n)

	var semvers []string
	for _, v := range slot[:3] {
		str := strconv.Itoa(int(v))
		semvers = append(semvers, str)
	}

	idx := 3
	for _, b := range slot[idx:] {
		if b == 0xff || b == 0x00 {
			break
		}
		idx++
	}

	semver := strings.Join(semvers, ".")
	if idx == 3 {
		return Semver(semver)
	}

	return Semver(semver + "-" + string(slot[3:idx]))
}

package model

import (
	"encoding/binary"
	"strconv"
	"strings"
)

// Semver https://semver.org/lang/zh-CN/
type Semver string

func (ver Semver) String() string {
	return string(ver)
}

func (ver Semver) Uint64() uint64 {
	// 25.3.11-170856 -> [25.3.11, 170856]
	before, after, _ := strings.Cut(string(ver), "-")
	sn := strings.SplitN(before, ".", 3)
	buf := make([]byte, 8) // uint64
	ver.parseVersion(buf[:3], sn)
	ver.parsePrerelease(buf[3:], after)

	return binary.BigEndian.Uint64(buf[:])
}

func (Semver) parseVersion(dest []byte, versions []string) {
	size := len(dest)
	for i, v := range versions {
		if i >= size {
			break
		}
		n, _ := strconv.ParseInt(v, 10, 64)
		dest[i] = byte(n)
	}
}

func (Semver) parsePrerelease(dest []byte, prerelease string) {
	if num := copy(dest, prerelease); num == 0 {
		for i := range dest {
			dest[i] = 0xff
		}
	}
}

func Uint64ToSemver(u uint64) Semver {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)

	versions := make([]string, 3)
	for n, b := range buf[:3] {
		v := strconv.FormatInt(int64(b), 10)
		versions[n] = v
	}
	version := strings.Join(versions, ".")

	after := buf[3:]
	if after[0] == 0xff {
		return Semver(version)
	}

	end := len(after)
	for i, b := range after {
		if b == 0 {
			end = i
			break
		}
	}
	version = version + "-" + string(after[:end])

	return Semver(version)
}

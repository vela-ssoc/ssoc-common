package model

type Checksum struct {
	MD5     string `json:"md5"`     // MD5
	SHA1    string `json:"sha1"`    // SHA1
	SHA256  string `json:"sha256"`  // SHA256
	BLAKE2b string `json:"blake2b"` // BLAKE2b
}

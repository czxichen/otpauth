package otpauth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// CompareCode 比较Code是否匹配,shift是向前漂移多少个30s
func CompareCode(shift int, code uint32, key string) bool {
	now := time.Now().Unix()
	if shift == 0 {
		shift = 1
	}
	for i := 0; i < shift; i++ {
		now -= int64(i) * 30
		if c, _, _ := GenerateCode(key, now); c == code {
			return true
		}
	}
	return false
}

// GenerateOTP GenerateOTP url
// otpauth://totp/dijielin@qq.com?secret=U4QWUHPI4JZNVXSC&issuer=GOOGLE
func GenerateOTP(issuer, tag string) string {
	if issuer == "" {
		issuer = "GOOGLE"
	}
	return fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s", tag, GenerateSecretKey(), issuer)
}

// GenerateSecretKey 生成随便密钥
func GenerateSecretKey() string {
	var bytes = make([]byte, 10)
	rand.Read(bytes)
	return base32.StdEncoding.EncodeToString(bytes)
}

// GenerateCode 生成动态code
func GenerateCode(secretKey string, epochSeconds int64) (uint32, int64, error) {
	inputNoSpacesUpper := strings.ToUpper(secretKey)
	key, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if err != nil {
		return 0, 0, err
	}
	if epochSeconds == 0 {
		epochSeconds = time.Now().Unix()
	}
	pwd := oneTimePassword(key, toBytes(epochSeconds/30))
	return pwd, epochSeconds, nil
}

func toBytes(value int64) []byte {
	var result = make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(value))
	return result
}

func toUint32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// 签名算法是: HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F

	return toUint32(hashParts) % 1000000
}

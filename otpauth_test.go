package otpauth

import (
	"testing"
)

func Test_GenerateOTP(t *testing.T) {
	str := GenerateOTP("Di", "dijielin@qq.com")
	t.Logf("Path: %s\n", str)
}

func Test_CompareCode(t *testing.T) {
	ok := CompareCode(3, 865946, "MNTIZ73RIWUUO2PJ")
	t.Log(ok)
}

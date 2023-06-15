package utils

import (
	"github.com/njylll/thirdparty_auxiliary_tool_go/setting"
	"testing"
)

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjIiLCJleHAiOjE2ODYwNDM3NDEsImlzcyI6ImRlcmkifQ.ED9Ru6USxS2PtUBmJZuGT9CkW9kM8RhF5UQPrCrMlsA"
	setting.Setup()
	Setup()
	got, err := ParseToken(token)
	if err != nil {
		t.Error(err)
		return
	}

	println(got.ID)

}

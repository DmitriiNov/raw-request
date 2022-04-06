package dns

import (
	"testing"
)

func TestDns(t *testing.T) {
	DoDnsRequest("google.com")
}

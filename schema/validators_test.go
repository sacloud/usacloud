package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateIPv4(t *testing.T) {

	expects := map[string]bool{
		"192.168.0.1":                  true,
		"192.168.0.1.1":                false,
		"192.168.0.0/24":               false,
		"2401:2500:10a:100e::1":        false,
		"fe::1":                        false,
		"2401:2500:10a:100e::xxxxxxxx": false,
	}

	for target, expect := range expects {

		errs := ValidateIPv4Address()("TEST", target)
		actual := len(errs) == 0

		assert.Equal(t, actual, expect, "Target: %s", target)
	}

}

func TestValidateIPv6(t *testing.T) {

	expects := map[string]bool{
		"192.168.0.1":                  false,
		"2401:2500:10a:100e::1":        true,
		"fe::1":                        true,
		"2401:2500:10a:100e::xxxxxxxx": false,
	}

	for target, expect := range expects {

		errs := ValidateIPv6Address()("TEST", target)
		actual := len(errs) == 0

		assert.Equal(t, actual, expect, "Target: %s", target)
	}

}

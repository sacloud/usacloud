package define

import (
	"fmt"

	"github.com/sacloud/usacloud/schema"
)

func validateMulti(validators ...schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateMulti(validators...)
}
func validateMultiOr(validators ...schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateMultiOr(validators...)
}
func validateStringSlice(validator schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateStringSlice(validator)
}

func validateIntSlice(validator schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateIntSlice(validator)
}

func validateStrLen(min int, max int) schema.ValidateFunc {
	return schema.ValidateStrLen(min, max)
}

func validateIntRange(min int, max int) schema.ValidateFunc {
	return schema.ValidateIntRange(min, max)
}

func validateInStrValues(allows ...string) schema.ValidateFunc {
	return schema.ValidateInStrValues(allows...)
}

func validateInIntValues(allows ...int) schema.ValidateFunc {
	return schema.ValidateInIntValues(allows...)
}

func validateSakuraID() schema.ValidateFunc {
	return schema.ValidateSakuraID()
}

func validateSakuraShortID(digit int) schema.ValidateFunc {
	return schema.ValidateSakuraShortID(digit)
}

func validateMemberCD() schema.ValidateFunc {
	return schema.ValidateMemberCD()
}

func validateFileExists() schema.ValidateFunc {
	return schema.ValidateFileExists()
}

func validateExistsFileOrStdIn() schema.ValidateFunc {
	return schema.ValidateMultiOr(schema.ValidateFileExists(), schema.ValidateStdinExists())
}

func validateIPv4Address() schema.ValidateFunc {
	return schema.ValidateIPv4Address()
}

func validateIPv4AddressWithPrefixOption() schema.ValidateFunc {
	return schema.ValidateIPv4AddressWithPrefixOption()
}

func validateIPv4AddressWithPrefix() schema.ValidateFunc {
	return schema.ValidateIPv4AddressWithPrefix()
}

func validateMACAddress() schema.ValidateFunc {
	return schema.ValidateMACAddress()
}

func validateDateTimeString() schema.ValidateFunc {
	return schema.ValidateDateTimeString()
}

func validateBackupTime() schema.ValidateFunc {
	timeStrings := []string{}

	minutes := []int{0, 15, 30, 45}

	// create list [00:00 ,,, 23:45]
	for hour := 0; hour <= 23; hour++ {
		for _, minute := range minutes {
			timeStrings = append(timeStrings, fmt.Sprintf("%02d:%02d", hour, minute))
		}
	}

	return schema.ValidateInStrValues(timeStrings...)
}

func mergeParameterMap(params ...map[string]*schema.Schema) map[string]*schema.Schema {
	dest := map[string]*schema.Schema{}
	for _, m := range params {
		for k, v := range m {
			dest[k] = v
		}
	}
	return dest
}

package define

import (
	"github.com/sacloud/usacloud/schema"
)

func validateMulti(validators ...schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateMulti(validators...)
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

func validateIPv4Address() schema.ValidateFunc {
	return schema.ValidateIPv4Address()
}

func validateDateTimeString() schema.ValidateFunc {
	return schema.ValidateDateTimeString()
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

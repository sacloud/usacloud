package define

import (
	"github.com/sacloud/usacloud/schema"
)

func validateMulti(validators ...schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return schema.ValidateMulti(validators...)
}

func validateStringSlice(validator schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return schema.ValidateStringSlice(validator)
}

func validateIntSlice(validator schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return schema.ValidateIntSlice(validator)
}

func validateStrLen(min int, max int) schema.SchemaValidateFunc {
	return schema.ValidateStrLen(min, max)
}

func validateIntRange(min int, max int) schema.SchemaValidateFunc {
	return schema.ValidateIntRange(min, max)
}

func validateInStrValues(allows ...string) schema.SchemaValidateFunc {
	return schema.ValidateInStrValues(allows...)
}

func validateInIntValues(allows ...int) schema.SchemaValidateFunc {
	return schema.ValidateInIntValues(allows...)
}

func validateSakuraID() schema.SchemaValidateFunc {
	return schema.ValidateSakuraID()
}

func validateSakuraShortID(digit int) schema.SchemaValidateFunc {
	return schema.ValidateSakuraShortID(digit)
}

func validateMemberCD() schema.SchemaValidateFunc {
	return schema.ValidateMemberCD()
}

func validateFileExists() schema.SchemaValidateFunc {
	return schema.ValidateFileExists()
}

func validateIPv4Address() schema.SchemaValidateFunc {
	return schema.ValidateIPv4Address()
}

func validateDateTimeString() schema.SchemaValidateFunc {
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

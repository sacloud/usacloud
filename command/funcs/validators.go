package funcs

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/output"
)

func validateSakuraID(fieldName string, object interface{}) []error {
	return command.ValidateSakuraID(fieldName, object)
}

func validateInStrValues(fieldName string, object interface{}, allowValues ...string) []error {
	return command.ValidateInStrValues(fieldName, object, allowValues...)
}

func validateRequired(fieldName string, object interface{}) []error {
	return command.ValidateRequired(fieldName, object)
}

func validateSetProhibited(fieldName string, object interface{}) []error {
	return command.ValidateSetProhibited(fieldName, object)
}

func validateConflicts(fieldName string, object interface{}, values map[string]interface{}) []error {
	return command.ValidateConflicts(fieldName, object, values)
}

func validateConflictValues(fieldName string, object interface{}, values map[string]interface{}) []error {
	return command.ValidateConflictValues(fieldName, object, values)
}

func validateBetween(fieldName string, object interface{}, min int, max int) []error {
	return command.ValidateBetween(fieldName, object, min, max)
}

func validateOutputOption(o output.Option) []error {
	return command.ValidateOutputOption(o)
}

func validateIPv4AddressArgs(ipaddr string) []error {
	return command.ValidateIPv4Address("Args", ipaddr)
}

func validateIPv6AddressArgs(ipaddr string) []error {
	return command.ValidateIPv6Address("Args", ipaddr)
}

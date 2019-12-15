package interfaces

import (
	"bytes"
	"encoding/json"
)

// Company and people check status
type CheckStatusEnum int

const (
	// Entity isn't on pep or sanction list
	OK CheckStatusEnum = iota
	//Entity is on pep list
	WARNING
	// Entity is on sanction list
	ISSUE
)

func (s CheckStatusEnum) String() string {
	return statusEnumToString[s]
}

var statusEnumToString = map[CheckStatusEnum]string{
	OK:      "OK",
	WARNING: "WARNING",
	ISSUE:   "ISSUE",
}

var statusEnumToID = map[string]CheckStatusEnum{
	"OK":      OK,
	"WARNING": WARNING,
	"ISSUE":   ISSUE,
}

var BadTypeToStatusMapping = map[string]CheckStatusEnum{
	"pep":      WARNING,
	"sanction": ISSUE,
}

// MarshalJSON marshals the enum as a quoted json string
func (s CheckStatusEnum) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(statusEnumToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *CheckStatusEnum) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = statusEnumToID[j]
	return nil
}

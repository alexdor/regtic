package interfaces

import (
	"bytes"
	"encoding/json"
)

type EntityTypeEnum int

const (
	COMPANY EntityTypeEnum = iota
	PERSON
)

func (s EntityTypeEnum) String() string {
	return entityTypeEnumToString[s]
}

var entityTypeEnumToString = map[EntityTypeEnum]string{
	COMPANY: "COMPANY",
	PERSON:  "PERSON",
}

var entityTypeEnumToID = map[string]EntityTypeEnum{
	"COMPANY": COMPANY,
	"PERSON":  PERSON,
}

// MarshalJSON marshals the enum as a quoted json string
func (s EntityTypeEnum) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(entityTypeEnumToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *EntityTypeEnum) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = entityTypeEnumToID[j]
	return nil
}

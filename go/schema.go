package notification

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/mehrdad-masoumi/broker-contract/schemas"
)

var (
	schemaOnce sync.Once
	schemaObj  *jsonschema.Schema
	schemaErr  error
)

func loadSchema() (*jsonschema.Schema, error) {
	schemaOnce.Do(func() {
		c := jsonschema.NewCompiler()
		c.Draft = jsonschema.Draft2020
		const url = "https://github.com/mehrdad-masoumi/broker-contract/schemas/notification.requested.v1.json"
		if err := c.AddResource(url, strings.NewReader(string(schemas.NotificationRequestedV1))); err != nil {
			schemaErr = err
			return
		}
		schemaObj, schemaErr = c.Compile(url)
	})
	return schemaObj, schemaErr
}

// ValidateJSONSchema validates raw JSON bytes against notification.requested.v1.json.
func ValidateJSONSchema(raw []byte) error {
	var doc any
	if err := json.Unmarshal(raw, &doc); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}
	s, err := loadSchema()
	if err != nil {
		return err
	}
	if err := s.Validate(doc); err != nil {
		return fmt.Errorf("json schema: %w", err)
	}
	return nil
}

// ParseAndValidateJSON unmarshals JSON into NotificationRequested and runs
// both JSON Schema and structural ValidateNotificationRequested checks.
func ParseAndValidateJSON(raw []byte) (NotificationRequested, error) {
	if err := ValidateJSONSchema(raw); err != nil {
		return NotificationRequested{}, err
	}
	var n NotificationRequested
	if err := json.Unmarshal(raw, &n); err != nil {
		return NotificationRequested{}, err
	}
	EnsureVersion(&n)
	n.Channels = NormalizeChannels(n.Channels)
	if err := ValidateNotificationRequested(n); err != nil {
		return NotificationRequested{}, err
	}
	return n, nil
}

package domain

import (
	"encoding/json"
	"time"
)

type Application struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      string
	Name      string
	SpaceGUID string
	Diego     bool
}

func NewApplication(guid string) Application {
	return Application{
		GUID: guid,
	}
}

func (app Application) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid":       app.GUID,
			"created_at": app.CreatedAt,
			"updated_at": app.UpdatedAt,
		},
		"entity": map[string]interface{}{
			"name":                 app.Name,
			"space_guid":           app.SpaceGUID,
			"diego":                app.Diego,
			"space_url":            "/v2/spaces/" + app.SpaceGUID,
			"stack_url":            "/v2/stacks/some-not-implemented-stack-guid",
			"events_url":           "/v2/apps/" + app.GUID + "/events",
			"service_bindings_url": "/v2/apps/" + app.GUID + "/service_bindings",
			"routes_url":           "/v2/apps/" + app.GUID + "/routes",
		},
	})
}

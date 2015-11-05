package domain

import (
	"encoding/json"
	"time"
)

type Buildpack struct {
	Name      string
	GUID      string
	CreatedAt time.Time
}

func NewBuildpack(s string) Buildpack {
	return Buildpack{
		GUID:      s,
		CreatedAt: time.Now(),
	}
}

func (buildpack Buildpack) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid":       buildpack.GUID,
			"url":        "/v2/buildpacks/" + buildpack.GUID,
			"created_at": buildpack.CreatedAt,
		},
		"entity": map[string]interface{}{
			"name":     buildpack.Name,
			"position": 1,
			"enabled":  true,
			"locked":   false,
			"filename": "",
		},
	})
}

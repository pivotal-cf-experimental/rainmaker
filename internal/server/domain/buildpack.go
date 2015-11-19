package domain

import (
	"encoding/json"
	"time"
)

type Buildpack struct {
	Name     string
	GUID     string
	Position int
	Enabled  bool
	Locked   bool
	Filename string
}

func NewBuildpack(s string) Buildpack {
	return Buildpack{
		GUID: s,
	}
}

func (buildpack Buildpack) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid":       buildpack.GUID,
			"url":        "/v2/buildpacks/" + buildpack.GUID,
			"created_at": time.Time{},
			"updated_at": time.Time{},
		},
		"entity": map[string]interface{}{
			"name":     buildpack.Name,
			"position": buildpack.Position,
			"enabled":  buildpack.Enabled,
			"locked":   buildpack.Locked,
			"filename": buildpack.Filename,
		},
	})
}

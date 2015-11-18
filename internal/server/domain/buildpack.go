package domain

import "encoding/json"

type Buildpack struct {
	Name string
	GUID string
}

func NewBuildpack(s string) Buildpack {
	return Buildpack{
		GUID: s,
	}
}

func (buildpack Buildpack) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": buildpack.GUID,
			"url":  "/v2/buildpacks/" + buildpack.GUID,
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

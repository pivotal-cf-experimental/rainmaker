package rainmaker

import (
	"net/url"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
)

func NewRequestPlan(path string, query url.Values) requestPlan {
	return newRequestPlan(path, query)
}

func NewOrganizationFromResponse(config Config, document documents.OrganizationResponse) Organization {
	return newOrganizationFromResponse(config, document)
}

func NewSpaceFromResponse(config Config, document documents.SpaceResponse) Space {
	return newSpaceFromResponse(config, document)
}

func NewApplicationFromResponse(config Config, document documents.ApplicationResponse) Application {
	return newApplicationFromResponse(config, document)
}

package rainmaker

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pivotal-cf-experimental/rainmaker/internal/documents"
	"github.com/pivotal-cf-experimental/rainmaker/internal/network"
)

type OrganizationsList struct {
	config        Config
	plan          requestPlan
	TotalResults  int
	TotalPages    int
	NextURL       string
	PrevURL       string
	Organizations []Organization
}

func NewOrganizationsList(config Config, plan requestPlan) OrganizationsList {
	return OrganizationsList{
		config: config,
		plan:   plan,
	}
}

func (list OrganizationsList) Create(org Organization, token string) (Organization, error) {
	var document documents.OrganizationResponse
	resp, err := newNetworkClient(list.config).MakeRequest(network.Request{
		Method:        "POST",
		Path:          list.plan.Path,
		Authorization: network.NewTokenAuthorization(token),
		Body:          network.NewJSONRequestBody(org),
		AcceptableStatusCodes: []int{http.StatusCreated},
	})
	if err != nil {
		return Organization{}, err
	}

	err = json.Unmarshal(resp.Body, &document)
	if err != nil {
		panic(err)
	}

	return newOrganizationFromResponse(list.config, document), nil
}

func (list OrganizationsList) Next(token string) (OrganizationsList, error) {
	nextURL, err := url.Parse("http://example.com" + list.NextURL)
	if err != nil {
		return OrganizationsList{}, err
	}

	nextList := NewOrganizationsList(list.config, newRequestPlan(nextURL.Path, nextURL.Query()))
	err = nextList.Fetch(token)

	return nextList, err
}

func (list OrganizationsList) Prev(token string) (OrganizationsList, error) {
	prevURL, err := url.Parse("http://example.com" + list.PrevURL)
	if err != nil {
		return OrganizationsList{}, err
	}

	prevList := NewOrganizationsList(list.config, newRequestPlan(prevURL.Path, prevURL.Query()))
	err = prevList.Fetch(token)

	return prevList, err
}

func (list OrganizationsList) HasNextPage() bool {
	return list.NextURL != ""
}

func (list OrganizationsList) HasPrevPage() bool {
	return list.PrevURL != ""
}

func (list OrganizationsList) AllOrganizations(token string) ([]Organization, error) {
	l := list
	var orgs []Organization

	for l.HasPrevPage() {
		var err error
		l, err = l.Prev(token)
		if err != nil {
			return []Organization{}, err
		}

		orgs = append(l.Organizations, orgs...)
	}

	orgs = append(orgs, list.Organizations...)

	l = list
	for l.HasNextPage() {
		var err error
		l, err = l.Next(token)
		if err != nil {
			return []Organization{}, err
		}

		orgs = append(orgs, l.Organizations...)
	}

	return orgs, nil
}

func (list *OrganizationsList) Fetch(token string) error {
	u := url.URL{
		Path:     list.plan.Path,
		RawQuery: list.plan.Query.Encode(),
	}

	resp, err := newNetworkClient(list.config).MakeRequest(network.Request{
		Method:                "GET",
		Path:                  u.String(),
		Authorization:         network.NewTokenAuthorization(token),
		AcceptableStatusCodes: []int{http.StatusOK},
	})
	if err != nil {
		return err
	}

	var response documents.OrganizationsListResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		panic(err)
	}

	updatedList := newOrganizationsListFromResponse(list.config, list.plan, response)
	list.TotalResults = updatedList.TotalResults
	list.TotalPages = updatedList.TotalPages
	list.NextURL = updatedList.NextURL
	list.PrevURL = updatedList.PrevURL
	list.Organizations = updatedList.Organizations

	return nil
}

func newOrganizationsListFromResponse(config Config, plan requestPlan, response documents.OrganizationsListResponse) OrganizationsList {
	list := NewOrganizationsList(config, plan)
	list.TotalResults = response.TotalResults
	list.TotalPages = response.TotalPages
	list.PrevURL = response.PrevURL
	list.NextURL = response.NextURL
	list.Organizations = make([]Organization, 0)

	for _, orgResponse := range response.Resources {
		list.Organizations = append(list.Organizations, newOrganizationFromResponse(config, orgResponse))
	}

	return list
}

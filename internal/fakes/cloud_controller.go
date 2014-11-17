package fakes

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
)

type CloudController struct {
	server *httptest.Server
}

func NewCloudController() *CloudController {
	fake := &CloudController{}
	router := mux.NewRouter()
	router.HandleFunc("/v2/organizations/{guid}", fake.GetOrganization)
	router.HandleFunc("/v2/organizations/{guid}/users", fake.GetOrganizationUsers)
	router.HandleFunc("/v2/organizations/{guid}/billing_managers", fake.GetOrganizationBillingManagers)
	router.HandleFunc("/v2/spaces/{guid}", fake.GetSpace)
	router.HandleFunc("/v2/users", fake.GetUsers)

	fake.server = httptest.NewUnstartedServer(router)
	return fake
}

func (fake *CloudController) Start() {
	fake.server.Start()
}

func (fake *CloudController) Close() {
	fake.server.Close()
}

func (fake *CloudController) URL() string {
	return fake.server.URL
}

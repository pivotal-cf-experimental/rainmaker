package fakes

import (
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/gorilla/mux"
)

type CloudController struct {
	server           *httptest.Server
	Organizations    *Organizations
	Spaces           *Spaces
	Users            *Users
	ServiceInstances *ServiceInstances
}

func NewCloudController() *CloudController {
	fake := &CloudController{
		Organizations:    NewOrganizations(),
		Spaces:           NewSpaces(),
		Users:            NewUsers(),
		ServiceInstances: NewServiceInstances(),
	}

	router := mux.NewRouter()
	router.HandleFunc("/v2/organizations", fake.createOrganization).Methods("POST")
	router.HandleFunc("/v2/organizations/{guid}", fake.getOrganization).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/users", fake.getOrganizationUsers).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/users/{user_guid}", fake.associateUserToOrganization).Methods("PUT")
	router.HandleFunc("/v2/organizations/{guid}/billing_managers", fake.getOrganizationBillingManagers).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/billing_managers/{billing_manager_guid}", fake.associateBillingManagerToOrganization).Methods("PUT")
	router.HandleFunc("/v2/organizations/{guid}/auditors", fake.getOrganizationAuditors).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/auditors/{auditor_guid}", fake.associateAuditorToOrganization).Methods("PUT")
	router.HandleFunc("/v2/organizations/{guid}/managers", fake.getOrganizationManagers).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/managers/{manager_guid}", fake.associateManagerToOrganization).Methods("PUT")
	router.HandleFunc("/v2/spaces", fake.createSpace).Methods("POST")
	router.HandleFunc("/v2/spaces/{guid}", fake.getSpace).Methods("GET")
	router.HandleFunc("/v2/spaces/{guid}/developers/{developer_guid}", fake.associateDeveloperToSpace).Methods("PUT")
	router.HandleFunc("/v2/spaces/{guid}/developers", fake.getSpaceDevelopers).Methods("GET")
	router.HandleFunc("/v2/users", fake.getUsers).Methods("GET")
	router.HandleFunc("/v2/users", fake.createUser).Methods("POST")
	router.HandleFunc("/v2/users/{guid}", fake.getUser).Methods("GET")
	router.HandleFunc("/v2/service_instances", fake.createServiceInstance).Methods("POST")
	router.HandleFunc("/v2/service_instances/{guid}", fake.getServiceInstance).Methods("GET")

	handler := fake.RequireToken(router)
	fake.server = httptest.NewUnstartedServer(handler)
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

func (fake *CloudController) Reset() {
	fake.Organizations.Clear()
	fake.Spaces.Clear()
	fake.Users.Clear()
}

func (fake *CloudController) RequireToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ok, err := regexp.MatchString(`Bearer .+`, req.Header.Get("Authorization"))
		if err != nil {
			panic(err)
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Not Authorized"))
			return
		}

		handler.ServeHTTP(w, req)
	})
}

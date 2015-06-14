package fakes

import (
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/organizations"
)

type CloudController struct {
	server           *httptest.Server
	Organizations    *domain.Organizations
	Spaces           *domain.Spaces
	Users            *domain.Users
	ServiceInstances *domain.ServiceInstances
}

func NewCloudController() *CloudController {
	fake := &CloudController{
		Organizations:    domain.NewOrganizations(),
		Spaces:           domain.NewSpaces(),
		Users:            domain.NewUsers(),
		ServiceInstances: domain.NewServiceInstances(),
	}

	router := mux.NewRouter()
	router.Handle("/v2/organizations{anything:.*}", organizations.NewRouter(fake.Organizations, fake.Users))
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

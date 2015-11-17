package testserver

import (
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/applications"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/buildpacks"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/domain"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/organizations"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/serviceinstances"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/spaces"
	"github.com/pivotal-cf-experimental/rainmaker/internal/server/users"
)

type CloudController struct {
	server           *httptest.Server
	Organizations    *domain.Organizations
	Spaces           *domain.Spaces
	Users            *domain.Users
	ServiceInstances *domain.ServiceInstances
	Applications     *domain.Applications
	Buildpacks       *domain.Buildpacks
}

func NewCloudController() *CloudController {
	cc := &CloudController{
		Organizations:    domain.NewOrganizations(),
		Spaces:           domain.NewSpaces(),
		Users:            domain.NewUsers(),
		ServiceInstances: domain.NewServiceInstances(),
		Applications:     domain.NewApplications(),
	}

	router := mux.NewRouter()
	router.Handle("/v2/organizations{anything:.*}", organizations.NewRouter(cc.Organizations, cc.Users))
	router.Handle("/v2/spaces{anything:.*}", spaces.NewRouter(cc.Organizations, cc.Spaces, cc.Users))
	router.Handle("/v2/users{anything:.*}", users.NewRouter(cc.Users, cc.Spaces))
	router.Handle("/v2/service_instances{anything:.*}", serviceinstances.NewRouter(cc.ServiceInstances))
	router.Handle("/v2/apps{anything:.*}", applications.NewRouter(domain.NewGUID, cc.Applications))
	router.Handle("/v2/buildpacks{anything:.*}", buildpacks.NewRouter(cc.Buildpacks))

	handler := cc.requireToken(router)
	cc.server = httptest.NewUnstartedServer(handler)

	return cc
}

func (cc *CloudController) Start() {
	cc.server.Start()
}

func (cc *CloudController) Close() {
	cc.server.Close()
}

func (cc *CloudController) URL() string {
	return cc.server.URL
}

func (cc *CloudController) Reset() {
	cc.Organizations.Clear()
	cc.Spaces.Clear()
	cc.Users.Clear()
	cc.ServiceInstances.Clear()
	cc.Applications.Clear()
}

func (cc *CloudController) requireToken(handler http.Handler) http.Handler {
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

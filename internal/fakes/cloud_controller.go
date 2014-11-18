package fakes

import (
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/gorilla/mux"
)

type CloudController struct {
	server *httptest.Server
}

func NewCloudController() *CloudController {
	fake := &CloudController{}
	router := mux.NewRouter()
	router.HandleFunc("/v2/organizations/{guid}", fake.GetOrganization).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/users", fake.GetOrganizationUsers).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/billing_managers", fake.GetOrganizationBillingManagers).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/auditors", fake.GetOrganizationAuditors).Methods("GET")
	router.HandleFunc("/v2/organizations/{guid}/managers", fake.GetOrganizationManagers).Methods("GET")
	router.HandleFunc("/v2/spaces/{guid}", fake.GetSpace).Methods("GET")
	router.HandleFunc("/v2/users", fake.GetUsers).Methods("GET")

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

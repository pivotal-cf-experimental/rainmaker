package organizations

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf-experimental/rainmaker/internal/fakes/domain"
)

func NewRouter(orgs *domain.Organizations, users *domain.Users) http.Handler {
	router := mux.NewRouter()

	router.Handle("/v2/organizations", createHandler{orgs}).Methods("POST")
	router.Handle("/v2/organizations/{guid}", getHandler{orgs}).Methods("GET")

	router.Handle("/v2/organizations/{guid}/users", getUsersHandler{orgs}).Methods("GET")
	router.Handle("/v2/organizations/{guid}/users/{user_guid}", associateUserHandler{orgs, users}).Methods("PUT")

	router.Handle("/v2/organizations/{guid}/billing_managers", getBillingManagersHandler{orgs}).Methods("GET")
	router.Handle("/v2/organizations/{guid}/billing_managers/{billing_manager_guid}", associateBillingManagerHandler{orgs, users}).Methods("PUT")

	router.Handle("/v2/organizations/{guid}/auditors", getAuditorsHandler{orgs}).Methods("GET")
	router.Handle("/v2/organizations/{guid}/auditors/{auditor_guid}", associateAuditorHandler{orgs, users}).Methods("PUT")

	router.Handle("/v2/organizations/{guid}/managers", getManagersHandler{orgs}).Methods("GET")
	router.Handle("/v2/organizations/{guid}/managers/{manager_guid}", associateManagerHandler{orgs, users}).Methods("PUT")

	return router
}
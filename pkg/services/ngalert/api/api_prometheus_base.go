/*Package api contains base API implementation of unified alerting
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 * Need to remove unused imports.
 */
package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
)

type PrometheusApiService interface {
	RouteGetAlertStatuses(*models.ReqContext) response.Response
	RouteGetRuleStatuses(*models.ReqContext) response.Response
}

type PrometheusApiBase struct {
	log log.Logger
}

func (api *API) RegisterPrometheusApiEndpoints(srv PrometheusApiService) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Get(toMacaronPath("/prometheus/{Recipient}/api/v1/alerts"), routing.Wrap(srv.RouteGetAlertStatuses))
		group.Get(toMacaronPath("/prometheus/{Recipient}/api/v1/rules"), routing.Wrap(srv.RouteGetRuleStatuses))
	})
}

func (base PrometheusApiBase) RouteGetAlertStatuses(c *models.ReqContext) response.Response {
	recipient := c.Params(":Recipient")
	base.log.Info("RouteGetAlertStatuses: ", "Recipient", recipient)
	return response.Error(http.StatusNotImplemented, "", nil)
}

func (base PrometheusApiBase) RouteGetRuleStatuses(c *models.ReqContext) response.Response {
	recipient := c.Params(":Recipient")
	base.log.Info("RouteGetRuleStatuses: ", "Recipient", recipient)
	return response.Error(http.StatusNotImplemented, "", nil)
}

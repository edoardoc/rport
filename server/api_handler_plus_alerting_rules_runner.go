package chserver

import (
	"context"
	"net/http"

	rportplus "github.com/realvnc-labs/rport/plus"
	alertingcap "github.com/realvnc-labs/rport/plus/capabilities/alerting"
	"github.com/realvnc-labs/rport/plus/capabilities/alerting/entities/rundata"
	"github.com/realvnc-labs/rport/server/api"
)

func (al *APIListener) getAlertingCapability() (capEx alertingcap.CapabilityEx, statusCode int, err error) {
	plusManager := al.Server.plusManager
	if plusManager == nil {
		return nil, http.StatusUnauthorized, rportplus.ErrPlusNotAvailable
	}

	capEx = plusManager.GetAlertingCapabilityEx()
	if capEx == nil {
		return nil, http.StatusForbidden, rportplus.ErrCapabilityNotAvailable(rportplus.PlusAlertingCapability)
	}

	return capEx, 0, nil
}

func (al *APIListener) handleTestRules(w http.ResponseWriter, r *http.Request) {
	asCap, status, err := al.getAlertingCapability()
	if err != nil {
		al.jsonErrorResponse(w, status, err)
	}

	runData := rundata.RunData{}

	err = parseRequestBody(r.Body, &runData)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	ctx := context.Background()

	results, errs, err := asCap.RunRulesTest(ctx, &runData, al.Logger)

	if err != nil {
		if errs != nil {
			errPayload := makeValidationErrorPayload(errs)
			al.writeJSONResponse(w, http.StatusBadRequest, errPayload)
			return
		}
		al.jsonErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response := api.NewSuccessPayload(results)

	al.writeJSONResponse(w, http.StatusOK, response)
}

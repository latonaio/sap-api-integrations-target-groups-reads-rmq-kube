package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-target-groups-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToTargetGroupCollection(raw []byte, l *logger.Logger) ([]TargetGroupCollection, error) {
	pm := &responses.TargetGroupCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to TargetGroupCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	targetGroupCollection := make([]TargetGroupCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		targetGroupCollection = append(targetGroupCollection, TargetGroupCollection{
			ObjectID:            data.ObjectID,
			ETag:                data.ETag,
			ID:                  data.ID,
			LifeCycleStatusCode: data.LifeCycleStatusCode,
			Description:         data.Description,
			MemberNumberValue:   data.MemberNumberValue,
			SalesOrganizationID: data.SalesOrganizationID,
			EntityLastChangedOn: data.EntityLastChangedOn,
		})
	}

	return targetGroupCollection, nil
}

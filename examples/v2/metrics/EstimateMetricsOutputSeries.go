// Tag Configuration Cardinality Estimator returns "Success" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewMetricsApi(apiClient)
	resp, r, err := api.EstimateMetricsOutputSeries(ctx, "system.cpu.idle", *datadog.NewEstimateMetricsOutputSeriesOptionalParameters().WithFilterGroups("app,host").WithFilterNumAggregations(4))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.EstimateMetricsOutputSeries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.EstimateMetricsOutputSeries`:\n%s\n", responseContent)
}

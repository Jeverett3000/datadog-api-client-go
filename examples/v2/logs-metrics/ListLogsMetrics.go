// Get all log-based metrics returns "OK" response

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
	api := datadog.NewLogsMetricsApi(apiClient)
	resp, r, err := api.ListLogsMetrics(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.ListLogsMetrics`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsMetricsApi.ListLogsMetrics`:\n%s\n", responseContent)
}

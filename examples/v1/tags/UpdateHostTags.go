// Update host tags returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := datadog.HostTags{
		Host: common.PtrString("test.host"),
		Tags: []string{
			"environment:production",
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewTagsApi(apiClient)
	resp, r, err := api.UpdateHostTags(ctx, "host_name", body, *datadog.NewUpdateHostTagsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateHostTags`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateHostTags`:\n%s\n", responseContent)
}

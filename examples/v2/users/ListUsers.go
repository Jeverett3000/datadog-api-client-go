// List all users returns "OK" response

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
	// there is a valid "user" in the system
	UserDataAttributesEmail := os.Getenv("USER_DATA_ATTRIBUTES_EMAIL")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewUsersApi(apiClient)
	resp, r, err := api.ListUsers(ctx, *datadog.NewListUsersOptionalParameters().WithFilter(UserDataAttributesEmail))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsers`:\n%s\n", responseContent)
}

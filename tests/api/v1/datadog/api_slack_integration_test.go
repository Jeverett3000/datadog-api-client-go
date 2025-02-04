package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	"github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests"

	"gopkg.in/h2non/gock.v1"
	is "gotest.tools/assert/cmp"
)

const (
	staticAccountName = "slack_integration_test_account"
	staticChannelName = "#test-channel"
)

func TestSlackIntegrationGetAllChannelsMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/get_all_channels.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.GetSlackIntegrationChannels")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels", staticAccountName)).
		Reply(200).
		JSON(data)

	var expected []datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := datadog.NewSlackIntegrationApi(Client(ctx))
	slackChannelsResp, httpResp, err := api.GetSlackIntegrationChannels(ctx, staticAccountName)
	if err != nil {
		t.Errorf("Failed to get slack integration all channels: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(3, len(slackChannelsResp))
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationGetAllChannelsErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewSlackIntegrationApi(Client(ctx))

			uniqueAccountName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := api.GetSlackIntegrationChannels(ctx, uniqueAccountName)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationCreateChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/create_channel.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.CreateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Post(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels", staticAccountName)).
		Reply(200).
		JSON(data)

	var expected datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := datadog.NewSlackIntegrationApi(Client(ctx))

	channelPayload := datadog.NewSlackIntegrationChannel()
	channelPayload.SetName(staticChannelName)
	channelPayload.Display = &datadog.SlackIntegrationChannelDisplay{
		Message:  common.PtrBool(false),
		Notified: common.PtrBool(true),
		Snapshot: common.PtrBool(true),
		Tags:     common.PtrBool(false),
	}

	slackChannelsResp, httpResp, err := api.CreateSlackIntegrationChannel(ctx, staticAccountName, *channelPayload)
	if err != nil {
		t.Errorf("Failed to create slack integration channel: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationCreateChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.SlackIntegrationChannel
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.SlackIntegrationChannel{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.SlackIntegrationChannel{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewSlackIntegrationApi(Client(ctx))

			uniqueAccountName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := api.CreateSlackIntegrationChannel(ctx, uniqueAccountName, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationCreateChannel404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadog.NewSlackIntegrationApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/slack-integration/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because an existing slack account is needed for 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.CreateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Post(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels", staticAccountName)).
		Reply(404).JSON(data)
	defer gock.Off()

	_, httpresp, err := api.CreateSlackIntegrationChannel(ctx, staticAccountName, datadog.SlackIntegrationChannel{})
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSlackIntegrationGetChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/get_channel.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.GetSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticAccountName, staticChannelName)).
		Reply(200).
		JSON(data)

	var expected datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := datadog.NewSlackIntegrationApi(Client(ctx))
	slackChannelsResp, httpResp, err := api.GetSlackIntegrationChannel(ctx, staticAccountName, staticChannelName)
	if err != nil {
		t.Errorf("Failed to get slack integration channel: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationGetChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewSlackIntegrationApi(Client(ctx))

			uniqueAccountName := *tests.UniqueEntityName(ctx, t)
			uniqueChannelName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := api.GetSlackIntegrationChannel(ctx, uniqueAccountName, uniqueChannelName)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationGetChannel404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadog.NewSlackIntegrationApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/slack-integration/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because an existing slack account is needed for 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.GetSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Get(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticAccountName, staticChannelName)).
		Reply(404).JSON(data)
	defer gock.Off()

	_, httpresp, err := api.GetSlackIntegrationChannel(ctx, staticAccountName, staticChannelName)
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSlackIntegrationUpdateChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	fixturePath, err := filepath.Abs("fixtures/slack-integration/update_channel.json")
	if err != nil {
		t.Errorf("Failed to get fixture file path: %s", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Errorf("Failed to open fixture file: %s", err)
	}

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.UpdateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Patch(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticAccountName, staticChannelName)).
		Reply(200).
		JSON(data)

	var expected datadog.SlackIntegrationChannel
	json.Unmarshal([]byte(data), &expected)

	api := datadog.NewSlackIntegrationApi(Client(ctx))
	channelPayload := datadog.NewSlackIntegrationChannel()
	channelPayload.SetName(staticChannelName)
	channelPayload.Display = &datadog.SlackIntegrationChannelDisplay{
		Message:  common.PtrBool(false),
		Notified: common.PtrBool(false),
		Snapshot: common.PtrBool(false),
		Tags:     common.PtrBool(false),
	}

	slackChannelsResp, httpResp, err := api.UpdateSlackIntegrationChannel(ctx, staticAccountName, staticChannelName, *channelPayload)
	if err != nil {
		t.Errorf("Failed to update slack integration channel: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.True(is.DeepEqual(expected, slackChannelsResp)().Success())
}

func TestSlackIntegrationUpdateChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.SlackIntegrationChannel
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.SlackIntegrationChannel{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.SlackIntegrationChannel{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewSlackIntegrationApi(Client(ctx))

			uniqueAccountName := *tests.UniqueEntityName(ctx, t)
			uniqueChannelName := *tests.UniqueEntityName(ctx, t)
			_, httpresp, err := api.UpdateSlackIntegrationChannel(ctx, uniqueAccountName, uniqueChannelName, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestSlackIntegrationUpdateChannel404Error(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	api := datadog.NewSlackIntegrationApi(Client(ctx))

	data, err := tests.ReadFixture("fixtures/slack-integration/error_404.json")
	if err != nil {
		t.Fatalf("Failed to read fixture: %s", err)
	}
	// Mocked because an existing slack account is needed for 404 response
	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.UpdateSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Patch(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticAccountName, staticChannelName)).
		Reply(404).JSON(data)
	defer gock.Off()

	_, httpresp, err := api.UpdateSlackIntegrationChannel(ctx, staticAccountName, staticChannelName, datadog.SlackIntegrationChannel{})
	assert.Equal(404, httpresp.StatusCode)
	apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
	assert.True(ok)
	assert.NotEmpty(apiError.GetErrors())
}

func TestSlackIntegrationRemoveChannelMocked(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx = WithClient(WithFakeAuth(ctx))
	assert := tests.Assert(ctx, t)
	defer gock.Off()

	URL, err := Client(ctx).GetConfig().ServerURLWithContext(ctx, "SlackIntegrationApi.RemoveSlackIntegrationChannel")
	assert.NoError(err)
	gock.New(URL).
		Delete(fmt.Sprintf("/api/v1/integration/slack/configuration/accounts/%s/channels/%s", staticAccountName, staticChannelName)).
		Reply(204)

	api := datadog.NewSlackIntegrationApi(Client(ctx))
	httpResp, err := api.RemoveSlackIntegrationChannel(ctx, staticAccountName, staticChannelName)
	if err != nil {
		t.Errorf("Failed to remove slack integration channel: %v", err)
	}
	assert.Equal(204, httpResp.StatusCode)
}

func TestSlackIntegrationRemoveChannelErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, 400},
		"403 Forbidden":   {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewSlackIntegrationApi(Client(ctx))

			uniqueAccountName := *tests.UniqueEntityName(ctx, t)
			uniqueChannelName := *tests.UniqueEntityName(ctx, t)
			httpresp, err := api.RemoveSlackIntegrationChannel(ctx, uniqueAccountName, uniqueChannelName)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError := err.(common.GenericOpenAPIError)
			assert.NotEmpty(apiError)
		})
	}
}

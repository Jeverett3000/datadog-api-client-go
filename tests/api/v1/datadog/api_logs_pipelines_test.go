/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	"github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestLogsPipelinesLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadog.NewLogsPipelinesApi(Client(ctx))

	// Create a pipeline
	grokParser := datadog.NewLogsGrokParserWithDefaults()
	grokParser.SetSource("src")
	grokParser.SetSamples([]string{"sample"})
	grokParser.SetGrok(datadog.LogsGrokParserRules{
		MatchRules:   "rule1 foo\nrule2 bar",
		SupportRules: common.PtrString("support baz"),
	})
	grokParser.SetName("grok parser")

	logDateRemapper := datadog.NewLogsDateRemapperWithDefaults()
	logDateRemapper.SetSources([]string{"source"})
	logDateRemapper.SetName("log date remapper")

	logStatusRemapper := datadog.NewLogsStatusRemapperWithDefaults()
	logStatusRemapper.SetSources([]string{"source"})
	logStatusRemapper.SetName("log status remapper")

	serviceRemapper := datadog.NewLogsServiceRemapperWithDefaults()
	serviceRemapper.SetSources([]string{"source"})
	serviceRemapper.SetName("service remapper")

	logMessageRemapper := datadog.NewLogsMessageRemapperWithDefaults()
	logMessageRemapper.SetSources([]string{"source"})
	logMessageRemapper.SetName("log message remapper")

	remapperToAttribute := datadog.NewLogsAttributeRemapperWithDefaults()
	remapperToAttribute.SetSources([]string{"source"})
	remapperToAttribute.SetSourceType("tag")
	remapperToAttribute.SetTarget("target")
	remapperToAttribute.SetTargetType("attribute")
	remapperToAttribute.SetPreserveSource(true)
	remapperToAttribute.SetOverrideOnConflict(true)
	remapperToAttribute.SetName("log attribute remapper to attribute target type")
	remapperToAttribute.SetTargetFormat(datadog.TARGETFORMATTYPE_STRING)

	remapperToTag := datadog.NewLogsAttributeRemapperWithDefaults()
	remapperToTag.SetSources([]string{"source"})
	remapperToTag.SetSourceType("tag")
	remapperToTag.SetTarget("target")
	remapperToTag.SetTargetType("tag")
	remapperToTag.SetPreserveSource(true)
	remapperToTag.SetOverrideOnConflict(true)
	remapperToTag.SetName("log attribute remapper to tag target type")

	urlParser := datadog.NewLogsURLParserWithDefaults()
	urlParser.SetSources([]string{"source"})
	urlParser.SetTarget("target")
	urlParser.SetName("URL parser")

	userAgentParser := datadog.NewLogsUserAgentParserWithDefaults()
	userAgentParser.SetSources([]string{"source"})
	userAgentParser.SetTarget("target")
	userAgentParser.SetIsEncoded(true)
	userAgentParser.SetName("user agent parser")

	categoryProcessor := datadog.NewLogsCategoryProcessorWithDefaults()
	categoryProcessor.SetCategories([]datadog.LogsCategoryProcessorCategory{{
		Name: common.PtrString("category"),
		Filter: &datadog.LogsFilter{
			Query: common.PtrString("query"),
		},
	},
	})
	categoryProcessor.SetTarget("target")
	categoryProcessor.SetName("category processor")

	arithmeticProcessor := datadog.NewLogsArithmeticProcessorWithDefaults()
	arithmeticProcessor.SetExpression("foo + bar")
	arithmeticProcessor.SetTarget("target")
	arithmeticProcessor.SetIsReplaceMissing(true)
	arithmeticProcessor.SetName("arithmetic processor")

	stringBuilderProcessor := datadog.NewLogsStringBuilderProcessorWithDefaults()
	stringBuilderProcessor.SetTemplate("template")
	stringBuilderProcessor.SetTarget("target")
	stringBuilderProcessor.SetIsReplaceMissing(true)
	stringBuilderProcessor.SetName("string builder processor")

	geoIPParser := datadog.NewLogsGeoIPParserWithDefaults()
	geoIPParser.SetSources([]string{"source"})
	geoIPParser.SetTarget("target")
	geoIPParser.SetName("geo IP parser")

	lookupProcessor := datadog.NewLogsLookupProcessorWithDefaults()
	lookupProcessor.SetSource("source")
	lookupProcessor.SetTarget("target")
	lookupProcessor.SetLookupTable([]string{"key,value"})
	lookupProcessor.SetDefaultLookup("default_lookup")
	lookupProcessor.SetName("lookup processor")

	traceRemapper := datadog.NewLogsTraceRemapperWithDefaults()
	traceRemapper.SetSources([]string{"source"})
	traceRemapper.SetName("trace remapper")

	// Nested Pipelines
	pipelineProcessor := datadog.NewLogsPipelineProcessorWithDefaults()
	pipelineProcessor.SetName("pipeline processor")
	pipelineProcessor.SetFilter(datadog.LogsFilter{
		Query: common.PtrString("query"),
	})
	pipelineProcessor.SetProcessors([]datadog.LogsProcessor{
		datadog.LogsGrokParserAsLogsProcessor(grokParser),
		datadog.LogsDateRemapperAsLogsProcessor(logDateRemapper),
	})

	pipeline := datadog.LogsPipeline{}
	pipeline.SetIsEnabled(true)
	pipeline.SetFilter(datadog.LogsFilter{Query: common.PtrString("query")})
	pipeline.SetProcessors([]datadog.LogsProcessor{
		datadog.LogsGrokParserAsLogsProcessor(grokParser),
		datadog.LogsDateRemapperAsLogsProcessor(logDateRemapper),
		datadog.LogsStatusRemapperAsLogsProcessor(logStatusRemapper),
		datadog.LogsServiceRemapperAsLogsProcessor(serviceRemapper),
		datadog.LogsMessageRemapperAsLogsProcessor(logMessageRemapper),
		datadog.LogsAttributeRemapperAsLogsProcessor(remapperToAttribute),
		datadog.LogsAttributeRemapperAsLogsProcessor(remapperToTag),
		datadog.LogsURLParserAsLogsProcessor(urlParser),
		datadog.LogsUserAgentParserAsLogsProcessor(userAgentParser),
		datadog.LogsCategoryProcessorAsLogsProcessor(categoryProcessor),
		datadog.LogsArithmeticProcessorAsLogsProcessor(arithmeticProcessor),
		datadog.LogsStringBuilderProcessorAsLogsProcessor(stringBuilderProcessor),
		datadog.LogsGeoIPParserAsLogsProcessor(geoIPParser),
		datadog.LogsLookupProcessorAsLogsProcessor(lookupProcessor),
		datadog.LogsTraceRemapperAsLogsProcessor(traceRemapper),
		datadog.LogsPipelineProcessorAsLogsProcessor(pipelineProcessor),
	})
	pipelineName := *tests.UniqueEntityName(ctx, t)
	pipeline.SetName(pipelineName)

	createdPipeline, httpresp, err := api.CreateLogsPipeline(ctx, pipeline)
	if err != nil {
		t.Fatalf("Error creating logs pipeline: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	defer deleteLogsPipeline(ctx, t, createdPipeline.GetId())
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(pipelineName, createdPipeline.GetName())
	assert.True(createdPipeline.GetIsEnabled())
	filter := createdPipeline.GetFilter()
	assert.Equal("query", filter.GetQuery())
	processors := createdPipeline.GetProcessors()
	_, ok := processors[0].GetActualInstance().(*datadog.LogsGrokParser)
	assert.True(ok)
	_, ok = processors[1].GetActualInstance().(*datadog.LogsDateRemapper)
	assert.True(ok)
	_, ok = processors[2].GetActualInstance().(*datadog.LogsStatusRemapper)
	assert.True(ok)
	_, ok = processors[3].GetActualInstance().(*datadog.LogsServiceRemapper)
	assert.True(ok)
	_, ok = processors[4].GetActualInstance().(*datadog.LogsMessageRemapper)
	assert.True(ok)
	_, ok = processors[5].GetActualInstance().(*datadog.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[6].GetActualInstance().(*datadog.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[7].GetActualInstance().(*datadog.LogsURLParser)
	assert.True(ok)
	_, ok = processors[8].GetActualInstance().(*datadog.LogsUserAgentParser)
	assert.True(ok)
	_, ok = processors[9].GetActualInstance().(*datadog.LogsCategoryProcessor)
	assert.True(ok)
	_, ok = processors[10].GetActualInstance().(*datadog.LogsArithmeticProcessor)
	assert.True(ok)
	_, ok = processors[11].GetActualInstance().(*datadog.LogsStringBuilderProcessor)
	assert.True(ok)
	_, ok = processors[12].GetActualInstance().(*datadog.LogsGeoIPParser)
	assert.True(ok)
	_, ok = processors[13].GetActualInstance().(*datadog.LogsLookupProcessor)
	assert.True(ok)
	_, ok = processors[14].GetActualInstance().(*datadog.LogsTraceRemapper)
	assert.True(ok)
	var nestedPipeline *datadog.LogsPipelineProcessor
	nestedPipeline, ok = processors[15].GetActualInstance().(*datadog.LogsPipelineProcessor)
	assert.True(ok)

	// Nested Pipeline Assertions
	nestedPipelineFitler := nestedPipeline.GetFilter()
	assert.Equal("query", nestedPipelineFitler.GetQuery())
	_, ok = nestedPipeline.GetProcessors()[0].GetActualInstance().(*datadog.LogsGrokParser)
	assert.True(ok)
	_, ok = nestedPipeline.GetProcessors()[1].GetActualInstance().(*datadog.LogsDateRemapper)
	assert.True(ok)

	// Get all pipelines and assert our freshly created one is part of the result
	pipelines, httpresp, err := api.ListLogsPipelines(ctx)
	if err != nil {
		t.Fatalf("Error getting all logs pipelines: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(pipelines, createdPipeline)

	// Get the freshly created pipeline
	pipe, httpresp, err := api.GetLogsPipeline(ctx, createdPipeline.GetId())
	if err != nil {
		t.Fatalf("Error getting log pipeline: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(createdPipeline.GetName(), pipe.GetName())

	// Update the pipeline
	processors = pipeline.GetProcessors()
	processors = append(processors[1:], processors[:1]...)
	pipeline.SetProcessors(processors)
	pipeline.SetIsEnabled(false)
	pipeline.SetFilter(datadog.LogsFilter{Query: common.PtrString("updated query")})
	pipeline.SetName(pipeline.GetName() + "-updated")

	updatedPipeline, httpresp, err := api.UpdateLogsPipeline(ctx, createdPipeline.GetId(), pipeline)
	if err != nil {
		t.Fatalf("Error updating logs pipeline: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(pipelineName+"-updated", updatedPipeline.GetName())
	assert.False(updatedPipeline.GetIsEnabled())
	filter = updatedPipeline.GetFilter()
	assert.Equal("updated query", filter.GetQuery())
	processors = updatedPipeline.GetProcessors()

	_, ok = processors[15].GetActualInstance().(*datadog.LogsGrokParser)
	assert.True(ok)
	_, ok = processors[14].GetActualInstance().(*datadog.LogsPipelineProcessor)
	assert.True(ok)
	_, ok = processors[0].GetActualInstance().(*datadog.LogsDateRemapper)
	assert.True(ok)
	_, ok = processors[1].GetActualInstance().(*datadog.LogsStatusRemapper)
	assert.True(ok)
	_, ok = processors[2].GetActualInstance().(*datadog.LogsServiceRemapper)
	assert.True(ok)
	_, ok = processors[3].GetActualInstance().(*datadog.LogsMessageRemapper)
	assert.True(ok)
	_, ok = processors[4].GetActualInstance().(*datadog.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[5].GetActualInstance().(*datadog.LogsAttributeRemapper)
	assert.True(ok)
	_, ok = processors[6].GetActualInstance().(*datadog.LogsURLParser)
	assert.True(ok)
	_, ok = processors[7].GetActualInstance().(*datadog.LogsUserAgentParser)
	assert.True(ok)
	_, ok = processors[8].GetActualInstance().(*datadog.LogsCategoryProcessor)
	assert.True(ok)
	_, ok = processors[9].GetActualInstance().(*datadog.LogsArithmeticProcessor)
	assert.True(ok)
	_, ok = processors[10].GetActualInstance().(*datadog.LogsStringBuilderProcessor)
	assert.True(ok)
	_, ok = processors[11].GetActualInstance().(*datadog.LogsGeoIPParser)
	assert.True(ok)
	_, ok = processors[12].GetActualInstance().(*datadog.LogsLookupProcessor)
	assert.True(ok)
	_, ok = processors[13].GetActualInstance().(*datadog.LogsTraceRemapper)
	assert.True(ok)

	// Delete the pipeline
	httpresp, err = api.DeleteLogsPipeline(ctx, createdPipeline.GetId())
	assert.Nil(err)
}

func TestUpdateLogsPipelineOrder(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadog.NewLogsPipelinesApi(Client(ctx))

	pipelineOrder, httpresp, err := api.GetLogsPipelineOrder(ctx)
	if err != nil {
		t.Fatalf("Error getting pipeline order: Response %s: %v", err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	newOrder := pipelineOrder.GetPipelineIds()
	newOrder = append(newOrder[1:], newOrder[:1]...)
	pipelineOrder.SetPipelineIds(newOrder)

	newPipelineOrder, httpresp, err := api.UpdateLogsPipelineOrder(ctx, pipelineOrder)
	if err != nil {
		t.Fatalf("Error updating with new order %v: Response %s: %v", newOrder, err.(common.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(pipelineOrder.GetPipelineIds(), newPipelineOrder.GetPipelineIds())
}

func TestLogsPipelinesOrderGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.GetLogsPipelineOrder(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsPipelinesOrderUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsPipelinesOrder
		ExpectedStatusCode int
	}{
		"400 Bad Request":          {WithTestAuth, datadog.LogsPipelinesOrder{}, 400},
		"403 Forbidden":            {WithFakeAuth, datadog.LogsPipelinesOrder{}, 403},
		"422 Unprocessable Entity": {WithTestAuth, datadog.LogsPipelinesOrder{PipelineIds: []string{"id"}}, 422},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.UpdateLogsPipelineOrder(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.ListLogsPipelines(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestLogsPipelinesCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsPipeline
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.LogsPipeline{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.LogsPipeline{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.CreateLogsPipeline(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.GetLogsPipeline(ctx, "id")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

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
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			httpresp, err := api.DeleteLogsPipeline(ctx, "id")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func TestLogsPipelinesUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.LogsPipeline
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.LogsPipeline{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.LogsPipeline{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadog.NewLogsPipelinesApi(Client(ctx))

			_, httpresp, err := api.UpdateLogsPipeline(ctx, "id", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			if tc.ExpectedStatusCode == 403 {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetErrors())
			} else {
				apiError, ok := err.(common.GenericOpenAPIError).Model().(datadog.LogsAPIErrorResponse)
				assert.True(ok)
				assert.NotEmpty(apiError.GetError())
			}
		})
	}
}

func deleteLogsPipeline(ctx context.Context, t *testing.T, pipelineID string) {
	api := datadog.NewLogsPipelinesApi(Client(ctx))

	httpresp, err := api.DeleteLogsPipeline(ctx, pipelineID)
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting Logs Pipeline: %v, Another test may have already deleted this pipeline.", pipelineID)
	}
}

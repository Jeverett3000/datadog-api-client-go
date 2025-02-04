// Update a notebook returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	// there is a valid "notebook" in the system
	NotebookDataID, _ := strconv.ParseInt(os.Getenv("NOTEBOOK_DATA_ID"), 10, 64)

	body := datadog.NotebookUpdateRequest{
		Data: datadog.NotebookUpdateData{
			Attributes: datadog.NotebookUpdateDataAttributes{
				Cells: []datadog.NotebookUpdateCell{
					datadog.NotebookUpdateCell{
						NotebookCellCreateRequest: &datadog.NotebookCellCreateRequest{
							Attributes: datadog.NotebookCellCreateRequestAttributes{
								NotebookMarkdownCellAttributes: &datadog.NotebookMarkdownCellAttributes{
									Definition: datadog.NotebookMarkdownCellDefinition{
										Text: `## Some test markdown

` + "```" + `js
var x, y;
x = 5;
y = 6;
` + "```",
										Type: datadog.NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN,
									},
								}},
							Type: datadog.NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
						}},
					datadog.NotebookUpdateCell{
						NotebookCellCreateRequest: &datadog.NotebookCellCreateRequest{
							Attributes: datadog.NotebookCellCreateRequestAttributes{
								NotebookTimeseriesCellAttributes: &datadog.NotebookTimeseriesCellAttributes{
									Definition: datadog.TimeseriesWidgetDefinition{
										Requests: []datadog.TimeseriesWidgetRequest{
											{
												DisplayType: datadog.WIDGETDISPLAYTYPE_LINE.Ptr(),
												Q:           common.PtrString("avg:system.load.1{*}"),
												Style: &datadog.WidgetRequestStyle{
													LineType:  datadog.WIDGETLINETYPE_SOLID.Ptr(),
													LineWidth: datadog.WIDGETLINEWIDTH_NORMAL.Ptr(),
													Palette:   common.PtrString("dog_classic"),
												},
											},
										},
										ShowLegend: common.PtrBool(true),
										Type:       datadog.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
										Yaxis: &datadog.WidgetAxis{
											Scale: common.PtrString("linear"),
										},
									},
									GraphSize: datadog.NOTEBOOKGRAPHSIZE_MEDIUM.Ptr(),
									SplitBy: &datadog.NotebookSplitBy{
										Keys: []string{},
										Tags: []string{},
									},
									Time: *datadog.NewNullableNotebookCellTime(nil),
								}},
							Type: datadog.NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
						}},
				},
				Name:   "Example-Update_a_notebook_returns_OK_response-updated",
				Status: datadog.NOTEBOOKSTATUS_PUBLISHED.Ptr(),
				Time: datadog.NotebookGlobalTime{
					NotebookRelativeTime: &datadog.NotebookRelativeTime{
						LiveSpan: datadog.WIDGETLIVESPAN_PAST_ONE_HOUR,
					}},
			},
			Type: datadog.NOTEBOOKRESOURCETYPE_NOTEBOOKS,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewNotebooksApi(apiClient)
	resp, r, err := api.UpdateNotebook(ctx, NotebookDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.UpdateNotebook`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.UpdateNotebook`:\n%s\n", responseContent)
}

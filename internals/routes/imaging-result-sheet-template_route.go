package routes

import (
	"telerad-core-module/internals/controllers"
	"telerad-core-module/internals/secure"
)

func ImagingResultSheetTemplateRoutes() {
	imagingResultSheetTemplate := "/imaging-result-sheet-template"
	staffCollection := getRoute(v1, staff).Group(imagingResultSheetTemplate, secure.CheckAuthorization())

	staffCollection.Get("", controllers.StaffGetPaginatedImagingResultSheetTemplates)
	staffCollection.Post("", controllers.StaffCreateImagingResultSheetTemplate)
	staffCollection.Get("/:objectId", controllers.StaffGetAImagingResultSheetTemplate)
	staffCollection.Put("/:objectId", controllers.StaffUpdateImagingResultSheetTemplate)
	staffCollection.Patch("/:objectId/activate", controllers.StaffActivateImagingResultSheetTemplate)
	staffCollection.Patch("/:objectId/deactivate", controllers.StaffDeactivateImagingResultSheetTemplate)
}

package routes

import (
	"telerad-core-module/internals/controllers"
	"telerad-core-module/internals/secure"
)

func ImagingResultTemplateRoutes() {
	imagingResultTemplate := "/imaging-result-template"
	staffCollection := getRoute(v1, staff).Group(imagingResultTemplate, secure.CheckAuthorization())

	staffCollection.Get("", controllers.StaffGetPaginatedImagingResultTemplates)
	staffCollection.Post("", controllers.StaffCreateImagingResultTemplate)
	// option cho form: loại chụp + bộ phận chụp
	staffCollection.Get("/actions/get-form-options", controllers.StaffGetImagingResultTemplateFormOptions)
	staffCollection.Get("/:objectId", controllers.StaffGetAImagingResultTemplate)
	staffCollection.Put("/:objectId", controllers.StaffUpdateImagingResultTemplate)
	staffCollection.Patch("/:objectId/activate", controllers.StaffActivateImagingResultTemplate)
	staffCollection.Patch("/:objectId/deactivate", controllers.StaffDeactivateImagingResultTemplate)
}

package location_controller

import (
	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/models"
	"github.com/kataras/iris/v12"
)

type GenericListRequest struct {
	Query  string `json:"query" example:""`
	Limit  int    `json:"limit" example:"10"`
	Offset int    `json:"offset" example:"0"`
}

// @Summary Get states
// @Description Retrieve states
// @Tags locations
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} controllers.GenericResponse{data=[]models.MGeoState} "States retrieved successfully"
// @Failure 404 {object} controllers.GenericResponse "States not found"
// @Router /api/locations/states [get]
func (c *controller) GetStates(ctx iris.Context) (*controllers.GenericResponse, error) {
	states, err := c.geoService.ListStates()
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: states}, nil
}

// @Summary Update location
// @Description Update location
// @Tags locations
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param location body models.MLocation true "Updated location"
// @Success 200 {object} controllers.GenericResponse{data=models.MLocation} "Location updated successfully"
// @Failure 400 {object} controllers.GenericResponse "Invalid request body"
// @Router /api/locations [put]
func (c *controller) Put(ctx iris.Context) (*controllers.GenericResponse, error) {
	location := &models.MLocation{}
	err := ctx.ReadJSON(&location)
	if err != nil {
		return nil, err
	}
	err = c.geoService.UpdateLocation(location)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: location}, nil
}

// @Summary Create a new location
// @Description Create a new location
// @Tags locations
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param location body models.MLocation true "Updated location"
// @Success 200 {object} controllers.GenericResponse{data=models.MLocation} "Location updated successfully"
// @Failure 400 {object} controllers.GenericResponse "Invalid request body"
// @Router /api/locations [post]
func (c *controller) Post(ctx iris.Context) (*controllers.GenericResponse, error) {
	location := &models.MLocation{}
	err := ctx.ReadJSON(&location)
	if err != nil {
		return nil, err
	}
	err = c.geoService.CreateLocation(location)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: location}, nil
}

// GetLocationsRequest represents the get locations request payload
type GetLocationsRequest struct {
	StateID int64  `json:"state_id" url:"state_id" example:"1"`
	Query   string `json:"query" url:"query" example:"New York"`
	Limit   int    `json:"limit" url:"limit" example:"10"`
	Offset  int    `json:"offset" url:"offset" example:"0"`
}

// @Sumgmary Get locations
// @Description Get locations by state or query name
// @Tags locations
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param state_id query int64 false "State ID" example:"1"
// @Param query query string false "Query" example:"تهران"
// @Param limit query int false "Limit" example:"10"
// @Param offset query int false "Offset" example:"0"
// @Success 200 {object} controllers.GenericResponse{data=[]models.MLocation} "Locations retrieved successfully"
// @Failure 400 {object} controllers.GenericResponse "Invalid request body"
// @Router /api/locations [get]
func (c *controller) Get(ctx iris.Context) (*controllers.GenericResponse, error) {

	request := GetLocationsRequest{}
	err := ctx.ReadQuery(&request)
	if err != nil {
		return nil, err
	}
	if request.Limit == 0 {
		request.Limit = 10
	}

	var locations []*models.MLocation
	if request.StateID != 0 {
		locations, err = c.geoService.ListLocationsByState(request.StateID, request.Limit, request.Offset, "")
	} else if request.Query != "" {
		locations, err = c.geoService.QueryLocations(request.Query, request.Limit, request.Offset, "")
	} else {
		locations, err = c.geoService.ListLocations(request.Limit, request.Offset, "")
	}
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: locations}, nil
}

func (c *controller) DeleteBy(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	err := c.geoService.DeleteLocation(id)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true}, nil
}

/*
func (c *controller) GetMilateryBase(ctx iris.Context) (*controllers.GenericResponse, error) {
	request := GenericListRequest{}
	err := ctx.ReadQuery(&request)
	if err != nil {
		return nil, err
	}
	if request.Limit == 0 {
		request.Limit = 10
	}
	var units []*models.MBranch
	if request.Query != "" {
		units, err = c.branchService.QueryBranches(request.Query, request.Limit, request.Offset, "")
	} else {
		units, err = c.geoService.ListMilateryBases()
	}
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: units}, nil
}

func (c *controller) GetByMilateryBase(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	units, err := c.geoService.GetMilateryBase(id)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: units}, nil
}

func (c *controller) PostMilateryBase(ctx iris.Context) (*controllers.GenericResponse, error) {
	unit := &models.MMilateryBase{}
	err := ctx.ReadJSON(&unit)
	if err != nil {
		return nil, err
	}
	if unit.ID == 0 {
		err = c.geoService.CreateMilateryBase(unit)
		if err != nil {
			return nil, err
		}
	} else {
		err = c.geoService.UpdateMilateryBase(unit)
		if err != nil {
			return nil, err
		}
	}

	return &controllers.GenericResponse{Success: true, Data: unit}, nil
}

func (c *controller) PutMilateryBase(ctx iris.Context) (*controllers.GenericResponse, error) {
	unit := &models.MMilateryBase{}
	err := ctx.ReadJSON(&unit)
	if err != nil {
		return nil, err
	}
	err = c.geoService.UpdateMilateryBase(unit)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: unit}, nil
}

func (c *controller) DeleteMilateryBaseBy(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	err := c.geoService.DeleteMilateryBase(id)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true}, nil
}
*/

func (c *controller) Options(ctx iris.Context) {
	ctx.StatusCode(200)
}

func (c *controller) OptionsBy(ctx iris.Context) {
	ctx.StatusCode(200)
}

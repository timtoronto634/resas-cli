// Package entity defines data entities that is widely used within product
package entity

// PopulationResponse is type of the resposne from RESAS api
type PopulationResponse struct {
	Message     string                  `json:"message"`
	Result      *PopulationResponseBody `json:"result"`
	StatusCode  int                     `json:"statusCode"`
	Description string                  `json:"description"`
}

// PopulationResponseBody is the body part of the resposne when the request was successful
type PopulationResponseBody struct {
	BoundaryYear int                `json:"boundaryYear"`
	Data         []*PopulationGroup `json:"data"`
}

// PopulationGroup is the group of data with the label of classification
type PopulationGroup struct {
	Label string        `json:"label"`
	Data  []*Population `json:"data"`
}

// Population is the each data of population
type Population struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

package request

import (
	"fmt"
)

// TODO: Create product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Company     string  `json:"company"`
	Location    string  `json:"location"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (req *CreateProductRequest) Validate() error {
	if req.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if req.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if req.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if req.Price <= 0 {
		return errParamIsRequired("price", "float64")
	}
	if req.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

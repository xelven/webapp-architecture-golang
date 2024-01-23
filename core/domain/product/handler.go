package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	err "webapp-core/core/domain/common/err"
)

type API struct {
	logger     *zerolog.Logger
	repository *Repository
	querier    *Querier
}

func New(logger *zerolog.Logger, db *gorm.DB) *API {
	return &API{
		logger:     logger,
		repository: NewRepository(db),
		querier:    NewQuerier(db),
	}
}

func (api *API) Create(response http.ResponseWriter, request *http.Request) {
	productForm := &PostProductForm{}
	if error := json.NewDecoder(request.Body).Decode(productForm); error != nil {

		err.BadRequest(response, err.RespJSONDecodeFailure)
		return
	}

	newProduct := productForm.FormToModel()
	newProduct.ID = uuid.New()

	newProduct, error := api.repository.Create(newProduct)
	if error != nil {
		err.ServerError(response, err.RespDBDataInsertFailure)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (api *API) GetList(w http.ResponseWriter, r *http.Request) {

	products, error := api.querier.GetList()
	if error != nil {
		err.ServerError(w, err.RespDBDataAccessFailure)
		return
	}

	if len(products) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if error := json.NewEncoder(w).Encode(products.ProductsToResponse()); error != nil {
		err.ServerError(w, err.RespJSONEncodeFailure)
		return
	}
}

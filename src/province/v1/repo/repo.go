package repo

import (
	"context"

	"github.com/moch-sudharmono/go-clean/tree/main/src/province/v1/model"
)

//Type Result repository
type ResultRepository struct {
	Result interface{}
	Error  error
}

type ProvinceRepository interface {
	Save(ctxReq context.Context, data model.Province) <-chan ResultRepository
	Update(ctxReq context.Context, data model.Province) <-chan ResultRepository
	FindProvinceById(ctxReq context.Context, id int) <-chan ResultRepository
	Delete(ctxReq context.Context, id int) <-chan ResultRepository
	GetProvinces(ctxReq context.Context, params *model.ParametersApplication) <-chan ResultRepository
	GetTotalProvices(ctxReq context.Context, params *model.ParametersApplication) <-chan ResultRepository
}

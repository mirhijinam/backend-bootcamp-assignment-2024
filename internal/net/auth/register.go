package auth

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	ht "github.com/ogen-go/ogen/http"
)

func (api API) RegisterPost(ctx context.Context, req generated.OptRegisterPostReq) (r generated.RegisterPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

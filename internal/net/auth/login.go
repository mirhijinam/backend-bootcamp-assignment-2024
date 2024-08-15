package auth

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	ht "github.com/ogen-go/ogen/http"
)

func (api API) LoginPost(ctx context.Context, req generated.OptLoginPostReq) (r generated.LoginPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

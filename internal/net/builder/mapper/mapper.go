package mapper

import (
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
)

var ReqStatus = map[generated.Status]models.Status{
	generated.StatusCreated:      models.StatusCreated,
	generated.StatusOnModeration: models.StatusOnModeration,
	generated.StatusApproved:     models.StatusApproved,
	generated.StatusDeclined:     models.StatusDeclined,
}

var RespStatus = map[models.Status]generated.Status{
	models.StatusCreated:      generated.StatusCreated,
	models.StatusOnModeration: generated.StatusOnModeration,
	models.StatusApproved:     generated.StatusApproved,
	models.StatusDeclined:     generated.StatusDeclined,
}

package houses

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"go.uber.org/zap"
)

func TestService_GetFlatsByHouseId(t *testing.T) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		ctx     context.Context
		houseId int
		role    models.Role
	}
	tests := []struct {
		name       string
		args       args
		want       []models.Flat
		wantErr    bool
		repoBehave func(repo MockhouseRepo)
	}{
		{
			name: "Success for role moderator",
			args: args{
				ctx:     context.Background(),
				houseId: 1,
				role:    models.RoleModerator,
			},
			repoBehave: func(repo MockhouseRepo) {
				repo.EXPECT().GetFlatsByHouseId(gomock.Any(), 1, models.RoleModerator).Return([]models.Flat{
					{
						Number:      1,
						HouseID:     1,
						Price:       1000,
						Rooms:       3,
						Status:      models.StatusApproved,
						ModeratorID: &uuid,
					},
				}, nil)
			},
			want: []models.Flat{
				{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusApproved,
					ModeratorID: &uuid,
				},
			},
			wantErr: false,
		},
		{
			name: "Fail for role moderator",
			args: args{
				ctx:     context.Background(),
				houseId: 1,
				role:    models.RoleModerator,
			},
			repoBehave: func(repo MockhouseRepo) {
				repo.EXPECT().GetFlatsByHouseId(gomock.Any(), 1, models.RoleModerator).Return(nil, errors.New("test"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := NewMockhouseRepo(ctrl)
			if tt.repoBehave != nil {
				tt.repoBehave(*repoMock)
			}

			s := Service{
				houseRepo: repoMock,
				logger:    zap.NewNop(),
			}

			got, err := s.GetFlatsByHouseId(tt.args.ctx, tt.args.houseId, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetFlatsByHouseId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetFlatsByHouseId() = %v, want %v", got, tt.want)
			}
		})
	}
}

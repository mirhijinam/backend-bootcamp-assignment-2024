package flats

import (
	"context"
	"errors"
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/pointer"
	"go.uber.org/zap"
)

func TestService_UpdateStatus(t *testing.T) {

	type args struct {
		ctx    context.Context
		params dto.FlatUpdateParams
	}

	tests := []struct {
		name        string
		args        args
		want        models.Flat
		wantErr     bool
		wantErrType error
		repoBehave  func(repoFlats MockflatsRepo, transactor Mocktransactor)
	}{
		{
			name: "Successful: created -> on_moderation",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusOnModeration,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusCreated,
					ModeratorID: nil,
				}, nil)

				repoFlats.EXPECT().UpdateStatus(gomock.Any(), dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusOnModeration,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				})
			},
			want: models.Flat{
				Number:      1,
				HouseID:     1,
				Price:       1000,
				Rooms:       3,
				Status:      models.StatusOnModeration,
				ModeratorID: pointer.ToPtr(uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9")),
			},
			wantErr:     false,
			wantErrType: nil,
		},
		{
			name: "Invalid transition: created -> approved",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusApproved,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusCreated,
					ModeratorID: nil,
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrInvalidTransition,
		},
		{
			name: "Invalid transition: created -> declined",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusDeclined,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusCreated,
					ModeratorID: nil,
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrInvalidTransition,
		},
		{
			name: "Already moderated: approved -> approved",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusApproved,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusApproved,
					ModeratorID: nil,
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrAlreadyModerated,
		},
		{
			name: "Already moderated: approved -> declined",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusDeclined,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusApproved,
					ModeratorID: nil,
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrAlreadyModerated,
		},
		{
			name: "Already moderated: declined -> approved",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusApproved,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusDeclined,
					ModeratorID: nil,
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrAlreadyModerated,
		},
		{
			name: "Already moderated: declined -> declined",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusDeclined,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusDeclined,
					ModeratorID: nil,
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrAlreadyModerated,
		},
		{
			name: "Another moderate",
			args: args{
				ctx: context.Background(),
				params: dto.FlatUpdateParams{
					Number:      1,
					HouseID:     1,
					Status:      models.StatusOnModeration,
					ModeratorID: uuid.MustParse("c3b3297b-7046-41bc-bdbd-7d405cade9d9"),
				},
			},
			repoBehave: func(repoFlats MockflatsRepo, transactor Mocktransactor) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					},
				)
				repoFlats.EXPECT().Get(gomock.Any(), 1, 1).Return(models.Flat{
					Number:      1,
					HouseID:     1,
					Price:       1000,
					Rooms:       3,
					Status:      models.StatusOnModeration,
					ModeratorID: pointer.ToPtr(uuid.MustParse("8ae6dd9d-88b8-4bce-b1ac-3252282b0225")),
				}, nil)
			},
			want:        models.Flat{},
			wantErr:     true,
			wantErrType: ErrAnotherModerating,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := NewMockflatsRepo(ctrl)
			transactorMock := NewMocktransactor(ctrl)

			if tt.repoBehave != nil {
				tt.repoBehave(*repoMock, *transactorMock)
			}

			s := Service{
				flatsRepo:  repoMock,
				logger:     zap.NewNop(),
				transactor: transactorMock,
			}

			got, err := s.UpdateStatus(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr && !errors.Is(err, tt.wantErrType) {
				t.Errorf("Service.UpdateStatus() error = %v, wantErr %v, wantErrType=%v", err, tt.wantErr, tt.wantErrType)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UpdateStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

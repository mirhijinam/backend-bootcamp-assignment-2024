package houses_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/houses"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetFlatsByHouseId(t *testing.T) {
	// Define test cases using table-driven approach
	tests := []struct {
		name          string
		houseID       int
		role          models.Role
		expectedFlats []models.Flat
		expectedErr   error
		mockSetup     func(mock pgxmock.PgxPoolIface)
	}{
		{
			name:    "Client role, approved flats",
			houseID: 1,
			role:    models.RoleClient,
			expectedFlats: []models.Flat{
				{Number: 4, HouseID: 1, Price: 100000, Rooms: 3, Status: models.StatusApproved},
			},
			expectedErr: nil,
			mockSetup: func(mock pgxmock.PgxPoolIface) {
				rows := pgxmock.NewRows([]string{"flat_number", "house_id", "price", "rooms", "status"}).
					AddRow(4, 1, 100000, 3, models.StatusApproved)
				mock.ExpectQuery("SELECT flat_number, house_id, price, rooms, status FROM flats WHERE").
					WithArgs(1, models.StatusApproved).
					WillReturnRows(rows)
			},
		},
		{
			name:    "Moderator role, all flats",
			houseID: 2,
			role:    models.RoleModerator,
			expectedFlats: []models.Flat{
				{Number: 2, HouseID: 2, Price: 150000, Rooms: 4, Status: models.StatusOnModeration},
				{Number: 3, HouseID: 2, Price: 200000, Rooms: 5, Status: models.StatusApproved},
			},
			expectedErr: nil,
			mockSetup: func(mock pgxmock.PgxPoolIface) {
				rows := pgxmock.NewRows([]string{"flat_number", "house_id", "price", "rooms", "status"}).
					AddRow(2, 2, 150000, 4, models.StatusOnModeration).
					AddRow(3, 2, 200000, 5, models.StatusApproved)
				mock.ExpectQuery("SELECT flat_number, house_id, price, rooms, status FROM flats WHERE").
					WithArgs(2).
					WillReturnRows(rows)
			},
		},
		{
			name:          "No flats found",
			houseID:       3,
			role:          models.RoleModerator,
			expectedFlats: nil,
			expectedErr:   nil,
			mockSetup: func(mock pgxmock.PgxPoolIface) {
				rows := pgxmock.NewRows([]string{"flat_number", "house_id", "price", "rooms", "status"})
				mock.ExpectQuery("SELECT flat_number, house_id, price, rooms, status FROM flats WHERE").
					WithArgs(3).
					WillReturnRows(rows)
			},
		},
		{
			name:          "SQL error",
			houseID:       4,
			role:          models.RoleModerator,
			expectedFlats: nil,
			expectedErr:   fmt.Errorf("repo.House.GetFlatsByHouseId: mock error"),
			mockSetup: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectQuery("SELECT flat_number, house_id, price, rooms, status FROM flats WHERE").
					WithArgs(4).
					WillReturnError(fmt.Errorf("mock error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockPool, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mockPool.Close()

			r := houses.New(mockPool)

			tt.mockSetup(mockPool)

			ctx := context.Background()

			// Act
			flats, err := r.GetFlatsByHouseId(ctx, tt.houseID, tt.role)

			// Assert
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedFlats, flats)
			}

			assert.NoError(t, mockPool.ExpectationsWereMet())
		})
	}
}

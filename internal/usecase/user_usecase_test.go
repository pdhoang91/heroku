package usecase

import (
	"fmt"
	"heroku/internal/delivery/http/model"
	"heroku/internal/entities"
	repo "heroku/internal/repository"
	"heroku/internal/usecase/mocks"
	"reflect"
	"testing"
)

func TestIUserUseCase_GetUserInfo(t *testing.T) {

	type fields struct {
		userRepository repo.UserRepository
	}
	type args struct {
		user_id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserInfo
		wantErr bool
	}{
		{
			name: "Happy case",
			fields: fields{
				userRepository: func() repo.UserRepository {
					mockRepo := &mocks.UserRepository{}
					// mock result here
					mockRepo.On("FindUserByID", 1).Return(&entities.User{
						ID:   1,
						Name: "Alice",
					}, nil)
					var mockAccount []*entities.Account
					mockAccount = append(mockAccount, &entities.Account{ID: 1, Name: "A銀行", Balance: 20000})
					mockRepo.On("GetUserAccounts", 1).Return(mockAccount, nil)
					return mockRepo
				}(),
			},
			args: args{
				user_id: 1,
			},
			want: &model.UserInfo{
				UserID:   1,
				Name:     "Alice",
				Accounts: []*entities.Account{{ID: 1, Name: "A銀行", Balance: 20000}},
				Balance:  20000,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IUserUseCase{
				UserRepository: tt.fields.userRepository,
				BalanceCalc:    &SimpleBalanceCalculator{},
			}

			got, err := s.GetUserInfo(tt.args.user_id)
			fmt.Println("got", got)
			fmt.Println("err", err)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}

}

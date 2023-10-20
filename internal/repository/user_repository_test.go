package repo

import (
	"fmt"
	"heroku/internal/entities"
	"reflect"
	"testing"
)

func TestUserRepository_FindUserByID(t *testing.T) {

	type args struct {
		user_id int
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.User
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				user_id: 1,
			},
			want: &entities.User{
				ID:         1,
				Name:       "Alice",
				AccountIDs: []int{1, 3, 5},
			},
			wantErr: false,
		},
		{
			name: "Case empty value",
			args: args{
				user_id: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewUserRepository()
			got, err := repo.FindUserByID(tt.args.user_id)
			fmt.Println("got", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUserRepository_GetUserAccounts(t *testing.T) {

	type args struct {
		user_id int
	}
	tests := []struct {
		name    string
		args    args
		want    []*entities.Account
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				user_id: 1,
			},
			want: []*entities.Account{
				{
					ID:      1,
					UserId:  1,
					Name:    "A銀行",
					Balance: 20000,
				},
				{
					ID:      3,
					UserId:  1,
					Name:    "C信用金庫",
					Balance: 120000,
				},
				{
					ID:      5,
					UserId:  1,
					Name:    "E銀行",
					Balance: 5000,
				},
			},
			wantErr: false,
		},
		{
			name: "Case empty value",
			args: args{
				user_id: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewUserRepository()
			got, err := repo.GetUserAccounts(tt.args.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package redis_test

import (
	"testing"

	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/internal/services/database/mock"
	"github.com/meoera/doorman/internal/services/database/redis"
	"github.com/meoera/doorman/pkg/hasher"
	"github.com/meoera/doorman/pkg/models"
)

func TestRedis_Connect(t *testing.T) {

	//making mockdb
	db := &mock.MockDatabase{}
	testpw, testsalt, err := hasher.HashPassword("testpassword123", nil)
	if err != nil {
		t.Errorf("An error occured in the preparation for the unit test: %v", err)
	}
	db.Connect([]*models.DatabaseUser{
		{
			Username: "testuser123",
			Email: "test@email123.com",
			PasswordHash: testpw,
			Salt: testsalt,
		},
	})

	type args struct {
		maindb      database.Database
		credentials []interface{}
	}
	tests := []struct {
		name    string
		args    args
		expect interface{}
		wantErr bool
	}{
		{
			name: "cfg nil input",
			args: args{
				db, nil,
			},
			wantErr: true,
			expect: redis.ErrUnspecifiedConfig,
		},
		{
			name: "maindb nil input",
			args: args{
				nil, []interface{}{&config.Redis{
					Host:               "127.0.0.1",
					Port:               6379,
					Username:           "default",
					Password:           "",
					Database:           0,
					StandartExpiration: 10,
				}},
			},
			wantErr: true,
			expect: redis.ErrUnspecifiedMainDatabase,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &redis.Redis{}
			err := db.Connect(tt.args.maindb, tt.args.credentials...)
			if (err != nil) && tt.wantErr == false || tt.expect != err  {
				t.Errorf("Redis.Connect() error = %v, wantErr %v, expeted: %v", err, tt.wantErr, tt.expect)
			}
		})
	}
}

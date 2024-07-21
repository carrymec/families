package person

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github/carrymec/families/mocks/mock_session"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func getErr(b bool) error {
	if b {
		return errors.New("500")
	}
	return nil
}

// Create a test logger
func NewTestLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}

func TestDao_CheckExistByName(t *testing.T) {
	type fields struct {
		lg            *zap.Logger
		sessionClient neo4j.SessionWithContext
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		runErr  bool
		wantErr bool
	}{
		{
			name: "run err",
			args: args{
				ctx:  context.Background(),
				name: "test",
			},
			runErr:  true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			withContext := mock_session.NewMockPersonSessionWithContext(ctrl)

			transaction := mock_session.NewMockManagedTransaction(ctrl)
			transaction.EXPECT().Run(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil, getErr(tt.runErr))

			// TODO ERR
			//withContext.EXPECT().ExecuteRead(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, work neo4j.ManagedTransactionWork) (interface{}, error) {
			//	return work(transaction)
			//})
			d := &Dao{
				lg:            NewTestLogger(),
				sessionClient: withContext,
			}
			got, err := d.CheckExistByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckExistByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckExistByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestCheckExistByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionClient := mock_person.NewMockSessionClient(ctrl)
	logger := zaptest.NewLogger(t)

	dao := &person.Dao{
		sessionClient: mockSessionClient,
		lg:            logger,
	}

	ctx := context.Background()
	name := "testname"

	mockTransaction := mock_person.NewMockManagedTransaction(ctrl)
	mockResult := mock_person.NewMockResultWithContext(ctrl)

	mockSessionClient.EXPECT().ExecuteRead(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, work neo4j.ManagedTransactionWork) (interface{}, error) {
			return work(mockTransaction)
		},
	)

	mockTransaction.EXPECT().Run(ctx, gomock.Any(), gomock.Any()).Return(mockResult, nil)
	mockResult.EXPECT().Next(ctx).Return(true)
	mockResult.EXPECT().Record().Return(neo4j.Record{
		Values: []any{int64(1)},
	})

	exists, err := dao.CheckExistByName(ctx, name)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !exists {
		t.Errorf("expected true, got false")
	}
}
*/

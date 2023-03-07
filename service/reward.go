package service

import (
	"context"
	"encoding/csv"
	"errors"
	"github.com/go-water/water"
	"io"
	"os"
)

type RewardRequest struct{}

type RewardService struct {
	*water.ServerBase
}

func (srv *RewardService) Handle(ctx context.Context, req *RewardRequest) (interface{}, error) {
	csvFile, err := os.Open("./content/reward.csv")
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	result := make([][]string, 0)
	for {
		row, er := csvReader.Read()
		if er == io.EOF {
			break
		} else if er != nil {
			return nil, er
		}

		result = append(result, row)
	}

	return result, nil
}

func (srv *RewardService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*RewardRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}

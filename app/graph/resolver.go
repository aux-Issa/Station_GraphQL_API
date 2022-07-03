package graph

import (
	"context"
	"errors"

	"github.com/aux-Issa/Station_GraphQL_API/graph/model"
	"github.com/aux-Issa/Station_GraphQL_API/models"
	"github.com/aux-Issa/Station_GraphQL_API/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *Resolver) getStationByCD(ctx context.Context, stationCd *int) (*model.Station, error) {
	db, err := utils.SetupDB()
	// models.StationConnsByStationCD(*stationCd)
	stations, err := models.StationConnsByStationCD(ctx, db, *stationCd)
	if err != nil {
		return nil, err
	}
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}
	first := stations[0]
	var beforeStation *model.Station
	if first.BeforeStationName != "" {
		beforeStation = &model.Station{
			LineName:    &first.LineName,
			StationCd:   first.BeforeStationCd,
			StationName: first.BeforeStationName,
			Address:     nil,
		}
	}
	var afterStation *model.Station
	if first.AfterStationName != "" {
		afterStation = &model.Station{
			LineName:    &first.LineName,
			StationCd:   first.AfterStationCd,
			StationName: first.AfterStationName,
			Address:     nil,
		}
	}
	transfers := make([]*model.Station, 0, len(stations))
	for _, v := range stations {
		if v.TransferStationName == "" {
			continue
		}
		transfers = append(transfers, &model.Station{
			LineName:    &v.TransferLineName,
			StationCd:   v.TransferStationCd,
			StationName: v.TransferStationName,
			Address:     nil,
		})
	}
	return &model.Station{
		LineName:        &first.LineName,
		StationCd:       first.StationCd,
		StationName:     first.StationName,
		Address:         &first.Address,
		BeforeStation:   beforeStation,
		AfterStation:    afterStation,
		TransferStation: transfers,
	}, nil
}

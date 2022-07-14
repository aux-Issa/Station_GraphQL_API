package graph

import (
	"context"
	"errors"
	"fmt"

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
	// stations, err := models.StationConnsByStationCD(ctx, db, *stationCd)
	stations, err := models.StationByCDsByStationCD(ctx, db, *stationCd)
	if err != nil {
		return nil, err
	}
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}
	first := stations[0]

	return &model.Station{
		LineName:    &first.LineName,
		StationCd:   first.StationCd,
		StationName: first.StationName,
		Address:     &first.Address,
	}, nil
}

func (r *Resolver) getStationByName(ctx context.Context, stationName *string) ([]*model.Station, error) {
	db, err := utils.SetupDB()
	if err != nil {
		return nil, err
	}

	stations, err := models.StationByNamesByStationName(ctx, db, *stationName)
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}
	res := make([]*model.Station, len(stations))
	for _, station := range stations {
		res = append(res, &model.Station{
			StationCd:   station.StationCd,
			StationName: station.StationName,
			LineName:    &station.LineName,
			Address:     &station.Address,
		})
	}
	return res, nil
}

func (r *Resolver) getTransferStaion(ctx context.Context, obj *model.Station) ([]*model.Station, error) {
	db, err := utils.SetupDB()
	if err != nil {
		return nil, err
	}
	stationCd := obj.StationCd
	records, err := models.TransfersByStationCD(ctx, db, stationCd)
	if err != nil {
		return nil, err
	}

	transferStations := make([]*model.Station, len(records))
	for _, record := range records {
		// &: メモリのaddressを示すので，型の＊と&はセットで使う
		fmt.Println(record, record.TransferStationCd, &record.TransferStationCd, record.TransferLineName, &record.TransferStationName)
		transferStations = append(transferStations, &model.Station{
			StationCd:   record.TransferStationCd,
			StationName: record.TransferStationName,
			LineName:    &record.TransferLineName,
			Address:     &record.TransferAddress,
		})
	}
	return transferStations, nil
}

func (r *Resolver) getBeforeStation(ctx context.Context, obj *model.Station) (*model.Station, error) {
	// todo: xo query for before station
	db, err := utils.SetupDB()
	if err != nil {
		return nil, err
	}
	cd := obj.StationCd
	stations, err := models.BeforeStationsByStationCD(ctx, db, cd)
	if err != nil {
		return nil, err
	}
	fmt.Println(stations, "いくつあるのだろう")
	station := stations[0]

	return &model.Station{
		StationCd:   station.StationCd,
		LineName:    &station.LineName,
		StationName: station.StationName,
		Address:     &station.Address,
	}, nil
}

func (r *Resolver) getAfterStation(ctx context.Context, obj *model.Station) []*model.Station {
	// todo: xo query for after station
	return nil
}

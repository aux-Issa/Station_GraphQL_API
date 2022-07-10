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
	stations, err := models.StationConnsByStationCD(ctx, db, *stationCd)
	// stations, err := models.StationByCDsByStationCD(ctx, db, *stationCd)
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
	// resp := make([]*model.Station, 0, len(records))
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

// TODO: afterStationのxo生成とresolverの実装
// TODO: bsforeStationnのxo生成とresolverの実装

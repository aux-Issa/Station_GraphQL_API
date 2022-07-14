// Package models contains generated code for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AfterStation represents a row from 'public.after_station'.
type AfterStation struct {
	StationCd            int    `json:"station_cd"`              // station_cd
	StationName          string `json:"station_name"`            // station_name
	StationGCd           int    `json:"station_g_cd"`            // station_g_cd
	Address              string `json:"address"`                 // address
	LineCd               int    `json:"line_cd"`                 // line_cd
	LineName             string `json:"line_name"`               // line_name
	AfterStationLineCd   int    `json:"after_station_line_cd"`   // after_station_line_cd
	AfterStationLineName string `json:"after_station_line_name"` // after_station_line_name
	AfterStationCd       int    `json:"after_station_cd"`        // after_station_cd
	AfterStationName     string `json:"after_station_name"`      // after_station_name
	AfterStationAddress  string `json:"after_station_address"`   // after_station_address
}

// AfterStationsByStationCD runs a custom query, returning results as AfterStation.
func AfterStationsByStationCD(ctx context.Context, db DB, stationCD int) ([]*AfterStation, error) {
	// query
	const sqlstr = ` ` +
		`select st3.station_cd, ` +
		`st3.station_name, ` +
		`st3.station_g_cd, ` +
		`st3.address, ` +
		`s2l.line_cd, ` +
		`s2l.line_name, ` +
		`COALESCE(s2l.line_cd, 0)     as after_station_line_cd, ` +
		`COALESCE(s2l.line_name, '')   as after_station_line_name, ` +
		`COALESCE(st3.station_cd, 0)   as after_station_cd, ` +
		`COALESCE(st3.station_name, '') as after_station_name, ` +
		`COALESCE(st3.address, '')      as after_station_address ` +
		` ` +
		`from station st ` +
		`inner join line li on st.line_cd = li.line_cd ` +
		`left outer join junction sjb on st.line_cd = sjb.line_cd and st.station_cd = sjb.station_cd2 ` +
		`left outer join junction sja on st.line_cd = sja.line_cd and st.station_cd = sja.station_cd1 ` +
		` ` +
		`left outer join station st2 on sjb.line_cd = st2.line_cd and sjb.station_cd1 = st2.station_cd ` +
		`left outer join line s2l on st2.line_cd = s2l.line_cd ` +
		` ` +
		`left outer join station st3 on sja.line_cd = st3.line_cd and sja.station_cd2 = st3.station_cd ` +
		` ` +
		`where st.station_cd = $1`
	// run
	logf(sqlstr, stationCD)
	rows, err := db.QueryContext(ctx, sqlstr, stationCD)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*AfterStation
	for rows.Next() {
		var as AfterStation
		// scan
		if err := rows.Scan(&as.StationCd, &as.StationName, &as.StationGCd, &as.Address, &as.LineCd, &as.LineName, &as.AfterStationLineCd, &as.AfterStationLineName, &as.AfterStationCd, &as.AfterStationName, &as.AfterStationAddress); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &as)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

package repository

import (
	"errors"
	"fmt"
	"github.com/anaabdi/location-mogo/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	databaseName   = "location"
	collectionName = "areas"
)

type Area interface {
	GetByLocation(lng, lat float64) (*model.Area, error)
}

type areaRepo struct {
	mgoSession *mgo.Session
}

func NewAreaRepo(mgoSession *mgo.Session) Area {
	return &areaRepo{
		mgoSession: mgoSession,
	}
}

func (areaRepo *areaRepo) GetByLocation(lng, lat float64) (*model.Area, error) {
	session := areaRepo.mgoSession.Clone()
	defer session.Close()

	locationDB := session.DB(databaseName)
	if locationDB == nil {
		return nil, errors.New("database location is not found")
	}

	collection := locationDB.C(collectionName)
	if collection == nil {
		return nil, errors.New("collection areas is not found")
	}

	var resp []model.Area
	err := collection.Find(queryCheckPointIntersecsPolygon(lng, lat)).All(&resp)
	if err != nil {
		return nil, fmt.Errorf("failed in query execution: %v", err)
	}

	if len(resp) == 0 {
		return nil, errors.New("No Area")
	}

	return &resp[0], nil
}

func queryCheckPointIntersecsPolygon(lng, lat float64) bson.M {
	return bson.M{
		"geometry": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
			},
		},
	}
}

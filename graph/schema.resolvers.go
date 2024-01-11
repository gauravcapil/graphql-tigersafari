package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"gaurav.kapil/tigerhall/auth"
	"gaurav.kapil/tigerhall/dbutils"
	"gaurav.kapil/tigerhall/graph/model"
	"github.com/99designs/gqlgen/graphql"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, userName string, password string, email string) (*model.UserData, error) {
	result := []*model.UserData{}
	dbutils.DbConn.Where(&model.UserData{UserName: userName}).First(&result)

	if len(result) != 0 {
		return nil, fmt.Errorf("user already exists with username: %s ", userName)
	}

	dbutils.DbConn.Where(model.UserData{Email: email}).First(&result)

	if len(result) != 0 {
		return nil, fmt.Errorf("user already exists with email: %s", email)
	}

	dbutils.DbConn.Create(&model.UserDataWithPassword{
		UserName: userName,
		Password: password,
	})
	value := dbutils.DbConn.Create(&model.UserData{
		UserName: userName,
		Email:    email,
	})
	return value.Statement.Model.(*model.UserData), nil
}

// CreateNewTiger is the resolver for the createNewTiger field.
func (r *mutationResolver) CreateNewTiger(ctx context.Context, userName string, name string, dateOfBirth string, lastSeen string, seenAtLat string, seenAtLon string, photo graphql.Upload) (int, error) {
	if err := auth.Authenticate(ctx); err != nil {
		return 0, err
	}
	photolink := name + "_" + lastSeen
	value := dbutils.DbConn.Create(&model.TigerData{
		UserName:    userName,
		Name:        name,
		DateOfBirth: dateOfBirth,
		Sightings: []*model.Sighting{
			{SeenAt: lastSeen, SeenAtLat: seenAtLat, SeenAtLon: seenAtLon,
				PhotoLocation: dbutils.PhotoFolder + photolink},
		},
	})
	log.Printf("about to upload the file")
	stream, readErr := ioutil.ReadAll(photo.File)
	if readErr != nil {
		fmt.Printf("error from file %v", readErr)
	}
	fileName := dbutils.GetPhotoDir() + "/" + photolink
	fileErr := ioutil.WriteFile(fileName, stream, 0644)
	if fileErr != nil {
		fmt.Printf("file err %v", fileErr)
	}
	log.Printf("file writted with name: %s", fileName)
	return value.Statement.Model.(*model.TigerData).ID, nil
}

// CreateNewSighting is the resolver for the createNewSighting field.
func (r *mutationResolver) CreateNewSighting(ctx context.Context, userName string, name string, seenAt string, seenAtLat string, seenAtLon string, photo graphql.Upload) (int, error) {
	if err := auth.Authenticate(ctx); err != nil {
		return 0, err
	}
	result := model.TigerData{}
	dbutils.DbConn.Where(model.TigerData{Name: name}).First(&result)
	photolink := name + "_" + seenAt
	value := dbutils.DbConn.Create(&model.Sighting{TigerID: result.ID, SeenAt: seenAt, SeenAtLat: seenAtLat, SeenAtLon: seenAtLon, PhotoLocation: dbutils.PhotoFolder + photolink})
	log.Printf("about to upload the file")
	stream, readErr := ioutil.ReadAll(photo.File)
	if readErr != nil {
		fmt.Printf("error from file %v", readErr)
	}
	fileName := dbutils.GetPhotoDir() + "/" + photolink
	fileErr := ioutil.WriteFile(fileName, stream, 0644)
	if fileErr != nil {
		fmt.Printf("file err %v", fileErr)
	}
	log.Printf("file writted with name: %s", fileName)
	newSightingID := value.Statement.Model.(*model.Sighting).ID
	return newSightingID, fileErr
}

// ListTigers is the resolver for the listTigers field.
func (r *queryResolver) ListTigers(ctx context.Context, offset *int, limit *int) ([]*model.TigerData, error) {
	result := []*model.TigerData{}
	dbutils.SetDefaults(&offset, &limit)
	dbutils.DbConn.Offset(*offset).Limit(*limit).Find(&result)
	log.Printf("listing: %d", len(result))
	return result, nil
}

// ListAllSightings is the resolver for the listAllSightings field.
func (r *queryResolver) ListAllSightings(ctx context.Context, tigerID int, offset *int, limit *int) ([]*model.Sighting, error) {
	result := []*model.Sighting{}
	dbutils.SetDefaults(&offset, &limit)
	dbutils.DbConn.Offset(*offset).Limit(*limit).Find(&result, model.Sighting{TigerID: tigerID})
	log.Printf("listing: %d", len(result))
	return result, nil
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, userName string, password *string) (*model.LoginData, error) {
	result := []*model.UserDataWithPassword{}
	dbutils.DbConn.Where(&model.UserDataWithPassword{UserName: userName, Password: *password}).First(&result)
	if len(result) == 0 {
		return nil, fmt.Errorf("authentication Failed")
	}
	token := auth.GenerateSecureToken(255)

	dbutils.DbConn.Where(&model.LoginData{Userid: result[0].ID}).Delete(&model.LoginData{})
	value := dbutils.DbConn.Create(&model.LoginData{
		Userid:     result[0].ID,
		Token:      token,
		Expiration: "Never", // Periodic Expiration is yet to be implemented
		Error:      nil,
	})
	return value.Statement.Model.(*model.LoginData), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

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
	"gaurav.kapil/tigerhall/emailserver"
	"gaurav.kapil/tigerhall/graph/model"
	"gaurav.kapil/tigerhall/models"
	"gaurav.kapil/tigerhall/utils"
	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm/clause"

	"github.com/umahmood/haversine"
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
func (r *mutationResolver) CreateNewTiger(ctx context.Context, userName string, name string, dateOfBirth string, lastSeen string, seenAtLat float64, seenAtLon float64, photo graphql.Upload) (int, error) {
	if err := auth.Authenticate(ctx); err != nil {
		return 0, err
	}
	photolink := utils.Generatephotofilename(name, lastSeen)
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
	fileName := dbutils.GetPhotoDir() + "\\" + photolink
	fileErr := ioutil.WriteFile(fileName, stream, 0644)
	if fileErr != nil {
		fmt.Printf("file err %v", fileErr)
	}
	log.Printf("file writted with name: %s", fileName)
	return value.Statement.Model.(*model.TigerData).ID, nil
}

// CreateNewSighting is the resolver for the createNewSighting field.
func (r *mutationResolver) CreateNewSighting(ctx context.Context, userName string, name string, seenAt string, seenAtLat float64, seenAtLon float64, photo graphql.Upload) (int, error) {
	if err := auth.Authenticate(ctx); err != nil {
		return 0, err
	}

	newCoordinates := haversine.Coord{Lat: seenAtLat, Lon: seenAtLon}

	result := model.TigerData{}
	dbutils.DbConn.Where(model.TigerData{Name: name}).First(&result)
	lastSeetAtresult := model.Sighting{}
	dbutils.DbConn.Where(model.Sighting{TigerID: result.ID}).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "seen_at"}, Desc: true}).First(&lastSeetAtresult)

	oldCoordinates := haversine.Coord{Lat: lastSeetAtresult.SeenAtLat, Lon: lastSeetAtresult.SeenAtLon}

	_, km := haversine.Distance(oldCoordinates, newCoordinates)

	log.Printf("The tiger is seen at : %f kms away", km)
	if km < dbutils.MinDistanceToConsider {
		return 0, fmt.Errorf("this Sightings was %f km from previous sighting,"+
			"and anything less than %f kms is ignored.", km, dbutils.MinDistanceToConsider)
	}
	photolink := utils.Generatephotofilename(name, seenAt)
	value := dbutils.DbConn.Create(&model.Sighting{TigerID: result.ID, SeenAt: seenAt, SeenAtLat: seenAtLat, SeenAtLon: seenAtLon, PhotoLocation: dbutils.PhotoFolder + "/" + photolink})
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
	emailserver.Notify(&models.MailData{User: result.UserName, Sighting: *value.Statement.Model.(*model.Sighting)})
	return newSightingID, fileErr
}

// ListTigers is the resolver for the listTigers field.
func (r *queryResolver) ListTigers(ctx context.Context, offset *int, limit *int) ([]*model.TigerDataResponse, error) {
	dbutils.SetDefaults(&offset, &limit)
	query := fmt.Sprintf("select tiger_id,seen_at,seen_at_lat,seen_at_lon, photo_location, user_name, name, date_of_birth"+
		" from (select distinct on (s.tiger_id) * from sightings s order by s.tiger_id, s.seen_at desc)"+
		" join tiger_data t on t.id = tiger_id limit %d offset %d;", *limit, *offset)
	log.Println(query)
	result := []*model.TigerDataResponse{}
	tx := dbutils.DbConn.Raw(query)
	tx.Scan(&result)
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

package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/mongo-tut/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	odm.AbstractRepository[models.Users]
}

func (u *UserRepository) FindNearestUserByEmail(useremail string) (*models.Users, error) {

	var user *models.Users
	currchan, errchan := u.FindOne(bson.M{"email": useremail})
	select {
	case finduser := <-currchan:
		user = finduser
	case err := <-errchan:
		return nil, err
	}

	var location = user.Location

	currchan, errchan = u.FindOne(
		bson.M{
			"email": bson.M{"$ne": useremail},
			"location": bson.M{
				"$nearSphere": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": location.Coordinates,
					},
				},
			},
		})

	select {
	case userfound := <-currchan:
		return userfound, nil
	case err := <-errchan:
		return nil, err
	}
}

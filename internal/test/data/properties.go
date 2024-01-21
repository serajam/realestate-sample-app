package data

import (
	"context"
	"embed"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/samber/lo"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/adapters/datastore/repositories"
	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

//go:embed geo-addresses.zip
var propsSource embed.FS

func GenProperties(
	repo repositories.PropertyRepository, uRepo repositories.UserRepository, logger *zap.SugaredLogger,
) {
	logger = logger.Named("data generator")
	err := repo.DeleteAll(context.Background())
	if err != nil {
		logger.Errorw("failed to delete all properties", zap.Error(err))
	}

	fs, err := propsSource.Open("geo-addresses.zip")
	if err != nil {
		logger.Errorw("failed to open file", zap.Error(err))
		return
	}
	defer fs.Close()

	logger.Debug("generate properties from file")
	propes, err := LoadProps(fs)
	if err != nil {
		logger.Errorw("failed to open file", zap.Error(err))
		return
	}
	logger.Debugw("loaded test properties", zap.Int("count", len(propes)))
	logger.Debug("inserting data")

	gofakeit.Seed(0)

	uRepo.CreateUser(
		context.Background(), &domain.User{
			Name:            "John",
			Surname:         "Doe",
			Email:           "john.doe@example.com",
			Password:        []byte(gofakeit.LoremIpsumWord()),
			IsVerifiedEmail: false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	)

	workersNum := 4 * runtime.GOMAXPROCS(0)
	workers := make(chan struct{}, workersNum)

	for i := 0; i <= len(propes)-1; i++ {
		testProp := propes[i]
		workers <- struct{}{}
		go func(prop *TestProperty) {
			err = repo.Create(context.Background(), DomainPropFromTestProp(prop))
			if err != nil {
				fmt.Println("smth went wrong", err)
			}

			err = repo.Create(context.Background(), DomainPropFromTestProp(prop))
			if err != nil {
				fmt.Println("smth went wrong", err)
			}
			<-workers
		}(&testProp)
	}

	logger.Debug("inserting done")
}

func DomainPropFromTestProp(prop *TestProperty) *properties.Property {
	lat, _ := strconv.ParseFloat(prop.Lat, 64)
	lon, _ := strconv.ParseFloat(prop.Lon, 64)

	return &properties.Property{
		UserID: 1,
		Location: properties.Point{
			lat, lon,
		},
		Price:         gofakeit.Float32Range(10000, 5000000),
		PriceCurrency: "EUR",
		Address:       prop.FullAddress,
		Country:       prop.Address.Country,
		City:          prop.City,
		State:         prop.Address.State,
		Street:        prop.Address.Neighbourhood,
		ZipCode:       prop.Address.Postcode,
		HouseNumber:   prop.Address.HouseNumber,
		Neighborhood:  prop.Address.Neighbourhood,
		HomeSize:      gofakeit.Float32Range(50, 500),
		LotSize:       gofakeit.Float32Range(50, 500),
		YearBuild:     uint16(gofakeit.Year()),
		Bedroom:       uint8(gofakeit.IntRange(1, 5)),
		Bathroom:      uint8(gofakeit.IntRange(1, 5)),
		LivingSize:    gofakeit.Float32Range(50, 500),
		Floor:         uint8(gofakeit.IntRange(1, 100)),
		TotalFloors:   uint8(gofakeit.IntRange(1, 100)),
		PropertyType:  uint8(gofakeit.IntRange(2, 3)),
		HomeType:      uint8(gofakeit.IntRange(1, 9)),
		Condition:     uint8(gofakeit.IntRange(1, 3)),
		BrokerName:    gofakeit.Name(),
		Description:   gofakeit.Sentence(100),
		Active:        lo.ToPtr(gofakeit.Bool()),
		HasImages:     lo.ToPtr(gofakeit.Bool()),
		HasGarage:     lo.ToPtr(gofakeit.Bool()),
		HasVideo:      lo.ToPtr(gofakeit.Bool()),
		Has3DTour:     lo.ToPtr(gofakeit.Bool()),
		TotalParking:  uint8(gofakeit.IntRange(0, 5)),
		HasAC:         lo.ToPtr(gofakeit.Bool()),
		Appliance:     lo.ToPtr(gofakeit.Bool()),
		Heating:       uint8(gofakeit.IntRange(1, 3)),
		PetsAllowed:   lo.ToPtr(gofakeit.Bool()),
		CreatedAt:     gofakeit.DateRange(time.Unix(1678995683, 0), time.Unix(1686944483, 0)),
		UpdatedAt:     gofakeit.DateRange(time.Unix(1678995683, 0), time.Unix(1686944483, 0)),
	}
}

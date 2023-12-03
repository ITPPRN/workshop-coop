package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/log"

	"service2/modules/entities/events"
	"service2/modules/entities/models"
)

type dogService struct {
	dogRepo         models.DogRepository
	userRepo        models.UserRepository
	Redis           *redis.Client
	producerUsecase models.UseCaseProducer
}

func NewDogService(
	dogRepo models.DogRepository,
	userRepo models.UserRepository,
	Redis *redis.Client,
	producerUsecase models.UseCaseProducer,
) models.DogUsecase {
	genData(dogRepo)
	return &dogService{dogRepo, userRepo, Redis, producerUsecase}
}

// ////////////////////////////////////////////////////////////////////////////////////
func (u *dogService) GetDogs() (dogs []models.Dog, err error) {
	key := "service:GetDogs"

	//redis get
	if dogsJson, err := u.Redis.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(dogsJson), &dogs) == nil {
			log.Debug("Read data from: redis")
			return dogs, nil
		}
	}

	//repository
	log.Debug("Read data from: database")
	dogDB, err := u.dogRepo.GetDogs()
	if err != nil {
		log.Error(err)
		return nil, errors.New("couldn't get drinks data")
	}

	for _, d := range dogDB {
		dogs = append(dogs, models.Dog{
			ID:               d.ID,
			Name:             d.Name,
			Temperament:      d.Temperament,
			LifeSpan:         d.LifeSpan,
			Origin:           d.Origin,
			WeightImperial:   d.WeightImperial,
			WeightMetric:     d.WeightMetric,
			HeightImperial:   d.HeightImperial,
			HeightMetric:     d.HeightMetric,
			BredFor:          d.BredFor,
			BreedGroup:       d.BreedGroup,
			ReferenceImageID: d.ReferenceImageID,
		})
	}

	//redis set
	if data, err := json.Marshal(dogs); err == nil {
		u.Redis.Set(context.Background(), key, string(data), time.Minute*1)
	}

	return dogs, nil
}

func dogToRawMessage(dog models.Dog) (json.RawMessage, error) {
	dogDB, err := json.Marshal(dog)
	if err != nil {
		return nil, err
	}

	rawMessage := json.RawMessage(dogDB)

	if !json.Valid(rawMessage) {
		return nil, errors.New("Not valid JSON")
	}

	return rawMessage, nil
}

// /////////////////////////////////////////////////////////////////////////////////////
func (u *dogService) UserReadData(userId, dogId uint) (json.RawMessage, error) {

	dogExist := u.dogRepo.DogExists(dogId)
	if !dogExist {
		return nil, errors.New("dog data not found")
	}

	userExist := u.userRepo.UserExists(userId)
	if !userExist {
		return nil, errors.New("user data not found")
	}

	dogDB, err := u.dogRepo.FindDogByID(dogId)
	if err != nil {
		log.Error(err)
		return nil, errors.New("get dog data faile")
	}

	rawMass, err := dogToRawMessage(*dogDB)
	if err != nil {
		return nil, err
	}
	log.Info(fmt.Sprintf("ui: %v,di: %v", userId,dogId))
	err = u.producerUsecase.UserReaded(&events.UserReadedEvent{
		UserId:     userId,
		DogId:      dogId,
		DogDetails: rawMass,
		TimeStamp:  time.Now(),
	})
	if err != nil {
		log.Error("con't not produce event UserReadData, error: ", err)
	}
	return rawMass, nil
}

// ///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////



func genData(dogRepo models.DogRepository) {
	// เรียกใช้ฟังก์ชัน genDataJson เพื่อดึงข้อมูลจาก API
	dogsData, err := genDataJson()
	if err != nil {
		fmt.Println("Error getting data from API:", err)
		return
	}

	// ตรวจสอบว่ามีข้อมูลหมาในฐานข้อมูลแล้วหรือไม่
	hasDog, err := dogRepo.HasDog()
	if err != nil {
		log.Debug("Error checking for HasDog:", err)
		return
	}

	if hasDog {
		log.Debug("Data already exists in the database")
		return
	}

	// Unmarshal JSON data into a slice of maps
	var rawData []map[string]interface{}
	err = json.Unmarshal(dogsData, &rawData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Check if there is at least one row of data
	if len(rawData) == 0 {
		fmt.Println("No data found")
		return
	}

	// Iterate through the data and save each dog to the database
	for _, data := range rawData {

		// Check if any required field is missing or nil
		if data["id"] == nil || data["name"] == nil || data["temperament"] == nil ||
			data["life_span"] == nil || data["origin"] == nil ||
			data["bred_for"] == nil || data["breed_group"] == nil ||
			data["reference_image_id"] == nil || data["weight"] == nil || data["height"] == nil {
			//fmt.Println("Incomplete or nil data for dog, skipping:", data)
			continue
		}
		// Convert data to the Dog struct

		dog := models.Dog{
			ID:               uint(data["id"].(float64)),
			Name:             data["name"].(string),
			Temperament:      data["temperament"].(string),
			LifeSpan:         data["life_span"].(string),
			Origin:           data["origin"].(string),
			WeightImperial:   data["weight"].(map[string]interface{})["imperial"].(string),
			WeightMetric:     data["weight"].(map[string]interface{})["metric"].(string),
			HeightImperial:   data["height"].(map[string]interface{})["imperial"].(string),
			HeightMetric:     data["height"].(map[string]interface{})["metric"].(string),
			BredFor:          data["bred_for"].(string),
			BreedGroup:       data["breed_group"].(string),
			ReferenceImageID: data["reference_image_id"].(string),
		}

		//fmt.Printf("Dog: %+v\n", dog)

		//Save the dog to the database
		if err := dogRepo.CreateDog(&dog); err != nil {
			fmt.Println("Error creating dog:", err)
		} else {
			fmt.Printf("Dog %s created successfully\n", dog.Name)
		}

	}

}

func genDataJson() ([]byte, error) {
	apiKey := "live_7kiR6U5DRd6GIZFj3Zw3g8qgsMw7CGrQx14TGkONDOTdoyk2cGIhtFAhOpsIhGMz"
	apiURL := "https://api.thedogapi.com/v1/breeds"

	// ทำ HTTP GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// เพิ่ม Header ที่ใส่ API Key
	req.Header.Add("x-api-key", apiKey)

	// สร้าง HTTP Client และส่ง request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// เพิ่มส่วนการตรวจสอบ HTTP Status Code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: Unexpected status code %d\n", response.StatusCode)
		return nil, fmt.Errorf("Unexpected status code %d", response.StatusCode)
	}

	// อ่านข้อมูล response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return body, nil
}

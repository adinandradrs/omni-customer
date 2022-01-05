package utility

import (
	"encoding/json"
	"fmt"
	"omni-customer/model/response"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/inconshreveable/log15.v2"
)

func GenerateToken(userId uint, customerLoginResponse *response.CustomerLoginResponse, tokenExpiration uint, tokenSecret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["data"] = customerLoginResponse
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenExpiration)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSecret))
}

func ValidateToken(cache *redis.Client, tokenString string, tokenSecret string) bool {
	customerLoginResponse := response.CustomerLoginResponse{}
	token, error := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, result := token.Method.(*jwt.SigningMethodHMAC); !result {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if error != nil {
		return false
	}
	claims, result := token.Claims.(jwt.MapClaims)
	mapstructure.Decode(claims["data"], &customerLoginResponse)
	dataCache, redisError := cache.Get(CACHE_CUSTOMER_LOGIN + customerLoginResponse.Email).Result()
	json.Unmarshal([]byte(dataCache), &customerLoginResponse)
	log15.Info("Redis token = ", customerLoginResponse.Token)
	log15.Info("Given token = ", tokenString)
	if customerLoginResponse.Token != tokenString {
		log15.Error("validate failed because token is mismatch")
		return false
	}
	if result && token.Valid && redisError == nil {
		return true
	} else {
		log15.Error("validate failed with a broken error")
		return false
	}

}

func GetCustomerInfo(cache *redis.Client, tokenString string, tokenSecret string) (response.CustomerLoginResponse, error) {
	customerLoginResponse := response.CustomerLoginResponse{}
	token, error := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, result := token.Method.(*jwt.SigningMethodHMAC); !result {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if error != nil {
		panic(ERR_MSG_UNAUTHORIZED)
	}
	claims, result := token.Claims.(jwt.MapClaims)
	if result && token.Valid {
		mapstructure.Decode(claims["data"], &customerLoginResponse)
		log15.Info("GetCustomerInfo.customerLoginResponse = ", customerLoginResponse)
	} else {
		return response.CustomerLoginResponse{}, fmt.Errorf(ERR_MSG_UNAUTHORIZED)
	}
	redisError := cache.Get(CACHE_CUSTOMER_LOGIN + customerLoginResponse.Email).Err()
	if redisError == redis.Nil || redisError != nil {
		log15.Error("Redis is empty for customer with email = ", customerLoginResponse.Email)
		return response.CustomerLoginResponse{}, fmt.Errorf(ERR_MSG_UNAUTHORIZED)
	}
	return customerLoginResponse, nil
}

package transport

import (
	"encoding/json"
	"go-rate-limit/errs"
	"go-rate-limit/redis"
	"time"
	"math"
	"log"
)

const maxCalls = 8

// APIUsage is use as redis value to check if a usage has exceed their limit
type APIUsage struct {
	Count     int
	LastUsage time.Time
}

// CheckRateLimit is use to check if a user has reached maximum API calls.
// This function utilizes the token bucket algorithm to support burst usage.
func CheckRateLimit(email string) error {
	val := redis.Client.Get(email)
	// Initialize new key if can't find given email
	if val.Err() != nil && val.Err().Error() == "redis: nil" {
		userUsage := APIUsage{
			Count: maxCalls,
			LastUsage: time.Now(),
		}

		data, err := json.Marshal(userUsage)
		if err != nil {
			return err
		}

		err = redis.Client.Set(email, data, 0).Err()
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		// Find time diff and add API call allocation given the difference
		data, err := redis.Client.Get(email).Result()
		if err != nil {
			return err
		}
		var userData APIUsage
		err = json.Unmarshal([]byte(data), &userData)
		if err != nil {
			return err
		}
		timeDiff := time.Now().Sub(userData.LastUsage).Seconds()
		roundedTimeDiff := math.Floor(timeDiff / 30)

		newCount := float64(userData.Count) + roundedTimeDiff - 1
		if newCount < 0 {
			return errs.MaxUsageErr
		}

		if newCount > maxCalls {
			newCount = maxCalls
		}

		userData.Count = int(newCount)
		userData.LastUsage = time.Now()
		newData, err := json.Marshal(userData)
		if err != nil {
			return err
		}

		err = redis.Client.Set(email, newData, 0).Err()
	}

	return nil
}

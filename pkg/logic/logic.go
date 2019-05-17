package logic

import (
	"fmt"
	"time"

	"github.com/getninjas/spot-advisor/pkg/api"
	"github.com/getninjas/spot-advisor/pkg/structure"
)

// Const to fix the warmup decreaseLogic. This number control if the requestID is valid
const timeStampCount = 2

// Start function work like a entrypoint to start all logic system
func Start() {
	spotID := api.GetSpotID()
	lock := make([]int64, len(spotID))

	for {
		for i, ID := range spotID {
			scaling := api.ScalingName(ID)
			timeNow := time.Now()
			unstablePoints := api.QueryDataRequest(structure.QueryData(
				structure.CloudDimension(ID),
			))

			trigger := structure.GetTrigger(unstablePoints)

			if !structure.MakeLock(lock[i]) {
				if trigger {
					scaleUp(scaling, unstablePoints)
					lock[i] = timeNow.Unix()
				}
			}

			if lock[i] > timeStampCount && !trigger {
				scaleDown := decreaseLogic(scaling, lock[i])
				if scaleDown {
					lock[i] = 0
				}
			} else {
				fmt.Println("normal -", scaling)
			}
		}
		time.Sleep(time.Duration((structure.Speed())) * time.Second)
	}
}

func scaleUp(scaling string, unstablePoints int64) bool {
	if api.IncreaseGroup(scaling, unstablePoints) {
		fmt.Println("enabling-fallback -", scaling)
		return true
	}
	fmt.Println("error-enabling-fallback -", scaling)
	return false
}

func decreaseLogic(scaling string, i int64) bool {
	if structure.CheckLock(i) {
		go api.DecreaseGroup(scaling)
		fmt.Println("disabling-fallback -", scaling)
		return true
	}
	fmt.Println("wainting-fallback -", scaling)
	return false
}

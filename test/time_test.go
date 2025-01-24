package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateTimeout(t *testing.T) {
	//This hold time now
	timeNow := time.Now()
	timeUnix := timeNow.Unix()

	//Time code expiration offsetting
	nextExpireTime := timeUnix + 5*60*1000

	fmt.Println(timeNow, timeUnix, nextExpireTime)
}

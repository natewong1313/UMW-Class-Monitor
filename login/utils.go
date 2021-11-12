package login

import (
	"time"

	"github.com/natewong1313/UMW-Class-Monitor/logger"
)

func (loginTask *LoginTask) handleErr(errMsg string, err error){
	logger.Err(errMsg, err)
	time.Sleep(time.Millisecond * time.Duration(loginTask.ErrDelay))
}
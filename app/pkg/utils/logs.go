package utils

import (
	"fmt"
	"time"
)

func LogSuccess(msg string, args ...interface{}) {
	currentTime := time.Now()
	formattedMsg := fmt.Sprintf(msg, args...)
	fmt.Printf(Gray+"%s"+Reset+"["+Cyan+"INFO"+Reset+"] -> %s\n", currentTime.Format("[2006-01-02|15:04:05.000000]"), formattedMsg)
}

func LogInfo(msg string, args ...interface{}) {
	currentTime := time.Now()
	formattedMsg := fmt.Sprintf(msg, args...)
	fmt.Printf(Gray+"%s"+Reset+"["+Cyan+"INFO"+Reset+"] -> %s\n", currentTime.Format("[2006-01-02|15:04:05.000000]"), formattedMsg)
}

func LogWarning(msg string, args ...interface{}) {
	currentTime := time.Now()
	formattedMsg := fmt.Sprintf(msg, args...)
	fmt.Printf(Gray+"%s"+Reset+"["+Yellow+"WARNING"+Reset+"] -> %s\n", currentTime.Format("[2006-01-02|15:04:05.000000]"), formattedMsg)
}

func LogError(msg string, args ...interface{}) {
	currentTime := time.Now()
	formattedMsg := fmt.Sprintf(msg, args...)
	fmt.Printf(Gray+"%s"+Reset+"["+Red+"ERROR"+Reset+"] -> %s\n", currentTime.Format("[2006-01-02|15:04:05.000000]"), formattedMsg)
}

func LogFatal(msg string, args ...interface{}) {
	currentTime := time.Now()
	formattedMsg := fmt.Sprintf(msg, args...)
	fmt.Printf(Gray+"%s"+Reset+"["+Red+"FATAL"+Reset+"] -> %s\n", currentTime.Format("[2006-01-02|15:04:05.000000]"), formattedMsg)
}

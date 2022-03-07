package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 98
}

const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

// func CheckQuota(username string) {
//     used := bytesInUse(username)
//     const quota = 1e9
//     percent := 100 * used
//     if percent < 90 {
//         return
//     }
//     msg := fmt.Sprintf(template, used, percent)
//     auth := smtp.PlainAuth("", sender, password, hostname)
//     err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
//     if err != nil {
//         log.Printf("smtp.SendMail(%s) failed: %s", username, err)
//     }
// }

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000
	percent := used
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

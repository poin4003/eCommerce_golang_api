package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprint("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// convert UUID to string, remove -
	uuidString := strings.ReplaceAll((newUUID).String(), "-", "")
	return strconv.Itoa(userId) + "clitoken" + uuidString
}

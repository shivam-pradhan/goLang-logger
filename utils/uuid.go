package utils

import "github.com/google/uuid"

// GetClientRequestUUID ...
/**
 * @return string a random uuid for each request
**/
func GetClientRequestUUID() string {
	return "clientreq_" + uuid.New().String()
}

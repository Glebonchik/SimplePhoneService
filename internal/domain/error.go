package domain

import "log"

func FormatErr(operation string, err error) error {
	log.Printf("%s failed cause of: %v", operation, err)
	return err
}

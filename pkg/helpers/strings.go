package helpers

import "github.com/gofrs/uuid"

func ListUUIDToListString(uuids []uuid.UUID) []string {
	var is []string
	for _, i := range uuids {
		is = append(is, i.String())
	}
	return is
}

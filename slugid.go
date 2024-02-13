package slugid

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

func EncodeUUID(uuid_ uuid.UUID) string {
	return base64.RawURLEncoding.EncodeToString(uuid_[:])
}

func EncodeNewUUID() string {
	u, _ := uuid.NewRandom()
	return EncodeUUID(u)
}

func DecodeUUID(slug string) (uuid.UUID, error) {
	uuid_, err := base64.RawURLEncoding.DecodeString(slug)
	if err != nil {
		return uuid.UUID{}, err
	}
	if len(uuid_) != 16 {
		return uuid.UUID{}, fmt.Errorf("incorrect length of uuid")
	}
	return uuid.UUID(uuid_), nil
}

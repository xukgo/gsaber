package randomUtil

import (
	"encoding/hex"
	"github.com/google/uuid"
	"strings"
)

func CreateTightUid() string {
	u, err := uuid.NewV7()
	if err != nil {
		u = uuid.New()
	}
	return FormatTightUid(u)
}

func FormatTightUid(u uuid.UUID) string {
	var sb strings.Builder
	sb.Grow(32)

	// Convert the UUID bytes to a hex string in one go, without hyphens
	hexString := make([]byte, hex.EncodedLen(len(u)))
	hex.Encode(hexString, u[:])
	sb.Write(hexString)

	return sb.String()
}

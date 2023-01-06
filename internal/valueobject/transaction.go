package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	// all values lowercase since they are immutable
	ID        uuid.UUID
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}

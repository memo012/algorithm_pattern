package pool

import (
	"errors"
	"math"
)

const (
	// DEFAULT_ANTS_POOL_SIZE is the default capacity for a default goroutine pool.
	DEFAULT_ANTS_POOL_SIZE = math.MaxInt32

	// DEFAULT_CLEAN_INTERVAL_TIME is the interval time to clean up goroutines.
	DefaultCleanIntervalTime = 1

	// CLOSED represents that the pool is closed.
	CLOSED = 1
)

var (
	// Error types for the Ants API.
	//---------------------------------------------------------------------------

	// ErrInvalidPoolSize will be returned when setting a negative number as pool capacity.
	ErrInvalidPoolSize = errors.New("invalid size for pool")

	// ErrInvalidPoolExpiry will be returned when setting a negative number as the periodic duration to purge goroutines.
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	// ErrPoolClosed will be returned when submitting task to a closed pool.
	ErrPoolClosed = errors.New("this pool has been closed")
)


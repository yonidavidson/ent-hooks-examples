package hook

import "context"

// Syncer is the interface that wraps Sync logic.
type Syncer interface {
	Sync(ctx context.Context, cacheID int)
}

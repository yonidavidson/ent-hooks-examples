package hook

import "context"

type Syncer interface {
	Sync(ctx context.Context, cloudID int)
}

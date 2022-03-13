package cloud

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/yonidavidson/ent-side-effect-hooks-example/ent"
)

// Syncer is used to sync cloud information in a different context.
type Syncer struct {
	client  *ent.Client
	wg      *sync.WaitGroup
	enabled bool
	sync.Mutex
	ctx context.Context
}

// NewSyncer returns a new disabled Syncer.
func NewSyncer() *Syncer {
	return &Syncer{
		wg:      &sync.WaitGroup{},
		enabled: false,
	}
}

// Start enables the SchemaSyncer to handle Sync requests.
func (s *Syncer) Start(ctx context.Context, c *ent.Client) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.client = c
	s.ctx = ctx
	s.enabled = true
}

// Sync implements a request for syncing the specific cloud ID.
func (s *Syncer) Sync(ctx context.Context, cloudID int) {
	s.start(cloudID)
}

// start refresh data for  a specific cloud id in a different context.
func (s *Syncer) start(cloudID int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if !s.enabled {
		return
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		fmt.Printf("start sync for id %d\n", cloudID)
		time.Sleep(2 * time.Second)
		rs := rand.NewSource(time.Now().UnixNano())
		r := rand.New(rs)
		if err := s.client.Cloud.UpdateOneID(cloudID).SetWalks(r.Intn(100)).Exec(s.ctx); err != nil {
			fmt.Printf("cloud/syncer: failed to sync %d, %e\n", cloudID, err)
		}
	}()
}

// Close provides a graceful shutdown by not allowing any new requests to start and waiting until all active requests complete.
func (s *Syncer) Close() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.enabled = false
	s.wg.Wait()
}

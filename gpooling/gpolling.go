package gpooling

import (
	"github.com/panjf2000/ants/v2"
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"google.golang.org/appengine/log"
)

// Pool - pooling struct
type Pool struct {
	antsPool *ants.Pool
}

// IPool - pooling interface
type IPool interface {
	Submit(task func())
	Release()
	Running() int
}

// Init - init pooling
func NewPooling(maxPoolSize int) (*Pool, error) {
	pool, err := ants.NewPool(maxPoolSize, ants.WithNonblocking(false), ants.WithPanicHandler(func(data interface{}) {
		log.Errorf(helpers.ContextWithTimeOut(), "%s", data)
	}))
	if err != nil {
		return nil, err
	}
	return &Pool{
		antsPool: pool,
	}, nil
}

// Release - release all gorotine
func (p *Pool) Release() {
	p.antsPool.Release()
}

// Running - returns the number of the currently running goroutines.
func (p *Pool) Running() int {
	return p.antsPool.Running()
}

// Submit - submit a task to this pool
func (p *Pool) Submit(task func()) {
	p.antsPool.Submit(task)
}

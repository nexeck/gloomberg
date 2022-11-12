package gloomberg

import (
	"github.com/benleb/gloomberg/internal/chainwatcher"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/models/wallet"
	"github.com/benleb/gloomberg/internal/nodes"
	ossw "github.com/benleb/gloomberg/internal/osstreamwatcher"
	"github.com/benleb/gloomberg/internal/web"
	"github.com/ethereum/go-ethereum/common"
)

type Gloomberg struct {
	ChainWatcher   *chainwatcher.ChainWatcher
	GloomWeb       *web.GloomWeb
	Nodes          *nodes.Nodes
	StreamWatcher  *ossw.OSStreamWatcher
	Watcher        *models.Watcher
	WebEventStream *web.EventStream

	CollectionDB *collections.CollectionDB
	OwnWallets   *wallet.Wallets

	WatchUsers *models.WatcherUsers

	// WatchGroups *models.WatchGroups
	// WatchRules  *models.WatchRules

	OutputQueues map[string]chan *collections.Event
	QueueSlugs   chan common.Address
	BasicMIWs    map[common.Address]int
}

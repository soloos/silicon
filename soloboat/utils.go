package soloboat

import (
	"soloos/common/snet"
	"sort"
	"strings"
)

func SortSNetPeerJSON(nodes []snet.PeerJSON) {
	sort.Slice(nodes, func(i, j int) bool {
		var r = strings.Compare(nodes[i].PeerID, nodes[j].PeerID)
		if r != 0 {
			return r == 1
		}
		return r == strings.Compare(nodes[i].PeerID, nodes[j].PeerID)
	})
}

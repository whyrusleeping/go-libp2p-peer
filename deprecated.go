// Deprecated: Use github.com/libp2p/go-libp2p-core/peer instead.
package peer

import moved "github.com/libp2p/go-libp2p-core/peer"

var (
	// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ErrEmptyPeerID instead.
	ErrEmptyPeerID = moved.ErrEmptyPeerID
	// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ErrNoPublicKey instead.
	ErrNoPublicKey = moved.ErrNoPublicKey
)

// Deprecated: need to break the interface it's impossible to alias a var.
// Use github.com/libp2p/go-libp2p-core/peer.AdvanceEnableInlining instead.
// var AdvancedEnableInlining = true

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.ID instead.
type ID = moved.ID

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromString instead.
var IDFromString = moved.IDFromString

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromBytes instead.
var IDFromBytes = moved.IDFromBytes

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDB58Decode instead.
var IDB58Decode = moved.IDB58Decode

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDB58Encode instead.
var IDB58Encode = moved.IDB58Encode

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDHexDecode instead.
var IDHexDecode = moved.IDHexDecode

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDHexEncode instead.
var IDHexEncode = moved.IDHexEncode

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromPublicKey instead.
var IDFromPublicKey = moved.IDFromPublicKey

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDFromPrivateKey instead.
var IDFromPrivateKey = moved.IDFromPrivateKey

// Deprecated: Use github.com/libp2p/go-libp2p-core/peer.IDSlice instead.
type IDSlice = moved.IDSlice

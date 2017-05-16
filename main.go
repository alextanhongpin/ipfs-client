// This program is the ipfs client that will connect to the server
package main

import (
	"fmt"
	"io"
	"os"

	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	// peer "github.com/ipfs/go-ipfs/p2p/peer"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	peer "gx/ipfs/QmdS9KpbDyPrieswibZhkod1oXqRwZJrUPzxCofAMWpFGq/go-libp2p-peer"

	// "code.google.com/p/go.net/context" to "context"
	"context"
)

func main() {
	// Check if the peer ID is provided as argument
	if len(os.Args) < 2 {
		fmt.Println("Please give a peer ID as an argument")
		return
	}

	peerID := os.Args[1]
	target, err := peer.IDB58Decode(peerID)
	if err != nil {
		panic(err)
	}

	// Basic ipfsnode setup
	r, err := fsrepo.Open("~/.ipfs2") // See README
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}

	// NewNode returns a node that will listen to the peer
	nd, err := core.NewNode(ctx, cfg)

	if err != nil {
		panic(err)
	}

	fmt.Printf("I am a peer %s dialing %s\n", nd.Identity, target)

	con, err := corenet.Dial(nd, target, "/app/johndoe")
	if err != nil {
		fmt.Print(err)
		return
	}

	io.Copy(os.Stdout, con)
}

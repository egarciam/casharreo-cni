package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	ID              int
	State           string
	VoteCount       int
	Mutex           sync.Mutex
	Peers           []*Node
	ElectionTimeout time.Duration
}

func (n *Node) RunElectionTimer() {
	for {
		time.Sleep(n.ElectionTimeout)
		n.InitiateElection()
	}
}

func (n *Node) InitiateElection() {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	n.State = "Candidate"
	n.VoteCount = 1
	for _, peer := range n.Peers {
		go n.RequestVote(peer)
	}
}

func (n *Node) RequestVote(peer *Node) {
	peer.Mutex.Lock()
	defer peer.Mutex.Unlock()

	if peer.State != "Leader" {
		n.VoteCount++
		if n.VoteCount > len(n.Peers)/2 {
			n.BecomeLeader()
		}
	}
}

// BecomeLeader
func (n *Node) BecomeLeader() {
	n.State = "Leader"
	fmt.Printf("Node %d became the leader\n", n.ID)
}

func CreateCluster(nodeCount int) []*Node {
	nodes := make([]*Node, nodeCount)
	for i := range nodes {
		nodes[i] = &Node{
			ID:              i,
			State:           "Follower",
			ElectionTimeout: time.Duration(rand.Intn(150)+150) * time.Millisecond,
		}
	}
	for i := range nodes {
		peers := make([]*Node, 0, nodeCount-1)
		for j := range nodes {
			if i != j {
				peers = append(peers, nodes[j])
			}
		}
		nodes[i].Peers = peers
	}
	return nodes
}

func main() {
	nodes := CreateCluster(5)
	for _, node := range nodes {
		go node.RunElectionTimer()
	}

	time.Sleep(5 * time.Second)
}

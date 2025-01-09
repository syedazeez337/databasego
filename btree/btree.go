package btree

import (
	"encoding/binary"
	"log"
)

const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	// assert(node1max <= BTREE_PAGE_SIZE) //maximum KV
	if node1max > BTREE_PAGE_SIZE {
		log.Fatalf("node1max (%d) exceeds BTREE_PAGE_SIZE (%d)", node1max, BTREE_PAGE_SIZE)
	}
}

// In memory datatypes
type BNode []byte

// decouple data structure from IO
type BTree struct {
	// pointer (a nonzero page number)
	root uint64
	// callbacks for managing on-disk pages
	get func(uint64) []byte // dereference a pointer
	new func([]byte) uint64 // allocate a new page
	del func(uint64)        // deallocate a new page
}

// Decode the node format

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}

// child pointers
func (node BNode) getPtr(idx uint16) uint64 {
	// assert(idx < node.nkeys())
	pos := HEADER + 8 * idx
	return binary.LittleEndian.Uint64(node[pos:])
}
package address_curve

import "crypto/sha3"


// deriveLeafNodes generates leaf nodes for the Merkle tree using the private key.
// It creates 512 leaf nodes by hashing the private key with an index from 0 to 511.
// Each leaf node is a 256-bit hash derived from the private key and the index.
func deriveLeafNodes(privKey []byte) [][]byte {
	leaves := make([][]byte, 512)
	for i := 0; i < 512; i++ {
		hasher := sha3.New256()
		hasher.Write(privKey)
		hasher.Write([]byte{byte(i)})
		leaves[i] = hasher.Sum(nil)
	}
	return leaves
}

// buildMerkleTree constructs a Merkle tree from the given leaves.
// It returns the root hash of the tree.
func buildMerkleTree(leaves [][]byte) []byte {
	if len(leaves) == 0 {
		return nil
	}
	nodes := leaves

	for len(nodes) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) {
				nodes = append(nodes, nodes[i])
			}
			h := sha3.New256()
			h.Write(nodes[i])
			h.Write(nodes[i+1])
			nextLevel = append(nextLevel, h.Sum(nil))
		}
		nodes = nextLevel
	}

	return nodes[0]
}

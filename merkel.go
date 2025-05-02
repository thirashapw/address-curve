package address_curve

import "crypto/sha3"

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

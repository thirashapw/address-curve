# address-curve
Creating secure private keys using symmetric SHA3 hashing + Ed25519 and Merkle trees

## Process flow
```bash PrivateKey (64 bytes)
                    ↓
    512 x (SHA3(privateKey + index))
                    ↓
Build Merkle Tree (pairwise SHA3 hashing)
                    ↓
            Get 1 Merkle Root
                    ↓
    Address =Base58(SHA3(Merkle Root))
```


## Visual Diagram
```bash           
          Merkle Root
               |
     ---------------------
    |                     |
   Hash1                 Hash2
  /    \                /     \
Leaf1 Leaf2         Leaf3  Leaf4
... and so on up from 512 leaves
```
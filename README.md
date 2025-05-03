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

## Example
```bash
account, err := address_curve.CreateAccount()  
if err != nil {  
    return fmt.Errorf("failed to generate private key: %v", err)  
}  

accountInfo := fmt.Sprintf(  
    "private key generated successfully: %s\n"+  
    "public key generated successfully: %s\n"+  
    "address generated successfully: %s",  
    account.PrivateKey,  
    account.PublicKey,  
    account.Address,  
)  

```


```bash
loaded_account, err := address_curve.ComputeAddress(account.PrivateKey)  
if err != nil {  
    return fmt.Errorf("failed to compute private key: %v", err)  
}  

loadedInfo := fmt.Sprintf(  
    "private key loaded successfully: %s\n"+  
    "public key loaded successfully: %s\n"+  
    "address loaded successfully: %s",  
    loaded_account.PrivateKey,  
    loaded_account.PublicKey,  
    loaded_account.Address,  
)  
```
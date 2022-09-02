
Hi Prof Buchanan, 

Thank you for the article!

On the section about Big Integers, I'd like to just propose some adjustments on the code:

Original:
```Go
fmt.Printf("p== %s\n", prime)
p := new(big.Int).Exp(x1, x2, prime) // x1^x2 (mod prime)
fmt.Printf("v1^v2 (mod p)= %s\n", prime)
```

Modified:
```Go
fmt.Printf("prime = %s\n", prime)
p := new(big.Int).Exp(x1, x2, prime) // x1^x2 (mod prime)
fmt.Printf("v1^v2 (mod prime)= %s\n", p)
```
- On the first `Printf`,  I'm changing the `p` to `prime`.
- On the second `Printf`,  I'm just changing the variable inside the `Printf` from `prime` to `p` and also changing the string inside from `(mod p)` to `(mod prime)`.



---

Hi Prof Buchanan, 
Thank you for the article!
I'd like to just propose some adjustments to the code:
```Go
fmt.Printf("p = %s\n", prime)
result := new(big.Int).Exp(x1, x2, prime) // x1^x2 (mod prime)
fmt.Printf("v1^v2 (mod p)= %s\n", result)
```
I changed the name of the variable to no confound the prime in the module, and changed the variables that is passed in the print


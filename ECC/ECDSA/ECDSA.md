# ECDSA

内容来自[wiki](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm)、[FIPS PUB 186-4](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-4.pdf)

## Signature generation algorithm

> Alice creates a key pair, consisting of a private key integer $d_{A}$, randomly selected in the interval ${\displaystyle [1,n-1]}$; and a public key curve point ${\displaystyle Q_{A}=d_{A}\times G}$. 
>
> For Alice to sign a message ${\displaystyle m}$, she follows these steps:
>
> 1. Calculate ${\displaystyle e={\textrm {HASH}}(m)}$, where HASH is a cryptographic hash function, such as SHA-2.
> 2. Let ${\displaystyle z}$ be the ${\displaystyle L_{n}}$ **leftmost** bits of ${\displaystyle e}$, where ${\displaystyle L_{n}}$ is the **bit length of the group order** ${\displaystyle n}$.
> 3. Select a **cryptographically secure random** integer ${\displaystyle k}$ from ${\displaystyle [1,n-1]}$.
> 4. Calculate the curve point ${\displaystyle (x_{1},y_{1})=k\times G}$.
> 5. Calculate ${\displaystyle r=x_{1}\,{\bmod {\,}}n}$. If ${\displaystyle r=0}$, go back to step 3.
> 6. Calculate ${\displaystyle s=k^{-1}(z+rd_{A})\,{\bmod {\,}}n}$. If ${\displaystyle s=0}$, go back to step 3.
> 7. The signature is the pair ${\displaystyle (r,s)}$.

没次签名，在第三步中，需要选择一个随机数k。也就是说，同样的messag和私钥进行签名，结果会是不一样的。而且这个k要保证随机，不然会有泄漏$d_A$的风险。

## Signature verification algorithm

> For Bob to authenticate Alice's signature, he must have a copy of her public-key curve point ${\displaystyle Q_{A}}$. Bob can verify ${\displaystyle Q_{A}}$ is a valid curve point as follows:
>
> 1. Check that ${\displaystyle Q_{A}}$ is not equal to the identity element ${\displaystyle O}$, and its coordinates are otherwise valid
> 2. Check that ${\displaystyle Q_{A}}$ lies on the curve
> 3. Check that ${\displaystyle n\times Q_{A}=O}$

先验证公钥的合法性。

> After that, Bob follows these steps:
>
> 1. Verify that ${\displaystyle r}$ and ${\displaystyle s}$ are integers in ${\displaystyle [1,n-1]}$. If not, the signature is invalid.
> 2. Calculate ${\displaystyle e={\textrm {HASH}}(m)}$, where HASH is the same function used in the signature generation.
> 3. Let ${\displaystyle z}$ be the ${\displaystyle L_{n}}$ leftmost bits of ${\displaystyle e}$.
> 4. Calculate ${\displaystyle w=s^{-1}\,{\bmod {\,}}n}$.
> 5. Calculate ${\displaystyle u_{1}=zw\,{\bmod {\,}}n}$ and ${\displaystyle u_{2}=rw\,{\bmod {\,}}n}$.
> 6. Calculate the curve point ${\displaystyle (x_{1},y_{1})=u_{1}\times G+u_{2}\times Q_{A}}$. If ${\displaystyle (x_{1},y_{1})=O}$ then the signature is invalid.
> 7. The signature is valid if ${\displaystyle r\equiv x_{1}{\pmod {n}}}$, invalid otherwise.

比较常规的验签的过程。

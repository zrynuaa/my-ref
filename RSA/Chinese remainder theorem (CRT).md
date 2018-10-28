# 中国剩余定理 Chinese remainder theorem

## 历史

中国剩余定理是一个求解一元线性同余方程组的方法，一元线性同余方程组问题最早可见于中国南北朝时期（公元5世纪）的数学著作《孙子算经》卷下第二十六题，叫做“物不知数”的问题中，原文如下：

> 有物不知其数，三三数之剩二，五五数之剩三，七七数之剩二。问物几何？
>

即，一个整数除以三余二，除以五余三，除以七余二，求这个整数。《孙子算经》中首次提到了同余方程组问题，以及以上具体问题的解法，因此在中文数学文献中也将 **中国剩余定理** 称为 **孙子定理**。



## 描述

中国剩余定理说明某一范围内的整数可通过它的一组剩余类数来重构，这组剩余类数是对该整数用一组两两互素的整数取模得到的。

给出以下的一元线性同余方程组：
$$
\begin{equation}
(S):\left\{
\begin{aligned}
x & \equiv & a_1 (mod\ m_1) \\
x & \equiv & a_2 (mod\ m_2) \\
\vdots \\
x & \equiv & a_n (mod\ m_n)
\end{aligned}
\right.
\end{equation}
$$
**有解的判定条件**：

> 如果整数 $$m_1, m_2, ... , m_n$$ 中任意两个数互质，则对任意的整数：$$ a_1, a_2, ... , a_n$$，方程组有解。

**通解**可以用如下方式构造得到：

> * 设 $$M=m_{1}\times m_{2}\times \cdots \times m_{n}=\prod _{i=1}^{n}m_{i} $$ 是整数$$m_1, m_2, ... , m_n$$ 的乘积，
>
>   并设$$M_{i}=M/m_{i},\;\;\forall i\in \{1,2,\cdots ,n\}$$ 即 $$M_{i}$$ 是除了 $$m_i$$ 以外的$$n − 1$$ 个整数的乘积。
>
> * 设 $$t_{i}=M_{i}^{-1}$$ 为 $$M_{i}$$ 模 $$m_{i}$$ 的数论倒数：$$t_{i}M_{i}\equiv 1{\pmod {m_{i}}},\;\;\forall i\in \{1,2,\cdots ,n\}.$$ 
>
> * 方程组 $$(S)$$ 的通解形式为： $$x=kM+a_{1}t_{1}M_{1}+a_{2}t_{2}M_{2}+\cdots +a_{n}t_{n}M_{n}=kM+\sum _{i=1}^{n}a_{i}t_{i}M_{i},\quad k\in \mathbb {Z} .$$ 
>
>   在模M的意义下，方程组 $$ (S)$$ 只有一个解：$$ x=a_{1}t_{1}M_{1}+a_{2}t_{2}M_{2}+\cdots +a_{n}t_{n}M_{n}=\sum _{i=1}^{n}a_{i}t_{i}M_{i}\ (mod\ M).$$ 



## 证明

**正确性：**

对任何 $$ i\in \{1,2,\cdots ,n\}$$ ，由于 $$ \forall j\in \{1,2,\cdots ,n\},\;j\neq i,\;\;\operatorname {gcd} (m_{i},m_{j})=1$$ ，

所以： 
$$
\operatorname {gcd} (m_{i},M_{i})=1.
$$
这说明存在整数 $$ t_{i}$$ 使得$$ t_{i}M_{i}\equiv 1{\pmod {m_{i}}}$$ ，即$$M_i$$ 存在关于$$m_i$$ 的逆$$t_i$$ 。考察乘积 $$a_{i}t_{i}M_{i}$$ 可知：
$$
a_i t_i M_i \equiv a_{i}*1 \equiv a_{i}\ (mod\ m_i)
$$

$$
a_i t_i M_i \equiv 0\ (mod\ m_j),\;\forall j\in \{1,2,\cdots ,n\},\;j\neq i
$$

所以$$ x_{0} = a_{1}t_{1}M_{1}+a_{2}t_{2}M_{2}+\cdots +a_{n}t_{n}M_{n}​$$ 满足：
$$
x_0 \equiv a_i t_i M_i + \sum _{j\neq i}^{n}a_{j}t_{j}M_{j} \equiv a_i + \sum _{j\neq i}^{n}0 \equiv a_i\ (mod\ m_i),\;\forall i\in \{1,2,\cdots ,n\}
$$
这说明$$x_0$$ 就是方程组$$ (S)$$ 的一个解。

**唯一性：**

假设$$x_{1}$$ 和$$ x_{2}$$ 都是方程组$$(S)$$ 的解，那么：
$$
x_1 - x_2 \equiv 0 \pmod {m_i}, \; \;   \forall i \in \{1, 2, \cdots , n\}.
$$
$$
x_1 - x_2 = k_im_i
$$

而又因为$$ m_{1},m_{2},\ldots ,m_{n}$$ 两两互质，这说明$$ M=\prod _{i=1}^{n}m_{i}$$ 整除$$ x_{1}-x_{2}$$ . 所以方程组$$ (S)$$ 的任何两个解之间必然相差$$ M$$ 的整数倍。所有形为：
$$
kM+\sum _{i=1}^{n}a_{i}t_{i}M_{i},\quad k\in \mathbb {Z} .
$$
方程组$$ (S)$$所有的解的集合就是：
$$
x=kM+a_{1}t_{1}M_{1}+a_{2}t_{2}M_{2}+\cdots +a_{n}t_{n}M_{n}=kM+\sum _{i=1}^{n}a_{i}t_{i}M_{i},\quad k\in \mathbb {Z} .
$$
在模$$M$$ 的意义下，方程组 $$ (S)$$ 只有一个解：
$$
x_0 = a_{1}t_{1}M_{1}+a_{2}t_{2}M_{2}+\cdots +a_{n}t_{n}M_{n}=\sum _{i=1}^{n}a_{i}t_{i}M_{i}.
$$

## 应用

### RSA-CRT

利用中国剩余定理快速解密和签名

1、将输入变换到CRT域
$$
x_p \equiv x\ mod\ p
$$

$$
x_q \equiv x\ mod\ q
$$

$$
d_p	\equiv d\ mod\ (p-1)
$$

$$
d_q	\equiv d\ mod\ (q-1)
$$

2、CRT域运算
$$
y = x^d
$$

$$
y_p \equiv y\ (mod p)\equiv x^d\ (mod\ p) \equiv (x\ mod\ p)^{d\ mod\ (p-1)}\ (mod\ p) \equiv x_p^{d_p}\ (mod\ p)
$$

$$
y_q \equiv y\ (mod q) \equiv x^d\ (mod\ q) \equiv (x\ mod\ q)^{d\ mod\ (q-1)}\ (mod\ q) \equiv x_q^{d_q}\ (mod\ q)
$$

线性同余方程组可表示为：
$$
\begin{equation}
(S):\left\{
\begin{aligned}
y & \equiv & y_p (mod\ p) \\
y & \equiv & y_q (mod\ q) 
\end{aligned}
\right.
\end{equation}
$$
那么根据中国剩余定理的结果：
$$
y \equiv y_p*(q^{-1}\ mod\ p)*q + y_q*(p^{-1}\ mod\ q)*p\ (mod\ n),\ n=pq
$$
辅助公式：
$$
y \equiv a_1t_1M_1 + a_2t_2M_2\ (mod\ n)
$$

运用中国剩余定理来加速加密和签名过程的计算，优点是整个过程可以提高4倍的速度；同时缺点也是明显的，需要存储秘密参数p、q，会加大泄漏密钥的风险。
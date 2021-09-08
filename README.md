# gocrypto
## 概述
代码采用PBKDF2加密算法，通常类似用户密码这样需要密文存储的串，使用该算法进行加密比较安全。

**PBKDF2**(Password-Based Key Derivation Function) 是一个用来导出密钥的函数，常用于生成加密的密码。原理是通过 string 和 salt 进行 hash 加密，然后将结果作为 salt 与 string 再进行 hash，多次重复此过程，生成最终的密文。如果重复的次数足够大（几千数万次），破解的成本就会变得很高。而盐值的添加也会增加“彩虹表”攻击的难度。

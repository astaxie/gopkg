# crypto/subtle包函数列表

Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly.  
subtle包是在加密的代码经常使用的一些方法，但是需要仔细考虑并正确使用。

- [func ConstantTimeByteEq(x, y uint8) int](ConstantTimeByteEq.md) 
- [func ConstantTimeCompare(x, y []byte) int](ConstantTimeCompare.md)  
- [func ConstantTimeCopy(v int, x, y []byte)](ConstantTimeCopy.md)  
- [func ConstantTimeEq(x, y int32) int](ConstantTimeEq.md)  
- [func ConstantTimeSelect(v, x, y int) int](ConstantTimeSelect.md)
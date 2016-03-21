# 旋转字符串
要求时间复杂度为O(n),空间复杂度为O(1)

## 三步反转法
把字符串前2个字节移动到后面去 “abcdef” -> "cdefab"
* 1.将原字符串分成2部分，X:ab，Y:cdef
* 2.将X反转X->X^T,"ab"->"ba";将Y反转Y->Y^T,"cdef"->"fedc"
* 3.(X^TY^T)^T,"bafedc"->"cdefab"


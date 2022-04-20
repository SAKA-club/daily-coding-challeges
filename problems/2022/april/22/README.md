# Arrays and String 1.8
#### April 22, 2022

**Zero Matrix**: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
column are set to 0.

```
// A 2x2 matrix
// Visual
//    11    =>    01
//    01          00
ZeroMatrix("1101") // should return ""0100""

// A 4x4 matrix
// Visual 
//    ab12          a012
//    e0fa    =>    0000
//    1111          1011
//    beef          b0ef
ZeroMatrix("ab12e0fa1111beef") // should return "a01200001011b0ef"

// A 2x6 matrix
// Visual 
//    123456          123450
//    543210    =>    000000
RotateMatrix("123456543210") // should return "123450000000"
```
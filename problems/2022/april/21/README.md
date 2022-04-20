# Arrays and String 1.7
#### April 21, 2022

**Rotate Matrix**: Given an image represented by an NxN matrix, where each pixel in the image is 4
bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

```
// A 1 pixel image rotated clockwise
// Visual
//    ab    =>    ca
//    cd          db
RotateMatrix("abcd", true) // should return ""cadb""

// A 1 pixel image rotated counter-clockwise
// Visual
//    ab    =>    bd
//    cd          ac
RotateMatrix("abcd", false) // should return "bdac"

// a 4 pixel image rotated clockwise
// Visual 
//    aabb          ccaa
//    aabb    =>    ccaa
//    ccdd          ddbb
//    ccdd          ddbb
RotateMatrix("aabbaabbccddccdd", false) // should return "ccccaaaaddddbbbb"

// a 4 pixel image rotated clockwise
// Visual 
//    aabb          bbdd
//    aabb    =>    bbdd
//    ccdd          aacc
//    ccdd          aacc
RotateMatrix("aaaabbbbccccdddd", true) // should return "bbbbddddaaaacccc"
```
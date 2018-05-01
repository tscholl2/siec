package ff

var (
	doublep = Element{0xfff3ff3cf6e79f3d, 0x7fffffffffffffff, 0xfffffffff9ffcf3e, 0x3fffffffffffffff}
)

// Note: only works if 2*a < 2^256. One way to garuntee this works is:
//
//    b := double(a)
//    if a[3]>>63 == 0 {
//      b = add(doublep, normalize(a))
//    }
//
func double(a Element) (b Element) {
	b[0] = a[0] << 1
	b[1] = (a[1] << 1) | (a[0] >> 63)
	b[2] = (a[2] << 1) | (a[1] >> 63)
	b[3] = (a[3] << 1) | (a[2] >> 63)
	return
}

func SwapBits(octet byte) byte {
	return ((octet&0x0F)<<4 | (octet&0xF0)>>4)
}
_________________________________
func SwapBits(octet byte) byte {
	var div, mod byte

	div = octet / 16
	mod = octet % 16
	return mod*16 + div
}
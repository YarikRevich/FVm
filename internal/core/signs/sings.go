package signs

/*
#include <stdbool.h>
#include <stdint.h>
uint16_t lshift(uint16_t x, uint16_t y){
	return x >> y;
}
*/
import (
	"C"
)

func LShift(x uint16, y uint16)int{
	return int(uint16(C.lshift(C.ushort(x), C.ushort(y-1))))
};

func GetExtendedSign(x uint16, length int)uint16{
	if ((LShift(x, uint16(length-1)) & length) != 0){
		x |= (0xFFFF << length)
	}
	return x
}

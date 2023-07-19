// +build darwin

#include "stddef.h"
#include "stdbool.h"

typedef unsigned long uint_t;
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned long long uint64_t;

struct Size {
	uint_t Width;
	uint_t Height;
	uint_t Depth;
};

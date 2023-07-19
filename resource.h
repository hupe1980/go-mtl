// +build darwin

#include "mtl.h"

struct TextureDescriptor {
	uint16_t PixelFormat;
	uint_t   Width;
	uint_t   Height;
	uint8_t  StorageMode;
};

void * Device_NewBufferWithLength(void * device, size_t length, uint16_t options);
void * Device_NewBufferWithBytes(void * device, const void * bytes, size_t length, uint16_t options);

void * Buffer_Contents(void * buffer);

void * Device_NewTextureWithDescriptor(void * device, struct TextureDescriptor descriptor);

void Texture_ReplaceRegion(void * texture, struct Region region, uint_t level, void * pixelBytes, size_t bytesPerRow);
void Texture_GetBytes(void * texture, void * pixelBytes, size_t bytesPerRow, struct Region region, uint_t level);
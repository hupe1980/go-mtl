// +build darwin

#import <Metal/Metal.h>
#include "resource.h"

void * Device_NewBufferWithLength(void * device, size_t length, uint16_t options) {
	return [(id<MTLDevice>)device newBufferWithLength:(NSUInteger)length
	                                         options:(MTLResourceOptions)options];
}

void * Device_NewBufferWithBytes(void * device, const void * bytes, size_t length, uint16_t options) {
	return [(id<MTLDevice>)device newBufferWithBytes:(const void *)bytes
	                                          length:(NSUInteger)length
	                                         options:(MTLResourceOptions)options];
}

void * Buffer_Contents(void * buffer) {
    return ((id<MTLBuffer>)buffer).contents;
}

void * Device_NewTextureWithDescriptor(void * device, struct TextureDescriptor descriptor) {
	MTLTextureDescriptor * textureDescriptor = [[MTLTextureDescriptor alloc] init];
	textureDescriptor.pixelFormat = descriptor.PixelFormat;
	textureDescriptor.width = descriptor.Width;
	textureDescriptor.height = descriptor.Height;
	textureDescriptor.storageMode = descriptor.StorageMode;
    
	return [(id<MTLDevice>)device newTextureWithDescriptor:textureDescriptor];
}

void Texture_ReplaceRegion(void * texture, struct Region region, uint_t level, void * pixelBytes, size_t bytesPerRow) {
	[(id<MTLTexture>)texture replaceRegion:(MTLRegion){{region.Origin.X, region.Origin.Y, region.Origin.Z}, {region.Size.Width, region.Size.Height, region.Size.Depth}}
	                           mipmapLevel:(NSUInteger)level
	                             withBytes:(void *)pixelBytes
	                           bytesPerRow:(NSUInteger)bytesPerRow];
}

void Texture_GetBytes(void * texture, void * pixelBytes, size_t bytesPerRow, struct Region region, uint_t level) {
	[(id<MTLTexture>)texture getBytes:(void *)pixelBytes
	                      bytesPerRow:(NSUInteger)bytesPerRow
	                       fromRegion:(MTLRegion){{region.Origin.X, region.Origin.Y, region.Origin.Z}, {region.Size.Width, region.Size.Height, region.Size.Depth}}
	                      mipmapLevel:(NSUInteger)level];
}
// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// PNG Reader and Writer (CGo glue)

#ifndef pngglue_h
#define pngglue_h

#include <png.h>
#include <setjmp.h>
#include <stdlib.h>

// Go-side functions
void pngErrorCallback(png_struct *png, png_const_charp msg);
void pngWarningCallback(png_struct *png, png_const_charp msg);
void *pngMallocCallback(png_struct *png, size_t size);
void pngFreeCallback(png_struct *png, void *p);
int  pngReadCallback(png_struct *png, png_bytep data, size_t size);
int  pngWriteCallback(png_struct *png, png_bytep data, size_t size);

// do_pngErrorCallback wraps pngErrorCallback.
// The wrapper is required, because we cannot call png_longjmp from Go.
static inline void
do_pngErrorCallback(png_struct *png, const char *message) {
    pngErrorCallback(png, message);
    png_longjmp(png, 1);
}

// do_pngReadCallback wraps pngReadCallback.
// It calls png_error() in a case of an error, as we can't do it from Go.
static inline void
do_pngReadCallback(png_struct *png, png_bytep data, size_t size) {
    if (!pngReadCallback(png, data, size)) {
        png_error(png, "");
    }
}

// do_pngWriteCallback wraps pngWriteCallback.
// It calls png_error() in a case of an error, as we can't do it from Go.
static inline void
do_pngWriteCallback(png_struct *png, png_bytep data, size_t size) {
    if (!pngWriteCallback(png, data, size)) {
        png_error(png, "");
    }
}

// do_png_create_read_struct wraps png_create_read_struct_2.
// This is the convenience wrapper.
static inline png_struct*
do_png_create_read_struct(void *p) {
    png_struct *png;

    png = png_create_read_struct_2(PNG_LIBPNG_VER_STRING,
        p, do_pngErrorCallback, pngWarningCallback,
        p, pngMallocCallback, pngFreeCallback);

    png_set_read_fn(png, p, do_pngReadCallback);

    return png;
}

// do_png_create_write_struct wraps png_create_write_struct_2.
// This is the convenience wrapper.
static inline png_struct*
do_png_create_write_struct(void *p) {
    png_struct *png;

    png = png_create_write_struct_2(PNG_LIBPNG_VER_STRING,
        p, do_pngErrorCallback, pngWarningCallback,
        p, pngMallocCallback, pngFreeCallback);

    png_set_write_fn(png, p, do_pngWriteCallback, NULL);

    return png;
}

// do_png_read_info wraps png_read_info.
// The wrapper is required to catch setjmp return as
// we can't do it from Go
static inline void
do_png_read_info(png_struct *png, png_info *info_ptr) {
    if (setjmp(png_jmpbuf(png))) {
        return;
    }

    png_read_info(png, info_ptr);
}

// do_png_write_info wraps png_write_info.
// The wrapper is required to catch setjmp return as
// we can't do it from Go
static inline void
do_png_write_info(png_struct *png, png_info *info_ptr) {
    if (setjmp(png_jmpbuf(png))) {
        return;
    }

    png_write_info(png, info_ptr);
}

// do_png_get_IHDR wraps png_get_IHDR.
// The wrapper is required to catch setjmp return as
// we can't do it from Go
static inline png_uint_32
do_png_get_IHDR(png_struct *png, png_info *info_ptr,
                png_uint_32 *width, png_uint_32 *height, int *bit_depth,
                int *color_type, int *interlace_type, int *compression_type,
                int *filter_type) {

    if (setjmp(png_jmpbuf(png))) {
        return 0;
    }

    return png_get_IHDR(png, info_ptr, width, height, bit_depth,
                 color_type, interlace_type, compression_type,
                 filter_type);
}

// do_png_set_IHDR wraps png_set_IHDR.
// The wrapper is required to catch setjmp return as
// we can't do it from Go
static inline void
do_png_set_IHDR(png_struct *png, png_info *info_ptr,
                png_uint_32 width, png_uint_32 height, int bit_depth,
                int color_type, int interlace_type, int compression_type,
                int filter_type) {

    if (setjmp(png_jmpbuf(png))) {
        return;
    }

    png_set_IHDR(png, info_ptr, width, height, bit_depth,
                 color_type, interlace_type, compression_type,
                 filter_type);
}

// do_png_read_row wraps png_read_row.
// The wrapper is required to catch setjmp return as we can't do it from Go
static inline void
do_png_read_row(png_struct *png, void *row, png_bytep display_row) {
    if (setjmp(png_jmpbuf(png))) {
        return;
    }

    png_read_row(png, row, display_row);
}

// do_png_write_row wraps png_write_row.
// The wrapper is required to catch setjmp return as we can't do it from Go
static inline void
do_png_write_row(png_struct *png, void *row) {
    if (setjmp(png_jmpbuf(png))) {
        return;
    }

    png_write_row(png, row);
}

// do_png_write_end wraps png_write_end.
// The wrapper is required to catch setjmp return as
// we can't do it from Go
static inline void
do_png_write_end(png_struct *png, png_info *info_ptr) {
    if (setjmp(png_jmpbuf(png))) {
        return;
    }

    png_write_end(png, info_ptr);
}

#endif

// vim:ts=8:sw=4:et

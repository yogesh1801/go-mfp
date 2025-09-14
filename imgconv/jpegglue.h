// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// JPEG Reader and Writer (CGo glue)

#ifndef jpegglue_h
#define jpegglue_h

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

#include <jpeglib.h>

// The following function are provided by Go side

void jpegInitSourceCallback (j_decompress_ptr);
boolean jpegFillInputBufferCallback (j_decompress_ptr);
void jpegSkipInputDataCallback (j_decompress_ptr, long);
void jpegTermSourceCallback (j_decompress_ptr);

void jpegInitDestination (j_compress_ptr cinfo);
boolean jpegEmptyOutputBuffer (j_compress_ptr cinfo);
void jpegTermDestination (j_compress_ptr cinfo);

void jpegErrorCallback (j_common_ptr);
void jpegOutputMessageCallback (j_common_ptr);

// do_jpeg_format_message wraps common->format_message,
// so Go cal call it.
static inline void
do_jpeg_format_message (j_common_ptr common, char *buf) {
    common->err->format_message(common, buf);
}

// do_jpeg_init_decompress initializes JPEG decompressor
static inline void
do_jpeg_init_decompress (struct jpeg_decompress_struct *jpeg,
                         struct jpeg_error_mgr *errmgr,
                         struct jpeg_source_mgr *srcmgr,
                         uintptr_t handle) {

    jpeg->client_data = (void*) handle;
    jpeg->err = jpeg_std_error(errmgr);
    errmgr->error_exit = jpegErrorCallback;
    errmgr->output_message = jpegOutputMessageCallback;

    jpeg_create_decompress(jpeg);

    jpeg->src = srcmgr;
    jpeg->src->init_source = jpegInitSourceCallback;
    jpeg->src->fill_input_buffer = jpegFillInputBufferCallback;
    jpeg->src->skip_input_data = jpegSkipInputDataCallback;
    jpeg->src->term_source = jpegTermSourceCallback;
    jpeg->src->resync_to_restart = jpeg_resync_to_restart;
}

// do_jpeg_init_compress initializes JPEG compressor
static inline void
do_jpeg_init_compress (struct jpeg_compress_struct *jpeg,
                       struct jpeg_error_mgr *errmgr,
                       struct jpeg_destination_mgr *dstmgr,
                       uintptr_t handle) {

    jpeg->client_data = (void*) handle;
    jpeg->err = jpeg_std_error(errmgr);
    errmgr->error_exit = jpegErrorCallback;
    errmgr->output_message = jpegOutputMessageCallback;

    jpeg_create_compress(jpeg);

    jpeg->dest = dstmgr;
    jpeg->dest->init_destination = jpegInitDestination;
    jpeg->dest->empty_output_buffer = jpegEmptyOutputBuffer;
    jpeg->dest->term_destination = jpegTermDestination;
}

// do_jpeg_read_scanlines calls jpeg_read_scanlines for a
// single scanline using the provided output buffer.
static inline JDIMENSION
do_jpeg_read_scanline (j_decompress_ptr jpeg, void *buf) {
    JSAMPROW           lines[1] = {buf};
    return jpeg_read_scanlines(jpeg, lines, 1);
}

// do_jpeg_write_scanline calls jpeg_write_scanlines for a
// single scanline using the provided input buffer.
static inline JDIMENSION
do_jpeg_write_scanline (j_compress_ptr jpeg, void *buf) {
    JSAMPROW           lines[1] = {buf};
    return jpeg_write_scanlines(jpeg, lines, 1);
}

#endif

// vim:ts=8:sw=4:et

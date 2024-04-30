package core

import (
	"fmt"
	"image"
	"unsafe"
)

// #cgo pkg-config: libavcodec libavutil libswscale libavformat
//// #cgo CFLAGS: -I /home/timothy/work/arm/ffmpeg_arm/include
//// #cgo LDFLAGS: -L /home/timothy/work/arm/ffmpeg_arm/lib -l:libavcodec.a -l:libavutil.a -l:libswscale.a -l:libavformat.a
// #include <libavcodec/avcodec.h>
// #include <libavutil/imgutils.h>
// #include <libswscale/swscale.h>
// #include <libavformat/avformat.h>
import "C"

// Do not put black line between commens above and import C

func frameData(frame *C.AVFrame) **C.uint8_t {
	return (**C.uint8_t)(unsafe.Pointer(&frame.data[0]))
}

func frameLineSize(frame *C.AVFrame) *C.int {
	return (*C.int)(unsafe.Pointer(&frame.linesize[0]))
}

type SnapshotDecoder struct {
	codecCtx    *C.AVCodecContext
	srcFrame    *C.AVFrame
	swsCtx      *C.struct_SwsContext
	dstFrame    *C.AVFrame
	dstFramePtr []uint8
}

func NewSnapshotDecoder(codecName string) (*SnapshotDecoder, error) {
	var codecID uint32

	switch codecName {
	case "MPEG4":
		codecID = C.AV_CODEC_ID_MPEG4 // TODO check payload type to see what MPEG4 is
	case "H264":
		codecID = C.AV_CODEC_ID_H264
	case "H265":
		codecID = C.AV_CODEC_ID_HEVC
	}

	C.av_log_set_level(C.AV_LOG_QUIET)

	codec := C.avcodec_find_decoder(codecID)
	if codec == nil {
		return nil, fmt.Errorf("avcodec_find_decoder failed")
	}

	codecCtx := C.avcodec_alloc_context3(codec)
	if codecCtx == nil {
		return nil, fmt.Errorf("avcodec_alloc_context3 failed")
	}

	res := C.avcodec_open2(codecCtx, codec, nil)
	if res < 0 {
		C.avcodec_close(codecCtx)
		return nil, fmt.Errorf("avcodec_open failed")
	}

	srcFrame := C.av_frame_alloc()
	if srcFrame == nil {
		C.avcodec_close(codecCtx)
		return nil, fmt.Errorf("av_frame_alloc() failed")
	}

	return &SnapshotDecoder{
		codecCtx: codecCtx,
		srcFrame: srcFrame,
	}, nil
}

func (vd *SnapshotDecoder) close() {
	if vd.dstFrame != nil {
		C.av_frame_free(&vd.dstFrame)
	}

	if vd.swsCtx != nil {
		C.sws_freeContext(vd.swsCtx)
	}

	C.av_frame_free(&vd.srcFrame)
	C.avcodec_close(vd.codecCtx)

}

func (vd *SnapshotDecoder) decodeSnapshotFrame(data []byte, outputWidth, outputHeight int) (image.Image, error) {
	data = append([]uint8{0x00, 0x00, 0x00, 0x01}, []uint8(data)...)

	var avPacket C.AVPacket
	avPacket.data = (*C.uint8_t)(C.CBytes(data))
	defer C.free(unsafe.Pointer(avPacket.data))

	avPacket.size = C.int(len(data))
	res := C.avcodec_send_packet(vd.codecCtx, &avPacket)
	if res < 0 {
		return nil, nil
	}
	res = C.avcodec_receive_frame(vd.codecCtx, vd.srcFrame)
	if res < 0 {
		return nil, nil
	}

	/* if vd.srcFrame.key_frame != 1 {
		return nil, nil
	} */

	if vd.dstFrame == nil || vd.dstFrame.width != vd.srcFrame.width || vd.dstFrame.height != vd.srcFrame.height {
		if vd.dstFrame != nil {
			C.av_frame_free(&vd.dstFrame)
		}

		if vd.swsCtx != nil {
			C.sws_freeContext(vd.swsCtx)
		}

		vd.dstFrame = C.av_frame_alloc()
		vd.dstFrame.format = C.AV_PIX_FMT_RGBA //|| C.AV_PIX_FMT_YUVJ420P

		if outputWidth > 0 {
			vd.dstFrame.width = C.int(outputWidth)
			vd.dstFrame.height = C.int(outputHeight)
		} else {
			vd.dstFrame.width = vd.srcFrame.width
			vd.dstFrame.height = vd.srcFrame.height
		}

		vd.dstFrame.color_range = C.AVCOL_RANGE_JPEG

		res = C.av_frame_get_buffer(vd.dstFrame, 1)
		if res < 0 {
			return nil, fmt.Errorf("av_frame_get_buffer() error")
		}

		vd.swsCtx = C.sws_getContext(vd.srcFrame.width, vd.srcFrame.height, (int32)(vd.srcFrame.format), vd.dstFrame.width, vd.dstFrame.height, (int32)(vd.dstFrame.format), C.SWS_BILINEAR, nil, nil, nil)
		if vd.swsCtx == nil {
			return nil, fmt.Errorf("sws_getContext() error")
		}

		dstFrameSize := C.av_image_get_buffer_size((int32)(vd.dstFrame.format), vd.dstFrame.width, vd.dstFrame.height, 1)
		vd.dstFramePtr = (*[1 << 30]uint8)(unsafe.Pointer(vd.dstFrame.data[0]))[:dstFrameSize:dstFrameSize]
	}

	res = C.sws_scale(vd.swsCtx, frameData(vd.srcFrame), frameLineSize(vd.srcFrame), 0, vd.srcFrame.height, frameData(vd.dstFrame), frameLineSize(vd.dstFrame))
	if res < 0 {
		return nil, fmt.Errorf("sws_scale() error")
	}

	return &image.RGBA{
		Pix:    vd.dstFramePtr,
		Stride: 4 * (int)(vd.dstFrame.width),
		Rect: image.Rectangle{
			Max: image.Point{(int)(vd.dstFrame.width), (int)(vd.dstFrame.height)},
		},
	}, nil
}

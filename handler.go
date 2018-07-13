package agqr

import (
	rtmp "github.com/nbuzard/gortmp"
	"github.com/zhangpeihao/goflv"
	"log"
)

var outBoundStream chan rtmp.OutboundStream

type AGQROutboundConnHandler struct {
	file *flv.File
}

func (handler *AGQROutboundConnHandler) OnStatus(conn rtmp.OutboundConn) {
	status, err := conn.Status()
	log.Printf("Stream OnStatus: %d, err: %v\n", status, err)
}

func (handler *AGQROutboundConnHandler) OnClosed(conn rtmp.Conn) {
	log.Println("Stream Closed")
}

func (handler *AGQROutboundConnHandler) OnReceived(conn rtmp.Conn, message *rtmp.Message) {
	switch message.Type {
	case rtmp.VIDEO_TYPE:
		handler.file.WriteVideoTag(message.Buf.Bytes(), message.AbsoluteTimestamp)
	case rtmp.AUDIO_TYPE:
		handler.file.WriteAudioTag(message.Buf.Bytes(), message.AbsoluteTimestamp)
	}
}

func (handler *AGQROutboundConnHandler) OnReceivedRtmpCommand(conn rtmp.Conn, command *rtmp.Command) {
	log.Printf("Stream OnReceived RTMP Command: %+v\n", command)
}

func (handler *AGQROutboundConnHandler) OnStreamCreated(conn rtmp.OutboundConn, stream rtmp.OutboundStream) {
	log.Printf("Stream Created: %d\n", stream.ID())
	outBoundStream <- stream
}

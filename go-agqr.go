package agqr

import (
	rtmp "github.com/nbuzard/gortmp"
	"github.com/zhangpeihao/goflv"
	"time"
)

type AGQR struct {
	conn    rtmp.OutboundConn
	handler *AGQROutboundConnHandler
	path    string
}

func New(path string) (*AGQR, error) {
	file, err := flv.CreateFile(path)
	if err != nil {
		return nil, err
	}
	outBoundStream = make(chan rtmp.OutboundStream)
	handler := &AGQROutboundConnHandler{
		file: file,
	}
	conn, err := rtmp.Dial("rtmp://fms-base2.mitene.ad.jp/agqr/", handler, 100)
	if err != nil {
		return nil, err
	}
	return &AGQR{conn: conn, handler: handler, path: path}, nil
}

func (s *AGQR) Start(d time.Duration, callback func(path string) error) error {
	err := s.conn.Connect()
	if err != nil {
		return err
	}

	for {
		select {
		case stream := <-outBoundStream:
			// Play
			err = stream.Play("aandg22", nil, nil, nil)
			if err != nil {
				return err
			}

		case <-time.After(d):
			return callback(s.path)
		}
	}
}

func (s *AGQR) Close() {
	s.conn.Close()
	s.handler.file.Close()
}

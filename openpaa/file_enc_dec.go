package openpaa

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"io"
)

// EncodeHeader 编码文件头部
// | head_len| head |
func EncodeHeader(w io.Writer, data any) error {
	headBody, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, uint32(len(headBody)))
	if err != nil {
		return err
	}
	_, err = w.Write(headBody)
	return err
}

// DecodeHeader 解码文件头部
// | head_len| head |
func DecodeHeader(r io.Reader, data any) error {
	var headLen uint32

	err := binary.Read(r, binary.BigEndian, &headLen)
	if err != nil {
		return err
	}
	headBody := make([]byte, headLen)
	_, err = io.ReadFull(r, headBody)
	if err != nil {
		return err
	}
	return xml.Unmarshal(headBody, data)
}

// EncodeHeaderData 编码文件头部
func EncodeHeaderData(data any) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := EncodeHeader(buffer, data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// EncodeContent 编码文件内容
// | content_len| content |
func EncodeContent(w io.Writer, b []byte) error {
	err := binary.Write(w, binary.BigEndian, uint32(len(b)))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// DecodeContent 解码文件内容
// | content_len| content |
func DecodeContent(r io.Reader) ([]byte, error) {
	var contentLen uint32

	err := binary.Read(r, binary.BigEndian, &contentLen)
	if err != nil {
		return nil, err
	}
	contentBody := make([]byte, int(contentLen))
	_, err = io.ReadFull(r, contentBody)
	if err != nil {
		return nil, err
	}
	return contentBody, nil
}

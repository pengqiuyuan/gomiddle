package codec

import (
	"bufio"
	"bytes"
    "encoding/binary"
)

const TcpHeaderSize = 23

type TcpMessageHeader struct {
	Flag     uint16
	Session  uint32
	Guid     uint64
	Proto    uint16
	Group    uint16
	Compress uint8
	Checksum uint16
	Size     uint16
}

type TcpMessage struct {
	Header  TcpMessageHeader
	Payload []byte
}

func Encode(m *TcpMessage) ([]byte, error) {
    // 读取消息的长度
    var pkg *bytes.Buffer = new(bytes.Buffer)
    // 写入消息头
    err := binary.Write(pkg, binary.LittleEndian, &(m.Header))
    if err != nil {
       return nil, err
    }
    // 写入消息实体
    err = binary.Write(pkg, binary.LittleEndian, m.Payload)
    if err != nil {
       return nil, err
    }
    return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	//fmt.Println(reader)
	var m TcpMessage
	var buf []byte
	lengthBuff := bytes.NewBuffer(buf[:TcpHeaderSize])
	err := binary.Read(lengthBuff, binary.LittleEndian, &(m.Header))
	if err != nil {
		return "", err
	}
	m.Payload = buf[TcpHeaderSize:]
	return string(m.Payload), nil
}


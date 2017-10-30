package proto

import (
	"encoding/binary"
)

const (
	CmdGet  = 0x01
	CmdSet  = 0x02
	CmdScan = 0x03
	CmdDel  = 0x04
)
const (
	MaxTableId   = 9 // TableId [0 ~ 9]
	HeadSize     = 17
	MaxUint8     = 255
	MaxUint16    = 65535
	MaxScanNum   = 10000
	MaxColKeyLen = 1024 * 8        // 8KB
	MaxValueLen  = 1024 * 256      // 256KB
	MaxPkgLen    = 1024 * 1024 * 2 // 2MB
)

type PkgHead struct {
	Crc    uint8
	Cmd    uint8
	Cid    uint16
	DbId   uint8
	Seq    uint64
	PkgLen uint32
}

type KeyValue struct {
	ErrCode int8
	DbId    uint8
	TableId uint8
	RowKey  []byte
	colKey  []byte
	value   []byte
}

func (head *PkgHead) Decode(pkg []byte) (int, error) {
	if len(pkg) < HeadSize {
		return 0, ErrInvPkgHeadLen
	}

	head.Crc = CalHeadCrc(pkg)
	if head.Crc != pkg[0] {
		return 0, ErrInvHeadCrc
	}

	head.Cmd = pkg[1]
	head.DbId = pkg[2]
	head.Seq = binary.BigEndian.Uint64(pkg[3:])
	head.Cid = binary.BigEndian.Uint16(pkg[11:])
	head.PkgLen = binary.BigEndian.Uint32(pkg[13:])

	return HeadSize, nil
}

func (head *PkgHead) Encode(pkg []byte) (int, error) {
	if len(pkg) < HeadSize {
		return 0, ErrInvPkgHeadLen
	}

	pkg[1] = head.Cmd
	pkg[2] = head.DbId
	binary.BigEndian.PutUint64(pkg[3:], head.Seq)
	binary.BigEndian.PutUint16(pkg[11:], head.Cid)
	binary.BigEndian.PutUint32(pkg[13:], head.PkgLen)
	pkg[0] = CalHeadCrc(pkg)

	return HeadSize, nil
}

func CalHeadCrc(pkg []byte) uint8 {
	var crc int8 = 10
	for i := 1; i < HeadSize; i++ {
		crc += int8(pkg[i])
	}
	return uint8(crc)
}

// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"fmt"

	"github.com/fagongzi/distributed-limiter/pkg/pb"
	"github.com/fagongzi/distributed-limiter/pkg/util"
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/log"
)

var (
	// Decoder decoder
	Decoder = goetty.NewIntLengthFieldBasedDecoder(&decoder{})
	// Encoder encoder
	Encoder = &encoder{}
)

var (
	msgTypePut       = byte(pb.MsgPut)
	msgTypePutRsp    = byte(pb.MsgPutRsp)
	msgTypeAccess    = byte(pb.MsgAccess)
	msgTypeAccessRsp = byte(pb.MsgAccessRsp)
)

type decoder struct {
}

type encoder struct {
}

// Marashal marashal interface
type Marashal interface {
	Size() int
	Marshal() ([]byte, error)
	MarshalTo(data []byte) (int, error)
}

// Decode returns a put or access msg
func (decoder *decoder) Decode(in *goetty.ByteBuf) (bool, interface{}, error) {
	data := in.GetMarkedRemindData()

	if data[0] == msgTypeAccess {
		msg := util.AcquireAccess()
		return true, msg, msg.Unmarshal(data[1:])
	} else if data[0] == msgTypeAccessRsp {
		msg := util.AcquireAccessRsp()
		return true, msg, msg.Unmarshal(data[1:])
	} else if data[0] == msgTypePut {
		msg := util.AcquirePut()
		return true, msg, msg.Unmarshal(data[1:])
	} else if data[0] == msgTypePutRsp {
		msg := util.AcquirePutRsp()
		return true, msg, msg.Unmarshal(data[1:])
	}

	return true, nil, fmt.Errorf("%d not support msg", data[0])
}

// Encode encode proxy message
func (encoder *encoder) Encode(data interface{}, out *goetty.ByteBuf) error {
	var t byte
	var m Marashal

	if msg, ok := data.(*pb.AccessRsp); ok {
		t = msgTypeAccessRsp
		m = msg
	} else if msg, ok := data.(*pb.Access); ok {
		t = msgTypeAccess
		m = msg
	} else if msg, ok := data.(*pb.PutRsp); ok {
		t = msgTypePutRsp
		m = msg
	} else if msg, ok := data.(*pb.Put); ok {
		t = msgTypePut
		m = msg
	} else {
		log.Fatalf("bug: unsupport msg: %+v", msg)
	}

	size := m.Size()
	out.WriteInt(size + 1)
	out.WriteByte(byte(t))

	if size > 0 {
		index := out.GetWriteIndex()
		out.Expansion(size)
		m.MarshalTo(out.RawBuf()[index : index+size])
		out.SetWriterIndex(index + size)
	}

	return nil
}

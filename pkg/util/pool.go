package util

import (
	"sync"

	"github.com/fagongzi/distributed-limiter/pkg/pb"
)

var (
	putsPool = &sync.Pool{
		New: func() interface{} {
			return &pb.Put{}
		},
	}

	putRspsPool = &sync.Pool{
		New: func() interface{} {
			return &pb.PutRsp{}
		},
	}

	accessesPool = &sync.Pool{
		New: func() interface{} {
			return &pb.Access{}
		},
	}

	accessRspsPool = &sync.Pool{
		New: func() interface{} {
			return &pb.AccessRsp{}
		},
	}
)

// AcquirePut returns an empty Put object from the pool
func AcquirePut() *pb.Put {
	return putsPool.Get().(*pb.Put)
}

// ReleasePut returns the object acquired via AcquirePut to the pool.
func ReleasePut(value *pb.Put) {
	value.Reset()
	putsPool.Put(value)
}

// AcquirePutRsp returns an empty PutRsp object from the pool
func AcquirePutRsp() *pb.PutRsp {
	return putRspsPool.Get().(*pb.PutRsp)
}

// ReleasePutRsp returns the object acquired via AcquirePutRsp to the pool.
func ReleasePutRsp(value *pb.PutRsp) {
	value.Reset()
	putRspsPool.Put(value)
}

// AcquireAccess returns an empty Access object from the pool
func AcquireAccess() *pb.Access {
	return accessesPool.Get().(*pb.Access)
}

// ReleaseAccess returns the object acquired via AcquireAccess to the pool.
func ReleaseAccess(value *pb.Access) {
	value.Reset()
	accessesPool.Put(value)
}

// AcquireAccessRsp returns an empty AccessRsp object from the pool
func AcquireAccessRsp() *pb.AccessRsp {
	return accessRspsPool.Get().(*pb.AccessRsp)
}

// ReleaseAccessRsp returns the object acquired via AcquireAccessRsp to the pool.
func ReleaseAccessRsp(value *pb.AccessRsp) {
	value.Reset()
	accessRspsPool.Put(value)
}

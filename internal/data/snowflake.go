package data

import (
	"errors"
	"sync"
	"time"
)

// make these constants be read from env...

const (
	// 9/21/2007 12:00:00AM GMT -- golang's inception
	epoch           int64 = 1190332800000
	workerIdBits    uint8 = 10
	maxWorkerId     int64 = -1 ^ (-1 << workerIdBits)
	sequenceBits    uint8 = 12
	sequenceBitMask int64 = -1 ^ (-1 << sequenceBits)

	// spacers: left = timestamp: middle = workerId: right = sequence
	timestampLeftShift uint8 = workerIdBits + sequenceBits
	workerIdLeftShift  uint8 = sequenceBits
)

// timestamp, worker number and sequence number. from https://blog.twitter.com/engineering/en_us/a/2010/announcing-snowflake
// strategy for worker number: assigned manually by me for now but probably in reality this would be set by a zookeeper...
// other options include: uuid, macaddress, ip address but these aren't great ideas tbh

type Snowflake struct {
	lastTimestamp int64
	workerId      int64
	sequence      int64
	mutex         sync.Mutex
}

func NewSnowFlake(workerId int64) (*Snowflake, error) {
	if workerId > maxWorkerId || workerId < 0 {
		return nil, errors.New("worker ID out of range")
	}
	return &Snowflake{
		// static so set it
		workerId: workerId,
		// only incremented with collision
		sequence: 0,
		// setting timestamp to 1s
		lastTimestamp: -1,
	}, nil
}

func (snowflake *Snowflake) GenerateId() (int64, error) {
	snowflake.mutex.Lock()
	defer snowflake.mutex.Unlock()

	timestamp := time.Now().UnixMilli()
	// to account for NTP issues
	if timestamp < snowflake.lastTimestamp {
		return 0, errors.New("clock went backwards. No new Id for you")
	}

	// collision
	if timestamp == snowflake.lastTimestamp {
		snowflake.sequence = (snowflake.sequence + 1) & sequenceBitMask
		// if sequence goes out of range i.e., > max snowflake sequence number allowed. Need to wait for next millisecond
		if snowflake.sequence == 0 {
			// need a method to wait for next millisecond
			snowflake.WaitNextMilli()
		} else {
			// no collision so reset sequence to 0
			snowflake.sequence = 0
		}
	}

	snowflake.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampLeftShift) | (snowflake.workerId << workerIdLeftShift) | snowflake.sequence
	return id, nil
}

func (snowflake *Snowflake) WaitNextMilli() {
	t := time.Now().UnixMilli()
	for t <= snowflake.lastTimestamp {
		t = time.Now().UnixMilli()
	}
	snowflake.lastTimestamp = t
}

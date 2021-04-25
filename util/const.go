package util

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type NumberInt int

func (n *NumberInt) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(*n))), nil
}

func (n *NumberInt) UnmarshalJSON(data []byte) (err error) {
	if v, e := strconv.ParseInt(string(data), 10, 64); e == nil {
		*n = NumberInt(v)
		return e
	}
	wrapper := struct {
		NumberInt string `json:"$numberInt"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseInt(wrapper.NumberInt, 10, 64)
	if err != nil {
		return
	}
	*n = NumberInt(v)
	return
}

type NumberDouble float64

func (n *NumberDouble) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (n *NumberDouble) UnmarshalJSON(data []byte) (err error) {
	if v, e := strconv.ParseFloat(string(data), 64); e == nil {
		*n = NumberDouble(v)
		return e
	}
	wrapper := struct {
		NumberDouble string `json:"$numberDouble"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseFloat(wrapper.NumberDouble, 64)
	if err != nil {
		return
	}
	*n = NumberDouble(v)
	return
}

type NumberLong int64

func (n *NumberLong) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (n *NumberLong) UnmarshalJSON(data []byte) (err error) {
	if v, e := strconv.ParseInt(string(data), 10, 64); e == nil {
		*n = NumberLong(v)
		return e
	}
	wrapper := struct {
		NumberLong string `json:"$numberLong"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseInt(wrapper.NumberLong, 10, 64)
	if err != nil {
		return
	}
	*n = NumberLong(v)
	return
}

type ISODate time.Time

func (t ISODate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`{"$date":{"$numberLong":"%d"}}`, time.Time(t).UnixNano()/1e6)
	return []byte(stamp), nil
}

func (t *ISODate) UnmarshalJSON(data []byte) (err error) {
	ti := time.Time(*t)
	if err = ti.UnmarshalJSON(data); err == nil {
		*t = ISODate(ti)
		return
	}
	wrapper := struct {
		NumberLong string `json:"$numberLong"`
	}{}
	if err = json.Unmarshal(data, &wrapper); err != nil {
		return
	}
	v, err := strconv.ParseInt(wrapper.NumberLong, 10, 64)
	if err != nil {
		return
	}
	*t = ISODate(time.Unix(0, v*1e6))
	return
}

func (t ISODate) Time() time.Time {
	return time.Time(t)
}

type ServerDate struct {
	Offset int64 `json:"offset"`
}

func (t *ServerDate) MarshalJSON() ([]byte, error) {
	wrapper := struct {
		TcbServerDate ServerDate `json:"$tcb_server_date"`
	}{*t}
	return json.Marshal(wrapper)
}

func (t ServerDate) Time() time.Time {
	return time.Now().Add(time.Duration(t.Offset * 1e6))
}

type Timestamp ISODate

func (t Timestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).UnixNano()/1e6)
	return []byte(stamp), nil
}

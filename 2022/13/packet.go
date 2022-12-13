package _13

import (
	"encoding/json"
	"math"
)

var (
	PacketDistressStart = Packet{{Single: 2}}
	PacketDistressEnd   = Packet{{Single: 6}}
)

type Order int64

const (
	Ordered Order = iota
	Unordered
	Unknown
)

func IsOrdered(left, right Packet) bool {
	for i := 0; i < int(math.Min(float64(len(left)), float64(len(right)))); i++ {
		if s := isOrdered(left[i], right[i]); s != Unknown {
			return s == Ordered
		}
	}
	return len(left) <= len(right)
}

func isOrdered(d1, d2 Data) Order {
	if d1.List == nil && d2.List == nil {
		if d1.Single == d2.Single {
			return Unknown
		}
		if d1.Single < d2.Single {
			return Ordered
		}
		return Unordered
	}
	if d1.List != nil && d2.List != nil {
		for i := 0; i < int(math.Max(float64(len(d1.List)), float64(len(d2.List)))); i++ {
			if i >= len(d2.List) {
				return Unordered
			}
			if i >= len(d1.List) {
				return Ordered
			}
			if s := isOrdered(d1.List[i], d2.List[i]); s != Unknown {
				return s
			}
		}
		return Unknown
	}
	if d1.List == nil {
		return isOrdered(Data{List: []Data{{Single: d1.Single}}}, d2)
	}
	return isOrdered(d1, Data{List: []Data{{Single: d2.Single}}})
}

type Packet []Data

type Data struct {
	Single int
	List   []Data
}

func (d *Data) UnmarshalJSON(data []byte) error {
	_ = json.Unmarshal(data, &d.Single)
	_ = json.Unmarshal(data, &d.List)
	return nil
}

func MustParsePacket(l string) Packet {
	var d Packet
	if err := json.Unmarshal([]byte(l), &d); err != nil {
		panic(err)
	}
	return d
}

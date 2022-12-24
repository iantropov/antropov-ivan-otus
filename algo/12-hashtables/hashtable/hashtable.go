package hashtable

import (
	"hash/fnv"
	"strconv"
)

type Hashtable[K, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
	Remove(key K)
}

func GetHashCode(value interface{}) int {
	h := fnv.New32a()
	stringValue := ""
	switch typedValue := value.(type) {
	case string:
		stringValue = typedValue
	case int:
		stringValue = strconv.Itoa(typedValue)
	default:
		panic("invalid key type")
	}
	h.Write([]byte(stringValue))
	return int(h.Sum32())
}

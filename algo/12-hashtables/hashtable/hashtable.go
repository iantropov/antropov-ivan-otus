package hashtable

type Hashtable[K, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
	Remove(key K)
}

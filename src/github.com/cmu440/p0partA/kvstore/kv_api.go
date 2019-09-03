// API of key value store to be used in server_impl

package kvstore

// KVStore -- Interface for Key/Value stores
type KVStore interface {
	Put(key string, value []byte)
	Get(key string) []([]byte)
	Clear(key string)
}

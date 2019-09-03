// Basic key value access functions to be used in server_impl

package kvstore

type impl struct {
	internal map[string][]([]byte)
}

// CreateWithBackdoor -- Used by tests to create a KVStore implementation while keeping a handle on inner storage
func CreateWithBackdoor() (KVStore, map[string][]([]byte)) {
	internal := make(map[string][]([]byte))
	return impl{internal}, internal
}

// Put inserts a new key value pair or updates the value for a
// given key in the store
func (im impl) Put(key string, value []byte) {
	im.internal[key] = append(im.internal[key], value)
}

// Get fetches the value associated with the key
func (im impl) Get(key string) []([]byte) {
	v, _ := im.internal[key]
	return v
}

// Clear removes all values associated with the key
func (im impl) Clear(key string) {
	delete(im.internal, key)
}

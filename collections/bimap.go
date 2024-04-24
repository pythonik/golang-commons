package collections

type BiMap[K, V comparable] struct {
	keyToValue map[K]V
	valueToKey map[V]K
}

func NewBiMap[K, V comparable]() *BiMap[K, V] {
	return &BiMap[K, V]{
		keyToValue: make(map[K]V),
		valueToKey: make(map[V]K),
	}
}

func (b *BiMap[K, V]) Put(key K, value V) {
	if oldVal, ok := b.keyToValue[key]; ok {
		delete(b.valueToKey, oldVal)
	}
	if oldKey, ok := b.valueToKey[value]; ok {
		delete(b.keyToValue, oldKey)
	}
	b.keyToValue[key] = value
	b.valueToKey[value] = key
}

func (b *BiMap[K, V]) GetByKey(key K) (V, bool) {
	val, ok := b.keyToValue[key]
	return val, ok
}
func (b *BiMap[K, V]) GetByKeyDefault(key K, defaultValue V) V {
	val, ok := b.keyToValue[key]
	if !ok {
		return defaultValue
	}
	return val
}

func (b *BiMap[K, V]) GetByValue(value V) (K, bool) {
	key, ok := b.valueToKey[value]
	return key, ok
}

func (b *BiMap[K, V]) GetByValueDefault(value V, defaultValue K) K {
	key, ok := b.valueToKey[value]
	if !ok {
		return defaultValue
	}
	return key
}

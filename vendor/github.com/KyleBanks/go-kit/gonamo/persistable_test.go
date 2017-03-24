package gonamo

type testPersistable struct {
	hashKey    interface{}
	rangeKey   interface{}
	attributes AttributeMap
}

func (p testPersistable) HashKey() interface{} {
	return p.hashKey
}

func (p testPersistable) RangeKey() interface{} {
	return p.rangeKey
}

func (p testPersistable) Attributes() AttributeMap {
	return p.attributes
}

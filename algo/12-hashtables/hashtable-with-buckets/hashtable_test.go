package hashtableWithBuckets

import (
	"fmt"
	"testing"
)

func TestHashtableEmptiness(t *testing.T) {
	table := NewHashtable()
	assertSize(t, table, 0)
}

func TestHashtablePut(t *testing.T) {
	table := NewHashtable()

	assertAbsence(t, table, "key1")
	table.Put("key1", "value1")
	assertPresence(t, table, "key1", "value1")

	assertSize(t, table, 1)
}

func TestHashtableMultiplePut(t *testing.T) {
	table := NewHashtable()

	table.Put("key1", "value1")
	table.Put("key2", "value2")
	table.Put("key3", "value3")
	table.Put("key4", "value4")

	assertPresence(t, table, "key1", "value1")
	assertPresence(t, table, "key2", "value2")
	assertPresence(t, table, "key3", "value3")
	assertPresence(t, table, "key4", "value4")

	assertSize(t, table, 4)
}

func TestHashtableRePut(t *testing.T) {
	table := NewHashtable()

	table.Put("key1", "value1")
	table.Put("key2", "value2")
	table.Put("key2", "value")

	assertPresence(t, table, "key2", "value")
	assertSize(t, table, 2)
}

func TestHashtableRemove(t *testing.T) {
	table := NewHashtable()

	table.Put("key1", "value1")

	table.Remove("key1")

	assertAbsence(t, table, "key1")
	assertSize(t, table, 0)
}

func TestHashtableAbsentRemove(t *testing.T) {
	table := NewHashtable()

	table.Put("key1", "value1")
	table.Put("key2", "value2")

	table.Remove("key0")
	table.Remove("key3")
	table.Remove("key4")

	assertSize(t, table, 2)
}

func TestHashtableMultipleRemove(t *testing.T) {
	table := NewHashtable()

	table.Put("key1", "value1")
	table.Put("key2", "value2")
	table.Put("key3", "value3")
	table.Put("key4", "value4")

	table.Remove("key2")
	table.Remove("key3")

	assertSize(t, table, 2)

	assertPresence(t, table, "key1", "value1")
	assertPresence(t, table, "key4", "value4")
	assertAbsence(t, table, "key2")
	assertAbsence(t, table, "key3")
}

func TestHashtableRehash(t *testing.T) {
	table := NewHashtable()

	for i := 0; i < 20; i++ {
		table.Put(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
	}

	assertSize(t, table, 20)
	for i := 0; i < 20; i++ {
		assertPresence(t, table, fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
	}
}

func assertPresence(t *testing.T, table *Hashtable, key, value string) {
	valueFromTable, pr := table.Get(key)
	if valueFromTable != value || !pr {
		t.Errorf("Hashtable should have key = %s with value = %s", key, value)
	}
}

func assertAbsence(t *testing.T, table *Hashtable, key string) {
	valueFromTable, pr := table.Get(key)
	if valueFromTable != "" || pr {
		t.Errorf("Hashtable shouldn't have key = %s", key)
	}
}

func assertSize(t *testing.T, table *Hashtable, size int) {
	if table.Size() != size {
		t.Errorf("Hashtable should have size = %d", size)
	}
}

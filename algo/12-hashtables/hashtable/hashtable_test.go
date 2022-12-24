package hashtable_test

import (
	"fmt"
	"hashtables/hashtable"

	hashtableWithBuckets "hashtables/hashtable-with-buckets"
	hashtableWithOpenAddressing "hashtables/hashtable-with-open-addressing"
	"testing"
)

func TestStringHashtables(t *testing.T) {
	tests := []struct {
		name       string
		buildTable func() hashtable.Hashtable[string, string]
	}{
		{
			name: "[string, string] with buckets",
			buildTable: func() hashtable.Hashtable[string, string] {
				return hashtableWithBuckets.NewHashtable[string, string]()
			},
		},
		{
			name: "[string, string] with open addressing",
			buildTable: func() hashtable.Hashtable[string, string] {
				return hashtableWithOpenAddressing.NewHashtable[string, string]()
			},
		},
		{
			name: "[string, string] with open addressing (quadratic probe)",
			buildTable: func() hashtable.Hashtable[string, string] {
				return hashtableWithOpenAddressing.NewHashtableWithQuadraticProbe[string, string]()
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("testHashtableEmptiness", func(t *testing.T) {
				table := tc.buildTable()
				assertSize(t, table, 0)
			})

			t.Run("testHashtablePut", func(t *testing.T) {
				table := tc.buildTable()

				assertAbsence(t, table, "key1", "")

				table.Put("key1", "value1")

				assertPresence(t, table, "key1", "value1")
				assertSize(t, table, 1)
			})

			t.Run("TestHashtableMultiplePut", func(t *testing.T) {
				table := tc.buildTable()

				table.Put("key1", "value1")
				table.Put("key2", "value2")
				table.Put("key3", "value3")
				table.Put("key4", "value4")

				assertPresence(t, table, "key1", "value1")
				assertPresence(t, table, "key2", "value2")
				assertPresence(t, table, "key3", "value3")
				assertPresence(t, table, "key4", "value4")

				assertSize(t, table, 4)
			})

			t.Run("TestHashtableRePut", func(t *testing.T) {
				table := tc.buildTable()

				table.Put("key1", "value1")
				table.Put("key2", "value2")
				table.Put("key2", "value")

				assertPresence(t, table, "key2", "value")
				assertSize(t, table, 2)
			})

			t.Run("TestHashtableRemove", func(t *testing.T) {
				table := tc.buildTable()

				table.Put("key1", "value1")

				table.Remove("key1")

				assertAbsence(t, table, "key1", "")
				assertSize(t, table, 0)
			})

			t.Run("TestHashtableAbsentRemove", func(t *testing.T) {
				table := tc.buildTable()

				table.Put("key1", "value1")
				table.Put("key2", "value2")

				table.Remove("key0")
				table.Remove("key3")
				table.Remove("key4")

				assertSize(t, table, 2)
			})

			t.Run("TestHashtableMultipleRemove", func(t *testing.T) {
				table := tc.buildTable()

				table.Put("key1", "value1")
				table.Put("key2", "value2")
				table.Put("key3", "value3")
				table.Put("key4", "value4")

				table.Remove("key2")
				table.Remove("key3")

				assertSize(t, table, 2)

				assertPresence(t, table, "key1", "value1")
				assertPresence(t, table, "key4", "value4")
				assertAbsence(t, table, "key2", "")
				assertAbsence(t, table, "key3", "")
			})

			t.Run("TestHashtableRehash", func(t *testing.T) {
				table := tc.buildTable()

				for i := 0; i < 20; i++ {
					table.Put(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
				}

				assertSize(t, table, 20)
				for i := 0; i < 20; i++ {
					assertPresence(t, table, fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
				}
			})
		})
	}
}

func TestIntHashtables(t *testing.T) {
	tests := []struct {
		name       string
		buildTable func() hashtable.Hashtable[int, string]
	}{
		{
			name: "[int, string] with buckets",
			buildTable: func() hashtable.Hashtable[int, string] {
				return hashtableWithBuckets.NewHashtable[int, string]()
			},
		},
		{
			name: "[int, string] with open addressing",
			buildTable: func() hashtable.Hashtable[int, string] {
				return hashtableWithOpenAddressing.NewHashtable[int, string]()
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("testHashtableEmptiness", func(t *testing.T) {
				table := tc.buildTable()
				assertSize(t, table, 0)
			})

			t.Run("testHashtablePut", func(t *testing.T) {
				table := tc.buildTable()

				assertAbsence(t, table, 154, "")

				table.Put(154, "value1")

				assertPresence(t, table, 154, "value1")
				assertSize(t, table, 1)
			})
		})
	}
}

func assertPresence[K, V comparable](t *testing.T, table hashtable.Hashtable[K, V], key K, value V) {
	valueFromTable, pr := table.Get(key)
	if valueFromTable != value || !pr {
		t.Errorf("Hashtable should have key = %v with value = %v", key, value)
	}
}

func assertAbsence[K, V comparable](t *testing.T, table hashtable.Hashtable[K, V], key K, emptyValue V) {
	valueFromTable, pr := table.Get(key)
	if valueFromTable != emptyValue || pr {
		t.Errorf("Hashtable shouldn't have key = %v", key)
	}
}

func assertSize[K, V comparable](t *testing.T, table hashtable.Hashtable[K, V], size int) {
	if table.Size() != size {
		t.Errorf("Hashtable should have size = %d", size)
	}
}

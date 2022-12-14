// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Persons", testPersons)
	t.Run("Stamps", testStamps)
}

func TestDelete(t *testing.T) {
	t.Run("Persons", testPersonsDelete)
	t.Run("Stamps", testStampsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Persons", testPersonsQueryDeleteAll)
	t.Run("Stamps", testStampsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Persons", testPersonsSliceDeleteAll)
	t.Run("Stamps", testStampsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Persons", testPersonsExists)
	t.Run("Stamps", testStampsExists)
}

func TestFind(t *testing.T) {
	t.Run("Persons", testPersonsFind)
	t.Run("Stamps", testStampsFind)
}

func TestBind(t *testing.T) {
	t.Run("Persons", testPersonsBind)
	t.Run("Stamps", testStampsBind)
}

func TestOne(t *testing.T) {
	t.Run("Persons", testPersonsOne)
	t.Run("Stamps", testStampsOne)
}

func TestAll(t *testing.T) {
	t.Run("Persons", testPersonsAll)
	t.Run("Stamps", testStampsAll)
}

func TestCount(t *testing.T) {
	t.Run("Persons", testPersonsCount)
	t.Run("Stamps", testStampsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Persons", testPersonsHooks)
	t.Run("Stamps", testStampsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Persons", testPersonsInsert)
	t.Run("Persons", testPersonsInsertWhitelist)
	t.Run("Stamps", testStampsInsert)
	t.Run("Stamps", testStampsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("StampToPersonUsingPerson", testStampToOnePersonUsingPerson)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PersonToStamps", testPersonToManyStamps)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("StampToPersonUsingStamps", testStampToOneSetOpPersonUsingPerson)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("StampToPersonUsingStamps", testStampToOneRemoveOpPersonUsingPerson)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PersonToStamps", testPersonToManyAddOpStamps)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PersonToStamps", testPersonToManySetOpStamps)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PersonToStamps", testPersonToManyRemoveOpStamps)
}

func TestReload(t *testing.T) {
	t.Run("Persons", testPersonsReload)
	t.Run("Stamps", testStampsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Persons", testPersonsReloadAll)
	t.Run("Stamps", testStampsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Persons", testPersonsSelect)
	t.Run("Stamps", testStampsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Persons", testPersonsUpdate)
	t.Run("Stamps", testStampsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Persons", testPersonsSliceUpdateAll)
	t.Run("Stamps", testStampsSliceUpdateAll)
}

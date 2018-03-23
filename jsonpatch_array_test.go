package jsonpatch

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var arraySrc = `
{
	"spec": {
		"loadBalancerSourceRanges": [
			"192.101.0.0/16",
			"192.0.0.0/24"
		]
	}
}
`

var arrayDst = `
{
	"spec": {
		"loadBalancerSourceRanges": [
			"192.101.0.0/24"
		]
	}
}
`

func TestArraySame(t *testing.T) {
	patch, e := CreatePatch([]byte(arraySrc), []byte(arraySrc))
	assert.NoError(t, e)
	assert.Equal(t, len(patch), 0, "they should be equal")
}

func TestArrayBoolReplace(t *testing.T) {
	patch, e := CreatePatch([]byte(arraySrc), []byte(arrayDst))
	assert.NoError(t, e)
	assert.Equal(t, 2, len(patch), "they should be equal")
	sort.Sort(ByPath(patch))

	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/spec/loadBalancerSourceRanges/0", change.Path, "they should be equal")
	assert.Equal(t, "192.101.0.0/24", change.Value, "they should be equal")
	change = patch[1]
	assert.Equal(t, change.Operation, "remove", "they should be equal")
	assert.Equal(t, change.Path, "/spec/loadBalancerSourceRanges/1", "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")
}

func TestArrayAlmostSame(t *testing.T) {
	src := `{"Lines":[1,2,3,4,5,6,7,8,9,10]}`
	to := `{"Lines":[2,3,4,5,6,7,8,9,10,11]}`
	patch, e := CreatePatch([]byte(src), []byte(to))
	assert.NoError(t, e)
	assert.Equal(t, 2, len(patch), "they should be equal")

	change := patch[0]
	assert.Equal(t, change.Operation, "add", "they should be equal")
	assert.Equal(t, change.Path, "/Lines/10", "they should be equal")
	assert.Equal(t, float64(11), change.Value, "they should be equal")
	change = patch[1]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/Lines/0", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")
}

func TestArrayReplacement(t *testing.T) {
	src := `{"letters":["A","B","C","D","E","F","G","H","I","J","K"]}`
	to := `{"letters":["L","M","N"]}`
	patch, e := CreatePatch([]byte(src), []byte(to))
	assert.NoError(t, e)
	assert.Equal(t, 11, len(patch), "they should be equal")

	change := patch[0]
	assert.Equal(t, change.Operation, "replace", "they should be equal")
	assert.Equal(t, change.Path, "/letters/0", "they should be equal")
	assert.Equal(t, "L", change.Value, "they should be equal")

	change = patch[1]
	assert.Equal(t, change.Operation, "replace", "they should be equal")
	assert.Equal(t, change.Path, "/letters/1", "they should be equal")
	assert.Equal(t, "M", change.Value, "they should be equal")

	change = patch[2]
	assert.Equal(t, change.Operation, "replace", "they should be equal")
	assert.Equal(t, change.Path, "/letters/2", "they should be equal")
	assert.Equal(t, "N", change.Value, "they should be equal")

	change = patch[3]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/10", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[4]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/9", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[5]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/8", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[6]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/7", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[7]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/6", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[8]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/5", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[9]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/4", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")

	change = patch[10]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/letters/3", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")
}

package simplejson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplejson(t *testing.T) {
	var err error

	js, err := NewJson([]byte(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"string_array_null": ["abc", null, "efg"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true,
			"sub_obj": {"a": 1}
		}
	}`))

	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	aws := js.Get("test").Get("arraywithsubs")
	assert.NotEqual(t, nil, aws)

	i, _ := js.Get("test").Get("int").Int()
	assert.Equal(t, 10, i)
}

/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2020/12/02 20:20
*/

package gopay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBodyMap_Set(t *testing.T) {
	var b = make(BodyMap)
	b.Set("test",func(b BodyMap){
		b.Set("sub", func(b BodyMap) {
			b.Set("sub-sub","value")
		})
	})

	assert.Equal(t, BodyMap{
		"test": BodyMap{
			"sub":BodyMap{
				"sub-sub":"value",
			},
		},
	}, b)

	var b2=make(BodyMap)
	b2.Set("test","value")
	assert.Equal(t, BodyMap{
		"test": "value",
	}, b2)
}
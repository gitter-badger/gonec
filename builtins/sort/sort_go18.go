// +build go1.8

package sort

import (
	s "sort"

	"github.com/covrom/gonec/vm"
)

func handleGo18(m *vm.Env) {
	m.DefineS("Slice", func(arr interface{}, less func(i, j int) bool) interface{} {
		s.Slice(arr, less)
		return arr
	})
}

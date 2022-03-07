package composite

import (
	"fmt"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	got := NewOrganization()
	tmp := got.(*Department)
	for _, v := range tmp.SubOrganizations {
		fmt.Println(v)
	}
	// assert.Equal(t, 20, got)
}

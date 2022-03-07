package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(new(User))

	err := proxy.Login("test", "123456")

	require.Nil(t, err)
}

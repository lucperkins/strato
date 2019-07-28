package strato

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGrpcClient(t *testing.T) {
	is := assert.New(t)

	cl, err := NewClient(goodClientCfg)

	t.Run("Instantiation", func(t *testing.T) {
		is.NoError(err)
		is.NotNil(cl)

		noAddressCfg := &ClientConfig{
			Address: "",
		}

		noClient, err := NewClient(noAddressCfg)
		is.Error(err, ErrNoAddress)
		is.Nil(noClient)

		badAddressCfg := &ClientConfig{
			Address: "1:2:3",
		}
		badCl, err := NewClient(badAddressCfg)
		is.NoError(err)
		is.NotNil(badCl)

		err = badCl.KVDelete(&Location{Bucket: "does-not-exist", Key: "does-not-exist"})
		stat, ok := status.FromError(err)
		is.True(ok)
		is.Equal(stat.Code(), codes.Unavailable)
	})
}

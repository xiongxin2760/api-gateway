package bootstrap

import (
	upstreamserver "api-gateway/business/upstreamServer"
	"context"
)

func initUpServer(ctx context.Context) {
	upstreamserver.ServerDiscoveryHandle = upstreamserver.NewServerDiscovery(ctx)
	err := upstreamserver.ServerDiscoveryHandle.Init(ctx)
	if err != nil {
		panic(err)
	}
}

package mongotesting

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
	"testing"
)

const (
	image         = "mongo:4.1"
	containerPort = "27017/tcp"
)

func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "27018", // 0 就是随机端口
				},
			},
		},
	}, nil, nil, "test")
	if err != nil {
		panic(err)
	}
	containerId := resp.ID
	defer func() {
		err := c.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			log.Fatalf("error remonving container:%v", err)
		}
	}()

	err = c.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	inspect, err := c.ContainerInspect(ctx, containerId)
	if err != nil {
		panic(err)
	}

	hostProt := inspect.NetworkSettings.Ports[containerPort][0]

	*mongoURI = fmt.Sprintf("mongodb://root:123456@%s:%s", hostProt.HostIP, hostProt.HostPort)

	return m.Run()
}

package app

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Parameters struct {
	Image string
	Url   string
	Users int
}

func RunContainer(parameter string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

}

func StartAndRunContainers(params Parameters) {
	names := []string{"Istio", "Axios", "Sidero", "Thanos", "Odigos", "Kiali", "Persis", "Cyrus", "Darius", "Dario"}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("here")
		panic(err)
	}
	defer cli.Close()
	numbers := params.Users / 100
	for i := 1; i <= numbers; i++ {
		containerName := "cont-" + strconv.Itoa(i)
		envCases := []string{"LINK_CLASS=" + params.Url, "ROBOT_NAME=" + names[i]}
		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Tty:   true,
			Image: params.Image,
			Env:   envCases,
		}, nil, nil, nil, containerName)
		if err != nil {
			panic(err)
		}

		fmt.Println(i)
		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}
		fmt.Println(resp.ID)
	}
	// envCases := []string{"LINK_CLASS=https://test.alocom.co/class/zaeem/e8a2f45f", "ROBOT_NAME=golang"}

}

// Stop and remove a container
func StopAndRemoveContainer(containername string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("here")
		panic(err)
	}
	defer cli.Close()

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}
	if err := cli.ContainerRemove(ctx, containername, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return err
	}

	return nil
}

package app

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

)
type Parameters struct{
    Image string
	Url   string
	Users int
}



func TestContainerApp(parameter Parameters) string {
	names := []string{"Axios", "Sidero", "Thanos", "Odigos", "Kiali", "Istio", "Persis", "Cyrus", "Darius", "Dario"}
	// fmt.Printf("%+v\n", Paramers{"my-selnium2", parameter, in})
	// envCases := []string{"LINK_CLASS="+parameter, "ROBOT_NAME="+names[0]}
	fmt.Printf("%+v\n", names)
	fmt.Printf("%+v\n", parameter.Url)
	return "response"
}

func RunContainer2(parameter string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All:true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

}

func RunContainer4(params Parameters) {
	names := []string{"Axios", "Sidero", "Thanos", "Odigos", "Kiali", "Istio", "Persis", "Cyrus", "Darius", "Dario"}
	envCases := []string{"LINK_CLASS=" + params.Url, "ROBOT_NAME="+names[0]}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("here")
		panic(err)
	}
	defer cli.Close()
	numbers = params.Users / 100 
	// envCases := []string{"LINK_CLASS=https://test.alocom.co/class/zaeem/e8a2f45f", "ROBOT_NAME=golang"}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Tty: true,
		Image: params.Image,
		Env:envCases,
		}, nil, nil, nil, "cont-2")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

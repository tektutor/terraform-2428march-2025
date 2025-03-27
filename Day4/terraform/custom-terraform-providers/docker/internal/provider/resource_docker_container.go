package provider

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerContainer() *schema.Resource {

	return &schema.Resource{

		Description: "This resource will support create, read, update and delete docker container resources via Terraform.",

		CreateContext: resourceCreateDockerContainer,
		ReadContext  : resourceReadDockerContainer,
		UpdateContext: resourceUpdateDockerContainer,
		DeleteContext: resourceDeleteDockerContainer,

		Schema: map[string]*schema.Schema{
			"container_name": {
				Description: "Name of the container.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"host_name": {
				Description: "Hostname of the container.",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"image_name": {
				Description: "Name of the docker image.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceCreateDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve docker container parameters provider by user from the terraform .tf scripts
	imageName 	  := d.Get("image_name").(string)
	containerName 	  := d.Get("container_name").(string)
	containerHostname := d.Get("host_name").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    	if err != nil {
        	panic(err)
    	}
    	defer cli.Close()

   	resp, err := cli.ContainerCreate(
			ctx, 
			&container.Config{
        					Image: imageName,
						Hostname: containerHostname,
    			}, 
			nil, 
			nil, 
			nil, 
			containerName,
	)
    	if err != nil {
        	panic(err)
    	}

    	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
        	panic(err)
    	}

	d.Set("container_name", containerName)
	d.Set("host_name", containerName)
	d.SetId(resp.ID)

	return resourceReadDockerContainer(ctx,d,meta) 
}

func resourceReadDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	return nil
}

func resourceUpdateDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve docker container parameters provider by user from the terraform .tf scripts
	containerName 	  := d.Get("container_name").(string)
	//containerHostname := d.Get("host_name").(string)
	containerID := d.Id()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    	if err != nil {
        	panic(err)
    	}
    	defer cli.Close()

	err = cli.ContainerRename(
			ctx, 
			containerID,
			containerName, 
	)
    	if err != nil {
		log.Printf("Unable to update container : %s", err )
		panic(err)
    	}

	return nil
}

func resourceDeleteDockerContainer(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    	if err != nil {
        	panic(err)
    	}
    	defer cli.Close()

	//Retrieve docker container parameters provider by user from the terraform .tf scripts
	containerName 	  := d.Get("container_name").(string)
	removeOptions	  := container.RemoveOptions{RemoveVolumes: true, Force: true}

	err = cli.ContainerStop( ctx, containerName, container.StopOptions{} )
    	if err != nil {
		log.Printf("Unable to stop container: %s", err)
		return nil
    	}

	err = cli.ContainerRemove( ctx, containerName, removeOptions )
    	if err != nil {
		log.Printf("Unable to delete container: %s", err)
		return nil
    	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// DockerClient wraps the Docker client
type DockerClient struct {
	cli *client.Client
	ctx context.Context
}

// NewDockerClient creates a new Docker client
func NewDockerClient() (*DockerClient, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %v", err)
	}
	return &DockerClient{cli: cli, ctx: ctx}, nil
}

// Close closes the Docker client
func (d *DockerClient) Close() {
	d.cli.Close()
}

// GetDockerVersion gets Docker version information
func (d *DockerClient) GetDockerVersion() error {
	version, err := d.cli.ServerVersion(d.ctx)
	if err != nil {
		return fmt.Errorf("failed to get Docker version: %v", err)
	}

	fmt.Printf("Docker Version: %s\n", version.Version)
	fmt.Printf("API Version: %s\n", version.APIVersion)
	fmt.Printf("OS/Arch: %s/%s\n", version.Os, version.Arch)
	return nil
}

// ListImages lists all Docker images
func (d *DockerClient) ListImages() error {
	images, err := d.cli.ImageList(d.ctx, types.ImageListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list images: %v", err)
	}

	fmt.Println("\nLocal Docker Images:")
	for _, image := range images {
		fmt.Printf("ID: %s\n", image.ID[:12])
		fmt.Printf("Tags: %v\n", image.RepoTags)
		fmt.Printf("Size: %.2f MB\n\n", float64(image.Size)/1024/1024)
	}
	return nil
}

// PullImage pulls a Docker image
func (d *DockerClient) PullImage(imageName string) error {
	reader, err := d.cli.ImagePull(d.ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image: %v", err)
	}
	defer reader.Close()

	// Print pull progress
	io.Copy(os.Stdout, reader)
	return nil
}

// RemoveImage removes a Docker image
func (d *DockerClient) RemoveImage(imageID string) error {
	_, err := d.cli.ImageRemove(d.ctx, imageID, types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: true,
	})
	if err != nil {
		return fmt.Errorf("failed to remove image: %v", err)
	}
	fmt.Printf("Successfully removed image: %s\n", imageID)
	return nil
}

func main() {
	// Create Docker client
	docker, err := NewDockerClient()
	if err != nil {
		fmt.Printf("Error creating Docker client: %v\n", err)
		return
	}
	defer docker.Close()

	// Get Docker version
	fmt.Println("Docker Version Info:")
	if err := docker.GetDockerVersion(); err != nil {
		fmt.Printf("Error getting Docker version: %v\n", err)
	}

	// List existing images
	if err := docker.ListImages(); err != nil {
		fmt.Printf("Error listing images: %v\n", err)
	}

	// Pull a new image
	fmt.Println("\nPulling Alpine image:")
	if err := docker.PullImage("alpine:latest"); err != nil {
		fmt.Printf("Error pulling image: %v\n", err)
	}

	// List images again to see the new image
	if err := docker.ListImages(); err != nil {
		fmt.Printf("Error listing images: %v\n", err)
	}

	// Example of removing an image (commented out for safety)
	/*
	   fmt.Println("\nRemoving Alpine image:")
	   if err := docker.RemoveImage("alpine:latest"); err != nil {
	       fmt.Printf("Error removing image: %v\n", err)
	   }
	*/
}

terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

# This is an instance of the provider with optional configurations
provider "docker" {
  # Configuration options
}

data "docker_image" "tektutor_docker_image" {
   name = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "my_ubuntu_container1" {
   image = data.docker_image.tektutor_docker_image.name
   name  = "ubuntu_container_1"
}

resource "docker_container" "my_ubuntu_container2" {
   image = data.docker_image.tektutor_docker_image.name
   name  = "ubuntu_container_2"
}



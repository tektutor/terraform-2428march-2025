data "docker_image" "tektutor_docker_image" {
   name = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "my_ubuntu_container1" {
   image = data.docker_image.tektutor_docker_image.name
   name  = var.container_name1 
}

resource "docker_container" "my_ubuntu_container2" {
   image = data.docker_image.tektutor_docker_image.name
   name  = var.container_name2 
}



data "docker_image" "tektutor_image" {
  name = var.image_name
}

resource "docker_container" "container" {
   count    = var.container_count

   name     = "container-${count.index+1}"
   hostname = "container-${count.index+1}"

   image    = data.docker_image.tektutor_image.name
   must_run = true

   ports {
     internal = "22"
     external = "200${count.index+1}"
   }
   ports {
     internal = "80"
     external = "800${count.index+1}"
   }
}

resource "local_file" "remote-execution" {
   count    = var.container_count
   content  = "Some file content - ${count.index+1}"
   filename = "some${count.index+1}.txt"

   connection {
      type = "ssh"
      user = "root"
      port = "200${count.index+1}"
      host = "localhost"
   }

   provisioner "remote-exec" {
     inline = [
        "hostname",
        "hostname -i",
     ]
   }
}

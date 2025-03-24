# Day 1

## Info - Hypervisor Overview
<pre>
- virtualization technology
- using Hypervisor, we can run multiple OS in a laptop/desktop/workstation/servers
- i.e many OS can run side by side on the same machine
- There are two types of Hypervisor
  - Type 1 and 
     - used in Servers & Workstations
     - Bare Metal Hypervisors
     - VMs can be created directly on top of Hardware 
     - examples
       - VMWare vSphere/vCenter
       - KVM
  - Type 2
    - used in laptops/desktops and Workstations
    - they can be installed on top a Host OS ( Windows, Linux, Mac )
    - examples
      - VMWare Workstation ( Linux & Windows )
      - VMWare Fusion ( Mac OS-X )
      - Oracle VirtualBox ( Windows, Linux & Mac )
      - Parallels ( Mac OS-X )
      - Microsoft Hyper-v ( Windows )
- the Operating System that runs within the Virtual Machine(VM) are called Guest OS
- each VM represents one fully functional Operating system
  - each OS has its own OS Kernel
  - its own hardware resources
  - its own network and graphics cards
  - has its own file system
  - has its own port range ( 0 to 65535 ports )
- this type of virtualization is called Heavy weight virtualization
  - the reason being, each VM(Guest OS) requires dedicated hardware resources
    - CPU Cores
    - RAM
    - Storage ( HDD/SDD )
    - Virtual Network Cards
    - Virtual Graphics Cards
</pre>

## Info - Containerization Technology
<pre>
- is a light weight virtualization technology
- all the containers that runs in the same OS/machine, shares the hardware resources available on the Host OS
- all the containers that runs in the same OS/machine, shares the OS Kernel on the Host OS
- each containers represents one application or an application component ( Frontend, Backend(DB Servers), etc.)
- containers are not Operating System
- containers are not a replacement for Virtual Machine or Operating Sytem
- technically containers can be created within a VM or OS installed on a Physical machine
- VMs and containers are used in combination, they are not competing technology they are complementing technology
- as each container is an application process, not a OS, technically comparing a container with a VMs is wrong
- containers and VMs share certain common behaviours/features
  - just like VMs, containers also get their own dedicated Network cards( virtual network card - software defined network cards )
  - just like VMs, containers also get its own dedicated network stack ( 7 OSI Layers )
  - just like VMs, containers also has their own file system
  - just like VMs, containers also get its own port range ( 0 to 65535 ports )
  - just like VMs, containers also get its own IP Address ( generally private IP addresses )
- each container will host one application and its dependent configuration and dependencies
- with containerization technology, one can run linux containerized applications on Linux OS, Windows OS or Mac OS
</pre>

## Info - Container Engine
<pre>
- is a high-level software that helps managing conainers and container images
- is very user-friendly, abstracts all the lowel-level OS features that enables the container technology
- end-users need not have to be an expert in Linux kernel or OS kernel features to create container or manage container images
- under the hood, container engines depends on Container Runtimes
- examples
  - Docker Container Engine depends on containerd which in turn depends on runC container runtime
  - Podman Container Engine depends on CRI-O container runtime
</pre>  

## Info - Container Runtime
<pre>
- is a low-level software that helps managing containers and container images
- it is not so user-friendly, hence end-users don't this use directly
- examples
  - runC container runtime
  - CRI-O container runtime
</pre>

## Info - Container Image
<pre>
- is a blueprint or specification of a container
- any software tool that need on a container has to installed on the image level
- when containers are created using a container image, whatever softwares are installed on the image are made available in the containers
- with One container image, multiple containers instances can be created
- Container Images are similar to Windows/Linux DVD ISO files
</pre>


## Info - Containers
<pre>
- container is a running instance of a Container Image
- each containers get its own IP address
- each container has its own file system, network card, port range, network stack, etc.,
- each container represents one application
- each containers should run only one main application
- though it is technically possible to run mutiple applications within a single container, it is not a best practice
- as per best practice, each container should host only one main application
- For example
  - a container can run a single REST API
  - a container can run a single SOAP API
  - a container can run a single Microservice
  - a container can run a single web server, application service, DB Server, Message Queue Server, etc.,
</pre>




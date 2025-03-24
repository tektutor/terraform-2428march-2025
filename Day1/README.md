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

## Containerization Technology
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
</pre>



## Please share your first day feedback from RPS Cloud machine
https://survey.zohopublic.com/zs/DkCYeg

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

## Info - Docker - Container Application Engine 
<pre>
- is developed in Go lang by a company called Docker Inc
- it follows client/server architecture
- it comes in 2 flavours
  1. Docker Community Edition - Docker CE
     - open source
     - hence no support can be expected
  2. Docker Enterprise Edition - Docker EE
     - requires buying commercial license
     - hence worldwide support is offered by Docker Inc organization
</pre>

## Info - Configuration Management Tools
<pre>
- is a DevOps tool that helps automating system administrative configuration management activities\
- Configuration Management
  - installing a software in Windows/Linux/Unix/Mac, network switches/routers
  - uninstalling a software in Windows/Linux/Unix/Mac, network switches/routers
  - managing Windows/Linux/Unix users
  - installing Oracle/MySql/MongoDB/CouchBase DB Servers, and then creating different DB users with various permissions
  - Importing Databases,tables and populating data into those tables, providing access to required users or denying access to certain users
  - already there is machine with Operating System pre-installed ( this is the assumption )
  - configuring Firewalls, application servers, web servers, connecting a machine to a domain, patching OS, etc.,
- Examples
  - Puppet
  - Chef
  - Salt/SaltStack
  - Ansible
</pre>

## Info - Puppet 
<pre>
- is a configuration management tool
- this is one of the oldest configuration management tool
- developed and maintained by a company called Perforce
- comes in 2 flavours
  1. Community Edition ( Free )
  2. Enterprise Edition ( requires commercial license ) 
- Domain Specific Language (DSL)
  - the language in which we write the automation script
  - uses Puppet's proprietary language to write the automation scripts
- follows client/server architecture ( based on Pull based architecture )
- software installation procedure is time consuming
- learning curve is also very steep
- each machine ( server ) managed by Puppet should have a Puppet agent running
</pre>

## Info - Chef
<pre>
- is a configuration management tool
- software installation procedure is time consuming
- learning curve is also very steep
- DSL used is Ruby scripting language
- follows client/server architecture ( based on Pull based architecture just like Puppet )
- each machine ( server ) managed by Chef should have a Chef agent running
- Chef comes with many tools for specialized automation purpose, hence mastering all these tools is very difficult
- comes in 2 flavours
  - a Free community edition with no support
  - a Paid Enterprise product with world-wide support
</pre>

## Info - Ansible
<pre>
- is developed in Python language by a company called Ansible Inc
- Ansible Inc company is founded by Michael Deehan
- Michael Deehan was an x employee of Red Hat
- easy to install and easy to learn 
- uses a very simple architecture
- agent-less
- uses YAML as the DSL
- comes in 2 flavours
  1. Ansible core 
     - open source and free
     - supports only command line interface
     - no user management, hence no role based access control
  2. Ansible Automation Platform  ( aka Ansible Tower )
     - AWX
       - open source and free
       - is developed on top of open source Ansible core
       - supports web console
       - no support
     - Red Hat Ansible Automation Platform 
       - an enterprise product that requires commercial license
       - developed on top of open source AWX
       - supports web console
       - also world-wide support is provided by Red Hat ( an IBM company )
  - configuration management tools are generally used to automate system administration activities ( like sofware installations and configurations ) on remote servers
  - we need already provisioned machines ( Machine with OS ) to perform configuration management
</pre>


## Info - Ansible High-Level Architecture
#### Ansible Jargons
<pre>
- Ansible Controller Machine(ACM)
- Ansible Inventory  
  - Static and Dynamic
- Ansible Modules
- Ansible Plugins
- Ansible Collections
- Ansible Nodes
- Ansible Roles
</pre>
![Ansible](AnsibleHighLevelArchitecture.png)

#### Ansible Inventory
<pre>
- it has connections details to Windows/Unix/Linux/Mac servers that must be managed by Ansible
- for instance
  - it has login credentials, IP address/hostnames, ports, password, public/private login authentication details, etc.,
  - in case, the remote servers happens to have a Unix/Linux/Mac OS
    - SSH Login credentials are captured in this file
  - in case, the remote servers happens to have a Windows Os
    - WinRM Login credentials are captured in this file
</pre>

#### Ansible Controller Machine (ACM)
<pre>
- this could be your work laptop/desktop that has Ansible installed in it
- ansible is officially supported on any Linux Distributions ( Ubuntu, CentOS Stream, Rocky Linux, Fedora, Suse or RHEL, any other Linux Distributions )
- Windows machine can't act as a Ansible Controller Machine
- but Windows machines can be managed by Ansible
</pre>  

#### Ansible Nodes
<pre>
- are remote servers managed by Ansible Configuration Management Tool
- Ansible nodes can be
  - a Windows server
  - a Unix/Linux/Mac/Network switch/router, a IOT device, etc
  - the server could be running in a on-prem data center or public/privatte/hybrid cloud
  - i.e can be an ec2 instance running in AWS, an Azure VM, etc.,
  - it can be a Physical machine with some OS pre-installed in it
- Software dependencies
  - Windows Servers 
    - Powershell must be installed ( comes out of the box with any Windows Server )
    - WinRM must be configured with some login authentication ( comes out of the box with any Windows Server )
  - Unix/Linux/Mac Server
    - Python must be installed ( comes out of the box with any unix/linux/mac server )
    - SSH Server must be running 
</pre>

#### Ansible Modules
<pre>
- Ansible modules are automation reused code either implemented as Powershell scripts or Python scripts
- they come out of the box when we install any flavour of Ansible 
- each Ansible modules supports one automation activity
  - For example
    - Copy module helps in copying a file from local machine(ACM) to remote server( Ansible Node ) or vice versa
    - File module helps in creating/deleting a file/folder on remote machine( Ansible Node )
    - Shell module that helps running linux shell commands on remote servers ( Ansible Unix/Linux/Mac nodes ) 
    - Similary Ansible supports many Windows modules ( Powershell scripts ) for similar automation
</pre>

#### Ansible Plugins
<pre>
- plugins alter the default behavior of ansible
- for example,
  - in order to install/uninstall/upgrade softwares, we have to perform the automation as an Adminstrator/root user
  - this can be achieved by an ansible plugins called become, sudo, etc.,
  - equivalent to "Run As Administrator" in Windows
</pre>

#### Ansible Playbooks
<pre>
- automation code/script written in YAML file that follows a specific structure and sections
- it invokes one or more Ansible modules, roles, plugins, etc.,
</pre> 

#### Ansible Roles
<pre>
- is a reusable code that can be invoked from Ansible Playbooks
- it follows a specific directory structure and set of Best practices recommended by Ansible
- roles can't be executed directly, they are normally invoked from Ansible Playbooks
</pre>  

#### Ansible Collections
<pre>
- an ansible package that has
  - one to many Ansible plugins
  - one to many Ansile Modules
  - one to many Ansible playbooks
  - one to many Ansible Roles
</pre>

## Lab - Checking Docker and Ansible versions ( Linux Terminal )
```
ansible --version
docker --version
```

Expected output
![image](https://github.com/user-attachments/assets/8b76fafb-155e-444e-94db-c33f6c3e67c7)
![image](https://github.com/user-attachments/assets/33239ceb-ddf2-419a-9710-1b5a7f4253d0)
                                                                                                                                                          
## Lab - Building our Custom Docker Images to create ansible nodes containers
Generate public/private key pair as rps(non-admin) user,accept all defaults by hitting enter each time it prompts something
```
ssh-keygen
```

Expected output
![image](https://github.com/user-attachments/assets/ff0afb46-3904-4d4e-9e62-a89e95b190a2)

Let's build a custom ubuntu ansible node docker image
```
cd ~/terraform-2428march-2025
git pull
cd Day1/CustomDockerImages/ubuntu
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/ubuntu-ansible-node:latest .
docker images
```

Expected output
![image](https://github.com/user-attachments/assets/15202c81-a2c6-4872-81e6-499ca0dee77f)
![image](https://github.com/user-attachments/assets/fdab8801-9ae0-4c6e-957d-dd374d57ed38)
![image](https://github.com/user-attachments/assets/a7bc08a1-3992-424f-a5e6-4de53146bab9)

## Lab - Let's create two containers using the Custom ubuntu ansible node image we build in the previous exercise

In case you already have containers with the same name or other containers using the same port, you can delete those containers
```
docker rm -f $(docker ps -aq)
```

If you wish to delete just specific containers
```
docker rm -f ubuntu1
docker rm -f ubuntu2
```

The below command will create two containers and starts the two containers
```
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
```
![Docker](PortForwarding.png)

Listing the currently running containers
```
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/b98ac94f-7308-4744-8416-b7481a74ee0b)

## Lab - Let's verify if we are able to SSH into ubuntu1 and ubuntu2 containers without providing password
```
docker ps

ssh -p 2001 root@localhost
exit

ssh -p 2002 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/6feda23c-249f-491a-96f8-507159185b9e)
![image](https://github.com/user-attachments/assets/61d93920-745e-4542-8ef6-dd7dcc8d363b)

## Lab - Building a custom rocky linux ansible node docker image to use it as an ansible node
```
cd ~/terraform-2428march-2025
git pull
cd Day1/CustomDockerImages/rocky
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/rocky-ansible-node:latest .
docker images
```

Expected output
![image](https://github.com/user-attachments/assets/289e644b-b174-4c07-b5b8-a3494086103d)
![image](https://github.com/user-attachments/assets/5fb40aaf-1b1f-4179-b0f9-00e26d8b912f)

## Lab - Create two rocky container to use them as ansible nodes
```
docker images | grep tektutor/rocky-ansible-node
docker run -d --name rocky1 --hostname rocky1 -p 2003:22 -p 8003:80 tektutor/rocky-ansible-node:latest
docker run -d --name rocky2 --hostname rocky2 -p 2004:22 -p 8004:80 tektutor/rocky-ansible-node:latest
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/6af972dc-6c5e-4283-858b-38745529e069)


## Lab - Let's verify if we are able to SSH into rocky1 and rocky2 ansible node containers with providing password
```
docker ps

ssh -p 2003 root@localhost
exit

ssh -p 2004 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/c9d1fe51-be31-4158-9513-f2cd29c0ee16)

## Lab - Finding IP address of ubuntu1, ubuntu2, rocky1 and rocky2 ansible nodes
```
docker inspect -f {{.NetworkSettings.IPAddress}} ubuntu1
docker inspect -f {{.NetworkSettings.IPAddress}} ubuntu2
docker inspect rocky1 | grep IPA
docker inspect rocky2 | grep IPA
```

Expected output
![image](https://github.com/user-attachments/assets/0f723f32-adc8-4b74-977d-a0b52f1e7a49)

## Lab - Running ansible ad-hoc command to test the ansible node connectivity with ansible ping module
In the below ansible ad-hoc command,
<pre>
i - switch indicates what follows is the name of the inventory file
all - indicates the group of machines you wish to ping
dev - indicates you wish to ping all the machines under dev group
prod - indicates you wish to ping all the machines under prod group
m - switch indicates, the module you wish to run on the ansible node/group of nodes
ubuntu1 - indicates you wish to ping only the ubuntu1 ansible node
rocky1 - indicates you wish to ping only the rocky1 ansible node 
ping - is the name of the ansible module (ping.py) you wish to run on the ansible node(s)
</pre>

```
cd ~/terraform-2428march-2025
git pull
cd Day1/ansible/static-inventory
ansible -i hosts all -m ping

ansible -i hosts dev -m ping
ansible -i hosts prod -m ping

ansible -i hosts ubuntu1 -m ping
ansible -i hosts rocky1 -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/84259fd0-faaa-4b3c-bd3b-b96f98276aa4)
![image](https://github.com/user-attachments/assets/880010ea-c1fc-49a3-a532-61d61b93a761)
![image](https://github.com/user-attachments/assets/08ca28dd-a8d8-4ee9-8a05-7ab5b688fded)

## Info - What happens behind the scenes when we issue ping ansible ad-hoc command
```
ansible -i hosts all -m ping
```

The below chain of activities happens
<pre>
- the below activities are done in parallel to all ansible nodes
- ansible will create a temp directory on the Ansible Controller Machine
- it copies the ping.py ansible module to the temp directory
- ansible creates a temp directory on the remote ansible node
- ansible copies the ping.py ansible module from Ansible Controller Machine(ACM) to the remote ansible nodes(s) using SCP/SFTP 
- ansible remotes assigns execute permission to the ping.py on the ansible nodes
- ansible runs the python script on the remote machine
- ansible captures the output of the ping.py python script
- ansible deletes the temp folder on the remote ansible nodes
- ansible then gives a summary of the outcomes about all the ansible nodes
</pre>

## Lab - Collecting facts using ansible setup module using an ad-hoc command
<pre>
- ansible ad-hoc command helps invoking one ansible module at a time to try the module before using them in ansible playbooks
- ansible facts are meta-data about remote ansible nodes  
- some intersting facts ansible collects are
  - machine hard details
  - python version
  - OS details
    - Linux distribution name
    - Linux distribution version
  - package manager 
</pre>

```
cd ~/terraform-2428march-2025
git pull
cd Day1/ansible/static-inventory
ansible -i hosts ubuntu1 -m setup
```

Expected output
![image](https://github.com/user-attachments/assets/1cf39738-9655-4519-8afd-95eec6877be9)
![image](https://github.com/user-attachments/assets/a86cfb80-4552-4ae2-81e7-288ae62d0ea1)
![image](https://github.com/user-attachments/assets/dc421a93-6708-42fe-ac36-721ac996512b)
![image](https://github.com/user-attachments/assets/51161be3-5805-4158-9f0e-398aa99ae5c5)
![image](https://github.com/user-attachments/assets/5fa8727e-a139-4e19-a2f5-e989b5c588c2)
![image](https://github.com/user-attachments/assets/dc412bc5-b1bc-44f4-a0c7-2ec1e28fea69)
![image](https://github.com/user-attachments/assets/824616ab-c88b-4e76-9fbc-b7c5d22205d5)
![image](https://github.com/user-attachments/assets/64b5f249-a85e-4f4b-987b-90f8fd47264f)

## Info - Ansible Playbook Structure
<pre>
- Ansible Playbook is a YAML file
- Each Ansible Playbook has one to many Play
- Each Play has
  - zero to many tasks( list of tasks )
  - zero to many roles( list of roles )
  - targets a list of ansible nodes or aleast a single ansible node to run the module/playbook
- Playbook invokes one or more ansible modules in a particular sequence one after the other to perform a configuration management automation
</pre>

![ansible](StructureofAnsiblePlaybook.png)

## Info - Getting help about ansible modules

Listing all the ansible modules supported, to scroll you could use the page up/down or cursor movements keys ( up/down )
```
ansible-doc -l
ansible-doc ping
ansible-doc setup
ansible-doc shell
ansible-doc copy
ansible-doc command
ansible-doc file
ansible-doc templates
ansible-doc service
```

Expected output
![image](https://github.com/user-attachments/assets/cc68fb8b-d0a1-4a3f-ae2c-f78b912c22ea)
![image](https://github.com/user-attachments/assets/ace4ba8a-4b05-4f26-ab73-9e92719456e1)
![image](https://github.com/user-attachments/assets/45486b29-f997-41e4-9c61-3dbc694a7ce5)
![image](https://github.com/user-attachments/assets/da15774f-aab4-4baa-ae3d-39eb85697b4b)

## Info - Ansible configuration file (ansible.cfg)
<pre>
- the ansible.cfg file has user specific overrides, system wide configurations, etc.,
- ansible checks an environment variable called ANSIBLE_CONFIG to locate the path and name of the ansible.cfg as a first attempt to locate ansible.cfg
- if the environment is not defined, ansible attempts to locate the ansible.cfg in the current directory, if it finds it uses that and stops any further search for ansible.cfg
- in case the ansible.cfg is not found in the current directory, then it attempts to locate the .ansible.cfg file under the user home directory
- in case the ansible.cfg is not found in the home directory, then as a last attempt it will search for /etc/ansible/ansible.cfg file
- wherever ansible is able to locate the ansible.cfg file first, it uses that and stops further searching

- in ansible.cfg file, you could configure many things
  - ssh timeout
  - colors 
  - default sudo user
  - can point the inventory file that must be used
  - suppress warnings, etc.,
</pre>
![image](https://github.com/user-attachments/assets/fb73bfa1-6f41-4598-be56-d4922bf2f94a)

## Lab - Running your first ansible playbook
```
cd ~/terraform-2428march-2025
git pull
cd Day1/ansible/ansible-playbooks
ansible-playbook ping-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/51c4950f-1746-4f9a-b652-9271b4a8ab8a)
![image](https://github.com/user-attachments/assets/1363bc71-b5aa-43f5-9c76-aeb8be47f209)

## Info - Ansible Host and Group variables in Ansible Inventory
```
cd ~/terraform-2428march-2025
git pull
cd Day1/ansible
cat hosts
cat hosts-refactored
```
<pre>
- In the refactored hosts inventory file, ansible_port is called Host variable as we had to mention the ansible_port for each ansible node.
- ansible_user, ansible_host, ansible_private_key_file are called Group variables as they are common for the ansible nodes listed under all group
- group variables can also be defined for each groups in the inventory file
</pre>

## Info - Ansible Idempotency Property
<pre>
- The DevOps engineer when they write an ansible playbook to put the ansible node into a desired state
- When ansible executes the playbook, it invokes the ansible modules, certain ansible modules are idempotent while others are not
- Whichever ansible module is idempotent, it will check the current state of the machine and compares that with the expected(desired) state of the machine, if the current state of the machine matches with the desired state then ansible will do nothing.  It simply reports the task as success with green color
- In case ansible finds the machine's current state is different from the desired state then ansible performs the necessary actions to ensure the current state of the machines matches the desired state and reports the status in Yellow color indicates it had to run the task and modify the machine in order to make it successful
- Subsequent executions, the current state of the machine when it matches the desired state it would simple report success with no change using green color
- this property is called Idempotency or Itempotent
</pre>

## Lab - Install nginx web server into Ubuntu ansible nodes using an ansible playbook
```
cd ~/terraform-2428march-2025
git pull
cd Day1/ansible/ansible-playbooks/nginx
cat install-nginx-playbook.yml
ansible-playbook install-nginx-playbook.yml
curl http://localhost:8001
curl http://localhost:8002
ansible ubuntu1 -m shell -a "service nginx status"

cat install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/5e6c9763-b3b2-46d5-a3db-260aef8fc587)
![image](https://github.com/user-attachments/assets/5e036125-d784-4ff9-a03b-ec386e505776)

In the above output
<pre>
- Green color indicates success with no change
- Yellow color indicates success with change, i.e during installation ansible would have created folders, copy certain files, updated existing files
- unreachable=0, means ansbile is able to connect to the ansible node via SSH
- failed=0, means all the tasks that were executed on individual anisble nodes were successful
- skipped=0, no conditional tasks were skipped on any ansible nodes
- rescued=0, indicates no tasks lead to executing exception(catch or recuse) code
- ignored=0, indicates no errors were ignored in any ansible node while executing the ansible playbook
</pre>

Now, let's copy the default configuration file from ubuntu1 container to local machine
```
docker cp ubuntu1:/etc/nginx/sites-available/default .
cat default
```
![image](https://github.com/user-attachments/assets/70e40817-102f-4995-8759-b2c3e0ee849c)
![image](https://github.com/user-attachments/assets/c00e87ce-3be4-4f93-8cd6-4b56d89d2b8d)

The updated playbook looks as shown below
![image](https://github.com/user-attachments/assets/d912d7a9-d7e7-49b2-8fc0-175f7ebce5ba)

Let's run the playbook
```
cd ~/terraform-2428march-2025
git pull
cd Day1/ansible/ansible-playbooks/nginx
cat install-nginx-playbook.yml
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/e1f8d2d3-d897-4eb3-8c9e-9d5faedc4fce)
![image](https://github.com/user-attachments/assets/44d89ebc-f16d-4afc-924b-7eb5fc019d95)
![image](https://github.com/user-attachments/assets/3a4d8279-383c-422d-b224-b2940a0dfa42)
![image](https://github.com/user-attachments/assets/dc4864cf-8d25-4e22-b2c3-22bd1d1ba9d9)

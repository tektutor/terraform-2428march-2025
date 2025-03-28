# Day 5

## Feedback
<pre>
https://survey.zohopublic.com/zs/Y7ftn4  
</pre>

## Lab - Upgrade Docker Community Edition in your lab machine 
```
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

docker --version
sudo service docker status
```

## Info - Terraform Provider Development best practices and recommended naming conventions
<pre>
- Provider name must be terraform-provider-nameoftheprovider, must be all lower case
- Resource names must start with nameoftheprovider-resource i.e docker_container, docker is the provider name while the resource managed is container
- resource and data source name must be all lowercase separated by underscore, and recommened to restrict to 2 or 3 words at the max
</pre>

## Info - Commons causes of Configuration Drift
<pre>
- manual changes
- inconsistent and manual deployments
- external dependencies
- difference in environments
- lack of version control
- poor or lack of documentation
</pre>

## Info - Risks associated with Configuration Drifts
<pre>
- Security Vulnerabilities
- Performance issues
- Compliance issues
- Makes troubleshooting very difficult
- increased down-time
- decreased reliability
- poor user-experience
</pre>

## Info - Solution to Configuration Drifts
<pre>
- Using Version Control
- Continuous Integration
- Use Configuration Management Tools to override manual changes in continuous fashion
- Use Infrastructure as Code Tools to override manual changes
- Regularly monitor and audit infrastructure
</pre>

## Info - Monitoring
<pre>
- is the process of collecting and analyzing data from IT infrastructure, system and processes
- using the data to improve business outcomes and drive value to the organization
- collects data to help keep your infrastructure up and running without any downtime
- Tools to
  - Data Collection (Logs)
  - Alerting
  - Remediation
  - Agent based monitoring
  - Agentless monitoring
  - Collecting Metrics
- Examples
  - Dynatrace
  - DataDog
  - Splunk
  - Prometheus & Grafana
</pre>

## Continuous Integration (CI) Overview
<pre>
- it is a fail-fast engineering process to find issues early 
- when bugs are detected early during development phase, they are easy to fix
- cost of fixing the bugs is also relatively cheaper
- it is similar SCRUM daily stand-up meeting, which is an inspect and adapt meeting
- the team that follows CI/CD, the engineers would be pushing code to version control several times a day
- code would be integrated many times a day
- Jenkins or similar CI/CD server will grab the latest code, they trigger a build, as part of the build, automated test cases would be executed to verify if the new code is as expected, if the new code is breaking any existing functionality.
- the build that was triggered due to new code integration succeeds, it means no functionality is broken, everything works as expected
- continous frequent feedback is given by CI/CD builds, eventually improving the code quaility and functional quality
- CI helps confidently deliver releases in a short amount of time
- Unit and Integration is the scope of CI
</pre>

## Continuous Deployment (CD) Overview
<pre>
- Once the dev release suceeds all the automated test cases added by dev team, it is automatically promoted for QA testing
- the dev certified release binaries are deployed automatically to QA environment for further automated QA testing
- the QA would automate, end to end functionality test, regression test, smoke test, performance test, stress test, component/API test, etc
</pre>

## Continuous Delivery (CD) Overview
<pre>
- the QA certified build(release) is automatically deployed into production or pre-prod environment
- in the pre-prod environment the customer or the Operations team would verify if the new release is working as expected
- especially fintech companies, after manual approval the binaries could go live in production environment
</pre>

## What is Jenkins?
- is a build automation server
- used mainly for CI/CD Builds
- this was developed in Java by Kohsuke Kawaguchi, former employee of Sun Microsystems
- Initially it was developed as Hudson is an opensource project
- Then Oracle acquired Sun Microsystems, then part of Hudson including Kohsuke Kawaguchi had quit Oracle
  created a new branch from Hudson called Jenkins
- The other part of the Hudson team they continue to develop the product as Hudson
- There is lot common code between Hudson and Jenkins
- More than 10000 active contributors are there for Jenkins
- Many other software vendors got inspired by Jenkins similar products came out in market like Bamboo, Team City, Microsoft TFS, etc.,
- Jenkins supports CI/CD build for products built in any software stack
  
## What is Cloudbees?
- Cloudbees is the enterprise paid variant of Jenkins
- Feature wise Jenkins and Cloudbees pretty much they are same
- We get support for Cloudbees while we only get community support for Jenkins
- Cloudbees is more stable as it is a paid version
  
## Jenkins Alternatives
- Bamboo
- Team City
- Cloudbees ( Enterprise Jenkins )
- Microsoft Team Foundation Server (TFS)

## Lab - Download Jenkins war file
Download the Generic Java Package (war) file from the left side (LTS)
<pre>
cd ~/Downloads
wget https://get.jenkins.io/war-stable/2.492.1/jenkins.war
</pre>

Expected output

## Lab - Launching Jenkins from terminal
```
cd ~/Downloads
java -jar ./jenkins.war
```

Expected output
![image](https://github.com/user-attachments/assets/8fe88b12-c923-4137-baa4-3737b1a50543)
![image](https://github.com/user-attachments/assets/517588c7-62d1-4938-b52d-3a88cfc7b35a)

## Lab - Accessing Jenkins Dashboard and configuring Jenkins from your RPS Cloud machine chrome web browser
<pre>
http://localhost:8080  
</pre>

Expected output
![image](https://github.com/user-attachments/assets/b0726299-8dcd-493b-9309-88181b4ca214)
![image](https://github.com/user-attachments/assets/7b406de9-076d-4c58-abd7-c48b07500670)
![image](https://github.com/user-attachments/assets/ae115b35-1f87-44d0-9e69-4e971b80d25e)
![image](https://github.com/user-attachments/assets/b9650839-c85d-4d25-96b6-c5240bf8e492)
Click "continue"
![image](https://github.com/user-attachments/assets/bfebad20-6ffa-4f57-8148-5dc5777ee691)
Select "Install Suggested Plugins"
![image](https://github.com/user-attachments/assets/d42c55d8-0c4f-4621-81d2-20b65ed7f36b)
![image](https://github.com/user-attachments/assets/3a603538-86f9-4df0-be7f-451744cf0fb8)
![image](https://github.com/user-attachments/assets/a7d38af7-a97c-4965-8ca4-257091edc999)
![image](https://github.com/user-attachments/assets/e1442fdc-81b3-4dac-a633-59a9e7adbfed)
![image](https://github.com/user-attachments/assets/d0bb1455-21ea-4b1e-97ad-3bd854969b82)
![image](https://github.com/user-attachments/assets/e33fafd8-54f5-4350-bdb8-e7d7761fca26)
![image](https://github.com/user-attachments/assets/f946ef9d-a3e5-41c1-9590-6ec818bb1d73)
Click "Save and Continue"
![image](https://github.com/user-attachments/assets/3b63e7b5-229d-43c5-83d0-dbcf728c3380)
click "Save and Finish"
![image](https://github.com/user-attachments/assets/7848db80-77c3-4e3a-b794-304455e7e95e)
Click "Start using Jenkins"
![image](https://github.com/user-attachments/assets/5346ecc3-204b-4c7c-a4fb-9b740d2265f5)

## Lab - Let's create a Jenkins Job to setup a CI/CD Pipeline for our Training Repository in GitHub

Let's launch the Jenkins Dashboard in chrome web browser
![image](https://github.com/user-attachments/assets/105786c8-5b44-446e-b79e-7bde879c424e)

Let's click on "Create a Job"
![image](https://github.com/user-attachments/assets/0fc6378d-c3cb-4e8c-b456-ce5760e3f382)

Under Enter an item name, type "DevOpsCICDPipeline" and select Pipeline under Select an item
![image](https://github.com/user-attachments/assets/e03adc52-2a90-4807-95a3-ed6d894c4ace)
Click "Ok"
![image](https://github.com/user-attachments/assets/4ca923ae-5dfd-4518-9010-ac7fc857a0d5)

General Section
![image](https://github.com/user-attachments/assets/06b1464b-4a61-4fb3-9510-4579ac0eb166)

Triggers Section
Select Poll SCM, Under the Schedule type "H/02 * * * *"
![image](https://github.com/user-attachments/assets/79cf9118-c3f3-4ce3-8614-ee203a665fc1)

Pipeline Section

Click Definition and Select "Pipeline script from SCM"
![image](https://github.com/user-attachments/assets/0de66a0e-3384-4ebc-a292-798320db4bb3)
Under the SCM, select "Git"
![image](https://github.com/user-attachments/assets/6494c9c5-a837-4115-a3b6-4a7cb03428a5)
Under the Repository URL, type "https://github.com/tektutor/terraform-2428march-2025.git" without quotes
Under the Branches to Build, replace "master" with "main"
![image](https://github.com/user-attachments/assets/68d51815-0dd8-4df5-9ae0-239adee48314)
![image](https://github.com/user-attachments/assets/cc0d1dac-ae9b-4d3d-a3c5-b1fd0a6ccbb7)
![image](https://github.com/user-attachments/assets/316597d8-bb2b-4a91-926c-9e6cf4ab76c9)
Under the Script Path Replace "Jenkinsfile" with "Day5/DevOpsCICDPipeline/Jenkinsfile" without quotes
![image](https://github.com/user-attachments/assets/356befcd-76b4-42b5-8a45-ec4f30885193)
Click "Save"
![image](https://github.com/user-attachments/assets/72bbfb19-6acc-40de-90f8-8f4ee4c06ee5)

Progress
![image](https://github.com/user-attachments/assets/55a96db3-3026-4b99-a68c-ef30e26c0e5c)
![image](https://github.com/user-attachments/assets/4f383eb4-5fec-4d65-84b8-61660c5536b2)
![image](https://github.com/user-attachments/assets/df4608f3-d8e2-45c1-ba6b-fb4c2870c6b1)
![image](https://github.com/user-attachments/assets/4a486b7e-63ad-44e3-9e64-eb9d98fc2836)
![image](https://github.com/user-attachments/assets/37af0926-3666-43c9-a206-2765be893587)
![image](https://github.com/user-attachments/assets/71ab5117-685a-4512-8f1f-73a18461d3d7)






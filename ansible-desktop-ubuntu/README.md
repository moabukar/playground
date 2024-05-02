# Ansible Playbook for configuring Ubuntu Desktop


## Introduction

This is a collection of roles I use for my Ubuntu desktop deployment.


## Requirements

1) Basic knowledge of Ansible.

2) Ubuntu (or alternative apt based Linux Distro would likely work, may require
minor changes)

3) ansible\
````pip3 install ansible````

4) jinja2 template\
````pip3 install jinja2````


## Quick Start

*This assumes a brand new installation and the execution of this Playbook is
is on the target machine.  In other words, the deployment server and client are
the same system.  Of course, this playbook can be run from a remote server, if
preferred.*

1. ````sudo apt update && sudo apt install git sshpass openssh-server -y````
** *Limit use of sshpass for early setup only, due to potential security issues.
Deploy ssh keys to target host(s) after this playbook has executed successfully.* **

2.````git clone https://github.com/moabukar/playground/````

3.````cd playground/ansible-desktop-ubuntu````

1. Amend inventory file if needed, default target is localhost.

2. Amend main.yml file for roles of software desired.

* Many of the third party packages are broken into separate roles, this was
setup this way to allow convenient inclusion/exclusion of roles as needed by
commenting/uncommenting roles in main.yml

6. ansible-playbook main.yml -bkKu <username>
  * enter SSH password
  * enter SUDO password. (assumes the user is a part of the sudo user group)


# go-netbox

[![Lisence](https://img.shields.io/badge/license-ISC-informational?style=flat-square)](https://github.com/smutel/go-netbox/blob/master/LICENSE)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-informational.svg?style=flat-square&logo=git)](https://conventionalcommits.org)
[![Build Status](https://img.shields.io/github/workflow/status/smutel/go-netbox/Master/master?style=flat-square&logo=github-actions)](https://github.com/smutel/go-netbox/actions)

Go library to interact with NetBox IPAM and DCIM service. 

## Compatibility with Netbox

The version for Netbox and go-netbox will the same except for the last digit.

Example:
* go-netbox v1.10.x is working with Netbox v1.10
* go-netbox v1.11.x is working with Netbox v1.11

## How to contribute to this project

* To contribute to this project, please follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0-beta.2/) rules.
* Most of the code of this project will be generated using the swagger spec of Netbox and the go-swagger(https://github.com/go-swagger/go-swagger) program.
* You can change the behavior of the generated library by pushing patchs in the `patchs` folder.
* The best is to see if the bug is due to a wrong swagger definition and to report also this bug to the Netbox(https://github.com/netbox-community/netbox) project.
* If the bug is due to the go-swagger program the best is to create a bug there also go-swagger(https://github.com/go-swagger/go-swagger).

## How to test your work locally

### Requirements

* docker
* docker-compose
* swagger(https://github.com/go-swagger/go-swagger) installed somewhere (/usr/local/bin)
* netbox-docker(https://github.com/netbox-community/netbox-docker.git) project installed somewhere

### Installing the go-netbox

```sh
$ mkdir -p ~/go/src/github.com/smutel
$ cd ~/go/src/github.com/smutel
$ git clone git@github.com:smutel/go-netbox.git
```

### Starting netbox and getting the swagger definitions

You can do this by executing the commands below:
```sh
$ cd ~/go/src/github.com/smutel/go-netbox
$ NETBOX_MAJOR_VERSION=$(cat netbox_major_version)
$ LAST_NETBOX_VERSION="$(./utils/netbox_get_last_version ${NETBOX_MAJOR_VERSION})"
$ export VERSION=${LAST_NETBOX_VERSION}
$ cd <path to netbox-docker>
$ mv docker-compose.override.yml.example docker-compose.override.yml
$ docker-compose up -d
$ while ! curl -s http://127.0.0.1:8000/api/swagger.json -o /tmp/swagger.json 2> /dev/null; do sleep 1; done
$ 
```

### Regenerating the library

```sh
$ cd ~/go/src/github.com/smutel/go-netbox
$ rm -rf client models
$ swagger generate client -f /tmp/swagger.json -A go-netbox -t ~/src/github.com/smutel/go-netbox
```

### Trying to apply patchs

```sh
$ cd ~/go/src/github.com/smutel/go-netbox
$ for p in patchs/*; do patch -p0 < $p; done
```


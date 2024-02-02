
# Sec1 - Network and Website Penetration Testing Tools

## Overview

Sec1 is a comprehensive repository that provides a collection of tools for network and website penetration testing. Whether you are a security professional, a penetration tester, or a curious enthusiast, this repository aims to equip you with a variety of powerful and effective tools to assess and enhance the security of networks and websites.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Usage](#usage)

## Introduction

Penetration testing is a critical aspect of securing information systems. Sec1 is designed to be a one-stop repository where you can find a curated selection of tools specifically chosen for network and website penetration testing. The tools included cover a wide range of functionalities, from network reconnaissance and vulnerability scanning to web application testing and exploitation.

## Features

- **Diverse Toolset**: Sec1 includes a diverse set of tools, each serving a specific purpose in the penetration testing lifecycle.
- **Up-to-date**: The repository is regularly updated to include the latest versions of popular tools, ensuring that you have access to the most recent features and security enhancements.

## Usage
 bash 
 ```
 git clone https://github.com/hacker301et/sec1.git
 cd sec1
 go mod tidy
 go run main.go
```

|         | Command                                                            | Example                                                           | Usage                                                                                     |
|---------|--------------------------------------------------------------------|-------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| network | ```go run main.go net port-scan -a ip-address -w [workers]```      | ```go run main.go net port-scan -a 192.168.0.1```                 | Scan for open ports. Use the `-w` flag to specify the number of workers (default is 100). |
| network | ```go run main.go net my-public-ip```                              | ```go run main.go net my-public-ip```                             | Retrieve your public IP address.                                                          |
| network | ```go run main.go net change-mac -i interfaceName -m macAddress``` | ```go run main.go net change-mac -i wlan0 -m 00:11:22:33:44:55``` | Change the MAC address of the specified interface. The `-m` flag is optional.             |
| web     | ```go run main.go web live-sub```                                  | ```go run main.go web live-sub```                                 | Use the live-sub domain finder tool.                                                      |

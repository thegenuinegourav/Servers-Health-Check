# Servers-Health-Check 
This program is wriiten in golang to check health status of different servers using the concept of ```go routines``` & ```channel```.

## Introduction
Server health check provides a comprehensive overview of the status and performance of individual servers. It also 
includes the monitoring of hardware parameters, external conditions in the server room, and various performance data. 
With a server health check, you can prevent downtime by keeping an eye on server elements such as CPU and memory use, fans, 
power supplies and consumption, the temperature of various components, and many other important gauges.
By monitoring server statuses, you can prevent costly crashes and ensure the optimal performance of all your servers.

## How It Works?
This program will prompt user to input different server urls for whom you want to monitor the health and the time that defines 
how frequently you want to check the status of each server. 
This program then periodically sends requests to different serevrs concurrently to test their status and continuously output 
the result of each server after every 't' seconds.

## Concepts Used
This program is following the concepts of ```go routines``` using which we are achieving concurrency of different routines to hit 
servers urls parallely. This is also using ```channels``` to communicate bamong different ```go routines```.
 
## Screenshots

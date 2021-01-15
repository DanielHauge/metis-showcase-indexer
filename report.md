# 02148: Repository indexer

- Daniel F. Hauge <s201686@student.dtu.dk>
- Team: 24
- Jan. 22 - 2021

##### Abstract
Metis is the name of a bigger collection of systems that has been built mostly as a hobby. 
This project implements a system made to enable search features in other application in Metis, 
the resulting application is able to index and analyse files in selected repositories on services such as github. 
The application has been constructed with a goal of high scalability, achieved by an abundant amount of distribution and concurrency. 
Go's excellent concurrency support and "goSpaces" library has been utilized to provide the distribution and concurrency. 
The application works as follows.
The system can be thought of as 3 different applications, but is a single application which can be run in 3 different modes: Coordinator, Manager and Worker. 
The Coordinator mode's purpose is to coordinate other applications by providing tuple spaces, status tracking, logging, interface for control and triggering work flows.
The Manager mode's purpose is to manage work flows by delegating tasks to worker.
The Worker mode's purpose is to receive and execute tasks, primarily index code files to a search optimized database (elasticsearch).


<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>


# Content
1. Protocol
    1. textual description
    2. graphical description
    3. informal description
    4. implementation
2. Petri Net
    1. Visual
    2. Informal
    3. Transition system
    4. Implementation
3. Concurrent Algorithm
    1. Code
    2. tup4fun model
    3. testing
    4. transition system
4. Project Highlights
    1. Best actual feature
    2. Best potential feature
5. Additional Information
    

<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>

    
## 1. Protocol
I have chosen to showcase "Update-Status" protocol. The protocol is used by applications running in worker and manager mode, to enable status tracking feature. 


### Protocol: textual description
(Use description language in lecture 4)


### Protocol: graphical description
(Sequence diagram, ala lecture 4)


### Protocol: informal description
Quick description. (Per line in text description)


### Protocol: implementation
Use projection function descripted in lecture 4


<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>


## 2. Petri Net
The index work flow which an application running in manager mode is managing is chosen, as it can be illustrated neatly with a petri net. 

### Petri Net: visual description
Use https://apo.adrian-jagusch.de/#!/Sample%20Net
 
### Petri Net: informal description
Descripe the visual diagram. aka. (Legend, what is p0-p5 and task a, b, c)

### Petri Net: transition system
Look at report template.

- Marking explanation
Initial state for triggering.

- Transitiuon system
State machine ish, how the petri net proceed (with initial marking as initial state)

- Maximal paths
Properties, of the petri net.

- Infinite paths
Properties of forever looping


### Petri Net: implementation
- Provide code snippets for how this works.

<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>


## 3. Concurrent Algorithm
Use status tracking as primary example.
Use Coordinator as an honarable mention with go (Control server, Scheduling triggers, tuple space servers)
- Status tracking, task delegation and more for coordination.

### Concurrent Algorithm: code
code can be found in file:
snippets.


### Concurrent Algorithm: tup4fun
tup4fun model

### Concurrent Algorithm: testing
tup4fun testing. 

### Concurrent Algorithm: transition system
tup4fun transition system

<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>

## 4. Project Highlights

### Best actual feature
- Scalability, to provide more throughput, simply just spin up more managers and workers.

### Best potential feature
- The foundation to a performant scalable solution is layed, additional tasks, analysis, monitoring and more can be put into the system to provide even more value, 
without much risk of performance issues.


<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>

## 5. Additional information
- github
- execution showcase video.

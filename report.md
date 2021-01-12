# 02148: Repository indexer

- Daniel F. Hauge <s201686@student.dtu.dk>
- Jan. 3 - 2021

##### Abstract


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

- Trigger protocol
  - Trigger a new reindex -> 
  - Put work to next available analyser ->
  - Analyser picks up and delegates tasks to available workers.

Status callback.

### Protocol: textual description

### Protocol: graphical description

### Protocol: informal description

### Protocol: implementation



<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>


## 2. Petri Net

- Analyser Task Delegation.
  - Take work -> 
  - Delegate work (goroutines?) ->
  - Wait for all work is complete
  - Put into idle

### Petri Net: visual description
 
### Petri Net: informal description

### Petri Net: transition system

### Petri Net: implementation


<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>


## 3. Concurrent Algorithm

- Status tracking, task delegation and more for coordinator.

### Concurrent Algorithm: code

### Concurrent Algorithm: tup4fun

### Concurrent Algorithm: testing

### Concurrent Algorithm: transition system


<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>

## 4. Project Highlights

### Best actual feature

### Best potential feature

<div style="page-break-after: always; visibility: hidden"> \pagebreak </div>

## 5. Additional information
- github
- live site
- execution showcase video.

# beanstalkd Commands

### Producer Commands

##### `put`
    put <pri> <delay> <ttr> <bytes>\r\n
    <data>\r\n
Inserts a job into the tube currently being used by the client.

##### `use`
    use <tube>\r\n
Changes the tube currently being used by the client.



### Worker Commands

##### `reserve`
    reserve\r\n
Returns a newly-reserved job.

##### `reserve-with-timeout`
    reserve-with-timeout <seconds>\r\n
Returns a newly-reserved job, or fails with a timeout.

##### `delete`
    delete <id>\r\n
Removes the specified job from the server entirely.

##### `release`
    release <id> <pri> <delay>\r\n
Puts the specified reserved job back into the ready queue.

##### `bury`
    bury <id> <pri>\r\n
Puts the specified job into the "buried" state.

##### `touch`
    touch <id>\r\n
Requests more time to work on the specified job.

##### `watch`
    watch <tube>\r\n
Adds a tube to the list of tubes currently being watched by the client.

##### `ignore`
    ignore <tube>\r\n
Removes a tube from the list of tubes currently being watched by the client.



### Other Commands

##### `peek`
    peek <id>\r\n
Inspect the specified job.

##### `peek-ready`
    peek-ready\r\n
Inspect the next ready job in the tube currently being used by the client.

##### `peek-delayed`
    peek-delayed\r\n
Inspect the delayed job with the shortest delay left in the tube currently being used by the client.

##### `peek-buried`
    peek-buried\r\n
Inspect the next job in the list of buried jobs in the tube currently being used by the client.

##### `kick`
    kick <bound>\r\n
Moves buried or delayed jobs into the ready queue in the tube currently being used by the client.

##### `kick-job`
    kick-job <id>\r\n
Moves the specified job into the ready queue.

##### `stats-job`
    stats-job <id>\r\n
Gives statistical information about the specified job.

##### `stats-tube`
    stats-tube <tube>\r\n
Gives statistical information about the specified tube.

##### `stats`
    stats\r\n
Gives statistical information about the system as a whole.

##### `list-tubes`
    list-tubes\r\n
Returns a list of all existing tubes.

##### `list-tube-used`
    list-tube-used\r\n
Returns the tube currently being used by the client.

##### `list-tubes-watched`
    list-tubes-watched\r\n
Returns the list of tubes currently being watched by the client.

##### `quit`
    quit\r\n
Closes the connection.

##### `pause-tube <tube-name> <delay>`
    pause-tube <tube-name> <delay>\r\n
Delays any new job being reserved for a given time.

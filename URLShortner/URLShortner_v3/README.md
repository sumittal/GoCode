# URLShortner version 3

* ADDING GOROUTINES for performance 

There is still a performance problem with the 2 nd version if too many clients attempt to add URLs
simultaneously: our map is safely updated for concurrent access thanks to the locking mechanism,
but the immediate writing of each new record to disk is a bottleneck. The disk writes may happen
simultaneously and depending on the characteristics of your OS, this may cause corruption.
Even if the writes do not collide, each client must wait for their data to be written to disk before
their Put function will return. Therefore, on a heavily I/O-loaded system, clients will wait longer
than necessary for their Add requests to go through.

To remedy these issues, we must decouple the Put and save processes: we do this by using Go’s
concurrency mechanism. Instead of saving records directly to disk, we send them to a channel,
which is a kind of buffer, so the sending function doesn’t have to wait for it.

The save process which writes to disk reads from that channel and is started on a separate thread
by launching it as a goroutine called saveloop . The main program and saveloop are now executed
concurrently, so there is no more blocking.

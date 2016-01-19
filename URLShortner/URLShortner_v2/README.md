# URLShortner version 2

* ADDING PERSISTENT STORAGE using Go’s encoding/gob package

When the goto process (the webserver on port 8080) ends, and this has to happen sooner or later,
the map with the shortened URLs in memory will be lost. To preserve our map data, we need to
save them in a disk file. We will modify URLStore so that it writes its data to a file, and restores
that data on goto start-up.

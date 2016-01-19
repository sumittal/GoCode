# URLShortner version 4

* PERSISTENT STORAGE WITH JSON 

If you are a keen tester you will perhaps have noticed that when goto is started 2 times, the 2 nd
time it has the short urls and works perfectly, however from the 3 rd start onwards, we get the
error: Error loading URLStore: extra data in buffer . This is because gob is a stream based
protocol that doesn’t support restarting. We will remedy this situation here by using json as storage
protocol, which stores the data as plain text, so it can also be read by processes written
in other languages than Go. This also shows how easy it is to change to a different persistent
protocol, because the code dealing with the storage is cleanly isolated in 2 methods, namely load
and saveLoop .

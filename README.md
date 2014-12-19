go-protobuf-clientserver
========================

This is a simple client and server that uses protocol buffers frame delimeted by a big endian uint32. Mostly PoC right now, might evolve into something useful...

* client - connects to localhost 8080 and sends packet
* server - listens on 8080 and receives, prints, packets
* protocol - the common protocol
* protocoder - common functionality to decode/encode packets

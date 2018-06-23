# socket_multiplexer
Tunneling multiple protocols through the same port. 
(UDP and TCP)
Every request body received by the server is converted to upper case and sent back as a response to the client.

# Steps
## Server:
- cd server/
- go build
- ./server host:port

## Client:
- cd client/
- go build
- ./client host:port



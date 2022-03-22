# ratelimit

Simple proof of concept for rate limiting the total number of calls made by multiple clients to a server.

Run with `docker compose up` and watch console output; "Rejected" messages are logged by the client to show when they were prevented from sending a request by the rate limiting, and "Request received" messages are logged by the server when called by a client.

NOTE: Depending on exactly when the containers are spun up, you may only see Rejected messages from _one_ of the clients. If you wait a while you'll eventually see them from the other client too, but if it's taking too long then you can just try quitting with Ctrl+C and restarting with `docker compose up` again.

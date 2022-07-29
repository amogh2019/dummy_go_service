# dummy_go_service
Dummy Go Microservice  // Go // gRPC // protobuf

### what are microservices and challenges?
- microservice // service for each business function
- buy // discount // delivery // user
    - each service can be written in diff language
    - they need to talk to one another
        - How do they talk to one another? Each service Agrees on a lot of things to build that API
        - What is the Checklist to build API? 
            - functional
                - whats the contract of the API
                - whats the data format?
                - error handling schemes and approches
                - ...
            - operational
                - what size of payload are we planning to handle?
                - latency of the response
                - how much concurrent requests can be handled?
                - loadbalancing?
                - fallbacks?
                - authorization
                - logging
                - monitoring
                - language interoperatablity

- why gRPC?
    - building api is hard, as we saw above, we need to address all those issues
    - Can we not just focus on data? and let the framework handle those concerns/charters
        - gRPC is that framework for interservice communication / service client comm // basically communication with our service
        - build over http2
        - open source // cloud native computation foundation like docker kuberenetes
        - mainly focus on low latency
        - supports streaming of data

- why http2 is better?
    - scene
        - consider a save where a webpage has to load // it has three resources to get // image css script
        - usually // browser will make three http calls to get these three things 
        - each call opens its own TCP connection
        - this also sends plaintext headers

    - features
        - http2 allows TCP connection reusing // long lasting tcp connections
        - there is server push 
        - binary headers (defined structure in packet frame, hence doesnt have to parse the header to determine the structure)

    - pros of the features
        - because of server push and long tcp connection, now we can make just one request and server can respond back with three messages in the same tcp connnection
        - this also decreases the no. of requests the client makes  // and so the no of tcp conn // hence less traffic




- 
        


        


    


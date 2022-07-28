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
    - building api is hard, as we saw above, we need to address all those issue
    - Can we not just focus on data? and let the framework handle those concerns/charters
        - gRPC is that framework for interservice communication
        - build over http2
        - open source // cloud native computation foundation like docker kuberenetes
        - mainly focus on low latency
        - supports streaming of data

        


    


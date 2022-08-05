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




- types of gRPC API
    - unary // single request message // single response message
    - server streaming // single request message // multiple response message
    - client streaming // multiple request message // single response message
    - bidirectional // multiple request message // multiple response message
        -   Example 
        ```
        service GreetService {
            //UNARY
            rpc Greet(GreetRequest) return (GreetResponse){};
             //SERVER STREAMING
            rpc GreetManyTimes(GreetRequest) return (stream GreetResponse){};
             //CLIENT STREAMING
            rpc LongGreet(stream GreetRequest) return (GreetResponse){};
             //BIDIRECTIONAL STREAMING
            rpc GreetEveryone(stream GreetRequest) return (stream GreetResponse){};
        }
    - 

- gRPC vs REST

|   grpc    |   rest   |
|-----------|----------|
|protocal buffers(strict schema(types and fields))| JSON free schema |
|   http2|  http1|
|   streaming |   unary |
|   bidirectional |   client -> server |
|   free design| GET POST ... etc actions|


- protobuf vs json
    - both are data schemes // data schema // ways we use to represent data // to send in  communication
    - json has like it own structure // send to any srver(on any language) // it will parse and read the message
    - payload is large since we are sending structure everytime
    - everyone's parsing // serializing and deserialzing // may not be optimal
    - hence comes protobuf
        - we fix a schema and tell to all clients // fix fields numbering(tags) // fix their types
        - along with structure // we also provide some sdks/libraries some api/code by ourserlves // for them to serialize and deserialize // they the clients dont have to worry about this
        - hence the message that we send is small in size // also not humanreadable, unless you decode it yourself
        - dont worry about changing schemes / changing fields  // backward compatbality is also handled by the api that it provides
        - since we have schema // that we share with clients // we can also write comments/documentation in the schemas that we share // unlike in json where documentation is external (and so can be obsolete)



### how to run?

1. make // this will build the go proto buf files from their proto definitions (if not already present)

2. run the respective service
    - running calculator
        - run server 
            `go run calculator/server/*.go`
        - run client
            `go run calculator/client/*.go`
    - running greet
        - run server 
            `go run greet/server/*.go`
        - run client
            `go run greet/client/*.go`
        

### is there a swagger to know api definitions on server?
an http server doesnt expose the definitions by itself, and so we need swagger to reflect the same.
similarly in go, we can reflect the grpc services, using reflection

### what about http clients like postman to check server endpoints?
try using grpc cli clients // to get service definitions // to make grpc calls to server
    - evans cli // evans --host hostname --port --reflection repl   ("read–evaluate–print loop")


### notes
1. Reference files taken from : https://github.com/Clement-Jean/grpc-go-course



        


    


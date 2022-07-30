## About Proto

### What are we doing with protobuf?
here we are using proto to define out data schema for communication

-  we will define out entites  // and the contract for the api // 
-  we will use protoc // protocompiler to generate GO's representation for the proto contract/entities
-  then we will test the go files generated // that its actually working in a go program



### What do we write in a proto file?

- write syntax to mention what proto version we are using
- write the packageing scheme for our proto contracts (refer the practice folder for the scheme that you can follow to define multiple message contracts)
- mention the go package you want the generated go classes to have (so that they can be imported in go program with ease)
    -   use the options flag go_package variable to let the protoc compiler know that /// thats the go pacakge value to be embedded in generate go files

- define the messages (entities/classes types for go program)
- define the service (apis for grpc)



### How to make protoc compile proto contracts and create files/models/services that we can use in our app

command goes like this
```
protoc -Igreet/proto --go_opt=module=github.com/amogh2019/dummy_go_service --go_out=. greet/proto/*.proto
```

- we run this at directory root
- -I tells compiler where to find the proto files to generate
- go_opt tells that we have some custom flags to honor while generating go code
- go_opt_module tells that which go module to place generated files in
- go_out tells the relative path to the go module to place the generated files in
- {greet/proto/*.proto} would be the specific files to compile





### Some more protoc compiler feature

- Since serialized proto is binary and not human read friendly, we can ask the compiler to decode it into tag:value wala human friendly readable print
    -   `cat ./greet/server/practice/messageInSerializedForm.bin | protoc --decode_raw`
    ```
    1: 20
    2: 1
    3: "Dummy21"
    4: "hoyee"
    4: "dumm"
    4: "sdfsadf"
    ```

- we can also ask the compiler to read the binary and parse into a specific proto
    - uss proto file ke uss message mein parse karo
    - standard input read karna hota hai usse, issliye example ke liye pipe karke input diya
    - `cat ./greet/server/practice/messageInSerializedForm.bin | protoc --decode=greet.practice.level3.SampleMessage greet/proto/practice/level3/sample1.proto `
    - 
    ```
    id: 20
    is_active: true
    name: "Dummy21"
    alt_names: "hoyee"
    alt_names: "dumm"
    alt_names: "sdfsadf"
    ```

- we can also encode the proto human friendly readable format into proto binary format
    - `cat ./greet/server/practice/messageInProtoHumanFriendly.txt | protoc --encode=greet.practice.level3.SampleMessage greet/proto/practice/level3/sample1.proto > temp.bin`



### About updating proto messages

- thoda dekh ke karna // sirf name changes mein no worries
- type change for a field mein issues ho skte // backward compatablity mein // pura data might break while deserializtion
- field ka tag change can also mess things up // since tag is the identifier in the actual binary form // fieldname is just a label
- tags update kar rahe toh future incremental tag dedo // aur purane ko block kardo
    - use reserved keyword // block tag and fieldname both



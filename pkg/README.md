# pkg

## Reading this Project

If the folder name has a suffix of `service`, then it is a GRPC service.
If the folder name has a suffix of `api`, then it is a REST api meant to be an entry point from possible web clients in using the GRPC services.

### GRPC services

GRPC services contain a main folder, setting up the GRPC server and initializing any other dependencies the folder may rely on. The rest of the code in the folder is the implementation of the service itself as defined by the corresponding protobuf file. Typically, all the interface definitions are found in the file with the same name as the folder itself without the suffix. (eg. color_service -> color.go)

Other files in the folder consist of the helper functions and necessary logic to implement the interface.

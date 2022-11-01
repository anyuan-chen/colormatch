# protos

This project has a schema-driven API. The IDL (protobuf) here serves as the universal source of truth for the application.

This folder contains the protobuf defnitions for all of the types and services in this repository. For every GRPC service, there is a corresponding folder except with the suffix changed from `_service` to `_protobufs` (eg. color_service to color_protobufs).

The shared_protobufs folder is special, containing all the shared data types used across the entire project.

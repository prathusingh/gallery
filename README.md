# Gallery

- <b> Gallery creates a personalized portfolio of photographers where they can showcase their work.</b>
- <b> A photographer can also search other photographer's work too. </b>
- <b> A customer can browse photographers and can interact with their work. </b>

#### Tech Stack

- Web Client uses React, GraphQl, Relay, Webpack
- Backend uses Go microservices
- Uses GRPC and CQRS driven Kafka based intercommunication mechanism
- Uses Postgres SQL and GCP.

##### Microservices

- gateway - endpoint for web client to interact with backend. communicates with user service using grpc.
- user - handles authentication, authorization, personalized url generation, subscribes to booking event. communicates with media and discovery service using grpc.
- media - handles images processing, loading, fetching from google cloud storage.
- discovery - handles search of new photographers as well as potential clients. supports browing by new cusotmers or photographers.
- booking - handles photographer book. communicates asynchronously with user service using event driven mechanism

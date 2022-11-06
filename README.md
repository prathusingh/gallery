# Gallery

- <b> Gallery creates a personalized portfolio of photographers where they can showcase their work.</b>
- <b> A photographer can have its own portal. In a portal, the photographer can showcase photos.
- <b> Customer can leave review on the photographer.
- <b> Customer can contact photgrapher by email.
- <b> Notification will let photographer to know that someone has reviewed or contacted them. 
- <b> A photographer can also search other photographer's work too. </b>

#### Tech Stack

- Web Client uses React, GraphQl, Relay, Webpack
- Backend uses Go microservices
- Uses GRPC and CQRS driven Kafka based intercommunication mechanism
- Uses Postgres SQL and GCP for data storage

##### Microservices

- gateway - endpoint for web client to interact with backend. communicates with user service using grpc.
- user - handles authentication, authorization, personalized url generation, subscribes to booking event. communicates with media and discovery service using grpc.
- media - handles images processing, loading, fetching from google cloud storage.
- discovery - handles search of new photographers as well as potential clients. supports browing by new cusotmers or photographers.
- notification - handles notifcation regarding booking, comments and upvotes. communicates asynchronously with user service using event driven mechanism.
- message - enables email communication.
- review - enables reviewing of photographer's work.

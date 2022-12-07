# Gallery

- <b> Gallery creates a personalized portfolio of photographers where they can showcase their work.</b>
- <b> There will be a write portal for the photographer and read portal for the user.
- <b> Customer can leave reviews on the photographer.
- <b> Customer can contact photgrapher by email.
- <b> Notification will let photographer to know that someone has reviewed or contacted them.
- <b> A photographer can also search other photographer's work too. </b>

#### Tech Stack

- Web Client uses React, GraphQl, Relay, Webpack.
- Backend uses Go microservices
- Uses GRPC and CQRS driven Kafka based intercommunication mechanism.
- Uses Postgres SQL and GCP for data storage.

##### Microservices

- gateway - endpoint for web client to interact with other microservices. Receives http request from
  the webclient.
- user - handles authentication, authorization, personalized url generation and review generation.
  stores all the information in PostgreSQL db. Commands logged in user in the event store
- media - handles images processing, loading, fetching from google cloud storage. stores meta infor
  mation inside media PostgreSQL db. Reads photographs based on the logged in user.
  Commands new photographs for the photographer.
- search - handles search of new photographs by the phtographs type. Read only view from the elastic
  store.
- notification - subscribe to new reviews and messages. read only view.
- message - enables email communication. Commands new message inside PostgreSQL db.

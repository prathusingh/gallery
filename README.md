## WINK SASSY

### Personalized blog for content creators

Home page consists of
1. Filter
2. Tiles showing relevant image

Description Page consists of
1. Rotational panel rotating the image.
2. Rotational panel rotating the description.

Microservices:
1. Image Loader that loads the image in the rotational panel.
2. Description Loader that loads the image in the rotational panel.
3. Query Loader that loads the default homepage or apply the filtered results to the top of page.
4. Profile Loader loads the profile of the content creator.

Tech specs:
1. Microservices interact with each other using gRPC.
2. Go kit framework is used for microservices.
3. Each microservice is containerized using docker.
4. Web and mobile client interact with microservices using GraphQL API gateway.
5. gqlgen is used for GraphQL.
6. Front end services are written using TypeScript.

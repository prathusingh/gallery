## WINK SASSY

### Personalized blog for content creators

Home page consists of

1. Filter/Search
2. Tiles showing relevant content
3. Profile icon button

Description Page consists of

1. Rotational panel rotating the image.
2. Rotational panel rotating the description.

Microservices:

1. "image" that loads the image in the rotational panel.
2. "description" that loads the description of the image in the rotational panel.
3. "content" that loads the default homepage or apply the searched results to the top of page.
4. "profile" loads the profile of the content creator.
5. "content" talks with "image" and "description".

Tech specs:

1. Microservices interact with each other using gRPC.
2. Go kit framework is used for microservices.
3. Each microservice is containerized using docker.
4. Web and mobile client interact with microservices using GraphQL API gateway.
5. gqlgen is used for GraphQL.
6. Front end services are written using TypeScript.

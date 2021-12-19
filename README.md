## WINK SASSY

### Personalized blog for instagrammers

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

Communication:
1. Microservices interact with each other using gRPC.
2. Web and mobile client interact with microservices using GraphQL API gateway.


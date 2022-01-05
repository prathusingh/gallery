## WINK SASSY

### Personalized blog for content creators

Client

Home page consists of

1. Filter/Search
2. Tiles shows relevant content
3. Profile icon button

Description Page consists of
1. Left Card shows image(s) in a rotational panel
2. Right Card shows description of the corresponding image.
3. User can share the image or the description or both either to other platforms (whatsapp, email or copy link)
4. If user is logged in then the number of likes and the comments appeared on the image/description will appear

Profile page consists of
1. Username if logged in
2. Total number of followers
3. Meta information of the content

Authorization page consists of
1. Google authentication
2. Instagram authentication
3. Pinterest authentication


Microservices:
1. Content service provides grpc endpoint for the content. A content includes image and text conveying information about the image.
2. Authentication service provides grpc endpoint for the User authentication. Authentication is done through Gmail, Instagram or Pinterest. No separate login.
3. Profile service provides the grpc endpoint for the User profile.
A profile consists of the user name, user image otherwise avatar, meta information of the content
4. Share service provides the grpc endpoint for sending the content across different platforms either email or whatsapp or copy link
5. Feed service provides grpc endpoint to display the content feed
Tech specs:

1. Microservices interact with each other using gRPC.
2. Microservice is containerized using docker.
3. Web and mobile client interact with microservices using GraphQL API gateway. Used gqlgen is used for GraphQL.
4. Front end services are written using TypeScript.
5. Currently web client is only supported.

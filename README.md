# User Service

The User Service section of Messenger is a crucial component responsible for managing user-related functionalities and interactions within the messaging platform. This service leverages master-slave replication and gRPC to ensure reliable and efficient performance.

## Features

The User Service section provides the following key features:

1. User Creation: Allow the creation of a new user profile by providing relevant information such as name, email, and password. Upon successful creation, assign a unique user ID to the user.

2. User ID Assignment and Retrieval: Generate and assign a unique user ID to each user upon registration. Provide functionality to retrieve the user ID associated with a specific user profile.

3. Friend Creation: Enable users to establish connections by sending friend requests to other users. Upon acceptance, create a friendship association between the two users. Include functionality to provide relevant details such as friend's name, email, or ID.

4. Friend ID Assignment and Retrieval: Assign a unique friend ID to each friend connection within the Messenger platform. Allow users to retrieve the friend ID associated with a particular friend.

5. User-to-Friend Association: Provide an option to associate a user with a friend connection. This association allows users to assign themselves to specific friends, indicating their connection or relationship. This information can be used for personalized features or recommendations within the messaging platform.

## Architecture

The User Service employs a robust architecture consisting of master-slave replication and gRPC communication protocol.

### Master-Slave Replication

The master-slave replication mechanism ensures high availability and fault tolerance for the User Service. It involves maintaining multiple instances of the service where one acts as the master while others serve as slaves. The master instance handles all write operations and synchronizes the data with slave instances in near real-time. In the event of a failure, one of the slave instances can be promoted to a master, ensuring uninterrupted service.

### gRPC Communication

The User Service leverages gRPC (Google Remote Procedure Call) as the communication protocol. gRPC is a high-performance, open-source framework that facilitates efficient and cross-platform communication between client and server applications. It enables the User Service to handle remote procedure calls, making it easier to implement inter-service communication, such as authentication, authorization, and data retrieval, with other components of the Messenger platform.

## Installation and Deployment

To deploy the User Service section of Messenger, follow the instructions below:

1. Clone the Messenger repository from the designated repository location.

2. Install the necessary dependencies specified in the project's requirements file.

3. Configure the database connection settings and ensure the required database schema is created.

4. Set up the necessary gRPC dependencies and ensure they are properly configured.

5. Start the master and slave instances of the User Service using the provided scripts or commands.

6. Monitor the service logs and ensure successful connectivity with other components of the Messenger platform.

For more detailed deployment instructions and troubleshooting, refer to the deployment documentation available in the Messenger repository.

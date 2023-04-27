package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// and "time"

const (
	masterNodeAddress = "localhost:50051" // Address of the master node gRPC server
)

// User represents a user in the system
type User struct {
	ID   int
	Name string
}

// Friend represents a friend of a user
type Friend struct {
	ID   int
	Name string
}

// UserServiceClient represents the user service client
type UserServiceClient struct {
	client UserServiceClient
}

// CreateUser creates a new user
func (c *UserServiceClient) CreateUser(name string) (int, error) {
	req := &CreateUserRequest{Name: name}
	res, err := c.client.CreateUser(context.Background(), req)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}

	return res.ID, nil
}

// GetUser returns the user details for a given user ID
func (c *UserServiceClient) GetUser(userID int) (*User, error) {
	req := &GetUserRequest{ID: userID}
	res, err := c.client.GetUser(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	user := &User{
		ID:   userID,
		Name: res.Name,
	}

	return user, nil
}

// CreateFriend creates a new friend
func (c *UserServiceClient) CreateFriend(name string) (int, error) {
	req := &CreateFriendRequest{Name: name}
	res, err := c.client.CreateFriend(context.Background(), req)
	if err != nil {
		return 0, fmt.Errorf("failed to create friend: %v", err)
	}

	return res.ID, nil
}

// GetFriend returns the friend details for a given friend ID
func (c *UserServiceClient) GetFriend(friendID int) (*Friend, error) {
	req := &GetFriendRequest{ID: friendID}
	res, err := c.client.GetFriend(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get friend: %v", err)
	}

	friend := &Friend{
		ID:   friendID,
		Name: res.Name,
	}

	return friend, nil
}

// AssignUserToFriend assigns a user to a friend
func (c *UserServiceClient) AssignUserToFriend(userID, friendID int) error {
	req := &AssignUserToFriendRequest{
		UserID:   userID,
		FriendID: friendID,
	}

	_, err := c.client.AssignUserToFriend(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to assign user to friend: %v", err)
	}

	return nil
}

func main() {
	// Create a gRPC connection to the master node
	conn, err := grpc.Dial(masterNodeAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the master node: %v", err)
	}
	defer conn.Close()

	// Create a user service client
	client := &UserServiceClient{
		client: NewUserServiceClient(conn),
	}

	// Example usage of the client functions
	userID, err := client.CreateUser("John Doe")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	user, err := client.GetUser(userID)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)

	friendID, err := client.CreateFriend("Jane Smith")
	if err != nil {
		log.Fatalf("Failed to create friend: %v", err)
	}

	friend, err := client.GetFriend(friendID)
	if err != nil {
		log.Fatalf("Failed to get friend: %v", err)
	}

	fmt.Printf("Friend ID: %d, Name: %s\n", friend.ID, friend.Name)

	err = client.AssignUserToFriend(userID, friendID)
	if err != nil {
		log.Fatalf("Failed to assign user to friend: %v", err)
	}

	userFriend, err := client.GetUserFriend(userID)
	if err != nil {
		log.Fatalf("Failed to get user's friend: %v", err)
	}

	fmt.Printf("User's Friend: %s\n", userFriend.Name)
}

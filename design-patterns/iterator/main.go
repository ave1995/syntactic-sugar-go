package main

import "fmt"

// --- 1. Define the Item Type ---

type User struct {
	Name string
	ID   int
}

// --- 2. The Iterator Interface ---

// UserIterator defines the traversal contract.
type UserIterator interface {
	HasNext() bool
	Next() *User
}

// --- 3. The Aggregate/Collection Interface ---

// UserAggregate defines the collection contract.
type UserAggregate interface {
	CreateIterator() UserIterator
}

// --- 4. The Concrete Collection (Aggregate) ---

// UserCollection is the complex data structure we want to iterate over.
type UserCollection struct {
	users []*User // This could be a map, linked list, or anything else internally
}

// CreateIterator provides the standardized way to get an iterator for this collection.
func (uc *UserCollection) CreateIterator() UserIterator {
	return &concreteUserIterator{
		collection: uc,
		index:      0, // Start at the beginning
	}
}

// --- 5. The Concrete Iterator ---

// concreteUserIterator holds the state of the traversal.
type concreteUserIterator struct {
	collection *UserCollection
	index      int // Tracks the current position in the collection
}

// HasNext checks if the index is within the bounds of the collection.
func (i *concreteUserIterator) HasNext() bool {
	return i.index < len(i.collection.users)
}

// Next returns the current element and advances the index.
func (i *concreteUserIterator) Next() *User {
	if !i.HasNext() {
		return nil // Should handle this case gracefully, but nil is simple for now
	}

	user := i.collection.users[i.index]
	i.index++ // Move to the next element
	return user
}

// --- 6. Client Code (Demonstration) ---

func main() {
	// Create the collection
	collection := &UserCollection{
		users: []*User{
			{Name: "Alice", ID: 101},
			{Name: "Bob", ID: 102},
			{Name: "Charlie", ID: 103},
			{Name: "Diana", ID: 104},
		},
	}

	// Get the iterator from the collection using the Aggregate interface
	iterator := collection.CreateIterator()

	fmt.Println("Starting Iteration:")

	// The client code only interacts with the Iterator interface (HasNext, Next).
	// It doesn't know that the UserCollection uses a slice internally.
	for iterator.HasNext() {
		user := iterator.Next()
		fmt.Printf("Processing User ID: %d, Name: %s\n", user.ID, user.Name)
	}

	fmt.Println("\nIteration Complete.")
}

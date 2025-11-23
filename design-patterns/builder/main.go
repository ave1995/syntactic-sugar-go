package main

import "fmt"

// --- 1. The Target Struct (The Complex Object) ---

// User represents a complex object with required and optional fields.
// All fields are unexported to ensure the struct is immutable after being built.
type User struct {
	firstName  string
	lastName   string
	email      string
	phone      string
	isVerified bool
	role       string
}

// Public accessors for the immutable fields (best practice)
func (u *User) String() string {
	return fmt.Sprintf("User: %s %s | Email: %s | Phone: %s | Verified: %t | Role: %s",
		u.firstName, u.lastName, u.email, u.phone, u.isVerified, u.role)
}

// --- 2. The Builder Struct ---

// UserBuilder holds the temporary state needed to construct the User.
type UserBuilder struct {
	// These fields map directly to the User fields
	user User

	// Flags to track required fields or setup defaults
	hasFirstName bool
	hasLastName  bool
}

// NewUserBuilder creates and returns a pointer to a new UserBuilder.
// This is the starting point for the fluent interface.
func NewUserBuilder(firstName, lastName string) *UserBuilder {
	// Initialize with required fields and defaults
	return &UserBuilder{
		user: User{
			firstName:  firstName,
			lastName:   lastName,
			role:       "guest", // Set a default role
			isVerified: false,
		},
		hasFirstName: true,
		hasLastName:  true,
	}
}

// --- 3. Optional Setter Methods (Method Chaining) ---

// WithEmail sets the user's email address. It returns the builder for chaining.
func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.email = email
	return b
}

// WithPhone sets the user's phone number.
func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.phone = phone
	return b
}

// Verified sets the user's verification status.
func (b *UserBuilder) Verified(verified bool) *UserBuilder {
	b.user.isVerified = verified
	return b
}

// WithRole sets the user's role.
func (b *UserBuilder) WithRole(role string) *UserBuilder {
	b.user.role = role
	return b
}

// --- 4. The Finalizer Method ---

// Build validates and constructs the final User struct.
func (b *UserBuilder) Build() (*User, error) {
	// Perform validation checks before finalizing the object
	if !b.hasFirstName || !b.hasLastName {
		return nil, fmt.Errorf("required fields (first name and last name) are missing")
	}

	// Optional: Check if email is required if role is 'admin'
	if b.user.role == "admin" && b.user.email == "" {
		return nil, fmt.Errorf("admin user must have an email address")
	}

	// Return a copy of the constructed user struct
	return &b.user, nil
}

func main() {
	// --- Example 1: Full Builder Usage (Fluent Interface) ---
	fmt.Println("--- Building User 1 (Admin) ---")
	user1, err := NewUserBuilder("Alice", "Smith").
		WithEmail("alice@company.com").
		WithPhone("555-1234").
		WithRole("admin").
		Verified(true).
		Build()

	if err != nil {
		fmt.Println("Error building user 1:", err)
	} else {
		fmt.Println(user1)
	}

	// --- Example 2: Minimal Usage (Required fields only) ---
	fmt.Println("\n--- Building User 2 (Guest) ---")
	user2, err := NewUserBuilder("Bob", "Jones").
		Build() // All other fields use default values (or zero values)

	if err != nil {
		fmt.Println("Error building user 2:", err)
	} else {
		fmt.Println(user2)
	}

	// --- Example 3: Validation Failure ---
	fmt.Println("\n--- Building User 3 (Validation Failure) ---")
	// Attempt to build an admin user without an email (fails validation in Build())
	user3, err := NewUserBuilder("Charlie", "Day").
		WithRole("admin"). // Set role to admin
		Build()

	if err != nil {
		fmt.Println("Validation Error:", err)
		if user3 != nil {
			fmt.Println("Error: User 3 should be nil here.")
		}
	}
}

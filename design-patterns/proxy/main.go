package main

import "fmt"

// --- 1. The Subject Interface ---

// StorageService defines the common interface for both the Real Subject and the Proxy.
type StorageService interface {
	DeleteFile(filename string, userID string) error
	ReadFile(filename string) (string, error)
}

// --- 2. The Real Subject ---

// RealStorageService is the actual, expensive-to-access service.
type RealStorageService struct{}

func (s *RealStorageService) DeleteFile(filename string, userID string) error {
	// In a real application, this would delete the file on disk or in the cloud.
	fmt.Printf("[Real Service] DELETING file '%s' for user %s...\n", filename, userID)
	return nil
}

func (s *RealStorageService) ReadFile(filename string) (string, error) {
	fmt.Printf("[Real Service] READING file '%s'...\n", filename)
	return fmt.Sprintf("Content of %s", filename), nil
}

// --- 3. The Proxy ---

// ProtectionProxy provides access control for the DeleteFile method.
type ProtectionProxy struct {
	realService *RealStorageService
}

// NewProtectionProxy creates a new proxy instance.
func NewProtectionProxy(service *RealStorageService) *ProtectionProxy {
	return &ProtectionProxy{realService: service}
}

// DeleteFile implements the StorageService interface.
// It includes protection logic before calling the real service.
func (p *ProtectionProxy) DeleteFile(filename string, userID string) error {
	// Protection Logic: Only allow deletion if the user is "admin"
	if userID != "admin" {
		return fmt.Errorf("[Proxy Check] Access Denied: User %s does not have admin privileges to delete files", userID)
	}

	// If access is granted, forward the request to the real service
	fmt.Println("[Proxy Check] Access Granted. Forwarding request to Real Service.")
	return p.realService.DeleteFile(filename, userID)
}

// ReadFile implements the StorageService interface.
// Reading is allowed for everyone, so the request is passed through immediately.
func (p *ProtectionProxy) ReadFile(filename string) (string, error) {
	return p.realService.ReadFile(filename)
}

func main() {
	// Create the Real Service instance
	real := &RealStorageService{}

	// Create the Proxy, giving it a reference to the Real Service
	proxy := NewProtectionProxy(real)

	fmt.Println("--- User 'guest' attempts to DELETE ---")
	err := proxy.DeleteFile("secret_data.txt", "guest")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- User 'admin' attempts to DELETE ---")
	err = proxy.DeleteFile("secret_data.txt", "admin")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- Any user attempts to READ ---")
	content, _ := proxy.ReadFile("public_log.txt")
	fmt.Println("Read Content:", content)
}

// Package auth is a lightweight credential store.
// It provides functionality for loading credentials, as well as validating credentials
package auth

import (
	"encoding/json"
	"io"
	"os"
)

const (
	// AllUsers is the username that indicates all users, even anonymous users (request without
	// any BasicAuth information).
	AllUsers = "*"

	// PermAll means all actions permitted.
	PermAll = "all"
	// PermJoin means user is permitted to join cluster.
	PermJoin = "join"
	// PermJoinReadOnly means user is permitted to join the cluster only as a read-only node
	PermJoinReadOnly = "join-read-only"
	// PermRemove means user is permitted to remove a node.
	PermRemove = "remove"
	// PermExecute means user can access execute endpoint.
	PermExecute = "execute"
	// PermQuery means user can access query endpoint
	PermQuery = "query"
)

// BasicAuther is the interface an object must support to return basic auth information.
type BasicAuther interface {
	BasicAuth() (string, string, bool)
}

// Credential represents authentication and authorization configuration for a single user.
type Credential struct {
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
	Perms    []string `json:"perms,omitempty"`
}

// CredentialsStore stores authentication and authorization information for all users.
type CredentialsStore struct {
	store map[string]string
	perms map[string]map[string]bool
}

// NewCredentialsStore returns a new instance of a CredentialStore
func NewCredentialsStore() *CredentialsStore {
	return &CredentialsStore{
		store: make(map[string]string),
		perms: make(map[string]map[string]bool),
	}
}

// NewCredentialsStoreFromFile returns a new instance of CredentialStore loaded from a file.
func NewCredentialsStoreFromFile(path string) (*CredentialsStore, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c := NewCredentialsStore()
	return c, c.Load(f)
}

// Load loads credential information from a reader.
func (c *CredentialsStore) Load(r io.Reader) error {
	dec := json.NewDecoder(r)
	// Read open bracket
	_, err := dec.Token()
	if err != nil {
		return err
	}

	var cred Credential
	for dec.More() {
		err := dec.Decode(&cred)
		if err != nil {
			return err
		}
		c.store[cred.Username] = cred.Password
		c.perms[cred.Username] = make(map[string]bool, len(cred.Perms))
		for _, p := range cred.Perms {
			c.perms[cred.Username][p] = true
		}
	}

	// Read closing bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	return nil
}

// Check returns turn if the password is correct for the given username.
func (c *CredentialsStore) Check(username, password string) bool {
	pw, ok := c.store[username]
	return ok && pw == password
}

// Password returns the password for the given user.
func (c *CredentialsStore) Password(username string) (string, bool) {
	pw, ok := c.store[username]
	return pw, ok
}

// CheckRequest returns true if b contains a valid username and password
func (c *CredentialsStore) CheckRequest(b BasicAuther) bool {
	username, password, ok := b.BasicAuth()
	if !ok || !c.Check(username, password) {
		return false
	}
	return true
}

// HasPerm returns true if username has the given perm, either 

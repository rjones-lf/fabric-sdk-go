/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package apifabclient

import (
	"errors"

	"github.com/hyperledger/fabric-sdk-go/api/apicryptosuite"
)

// User represents users that have been enrolled and represented by
// an enrollment certificate (ECert) and a signing key. The ECert must have
// been signed by one of the CAs the blockchain network has been configured to trust.
// An enrolled user (having a signing key and ECert) can conduct chaincode deployments,
// transactions and queries with the Chain.
//
// User ECerts can be obtained from a CA beforehand as part of deploying the application,
// or it can be obtained from the optional Fabric COP service via its enrollment process.
//
// Sometimes User identities are confused with Peer identities. User identities represent
// signing capability because it has access to the private key, while Peer identities in
// the context of the application/SDK only has the certificate for verifying signatures.
// An application cannot use the Peer identity to sign things because the application doesn’t
// have access to the Peer identity’s private key.
type User interface {
	IdentityContext

	Name() string
	EnrollmentCertificate() []byte
	Roles() []string
}

// IdentityContext supplies the serialized identity and key reference.
//
// TODO - refactor SigningIdentity and this interface.
type IdentityContext interface {
	MspID() string
	Identity() ([]byte, error)
	PrivateKey() apicryptosuite.Key
}

var (
	// ErrUserNotFound indicates the user was not found
	ErrUserNotFound = errors.New("user not found")
)

// UserKey is a lookup key in UserStore
type UserKey struct {
	MspID string
	Name  string
}

// UserStore is responsible for User persistence
type UserStore interface {
	Store(User) error
	Load(UserKey) (User, error)
}

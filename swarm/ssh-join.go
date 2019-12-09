package swarm

import (
	"github.com/marcboudreau/gaslight"
	"golang.org/x/crypto/ssh"
)

// SSHKeyStrategy is an interface used to abstract away how the SSHJoinService
// obtains the ssh.Signer object used as the ssh.AuthMethod for the connection.
type SSHKeyStrategy interface {
	Signer() (ssh.Signer, error)
}

type VaultSignedKeyStrategy struct {
	SSHKeyStrategy

	Address string
	
}

// SSHJoinService is an implementation that uses an SSH connection to remotely
// execute the docker swarm join-token command in order to obtain the join
// token.
type SSHJoinService struct {
	gaslight.SwarmJoinService

	Username string
}

// NewSSHJoinService creates a new instance of the SSHJoinService.
func NewSSHJoinService() *SSHJoinService {
	return &SSHJoinService{
		Username: "swarm",
	}
}

// Join attempts to
func (s *SSHJoinService) Join(instance *gaslight.Instance) error {
	signer := 

	config := ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(),
		},
	}
}

func (s *SSHJoinService) signKey(keypair ssh.Signer) (*ssh.Certificate, error) {

}
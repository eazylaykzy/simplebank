package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)

	testTable := []struct {
		name          string
		password      string
		wrongPassword string
	}{
		{
			name:     "Hash Password",
			password: password,
		},
		{
			name:     "Compare Right Password",
			password: password,
		},
		{
			name:          "Compare Wrong Password",
			password:      password,
			wrongPassword: RandomString(6),
		},
	}

	for i := range testTable {
		tt := testTable[i]

		t.Run(tt.name, func(t *testing.T) {
			hashedPassword1, err := HashPassword(password)
			require.NoError(t, err)
			require.NotEmpty(t, hashedPassword1)

			err = CheckPassword(password, hashedPassword1)
			require.NoError(t, err)

			hashedPassword2, err := HashPassword(tt.password)
			require.NoError(t, err)
			require.NotEmpty(t, hashedPassword2)

			err = CheckPassword(tt.wrongPassword, hashedPassword2)
			require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
			require.NotEqual(t, hashedPassword1, hashedPassword2)
		})
	}
}

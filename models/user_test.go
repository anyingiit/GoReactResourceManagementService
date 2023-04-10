package models_test

import (
	"testing"
	"time"

	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser_SetPassword(t *testing.T) {
	user := &models.User{}

	err := user.SetPassword("password")
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Salt)
}

func TestUser_SetPassword_InvalidPassword(t *testing.T) {
	user := &models.User{}

	err := user.SetPassword("pass")
	assert.Error(t, err)
	assert.Empty(t, user.Password)
	assert.Empty(t, user.Salt)
}

func TestUser_VaildPassword(t *testing.T) {
	user := &models.User{}

	password := "password"
	err := user.SetPassword(password)
	assert.NoError(t, err)

	assert.True(t, user.VaildPassword(password))
}

func TestUser_VaildPassword_InvalidPassword(t *testing.T) {
	user := &models.User{}

	password := "password"
	err := user.SetPassword(password)
	assert.NoError(t, err)

	assert.False(t, user.VaildPassword("invalidPassword"))
}

func TestUser_CheckMustChangePassword(t *testing.T) {
	user := &models.User{}

	user.MustChangePassword = true
	assert.True(t, user.CheckMustChangePassword())

	user.MustChangePassword = false
	assert.False(t, user.CheckMustChangePassword())
}

func TestUser_AccountError(t *testing.T) {
	user := &models.User{}

	err := user.AccountError()
	assert.NoError(t, err)

	user.MustChangePassword = true
	err = user.AccountError()
	assert.EqualError(t, err, models.ErrUserMustChangePassword.Error())
}

func TestUser_SetUsername(t *testing.T) {
	user := &models.User{}

	err := user.SetUsername("user")
	assert.NoError(t, err)
	assert.Equal(t, "user", user.Username)
}

func TestUser_SetUsername_InvalidUsername(t *testing.T) {
	user := &models.User{}

	err := user.SetUsername("use")
	assert.Error(t, err)
	assert.Empty(t, user.Username)
}

func TestUser_GenerateToken(t *testing.T) {
	globalVars.InitGlobalVars()
	globalVars.ProjectConfig.Set(&structs.ProjectConfig{
		Token: structs.TokenConfig{
			SigningKey:  "test-signing-key",
			ExpiredTime: time.Minute,
		},
	})

	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
	}

	token, err := user.GenerateToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestUser_ParseToken(t *testing.T) {
	globalVars.InitGlobalVars()
	globalVars.ProjectConfig.Set(&structs.ProjectConfig{
		Token: structs.TokenConfig{
			SigningKey:  "test-signing-key",
			ExpiredTime: time.Minute,
		},
	})

	user := &models.User{
		Model: gorm.Model{
			ID: 1,
		},
	}

	token, _ := user.GenerateToken()
	err := user.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
}

func TestUser_ParseToken_InvalidToken(t *testing.T) {
	user := &models.User{}

	err := user.ParseToken("invalidToken")
	assert.Error(t, err)
	assert.Equal(t, uint(0), user.ID)
}

func TestNewUser(t *testing.T) {
	role := models.Role{
		Name: "admin",
	}

	user, err := models.NewUser("user", "password", "John Doe", 30, role)
	assert.NoError(t, err)
	assert.Equal(t, "user", user.Username)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Salt)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, user.Age)
	assert.Equal(t, uint(0), user.RoleID)
	assert.Equal(t, role, user.Role)
}

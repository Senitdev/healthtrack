package repository

import (
	"fmt"
	"healthtrack/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository interface {
	GetUserByLoginAndPassword(username string, password string) bool
}
type loginRepository struct {
	DB *gorm.DB
}

// GetUserByLoginAndPassword implements LoginRepository.
func (l *loginRepository) GetUserByLoginAndPassword(username string, password string) bool {
	if username == "" || password == "" {
		return false
	}
	var user entity.User
	// 🔹 Récupération de l’utilisateur par username
	if err := l.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Utilisateur non trouvé")
			return false
		}
		fmt.Println("Erreur DB :", err)
		return false
	}
	// 🔹 Comparaison du mot de passe fourni avec le hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// si err != nil => les mots de passe ne correspondent pas
		fmt.Println("Mot de passe invalide")
		return false
	}
	fmt.Println("Authentification réussie !")
	return true
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{
		DB: db,
	}
}

package helpers


type SignedDetails struct{
	Email
}

func GenerateAllTokens(email, firstName, lastName, userType, userId string) (signedToken, signedRefreshToken string, err error) {

}
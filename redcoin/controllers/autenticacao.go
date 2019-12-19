package controllers

import(
	_ "fmt"
	_ "log"
	"net/http"
	"encoding/json"
	"time"
	"redcoin/modelos"
	"redcoin/conexoes"
	"github.com/dgrijalva/jwt-go"
)

// AutenticacaoApi método que realiza a autenticação da API
func AutenticacaoApi(w http.ResponseWriter, r *http.Request) {
	var creds modelos.Credenciais

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := modelos.UsuariosApi[creds.Usuario]

	if !ok || expectedPassword != creds.Senha {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	
	claims := &modelos.Claims{
		Usuario: creds.Usuario,
		StandardClaims: jwt.StandardClaims{	
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(conexoes.ChaveJwt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })

	w.Write([]byte(`
	{
		"msg":"Token criado com êxito, utilize-o como parametro Header nas chamadas das demais rotas",
		"token":"`+tokenString+`"
	}
	`))
}

func ValidarToken(w http.ResponseWriter, r *http.Request)(tokenValido bool) {
	// // We can obtain the session token from the requests cookies, which come with every request
	// c, err := r.Cookie("token")
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		// If the cookie is not set, return an unauthorized status
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	// For any other type of error, return a bad request status
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	tokenApi := r.Header.Get("token")
	
	// Get the JWT string from the cookie
	// tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &modelos.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenApi, claims, func(token *jwt.Token) (interface{}, error) {
		return conexoes.ChaveJwt, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	return true
}
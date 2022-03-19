package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReturnData struct {
	data map[string]interface{}
}

func NewReturnData() *ReturnData {
	return &ReturnData{
		data: make(map[string]interface{}),
	}
}

func ReturnError(err error) *ReturnData {
	return &ReturnData{
		data: map[string]interface{}{
			"error": err.Error(),
		},
	}
}

func ReturnDbg(err error) *ReturnData {
	return &ReturnData{
		data: map[string]interface{}{
			"sysdbg": err.Error(),
		},
	}
}

func (rd *ReturnData) AddField(key string, value interface{}) error {
	_, err := json.Marshal(value)
	if err != nil {
		return err
	}

	rd.data[key] = value
	return nil
}

func (rd *ReturnData) String() string {
	str, _ := json.Marshal(rd.data)
	return string(str)
}

func ValidateAccessToken(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	token := r.Form.Get("access_token")
	hash := sha256.Sum256([]byte(token))
	if hex.EncodeToString(hash[:]) != accessTokenHash {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s\n", ReturnError(fmt.Errorf("invalid access token: %s", token)).String())
		return false
	}

	delete(r.Form, "access_token")
	return true
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s\n", ReturnError(err).String())
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%s\n", ReturnDbg(err).String())
}

func ResponseOK(w http.ResponseWriter, kvs ...interface{}) {
	rd := NewReturnData()
	for i := 0; i < len(kvs)/2; i++ {
		rd.AddField(kvs[i*2].(string), kvs[i*2+1])
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", rd.String())
}

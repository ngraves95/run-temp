package tokens;

import (
	"github.com/ngraves95/util"
	"io/ioutil"
	"strconv"
	"strings"
)

type TokenManager struct {
	KeyDirectory string;
}

func NewTokenManager(rootDir string) *TokenManager {

	if !strings.HasSuffix(rootDir, "/") {
		rootDir += "/";
	}

	return &TokenManager{
		KeyDirectory: rootDir,
	}
}

func (t *TokenManager) pathTo(fname string) string {
	return t.KeyDirectory + fname;
}


func readSingleLineFile(fname string) string {
	raw, err := ioutil.ReadFile(fname);
	util.Check(err);
	return string(raw);
}

func (t *TokenManager) GetClientSecret() string {
	secret := readSingleLineFile(t.pathTo("strava_client_secret.key"));
	return secret;
}


func (t *TokenManager) GetClientId() int {
	id := readSingleLineFile(t.pathTo("strava_client_id.key"));
	i, err := strconv.Atoi(id);
	util.Check(err);
	return i;
}

func (t *TokenManager) GetStravaAccessToken() string {
	token := readSingleLineFile(t.pathTo("strava_access_token.key"));
	return token;
}

func (t *TokenManager) GetWeatherAccessToken() string {
	token := readSingleLineFile(t.pathTo("openweatherapi.key"));
	return token;
}

func (t *TokenManager) GetToken(name string) string {
	fname := name + ".key";
	token := readSingleLineFile(t.pathTo(fname));
	return token;

}

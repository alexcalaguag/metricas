package model
import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	config "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/config"

)

type SonarMetric struct {
	Component struct {
		ID          string `json:"id"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Qualifier   string `json:"qualifier"`
		Measures    []struct {
			Metric  string `json:"metric"`
			Value   string `json:"value"`
			Periods []struct {
				Index int    `json:"index"`
				Value string `json:"value"`
			} `json:"periods"`
		} `json:"measures"`
	} `json:"component"`
}


func (s *SonarMetric) Get(key string) (*SonarMetric, error) {

    url := buildUrlRequest(key)
	client := &http.Client{}
	log.Print(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", join("Basic ", basicAuthentication()))
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var sonarMetric SonarMetric
		json.Unmarshal(body, &sonarMetric)
		return &sonarMetric, nil
	}

	//TODO ugly
	return nil, nil


}

func basicAuthentication() string {
	auth := join(config.Cfg.User.Chave, ":", config.Cfg.User.Senha)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func buildUrlRequest(key string) string {
	return join(config.Cfg.Sonar.Host,":",config.Cfg.Sonar.Port,config.Cfg.Sonar.SufixTestCoverage, key)
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()

}
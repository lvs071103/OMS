package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type JenkinsBuild struct {
	Actions []struct {
		Parameters []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"parameters"`
	} `json:"actions"`
}

type JenkinsJob struct {
	LastBuild struct {
		Number int `json:"number"`
	} `json:"lastBuild"`
}

type JenkinsJobConfig struct {
	XMLName    xml.Name `xml:"flow-definition"`
	Properties struct {
		ParametersDefinitionProperty struct {
			ParameterDefinitions []Parameter `xml:"parameterDefinitions>hudson.model.ChoiceParameterDefinition"`
		} `xml:"hudson.model.ParametersDefinitionProperty"`
	} `xml:"properties"`
}

type Parameter struct {
	Name         string   `xml:"name"`
	Description  string   `xml:"description,omitempty"`
	DefaultValue string   `xml:"defaultValue,omitempty"`
	Choices      []string `xml:"choices>a>string,omitempty"`
}

func getLastBuildNumber(jenkinsURL, jobName, username, apiToken string) (int, error) {
	url := fmt.Sprintf("%s/job/%s/api/json", jenkinsURL, jobName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.SetBasicAuth(username, apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get job details: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var job JenkinsJob
	if err := json.Unmarshal(body, &job); err != nil {
		return 0, err
	}

	return job.LastBuild.Number, nil
}

func getJenkinsBuildParameters(jenkinsURL, jobName, buildNumber, username, apiToken string) (*JenkinsBuild, error) {
	url := fmt.Sprintf("%s/job/%s/%s/api/json", jenkinsURL, jobName, buildNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get build parameters: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var build JenkinsBuild
	if err := json.Unmarshal(body, &build); err != nil {
		return nil, err
	}

	return &build, nil
}

func getJenkinsJobConfig(jenkinsURL, jobName, username, apiToken string) (*JenkinsJobConfig, error) {
	url := fmt.Sprintf("%s/job/%s/config.xml", jenkinsURL, jobName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get job config: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 将 XML 1.1 版本转换为 XML 1.0 版本
	xmlContent := strings.Replace(string(body), `<?xml version="1.1"`, `<?xml version="1.0"`, 1)

	var config JenkinsJobConfig
	if err := xml.Unmarshal([]byte(xmlContent), &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	jenkinsURL := "http://10.50.48.35:8080"
	jobName := "dominos-gate2"
	username := "furong.zhou"
	apiToken := "1150d2403a07139440735f94615b2c7a8e"

	// 获取上次构建编号
	// lastBuildNumber, err := getLastBuildNumber(jenkinsURL, jobName, username, apiToken)
	// if err != nil {
	// 	fmt.Printf("Error getting last build number: %v\n", err)
	// 	os.Exit(1)
	// }

	// 使用上次构建编号获取构建参数
	// build, err := getJenkinsBuildParameters(jenkinsURL, jobName, fmt.Sprintf("%d", lastBuildNumber), username, apiToken)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Build Parameters:")
	// for _, action := range build.Actions {
	// 	for _, param := range action.Parameters {
	// 		fmt.Printf("Name: %s, Value: %s\n", param.Name, param.Value)
	// 	}
	// }

	// 获取 Jenkins 作业的配置
	config, err := getJenkinsJobConfig(jenkinsURL, jobName, username, apiToken)
	if err != nil {
		fmt.Printf("Error getting job config: %v\n", err)
		os.Exit(1)
	}

	// fmt.Println("Job Parameters:")
	// fmt.Println(config.Properties.ParametersDefinitionProperty.ParameterDefinitions)
	// 将切片转换为 JSON
	jsonData, err := json.MarshalIndent(config.Properties.ParametersDefinitionProperty.ParameterDefinitions, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

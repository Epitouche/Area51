package swagger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"area51/schemas"
)

type SwaggerFile interface {
	ResolvePath(relativePath string) string
	ImpactSwaggerFiles(routes []schemas.Route)
	ProcessFile(filepath string, route schemas.Route)
}

func ResolvePath(relativePath string) string {
	basePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(basePath, relativePath)
}

func ImpactSwaggerFiles(routes []schemas.Route) {
	filePathOfFiles := []string {
		ResolvePath("docs/docs.go"),
		ResolvePath("docs/swagger.json"),
		ResolvePath("docs/swagger.yaml"),
	}

	for _, route := range routes {
		for _, file := range filePathOfFiles {
			ProcessFile(file, route)
		}
	}
}

func ProcessFile(filepath string, route schemas.Route) {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", fileData, err)
		return
	}

	var paths map[string]interface{}
	var yamlPath interface{}

	if IsGOFile(filepath) {
		_, err := UpdateDocTemplate(filepath)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", fileData, err)
		}
	} else if IsJSONFile(filepath) {
		err = json.Unmarshal(fileData, &paths)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON file %s: %s\n", fileData, err)
			return
		}
	} else if IsYAMLFile(filepath) {
		err = yaml.Unmarshal(fileData, &yamlPath)
		if err != nil {
			fmt.Printf("Error unmarshalling YAML file %s: %s\n", fileData, err)
			return
		}
	} else {
		fmt.Printf("Unsupported file type %s\n", fileData)
		return
	}

	if paths == nil {
		paths = make(map[string]interface{})
	}
	if _, ok := paths["paths"]; !ok {
		paths["paths"] = make(map[string]interface{})
	}

	pathsMap := paths["paths"].(map[string]interface{})
	pathsMap[route.Path] = BuildRouteEntry(route)

	if IsGOFile(filepath) {
		_, err := json.MarshalIndent(paths, "", " ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filepath, err)
			return
		}

		newActualFilepath := "tmp.json"
		err = UpdateDocTemplateWithJSON(filepath, newActualFilepath)
		if err != nil {
			fmt.Printf("Error updating docTemplate in file %s: %v\n", filepath, err)
			return
		}
	} else if IsJSONFile(filepath) {
		updatedJSON, err := json.MarshalIndent(paths, "", " ")
		if err != nil {
			fmt.Printf("Error serializing JSON for file %s: %v\n", filepath, err)
			return
		}

		err = os.WriteFile(filepath, updatedJSON, 0644)
		if err != nil {
			fmt.Printf("Error writing JSON to file %s: %v\n", filepath, err)
			return
		}
	} else if IsYAMLFile(filepath) {
		updatedYAML, err := yaml.Marshal(paths)
		if err != nil {
			fmt.Printf("Error serializing YAML for file %s: %v\n", filepath, err)
			return
		}

		err = os.WriteFile(filepath, updatedYAML, 0644)
		if err != nil {
			fmt.Printf("Error writing YAML to file %s: %v\n", filepath, err)
			return
		}
	}
	fmt.Printf("Route added successfully to %s\n", filepath)
}
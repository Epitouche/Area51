package swagger

type SwaggerDetectedFileType interface {
	IsJSONFile(filepath string) bool
	IsYAMLFile(filepath string) bool
	IsGOFile(filepath string) bool
}

func IsJSONFile(filepath string) bool {
	return len(filepath) > 5 && filepath[len(filepath)-5:] == ".json"
}

func IsYAMLFile(filepath string) bool {
	return len(filepath) > 5 && filepath[len(filepath)-5:] == ".yaml"
}

func IsGOFile(filepath string) bool {
	return len(filepath) > 3 && filepath[len(filepath)-3:] == ".go"
}
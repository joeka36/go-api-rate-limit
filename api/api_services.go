package api

import (
    "math/rand"
)

var (
	nameMap = [...]string{
		"alisson",
		"roberto",
		"sadio",
		"mo",
		"van dijk",
		"joe",
		"dejan",
		"joel",
		"andrew",
		"trent",
		"jordan",
		"gini",
		"alex",
		"naby",
		"adam",
		"divock",
		"fabinho",
		"james",
		"takumi",
	}

	descriptionMap = [...]string{
		"adoring",
		"affectionate",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"beautiful",
		"blissful",
		"bold",
		"boring",
		"brave",
		"busy",
		"charming",
		"clever",
		"cool",
		"compassionate",
		"competent",
	}
)

type apiService struct {}

var (
	// APIService use to call service related function
	APIService apiService
)

// GetDockerName returns a randomize DockerName variable
func (a *apiService) GetDockerName() DockerName {
	randomNameIndex := rand.Intn(len(nameMap) - 1)
	randomDescriptionIndex := rand.Intn(len(descriptionMap) - 1)

	return DockerName{
		Name: nameMap[randomNameIndex],
		Description: descriptionMap[randomDescriptionIndex],
	}
}
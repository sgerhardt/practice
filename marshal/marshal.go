package marshal

import (
	"encoding/json"
	"strings"
)

func defaultJsonUnmarshal(b []byte) (string, error) {
	g := &Greeting{Message: string(b)}
	err := json.Unmarshal(b, g)
	if err != nil {
		return "", err
	}
	return g.Message, nil
}

type Greeting struct {
	Message string `json:"Message"`
}

type CustomGreeting struct {
	Message string `json:"Message"`
}

func JsonDecoderUnmarshal(b []byte) (string, error) {
	g := &Greeting{Message: string(b)}
	// unmarshal the json via a decoder
	reader := strings.NewReader(string(b))
	err := json.NewDecoder(reader).Decode(g)
	if err != nil {
		return "", err
	}
	return g.Message, nil
}

func (g *CustomGreeting) UnmarshalJSON(b []byte) error {
	// Create a temporary struct to unmarshal into
	type Alias Greeting
	temp := &struct {
		*Alias
	}{
		Alias: (*Alias)(g),
	}

	// Unmarshal JSON into the temporary struct
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	// Append " world!" to the Message field
	g.Message += " world!"

	return nil
}

func JsonCustomUnmarshal(b []byte) (string, error) {
	g := &CustomGreeting{Message: string(b)}
	// unmarshal the json via a decoder
	err := json.Unmarshal(b, g)
	if err != nil {
		return "", err
	}
	return g.Message, nil
}

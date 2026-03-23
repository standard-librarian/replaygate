package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type Incident struct {
	Request struct {
		Method string            `json:"method"`
		Path   string            `json:"path"`
		Body   string            `json:"body"`
		Header map[string]string `json:"header"`
	} `json:"request"`
	ExpectedResponse struct {
		Status int    `json:"status"`
		Body   string `json:"body"`
	} `json:"expected_response"`
	ActualResponse struct {
		Status int    `json:"status"`
		Body   string `json:"body"`
	} `json:"actual_response"`
	OutgoingHTTP []map[string]any `json:"outgoing_http"`
	Postgres     []map[string]any `json:"postgres"`
	Metadata     map[string]any   `json:"metadata"`
}

type Bundle struct {
	SchemaVersion int      `json:"schema_version"`
	Kind          string   `json:"kind"`
	Incident      Incident `json:"incident"`
}

func loadJSON(path string, target any) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(input, target)
}

func writeJSON(path string, payload any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	output, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	output = append(output, '\n')
	return os.WriteFile(path, output, 0o644)
}

func ingest(args []string) error {
	output, positionals, err := parseIngestArgs(args)
	if err != nil {
		return err
	}
	if len(positionals) != 1 || output == "" {
		return fmt.Errorf("usage: replaygate ingest <incident.json> --output <bundle.json>")
	}

	var incident Incident
	if err := loadJSON(positionals[0], &incident); err != nil {
		return err
	}

	bundle := Bundle{
		SchemaVersion: 1,
		Kind:          "replaygate.bundle",
		Incident:      incident,
	}
	if err := writeJSON(output, bundle); err != nil {
		return err
	}

	fmt.Printf("bundle written to %s\n", output)
	return nil
}

func parseIngestArgs(args []string) (string, []string, error) {
	positionals := []string{}
	output := ""

	for index := 0; index < len(args); index++ {
		switch args[index] {
		case "--output":
			if index+1 >= len(args) {
				return "", nil, fmt.Errorf("missing value for --output")
			}
			output = args[index+1]
			index++
		default:
			positionals = append(positionals, args[index])
		}
	}

	return output, positionals, nil
}

func appendDiff(lines *[]string, path string, left any, right any) {
	if reflect.DeepEqual(left, right) {
		return
	}
	*lines = append(*lines, fmt.Sprintf("%s: %v -> %v", path, left, right))
}

func replay(args []string) error {
	candidatePath, positionals, err := parseReplayArgs(args)
	if err != nil {
		return err
	}
	if len(positionals) != 1 {
		return fmt.Errorf("usage: replaygate replay <bundle.json> [--candidate <incident.json>]")
	}

	var bundle Bundle
	if err := loadJSON(positionals[0], &bundle); err != nil {
		return err
	}

	fmt.Println("bundle summary")
	fmt.Printf("- schema: %d\n", bundle.SchemaVersion)
	fmt.Printf("- request: %s %s\n", bundle.Incident.Request.Method, bundle.Incident.Request.Path)
	fmt.Printf("- outgoing http fixtures: %d\n", len(bundle.Incident.OutgoingHTTP))
	fmt.Printf("- postgres fixtures: %d\n", len(bundle.Incident.Postgres))

	if candidatePath == "" {
		return nil
	}

	var candidate Incident
	if err := loadJSON(candidatePath, &candidate); err != nil {
		return err
	}

	lines := []string{}
	appendDiff(&lines, "$.actual_response.status", bundle.Incident.ActualResponse.Status, candidate.ActualResponse.Status)
	appendDiff(&lines, "$.actual_response.body", bundle.Incident.ActualResponse.Body, candidate.ActualResponse.Body)
	appendDiff(&lines, "$.expected_response.status", bundle.Incident.ExpectedResponse.Status, candidate.ExpectedResponse.Status)
	appendDiff(&lines, "$.expected_response.body", bundle.Incident.ExpectedResponse.Body, candidate.ExpectedResponse.Body)

	if !reflect.DeepEqual(bundle.Incident.OutgoingHTTP, candidate.OutgoingHTTP) {
		lines = append(lines, "$.outgoing_http: fixtures changed")
	}
	if !reflect.DeepEqual(bundle.Incident.Postgres, candidate.Postgres) {
		lines = append(lines, "$.postgres: fixtures changed")
	}

	if len(lines) == 0 {
		fmt.Println("diff: no changes")
		return nil
	}

	fmt.Println("diff:")
	for _, line := range lines {
		fmt.Printf("- %s\n", line)
	}
	return nil
}

func parseReplayArgs(args []string) (string, []string, error) {
	positionals := []string{}
	candidate := ""

	for index := 0; index < len(args); index++ {
		switch args[index] {
		case "--candidate":
			if index+1 >= len(args) {
				return "", nil, fmt.Errorf("missing value for --candidate")
			}
			candidate = args[index+1]
			index++
		default:
			positionals = append(positionals, args[index])
		}
	}

	return candidate, positionals, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: replaygate <ingest|replay> ...")
		os.Exit(2)
	}

	var err error
	switch os.Args[1] {
	case "ingest":
		err = ingest(os.Args[2:])
	case "replay":
		err = replay(os.Args[2:])
	default:
		err = fmt.Errorf("unknown command: %s", os.Args[1])
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

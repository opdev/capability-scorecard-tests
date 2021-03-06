// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	scapiv1alpha3 "github.com/operator-framework/api/pkg/apis/scorecard/v1alpha3"
	apimanifests "github.com/operator-framework/api/pkg/manifests"
)

// This is the custom scorecard test example binary
// As with the Redhat scorecard test image, the bundle that is under
// test is expected to be mounted so that tests can inspect the
// bundle contents as part of their test implementations.
// The actual test is to be run is named and that name is passed
// as an argument to this binary.  This argument mechanism allows
// this binary to run various tests all from within a single
// test image.

const PodBundleRoot = "/bundle"

func main() {
	entrypoint := os.Args[1:]
	if len(entrypoint) == 0 {
		log.Fatal("Test name argument is required")
	}

	// Read the pod's untar'd bundle from a well-known path.
	cfg, err := apimanifests.GetBundleFromDir(PodBundleRoot)
	if err != nil {
		log.Fatal(err.Error())
	}

	var result scapiv1alpha3.TestStatus

	// Names of the custom tests which would be passed in the
	// `operator-sdk` command.
	switch entrypoint[0] {
	case CapabilityLevelOneTestName:
		result = CapabilityLevelOneTest(cfg)
	case CapabilityLevelTwoTestName:
		result = CapabilityLevelTwoTest(cfg)
	case CapabilityLevelThreeTestName:
		result = CapabilityLevelThreeTest(cfg)
	case CapabilityLevelFourTestName:
		result = CapabilityLevelFourTest(cfg)
	case CapabilityLevelFiveTestName:
		result = CapabilityLevelFiveTest(cfg)
	default:
		result = printValidTests()
	}

	// Convert scapiv1alpha3.TestResult to json.
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

}

// printValidTests will print out full list of test names to give a hint to the end user on what the valid tests are.
func printValidTests() scapiv1alpha3.TestStatus {
	result := scapiv1alpha3.TestResult{}
	result.State = scapiv1alpha3.FailState
	result.Errors = make([]string, 0)
	result.Suggestions = make([]string, 0)

	str := fmt.Sprintf("Valid tests for this image include: %s %s %s %s %s",
		CapabilityLevelOneTestName,
		CapabilityLevelTwoTestName,
		CapabilityLevelThreeTestName,
		CapabilityLevelFourTestName,
		CapabilityLevelFiveTestName,
		)
	result.Errors = append(result.Errors, str)
	return scapiv1alpha3.TestStatus{
		Results: []scapiv1alpha3.TestResult{result},
	}
}

const (
	CapabilityLevelOneTestName = "capability-level-one"
	CapabilityLevelTwoTestName = "capability-level-two"
	CapabilityLevelThreeTestName = "capability-level-three"
	CapabilityLevelFourTestName = "capability-level-four"
	CapabilityLevelFiveTestName = "capability-level-five"
)

// CapabilityLevelOneTest accepts a *apimanifests.Bundle and returns scapiv1alpha3.TestStatus
// This test verifies operator capability level one has been met.

func CapabilityLevelOneTest(bundle *apimanifests.Bundle) scapiv1alpha3.TestStatus {
	r := scapiv1alpha3.TestResult{}
	r.Name = CapabilityLevelOneTestName
	r.State = scapiv1alpha3.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)

	// Install operator
	output, err := exec.Command("/usr/local/bin/operator-sdk", "version").Output()
	if err != nil {
		log.Fatalf("Unable to run operator-sdk: %v", err)
	}

	log.Println(string(output))

	// Create CR from first object in alm-examples array

	// Get alm-examples
	almExamples := bundle.CSV.GetAnnotations()["alm-examples"]
	// TODO: should this be a fatal error?
	if almExamples == "" {
		fmt.Println("no alm-examples in the bundle CSV")
	}

	// Apply alm-example to cluster

	return wrapResult(r)
}

// CapabilityLevelTwoTest accepts a *apimanifests.Bundle and returns scapiv1alpha3.TestStatus
// This test verifies operator capability level two has been met.

func CapabilityLevelTwoTest(bundle *apimanifests.Bundle) scapiv1alpha3.TestStatus {
	r := scapiv1alpha3.TestResult{}
	r.Name = CapabilityLevelTwoTestName
	r.State = scapiv1alpha3.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)

	// Install operator
	output, err := exec.Command("/usr/local/bin/operator-sdk", "version").Output()
	if err != nil {
		log.Fatalf("Unable to run operator-sdk: %v", err)
	}

	log.Println(string(output))

	// Create CR from first object in alm-examples array

	// Get alm-examples
	almExamples := bundle.CSV.GetAnnotations()["alm-examples"]
	// TODO: should this be a fatal error?
	if almExamples == "" {
		fmt.Println("no alm-examples in the bundle CSV")
	}

	// Apply alm-example to cluster

	return wrapResult(r)
}

// CapabilityLevelThreeTest accepts a *apimanifests.Bundle and returns scapiv1alpha3.TestStatus
// This test verifies operator capability level three has been met.

func CapabilityLevelThreeTest(bundle *apimanifests.Bundle) scapiv1alpha3.TestStatus {
	r := scapiv1alpha3.TestResult{}
	r.Name = CapabilityLevelThreeTestName
	r.State = scapiv1alpha3.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)

	// Install operator
	output, err := exec.Command("/usr/local/bin/operator-sdk", "version").Output()
	if err != nil {
		log.Fatalf("Unable to run operator-sdk: %v", err)
	}

	log.Println(string(output))

	// Create CR from first object in alm-examples array

	// Get alm-examples
	almExamples := bundle.CSV.GetAnnotations()["alm-examples"]
	// TODO: should this be a fatal error?
	if almExamples == "" {
		fmt.Println("no alm-examples in the bundle CSV")
	}

	// Apply alm-example to cluster

	return wrapResult(r)
}

// CapabilityLevelFourTest accepts a *apimanifests.Bundle and returns scapiv1alpha3.TestStatus
// This test verifies operator capability level four has been met.

func CapabilityLevelFourTest(bundle *apimanifests.Bundle) scapiv1alpha3.TestStatus {
	r := scapiv1alpha3.TestResult{}
	r.Name = CapabilityLevelFourTestName
	r.State = scapiv1alpha3.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)

	// Install operator
	output, err := exec.Command("/usr/local/bin/operator-sdk", "version").Output()
	if err != nil {
		log.Fatalf("Unable to run operator-sdk: %v", err)
	}

	log.Println(string(output))

	// Create CR from first object in alm-examples array

	// Get alm-examples
	almExamples := bundle.CSV.GetAnnotations()["alm-examples"]
	// TODO: should this be a fatal error?
	if almExamples == "" {
		fmt.Println("no alm-examples in the bundle CSV")
	}

	// Apply alm-example to cluster

	return wrapResult(r)
}

// CapabilityLevelFiveTest accepts a *apimanifests.Bundle and returns scapiv1alpha3.TestStatus
// This test verifies operator capability level five has been met.

func CapabilityLevelFiveTest(bundle *apimanifests.Bundle) scapiv1alpha3.TestStatus {
	r := scapiv1alpha3.TestResult{}
	r.Name = CapabilityLevelFiveTestName
	r.State = scapiv1alpha3.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)

	// Install operator
	output, err := exec.Command("/usr/local/bin/operator-sdk", "version").Output()
	if err != nil {
		log.Fatalf("Unable to run operator-sdk: %v", err)
	}

	log.Println(string(output))

	// Create CR from first object in alm-examples array

	// Get alm-examples
	almExamples := bundle.CSV.GetAnnotations()["alm-examples"]
	// TODO: should this be a fatal error?
	if almExamples == "" {
		fmt.Println("no alm-examples in the bundle CSV")
	}

	// Apply alm-example to cluster

	return wrapResult(r)
}

func wrapResult(r scapiv1alpha3.TestResult) scapiv1alpha3.TestStatus {
	return scapiv1alpha3.TestStatus{
		Results: []scapiv1alpha3.TestResult{r},
	}
}

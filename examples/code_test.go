package examples

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"sync"
	"testing"

	"github.com/AdityaShakya-Treeleaf/go-examples/internal/dsa"
)

func TestMaximumEnergyFromDungeon(t *testing.T) {
	case1 := []int{5, 2, -10, -5, 1}
	k1 := 3
	expectedOutput1 := 3
	actualOutput1 := maximumEnergy(case1, k1)

	fmt.Printf("\nTest case 1: Success = %t", expectedOutput1 == actualOutput1)

	case2 := []int{-2, -3, -1}
	k2 := 2
	expectedOutput2 := -1
	actualOutput2 := maximumEnergy(case2, k2)

	fmt.Printf("\nTest case 1: Success = %t", expectedOutput2 == actualOutput2)
}

func maximumEnergy(energy []int, k int) int {
	n := len(energy)

	if n == 0 {
		return 0
	} else if n == 1 {
		return energy[0]
	}

	maxEnergy := math.MinInt

	for i := n - 1; i >= 0; i-- {
		if i+k < n {
			energy[i] += energy[i+k]
		}

		if energy[i] > maxEnergy {
			maxEnergy = energy[i]
		}
	}

	return maxEnergy
}

// Solution 2: Concurrent Processing with Goroutines
// Each chain (starting positions 0 to k-1) is processed in parallel
func maximumEnergyConcurrent(energy []int, k int) int {
	n := len(energy)

	if n == 0 {
		return 0
	} else if n == 1 {
		return energy[0]
	}

	// Create a copy to avoid modifying original
	energyCopy := make([]int, n)
	copy(energyCopy, energy)

	// Channel to collect results from each chain
	results := make(chan int, k)
	var wg sync.WaitGroup

	// Process each chain concurrently
	// Chain i: positions i, i+k, i+2k, ...
	for start := 0; start < k && start < n; start++ {
		wg.Add(1)

		go func(chainStart int) {
			defer wg.Done()

			maxInChain := energyCopy[chainStart]

			// Process this chain backwards
			for i := chainStart; i < n; i += k {
				if i+k < n {
					energyCopy[i] += energyCopy[i+k]
				}
				if energyCopy[i] > maxInChain {
					maxInChain = energyCopy[i]
				}
			}

			results <- maxInChain
		}(start)
	}

	// Close results channel after all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Find maximum across all chains
	maxEnergy := -1 << 31 // Min int
	for chainMax := range results {
		if chainMax > maxEnergy {
			maxEnergy = chainMax
		}
	}

	return maxEnergy
}

// Solution 2 Alternative: Using sync.Mutex for shared max
func maximumEnergyConcurrentMutex(energy []int, k int) int {
	n := len(energy)

	if n == 0 {
		return 0
	} else if n == 1 {
		return energy[0]
	}

	energyCopy := make([]int, n)
	copy(energyCopy, energy)

	var (
		maxEnergy int = -1 << 31
		mu        sync.Mutex
		wg        sync.WaitGroup
	)

	// Process each chain concurrently
	for start := 0; start < k && start < n; start++ {
		wg.Add(1)

		go func(chainStart int) {
			defer wg.Done()

			localMax := energyCopy[chainStart]

			// Process chain backwards
			for i := chainStart; i < n; i += k {
				if i+k < n {
					energyCopy[i] += energyCopy[i+k]
				}
				if energyCopy[i] > localMax {
					localMax = energyCopy[i]
				}
			}

			// Update global max with mutex
			mu.Lock()
			if localMax > maxEnergy {
				maxEnergy = localMax
			}
			mu.Unlock()
		}(start)
	}

	wg.Wait()
	return maxEnergy
}

func TestConcurrentRequests(t *testing.T) {
	url := "https://api.anydone.com/co-form/v1/form"
	numRequests := 3

	var wg sync.WaitGroup
	wg.Add(numRequests)

	requestBody :=
		`{
      		"name": "four",
			"type": 1,
			"formSection": [{
					"order": 1,
					"name": "one",
					"page": 1,
					"formGroups": [{
							"order": 1,
							"name": "First",
							"groupType": 3,
							"fieldId": "first_321551",
							"fields": [{
									"order": 1,
									"type": 19,
									"name": "First",
									"fieldId": "first_321551",
      								"fieldName": "First"
								}]
						}]
				}],
			"subProject": {
  				"subProjectId": "c8df81d1ffca4343be9c8d72ce718afd"
				}
		 }`

	authToken := "YjgwODNjNTVjNDgzNGNmYzkzYzQxZGUzOTkzYjA0YjcuNDA5MTU1YmY3NWU4NGZlMTkwYTQ5MWUxMmQwOWM3NDg=.cb7047b49c6723c56a2587bf49c2eb16bc730ba23f4b665014cea298e525593d9ec5667b52cd0b58b43b82f27b7ceb66b41ed219e38e8a3f53024ce6bd011c73"

	// Launch concurrent requests
	for i := 0; i < numRequests; i++ {
		go func(requestNum int) {
			defer wg.Done()

			// CREATE NEW REQUEST FOR EACH GOROUTINE
			req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
			if reqErr != nil {
				log.Printf("Error creating http request: %s", reqErr)
				t.Errorf("Request %d: Failed to create request: %v", requestNum, reqErr)
				return
			}

			// Set headers for each request
			req.Header.Set("Authorization", authToken)
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")

			// Use client to make request
			client := &http.Client{}
			resp, resErr := client.Do(req)
			if resErr != nil {
				log.Printf("Request %d Failed: %s", requestNum, resErr)
				t.Errorf("Request %d: %v", requestNum, resErr)
				return
			}
			defer resp.Body.Close() // Don't forget to close!

			t.Logf("Request %d: Status Code %d", requestNum, resp.StatusCode)

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Request %d: Expected 200, got %d", requestNum, resp.StatusCode)
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error extracting raw response: %s", err)
			} else {
				log.Printf("\n%d: %s", requestNum, string(data))
			}

		}(i)
	}

	// Wait for all requests to complete
	wg.Wait()
	t.Log("All concurrent requests completed")
}

func TestWaysDivideLongCorridor(t *testing.T) {
	wdlc := dsa.WaysDivideLongCoriddor{}
	input := "SSPPSPSPP"
	expected := 3
	actual := wdlc.NumberOfWays(input)
	if actual != expected {
		t.Fatalf("Failed. Expected: %d, Got: %d", expected, actual)
	}
}

func TestSmoothDescentPeriods(t *testing.T) {
	sdp := dsa.SmoothDescentPeriods{}
	input := []int{8, 6, 7, 7}
	expected := int64(4)
	actual := sdp.GetDescentPeriods(input)
	if actual != expected {
		t.Fatalf("Failed. Expected: %d, Got: %d", expected, actual)
	}
}

package prefixcollector_test

import (
	"cmd-exclude-prefixes-k8s/internal/prefixcollector"
	//"cmd-exclude-prefixes-k8s/internal/prefixcollector"
	"cmd-exclude-prefixes-k8s/internal/utils"
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

type dummyPrefixSource struct {
	prefixes []string
}

func (d *dummyPrefixSource) GetNotifyChannel() <-chan struct{} {
	return nil
}

func (d *dummyPrefixSource) GetPrefixes() []string {
	return d.prefixes
}

func newDummyPrefixSource(prefixes []string) *dummyPrefixSource {
	return &dummyPrefixSource{prefixes}
}

func getDummyCollectorOption(prefixes []string) prefixcollector.ExcludePrefixCollectorOption {
	return func(collector *prefixcollector.ExcludePrefixCollector, ctx context.Context) {
		collector.Sources = append(collector.Sources, newDummyPrefixSource(prefixes))
	}
}

func TestCollector(t *testing.T) {
	firstPrefixes := []string{
		"127.0.0.1/16",
		"127.0.2.1/16",
		"168.92.0.1/24",
	}
	firstOption := getDummyCollectorOption(firstPrefixes)

	secondPrefixes := []string{
		"127.0.3.1/16",
		"134.56.0.1/8",
		"168.92.0.1/16",
	}
	secondOption := getDummyCollectorOption(secondPrefixes)

	expectedResult := []string{
		"127.0.0.0/16",
		"168.92.0.0/16",
		"134.0.0.0/8",
	}

	testCollector(t, expectedResult, firstOption, secondOption)
}

func testCollector(t *testing.T, expectedResult []string, options ...prefixcollector.ExcludePrefixCollectorOption) {
	const path = "test.yaml"
	collector := prefixcollector.NewExcludePrefixCollector(path, nil, options...)
	collector.UpdateExcludedPrefixesConfigmap()
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("Error reading test file: ", err)
	}

	prefixes, err := utils.YamlToPrefixes(bytes)
	if err != nil {
		t.Fatal("Error transforming yaml to prefixes: ", err)
	}

	assert.ElementsMatch(t, expectedResult, prefixes.PrefixesList)
	_ = os.Remove(path)
}

package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestAlbExample(t *testing.T) {
	t.Parallel()

	opts := &terraform.Options{
		TerraformDir: "../live/examples/alb",

		Vars: map[string]interface{}{
			"alb_name": fmt.Sprintf("test-%s", random.UniqueId()),
		},
	}

	//clean up everything at the end of the test
	defer terraform.Destroy(t, opts)

	terraform.InitAndApply(t, opts)

	//Get the URL of the ALB
	albDnsName := terraform.OutputRequired(t, opts, "alb_dns_name")
	url := fmt.Sprintf("http://%s", albDnsName)

	//Test albs default action is working and returns a 404
	expectedStatus := 404
	expectedbody := "404: page not found"
	maxRetries := 10
	timeBetweenRetries := 10 * time.Second

	http_helper.HttpGetWithRetry(
		t,
		url,
		nil,
		expectedStatus,
		expectedbody,
		maxRetries,
		timeBetweenRetries,
	)
}

func TestAlbExamplePlan(t *testing.T) {
	t.Parallel()

	albName := fmt.Sprintf("test-%s", random.UniqueId())

	opts := &terraform.Options{
		TerraformDir: "../Live/examples/alb",
		Vars: map[string]interface{}{
			"alb_name": albName,
		},
	}

	planString := terraform.InitAndPlan(t, opts)

	// check plans outputs etc
	resourceCounts := terraform.GetResourceCount(t, planString)
	require.Equal(t, 5, resourceCounts.Add)
	require.Equal(t, 0, resourceCounts.Change)
	require.Equal(t, 0, resourceCounts.Destroy)

	planStruct := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, opts)
	alb, exists := planStruct.ResourcePlannedValuesMap["module.alb.aws_lb.example"]
	require.True(t, exists, "aws_lb resource must exist")

	name, exists := alb.AttributeValues["name"]
	require.True(t, exists, "missing name parameter")
	require.Equal(t, albName, name)
}

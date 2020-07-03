package alicloud

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAlicloudWafInstancesDataSource(t *testing.T) {
	rand := acctest.RandInt()
	wafInstanceId := os.Getenv("ALICLOUD_WAF_INSTANCE_ID")
	resourceGroupId := os.Getenv("ALICLOUD_RESOURCE_GROUP_ID")
	idsConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"ids": fmt.Sprintf(`["%s"]`, wafInstanceId),
		}),
		fakeConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"ids": `["fake"]`,
		}),
	}

	statusConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"status": `"1"`,
		}),
		fakeConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"status": `"0"`,
		}),
	}

	instanceSourceConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"instance_source": `"waf-cloud"`,
		}),
		fakeConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"ids":             `["fake"]`,
			"instance_source": `"waf-cloud"`,
		}),
	}

	resourceGroupIdConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"resource_group_id": fmt.Sprintf(`"%s"`, resourceGroupId),
		}),
		fakeConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"ids":               `["fake"]`,
			"resource_group_id": fmt.Sprintf(`"%s"`, resourceGroupId),
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"ids":               fmt.Sprintf(`["%s"]`, wafInstanceId),
			"status":            `"1"`,
			"instance_source":   `"waf-cloud"`,
			"resource_group_id": fmt.Sprintf(`"%s"`, resourceGroupId),
		}),
		fakeConfig: testAccCheckAlicloudWafInstanceDataSourceConfig(rand, map[string]string{
			"ids":               `["fake"]`,
			"status":            `"1"`,
			"instance_source":   `"waf-cloud"`,
			"resource_group_id": fmt.Sprintf(`"%s"`, resourceGroupId),
		}),
	}

	var existDnsRecordsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                         "1",
			"instances.#":                   "1",
			"instances.0.id":                wafInstanceId,
			"instances.0.instance_id":       wafInstanceId,
			"instances.0.end_date":          CHECKSET,
			"instances.0.in_debt":           CHECKSET,
			"instances.0.region":            "cn",
			"instances.0.remain_day":        CHECKSET,
			"instances.0.status":            "1",
			"instances.0.subscription_type": "Subscription",
			"instances.0.trial":             CHECKSET,
		}
	}

	var fakeDnsRecordsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":       "0",
			"instances.#": "0",
		}
	}

	var wafInstancesRecordsCheckInfo = dataSourceAttr{
		resourceId:   "data.alicloud_waf_instances.default",
		existMapFunc: existDnsRecordsMapFunc,
		fakeMapFunc:  fakeDnsRecordsMapFunc,
	}

	var perCheck = func() {
		testAccPreCheck(t)
		testAccPreCheckWithWafInstanceSetting(t)
		testAccPreCheckWithAlicloudResourceGroupIdSetting(t)
	}

	wafInstancesRecordsCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, perCheck, idsConf, statusConf, instanceSourceConf, resourceGroupIdConf, allConf)

}

func testAccCheckAlicloudWafInstanceDataSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}

	config := fmt.Sprintf(`
data "alicloud_waf_instances" "default" {
  %s
}
`, strings.Join(pairs, "\n  "))
	return config
}
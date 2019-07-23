package rancher2

import (
	"reflect"
	"testing"
)

var (
	testAwsCredsConfig                       *Config
	testClusterEKSConfigConf                 *AmazonElasticContainerServiceConfig
	testClusterEKSConfigInterface            []interface{}
	testProviderAwsCredsClusterEKSConfigConf *AmazonElasticContainerServiceConfig
	testNoAwsCredsClusterEKSConfigInterface  []interface{}
	testNoAwsTokenClusterEKSConfigConf       *AmazonElasticContainerServiceConfig
	testNoAwsTokenClusterEKSConfigInterface  []interface{}
)

func init() {
	testAwsCredsConfig = &Config{
		AwsCredentials: AwsCredentials{
			AccessKey:    "provider_XXXXXXXX",
			SecretKey:    "provider_YYYYYYYY",
			SessionToken: "provider_session_token",
		},
	}

	getBaseConfig := func() *AmazonElasticContainerServiceConfig {
		return &AmazonElasticContainerServiceConfig{
			AccessKey:                   "XXXXXXXX",
			SecretKey:                   "YYYYYYYY",
			AMI:                         "ami",
			AssociateWorkerNodePublicIP: newTrue(),
			DisplayName:                 "test",
			InstanceType:                "instance",
			KubernetesVersion:           "1.11",
			MaximumNodes:                5,
			MinimumNodes:                3,
			NodeVolumeSize:              40,
			Region:                      "region",
			SecurityGroups:              []string{"sg1", "sg2"},
			ServiceRole:                 "role",
			SessionToken:                "session_token",
			Subnets:                     []string{"subnet1", "subnet2"},
			UserData:                    "user_data",
			VirtualNetwork:              "network",
		}
	}

	getBaseEKSConfigInterface := func() map[string]interface{} {
		return map[string]interface{}{
			"access_key":                      "XXXXXXXX",
			"secret_key":                      "YYYYYYYY",
			"ami":                             "ami",
			"associate_worker_node_public_ip": true,
			"instance_type":                   "instance",
			"kubernetes_version":              "1.11",
			"maximum_nodes":                   5,
			"minimum_nodes":                   3,
			"node_volume_size":                40,
			"region":                          "region",
			"security_groups":                 []interface{}{"sg1", "sg2"},
			"service_role":                    "role",
			"session_token":                   "session_token",
			"subnets":                         []interface{}{"subnet1", "subnet2"},
			"user_data":                       "user_data",
			"virtual_network":                 "network",
		}
	}

	testClusterEKSConfigConf = getBaseConfig()

	testClusterEKSConfigInterface = []interface{}{
		getBaseEKSConfigInterface(),
	}

	// Cases for AWS creds in provider, not in eks_config
	testProviderAwsCredsClusterEKSConfigConf = getBaseConfig()
	testProviderAwsCredsClusterEKSConfigConf.AccessKey = "provider_XXXXXXXX"
	testProviderAwsCredsClusterEKSConfigConf.SecretKey = "provider_YYYYYYYY"
	testProviderAwsCredsClusterEKSConfigConf.SessionToken = "provider_session_token"

	copyBaseEKSInterface := getBaseEKSConfigInterface()
	delete(copyBaseEKSInterface, "access_key")
	delete(copyBaseEKSInterface, "secret_key")
	delete(copyBaseEKSInterface, "session_token")

	testNoAwsCredsClusterEKSConfigInterface = []interface{}{
		copyBaseEKSInterface,
	}

	// Test for AWS creds in provider, base in eks_config
	testNoAwsTokenClusterEKSConfigConf = getBaseConfig()
	testNoAwsTokenClusterEKSConfigConf.SessionToken = ""

	copyBaseEKSInterface = getBaseEKSConfigInterface()
	delete(copyBaseEKSInterface, "session_token")
	testNoAwsTokenClusterEKSConfigInterface = []interface{}{
		copyBaseEKSInterface,
	}
}

func TestFlattenClusterEKSConfig(t *testing.T) {

	cases := []struct {
		Input          *AmazonElasticContainerServiceConfig
		ExpectedOutput []interface{}
	}{
		{
			testClusterEKSConfigConf,
			testClusterEKSConfigInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterEKSConfig(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterEKSConfig(t *testing.T) {

	cases := []struct {
		Config         *Config
		Input          []interface{}
		ExpectedOutput *AmazonElasticContainerServiceConfig
	}{
		{
			&Config{},
			testClusterEKSConfigInterface,
			testClusterEKSConfigConf,
		},
		{
			testAwsCredsConfig,
			testClusterEKSConfigInterface,
			testClusterEKSConfigConf,
		},
		{
			testAwsCredsConfig,
			testNoAwsCredsClusterEKSConfigInterface,
			testProviderAwsCredsClusterEKSConfigConf,
		},
		{
			testAwsCredsConfig,
			testNoAwsTokenClusterEKSConfigInterface,
			testNoAwsTokenClusterEKSConfigConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterEKSConfig(tc.Input, "test", tc.Config)
		if err != nil {
			t.Fatalf("[ERROR] on expander: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

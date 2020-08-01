package concourse

import (
	"reflect"
	"testing"

	"github.com/concourse/concourse/atc"
)

func TestFlattenTeamAuthData(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		auth   atc.TeamAuth
		expect map[string][]string
	}{
		"owner groups exists": {
			auth: atc.TeamAuth{
				"owner": {
					"groups": []string{"testGroupID"},
				},
			},
			expect: map[string][]string{
				"owner_groups": {"testGroupID"},
			},
		},
		"owner users exists": {
			auth: atc.TeamAuth{
				"owner": {
					"users": []string{"testUserID"},
				},
			},
			expect: map[string][]string{
				"owner_users": {"testUserID"},
			},
		},
		"owner groups and users exist": {
			auth: atc.TeamAuth{
				"owner": {
					"groups": []string{"testGroupID"},
					"users":  []string{"testUserID"},
				},
			},
			expect: map[string][]string{
				"owner_groups": {"testGroupID"},
				"owner_users":  {"testUserID"},
			},
		},
		"pipeline-operator exists": {
			auth: atc.TeamAuth{
				"pipeline-operator": {
					"groups": []string{"testGroupID"},
					"users":  []string{"testUserID"},
				},
			},
			expect: map[string][]string{
				"pipeline_operator_groups": {"testGroupID"},
				"pipeline_operator_users":  {"testUserID"},
			},
		},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := flattenTeamAuthData(test.auth)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("map should be %v, but it is %v", test.expect, actual)
			}
		})
	}

}

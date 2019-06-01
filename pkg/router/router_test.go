package router

import (
	"testing"
)

func TestNewRoutingTable(t *testing.T) {

	wantRoutes := []Route{
		Route{
			Path:   "$.Test",
			Target: "http://example.com/e1",
		},
	}

	config := `{
	"routes": [
		{
			"path": "$.Test",
			"target": "http://example.com/e1"
		}
	]}`

	data := []byte(config)
	routeTable, err := NewRoutingTable(&data)
	if err != nil {
		t.Fatalf("Failed to parse the route table %s", err)
	}
	routes := routeTable.Routes

	for i, want := range wantRoutes {
		if want.Path != routes[i].Path {
			t.Fatalf("Failed got %s wanted %s", routes[i].Path, want.Path)
		}
		if want.Target != routes[i].Target {
			t.Fatalf("Failed got %s wanted %s", routes[i].Target, want.Target)
		}
	}
}

func TestFindMatches(t *testing.T) {

	routes := RoutingTable{
		Routes: []Route{
			Route{
				Path:   "$.MatchMe",
				Target: "http://example.com/e1",
			},
			Route{
				Path:   "$.MatchMeAlso",
				Target: "http://example.com/e1",
			},
			Route{
				Path:   "$.MatchMeAlso.b",
				Target: "http://example.com/e2",
			},
		},
	}

	json_input := `{
		"MatchMe" : {
			"a": 1
		},
		"MatchMeAlso" : {
			"b": 2
		},
		"DoNotMatchMe" : {
			"c": 3
		}
	}`

	wantTargets := map[string]bool{
		"http://example.com/e1": true,
		"http://example.com/e2": true,
	}

	targets, err := routes.FindMatches(&json_input)
	if err != nil {
		t.Errorf("Failed to unmarshal json %s", err)
	}
	if len(targets) != 2 {
		t.Errorf("Expecting 2 targets, got %d", len(targets))
	}
	for e, _ := range wantTargets {
		if !targets[e] {
			t.Errorf("Expecting endpoint %s", e)
		}
	}

}

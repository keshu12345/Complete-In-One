// Copyright The OpenTelemetry Authors
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

package propagation_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
)

type property struct {
	Key, Value string
}

type member struct {
	Key, Value string

	Properties []property
}

func (m member) Member(t *testing.T) baggage.Member {
	props := make([]baggage.Property, 0, len(m.Properties))
	for _, p := range m.Properties {
		p, err := baggage.NewKeyValueProperty(p.Key, p.Value)
		if err != nil {
			t.Fatal(err)
		}
		props = append(props, p)
	}

	bMember, err := baggage.NewMember(m.Key, m.Value, props...)
	if err != nil {
		t.Fatal(err)
	}
	return bMember
}

type members []member

func (m members) Baggage(t *testing.T) baggage.Baggage {
	bMembers := make([]baggage.Member, 0, len(m))
	for _, mem := range m {
		bMembers = append(bMembers, mem.Member(t))
	}
	bag, err := baggage.New(bMembers...)
	if err != nil {
		t.Fatal(err)
	}
	return bag
}

func TestExtractValidBaggageFromHTTPReq(t *testing.T) {
	prop := propagation.TextMapPropagator(propagation.Baggage{})
	tests := []struct {
		name   string
		header string
		want   members
	}{
		{
			name:   "valid w3cHeader",
			header: "key1=val1,key2=val2",
			want: members{
				{Key: "key1", Value: "val1"},
				{Key: "key2", Value: "val2"},
			},
		},
		{
			name:   "valid w3cHeader with spaces",
			header: "key1 =   val1,  key2 =val2   ",
			want: members{
				{Key: "key1", Value: "val1"},
				{Key: "key2", Value: "val2"},
			},
		},
		{
			name:   "valid w3cHeader with properties",
			header: "key1=val1,key2=val2;prop=1",
			want: members{
				{Key: "key1", Value: "val1"},
				{
					Key:   "key2",
					Value: "val2",
					Properties: []property{
						{Key: "prop", Value: "1"},
					},
				},
			},
		},
		{
			name:   "valid header with an invalid header",
			header: "key1=val1,key2=val2,a,val3",
			want:   members{},
		},
		{
			name:   "valid header with no value",
			header: "key1=,key2=val2",
			want: members{
				{Key: "key1", Value: ""},
				{Key: "key2", Value: "val2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			req.Header.Set("baggage", tt.header)

			ctx := context.Background()
			ctx = prop.Extract(ctx, propagation.HeaderCarrier(req.Header))
			expected := tt.want.Baggage(t)
			assert.Equal(t, expected, baggage.FromContext(ctx))
		})
	}
}

func TestExtractInvalidDistributedContextFromHTTPReq(t *testing.T) {
	prop := propagation.TextMapPropagator(propagation.Baggage{})
	tests := []struct {
		name   string
		header string
		has    members
	}{
		{
			name:   "no key values",
			header: "header1",
		},
		{
			name:   "invalid header with existing context",
			header: "header2",
			has: members{
				{Key: "key1", Value: "val1"},
				{Key: "key2", Value: "val2"},
			},
		},
		{
			name:   "empty header value",
			header: "",
			has: members{
				{Key: "key1", Value: "val1"},
				{Key: "key2", Value: "val2"},
			},
		},
		{
			name:   "with properties",
			header: "key1=val1,key2=val2;prop=1",
			has: members{
				{Key: "key1", Value: "val1"},
				{
					Key:   "key2",
					Value: "val2",
					Properties: []property{
						{Key: "prop", Value: "1"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			req.Header.Set("baggage", tt.header)

			expected := tt.has.Baggage(t)
			ctx := baggage.ContextWithBaggage(context.Background(), expected)
			ctx = prop.Extract(ctx, propagation.HeaderCarrier(req.Header))
			assert.Equal(t, expected, baggage.FromContext(ctx))
		})
	}
}

func TestInjectBaggageToHTTPReq(t *testing.T) {
	propagator := propagation.Baggage{}
	tests := []struct {
		name         string
		mems         members
		wantInHeader []string
	}{
		{
			name: "two simple values",
			mems: members{
				{Key: "key1", Value: "val1"},
				{Key: "key2", Value: "val2"},
			},
			wantInHeader: []string{"key1=val1", "key2=val2"},
		},
		{
			name: "values with escaped chars",
			mems: members{
				{Key: "key2", Value: "val3=4"},
			},
			wantInHeader: []string{"key2=val3%3D4"},
		},
		{
			name: "with properties",
			mems: members{
				{Key: "key1", Value: "val1"},
				{
					Key:   "key2",
					Value: "val2",
					Properties: []property{
						{Key: "prop", Value: "1"},
					},
				},
			},
			wantInHeader: []string{
				"key1=val1",
				"key2=val2;prop=1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			ctx := baggage.ContextWithBaggage(context.Background(), tt.mems.Baggage(t))
			propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

			got := strings.Split(req.Header.Get("baggage"), ",")
			assert.ElementsMatch(t, tt.wantInHeader, got)
		})
	}
}

func TestBaggagePropagatorGetAllKeys(t *testing.T) {
	var propagator propagation.Baggage
	want := []string{"baggage"}
	got := propagator.Fields()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("GetAllKeys: -got +want %s", diff)
	}
}
// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package roundtrip

import "testing"

// TestRevisionIsPositive proves R1: Revision reports a positive
// build marker, whichever one is current — a correction changes
// WHICH marker is current, never the requirement that it be
// positive.
func TestRevisionIsPositive(t *testing.T) {
	if Revision() <= 0 {
		t.Fatalf("Revision() = %d, want a positive marker", Revision())
	}
}

// TestRevisionIsStable proves R2: Revision reports the same
// marker on every call within one build — a constant, not
// something that drifts mid-run.
func TestRevisionIsStable(t *testing.T) {
	first := Revision()
	if second := Revision(); second != first {
		t.Fatalf("Revision() = %d then %d, want the same value both times", first, second)
	}
}

// TestRevisionEndToEnd exercises Revision the way a caller
// checking build provenance would: read it once and use it as
// an opaque marker.
func TestRevisionEndToEnd(t *testing.T) {
	if Revision() <= 0 {
		t.Fatalf("Revision() = %d, want a positive marker", Revision())
	}
}

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 *
 * Copyright 2019 gRPC authors.
 *
 */

// Package orca implements Open Request Cost Aggregation.
package orca

import (
	orcapb "github.com/cncf/xds/go/xds/data/orca/v3"

	"github.com/dubbogo/gost/log/logger"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc/metadata"
)

import (
	"dubbo.apache.org/dubbo-go/v3/xds/utils/balancerload"
)

const mdKey = "X-Endpoint-Load-Metrics-Bin"

// toBytes converts a orca load report into bytes.
func toBytes(r *orcapb.OrcaLoadReport) []byte {
	if r == nil {
		return nil
	}

	b, err := proto.Marshal(r)
	if err != nil {
		logger.Warnf("orca: failed to marshal load report: %v", err)
		return nil
	}
	return b
}

// ToMetadata converts a orca load report into grpc metadata.
func ToMetadata(r *orcapb.OrcaLoadReport) metadata.MD {
	b := toBytes(r)
	if b == nil {
		return nil
	}
	return metadata.Pairs(mdKey, string(b))
}

// fromBytes reads load report bytes and converts it to orca.
func fromBytes(b []byte) *orcapb.OrcaLoadReport {
	ret := new(orcapb.OrcaLoadReport)
	if err := proto.Unmarshal(b, ret); err != nil {
		logger.Warnf("orca: failed to unmarshal load report: %v", err)
		return nil
	}
	return ret
}

// FromMetadata reads load report from metadata and converts it to orca.
//
// It returns nil if report is not found in metadata.
func FromMetadata(md metadata.MD) *orcapb.OrcaLoadReport {
	vs := md.Get(mdKey)
	if len(vs) == 0 {
		return nil
	}
	return fromBytes([]byte(vs[0]))
}

type loadParser struct{}

func (*loadParser) Parse(md metadata.MD) any {
	return FromMetadata(md)
}

func init() {
	balancerload.SetParser(&loadParser{})
}

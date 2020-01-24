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

package config_center

import (
	"time"
)

import (
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/config_center/parser"
)

//////////////////////////////////////////
// DynamicConfiguration
//////////////////////////////////////////
const DEFAULT_GROUP = "dubbo"
const DEFAULT_CONFIG_TIMEOUT = "10s"

// DynamicConfiguration ...
type DynamicConfiguration interface {
	Parser() parser.ConfigurationParser
	SetParser(parser.ConfigurationParser)
	AddListener(string, ConfigurationListener, ...Option)
	RemoveListener(string, ConfigurationListener, ...Option)
	//GetProperties get properties file
	GetProperties(string, ...Option) (string, error)

	//GetRule get Router rule properties file
	GetRule(string, ...Option) (string, error)

	//GetInternalProperty get value by key in Default properties file(dubbo.properties)
	GetInternalProperty(string, ...Option) (string, error)
}

// Options ...
type Options struct {
	Group   string
	Timeout time.Duration
}

// Option ...
type Option func(*Options)

// WithGroup ...
func WithGroup(group string) Option {
	return func(opt *Options) {
		opt.Group = group
	}
}

// WithTimeout ...
func WithTimeout(time time.Duration) Option {
	return func(opt *Options) {
		opt.Timeout = time
	}
}

//GetRuleKey The format is '{interfaceName}:[version]:[group]'
func GetRuleKey(url common.URL) string {
	return url.ColonSeparatedKey()
}

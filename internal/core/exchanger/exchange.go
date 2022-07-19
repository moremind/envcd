/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exchanger

import (
	"github.com/acmestack/envcd/internal/core/exchanger/etcd"
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/godkits/gox/errorsx"
)

// Exchanger exchanger interface
type Exchanger interface {

	// Put put data into exchanger
	//  @param key data identity
	//  @param value data
	Put(key interface{}, value interface{}) error

	// Remove remove data from exchanger
	//  @param o data
	Remove(key interface{}) error
}

type Exchange struct {
	exchanger Exchanger
}

// Start envcd by envcd exchangerConnMetadata config
//  @param exchangerConnMetadata the config for envcd
func Start(exchangerConnMetadata *config.ConnMetadata) *Exchange {
	return &Exchange{exchanger: etcd.New(exchangerConnMetadata)}
}

// Put new data to Exchanger by key and value
func (exchange *Exchange) Put(key interface{}, value interface{}) error {
	if exchange == nil || isNil(exchange.exchanger) {
		return errorsx.Err("IIllegal state for exchange")
	}
	return exchange.exchanger.Put(key, value)
}

// Remove delete the data from Exchanger by key
func (exchange *Exchange) Remove(key interface{}) error {
	if exchange == nil || isNil(exchange.exchanger) {
		return errorsx.Err("IIllegal state for exchange")
	}
	return exchange.exchanger.Remove(key)
}

func isNil(exchanger Exchanger) bool {
	v, _ := exchanger.(*Exchange)
	return v == nil
}
//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package cache

import "context"

const logLevel = 3

type Cache struct {
	node    *CacheNodeManifest
}

type Cleaner func(ctx context.Context) error

func (c *Cache) Node() *CacheNodeManifest {
	return c.node
}

func NewCache() *Cache {
	c := new(Cache)
	c.node = NewCacheNodeManifest()
	return c
}

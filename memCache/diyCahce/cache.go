package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	// 不设置过期时间
	NoExpiration time.Duration = -1

	// cache默认的过期时间
	DefaultExpiration time.Duration = 0
)

type Cache struct {
	*cache
}

type cache struct {
	defaultExpiration time.Duration // 默认的过期间隔
	items             map[string]*Item
	mu                sync.RWMutex
	onEvicted         func(string, interface{}) // 回调函数，就是在操作key的时候需要执行的额外操作
	janitor           *janitor                  // 清理器
}

type Item struct {
	Object     interface{}
	Expiration int64 // 过期的时间戳
}

type keyAndValue struct {
	key   string
	value interface{} // value的值，item的object
}

func newCache(de time.Duration, m map[string]*Item) *cache {
	if de == 0 {
		de = -1
	}
	c := &cache{
		defaultExpiration: de,
		items:             m,
	}
	return c
}

/** newCacheWithJanitor
* @Description:
* @param de
* @param ci	触发清理的时间间隔
* @param m
* @return *Cache
 */
func newCacheWithJanitor(de time.Duration, ci time.Duration, m map[string]*Item) *Cache {
	c := newCache(de, m)
	C := &Cache{c}
	if ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}

/** New
* @Description: 创建cache
* @param defaultExpiration 可以为0或负数，输入0会被设为-1
* @param cleanupInterval
* @return *Cache
 */
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]*Item)
	return newCacheWithJanitor(defaultExpiration, cleanupInterval, items)
}

// ####################################### API #######################################

/** Delete
* @Description: 删除key
* @param key
 */
func (c *cache) Delete(key string) {
	c.mu.Lock()
	vo, evicted := c.delete(key)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(key, vo)
	}
}

/** delete
* @Description: 删除key，会返回需要执行回调函数的ov
* @param key
* @return interface{}
* @return bool
 */
func (c *cache) delete(key string) (interface{}, bool) {
	if c.onEvicted != nil {
		if v, found := c.items[key]; found {
			delete(c.items, key)
			return v.Object, true
		}
	}
	delete(c.items, key)
	return nil, false
}

/** DeleteExpired
* @Description: 删除过期的key
 */
func (c *cache) DeleteExpired() {
	var evictedItems []keyAndValue
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && v.Expiration < now {
			if vo, b := c.delete(k); b {
				evictedItems = append(evictedItems, keyAndValue{k, vo})
			}
		}
	}
	c.mu.Unlock()
	// 删除时需要执行的回调操作
	for _, item := range evictedItems {
		c.onEvicted(item.key, item.value)
	}
}

/** OnEvicted
* @Description: 注册cache的回调函数
* @param f
 */
func (c *cache) OnEvicted(f func(string, interface{})) {
	c.mu.Lock()
	c.onEvicted = f
	c.mu.Unlock()
}

/** Items
* @Description: 获取所有未过期的key信息
* @return map[string]*Item
 */
func (c *cache) Items() map[string]*Item {
	c.mu.Lock()
	defer c.mu.Unlock()

	var res = make(map[string]*Item, len(c.items))
	now := time.Now().UnixNano()
	for k, item := range c.items {
		if item.Expiration > 0 {
			if now > item.Expiration {
				continue
			}
		}
		res[k] = item
	}
	return res
}

/** ItemCount
* @Description: 返回当前key的总数，可能有过期的item还未来得及清理也计算在内
* @return int
 */
func (c *cache) ItemCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

/** Flush
* @Description: 清除所有的key
 */
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = map[string]*Item{}
	c.mu.Unlock()
}

/** Get
* @Description: 检索key，需要加锁避免读到的数据不一致
* @param key
* @return interface{}
* @return bool
 */
func (c *cache) Get(key string) (interface{}, bool) {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.get(key)
}

func (c *cache) get(key string) (interface{}, bool) {
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Object, true
}

/** Set
* @Description: 设置key，已存在则覆盖
* @param key
* @param val
* @param d
 */
func (c *cache) Set(key string, val interface{}, d time.Duration) {
	var expire int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		expire = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	c.items[key] = &Item{Object: val, Expiration: expire}

	c.mu.Unlock()
}

func (c *cache) set(key string, val interface{}, d time.Duration) {
	var expire int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		expire = time.Now().Add(d).UnixNano()
	}
	c.items[key] = &Item{Object: val, Expiration: expire}
}

/** Add
* @Description: 新增key，key已存在不能成功
* @param key
* @param val
* @param d
* @return error
 */
func (c *cache) Add(key string, val interface{}, d time.Duration) error {
	c.mu.Lock()
	if _, b := c.get(key); b {
		return fmt.Errorf("the key : [%s] already exists", key)
	}
	c.set(key, val, d)
	c.mu.Unlock()
	return nil
}

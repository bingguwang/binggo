package main

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	data  map[string]*cacheEntry
	close chan struct{}
}

type cacheEntry struct {
	value    interface{}
	expireAt time.Time
}

func NewCache(cleanInterval time.Duration) *Cache {
	c := &Cache{
		data:  make(map[string]*cacheEntry),
		close: make(chan struct{}),
	}

	if cleanInterval > 0 {
		go c.startCleaner(cleanInterval)
	}
	return c
}

// 设置键值对，ttl为生存时间
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = &cacheEntry{
		value:    value,
		expireAt: time.Now().Add(ttl),
	}
}

// 获取键值（自动处理过期删除）
func (c *Cache) Get(key string) interface{} {
	// 第一次检查（读锁保护）
	c.mu.RLock()
	entry, exists := c.data[key]
	if !exists {
		c.mu.RUnlock()
		return nil
	}
	if time.Now().Before(entry.expireAt) {
		value := entry.value
		c.mu.RUnlock()
		return value
	}
	c.mu.RUnlock() // 释放读锁

	// 获取写锁
	c.mu.Lock()
	defer c.mu.Unlock()

	/// 二次检查为了考虑到 在释放读锁和获取写锁的间隙期间，可能有其他协程更新了该键的过期时间

	// 第二次检查（写锁保护）
	// 惰性删除过期键, 因为定时清除无法保证准确的过期时间生效，所以需要再获取数据的时候手动删除过期数据
	if entry, exists = c.data[key]; exists && time.Now().After(entry.expireAt) {
		delete(c.data, key) // 确认过期后才删除
	}
	return nil
}

// 定期清理过期键
func (c *Cache) startCleaner(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.purgeExpired()
		case <-c.close:
			return
		}
	}
}

// 批量删除过期键
func (c *Cache) purgeExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		if now.After(entry.expireAt) {
			delete(c.data, key)
		}
	}
}

// 关闭缓存释放资源
func (c *Cache) Close() {
	close(c.close)
}
func main() {
	// 创建缓存（设置10秒清理间隔）
	cache := NewCache(10 * time.Second)
	defer cache.Close()

	// 设置键值（5秒过期）
	cache.Set("user:1001", "Alice", 5*time.Second)

	// 立即获取（返回Alice）
	println(cache.Get("user:1001").(string))

	// 6秒后获取（自动删除返回nil）
	time.Sleep(6 * time.Second)
	if val := cache.Get("user:1001"); val == nil {
		println("Key expired")
	}
}

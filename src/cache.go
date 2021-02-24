package src

type Cache struct {
	cache map[int]*BlockTotalData
}

func (c *Cache) CheckBlockData(blockId int) *BlockTotalData {
	if val, ok := c.cache[blockId]; ok {
		return val
	}
	return nil
}

func (c *Cache) InsertBlockData(blockId int, data *BlockTotalData) {
	c.cache[blockId] = data
}

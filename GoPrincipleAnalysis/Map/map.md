<!-- TOC -->
* [map](#map)
  * [Go语言中map的底层结构](#go语言中map的底层结构)
    * [map的访问原理](#map的访问原理)
    * [map的赋值原理](#map的赋值原理)
    * [map的扩容](#map的扩容)
      * [扩容过程](#扩容过程)
      * [迁移时机](#迁移时机)
    * [map的删除原理](#map的删除原理)
    * [map的遍历](#map的遍历)
      * [遍历过程](#遍历过程)
        * [迭代器](#迭代器)
<!-- TOC -->
# map
## Go语言中map的底层结构

1. 每个正常桶都有概率有一个溢出桶？
答：是的
2. tophash的作用是什么，为什么要选择用桶的键的哈希值的高8位做一个属性，怎么判定？
答：tophash的主要作用是快速确定一个键是否存在于某个桶中，以及它存储在桶中的哪个位置。
（1）快速比较：在查找键时，map会首先比较tophash数组中的值，而不是直接比较键的值。由于tophash数组较小且是直接访问的，可以大大加快查找速度。
（2）避免完整哈希计算：通过比较tophash，可以在不计算完整哈希值的情况下快速排除不匹配的键。
（3）指示空槽：tophash中的特殊值可以用来指示桶中的空槽，这样在插入时可以快速找到可用位置。
使用tophash的目的主要是为了提高map操作的效率，通过减少不必要的键比较和哈希计算，使得查找、插入和删除操作更快。
3. 为什么要先用tophash比而不用key直接比？
答：为了快速比对，hash前8位基本能确定两个key是否相同。（前8位相同key可同可不同，前8位不同key一定不同）

```go
type hmap struct {
	count      int              // map中元素个数，对应于len(map)的值
	flags      uint8            // 状态标志位，标记map的状态
	B          uint8            // 桶数以2为底的对数，如B = 3，那桶个数为2^3=8
	noverflow  uint16           // 溢出桶数量近似值
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr          // 扩容进度计数器
	extra      *mapextra  // 指向mapextra结构的指针，mapextra存储map中的溢出桶
}

type mapextra struct {
	overflow *[]*bmap       // 溢出桶链表地址
	oldoverflow *[]*bmap    // 老的溢出桶链表地址
	nextoverflow *bmap      // 下一个空闲溢出桶地址
}
```

### map的访问原理
```go
v     := map[key]
v, ok := map[key]
```
1. 判断map是否为空或者无数据，若为空或者无数据返回对应的空值。
2. map写检测，如果正处在写状态，表示此时不能进行操作。
3. 计算hash值和掩码。   掩码：掩码是一个位操作，通常用于哈希值的最后几位，以确定桶的位置。掩码的目的是将哈希值映射到一个较小的范围，通常是map桶的数量范围内。
```go
mask := uintptr(1)<<h.B - 1
bucketIndex := hashValue & mask
```
4. 判断当前map是否处于扩容状态，如果在扩容执行下面的步骤：
- 根据状态位（tophash[0]）判断当前桶是否发生迁移（通过将已经迁移的 bucket 的 tophash 设置为介于 emptyOne 和 minTopHash 之间的特殊值，Go 的 map 实现能够有效地区分已经迁移的 bucket 和正常的 bucket。evacuated 函数通过检查这个特殊的值范围来判断 bucket 是否已迁移。）
- 如果发生迁移，在新桶中查找
- 未被迁移，在旧桶中查找
- 根据掩码找到的位置
5. 依次遍历桶以及溢出桶来查找key
- 遍历桶内的8个槽位
- 比较该槽位的tophash和当前key的tophash是否相等
  - 相同，继续比较key是否相同，相同则直接返回对应value
  - 不相同，查看这个槽位的状态位是否为“后继空”状态
    - 是，key在以后的槽中也没有，这个key不存在，直接返回零值
    - 否，遍历下一个槽位
6. 当前桶没有找到，则遍历溢出桶，与上述方式一致

### map的赋值原理
```go
map[key] = value
```
1. map先初始化后赋值
2. map非线程安全，不支持并发读写操作

赋值大致流程：
1. map写检测，如果正处在写状态，表示此时不能进行读取，报fatal error
2. 计算hash值，将map置为写状态
3. 判断桶数组是否为空，若为空，初始化桶数组
4. 目标桶查找
- 根据hash值找到桶的位置
- 判断当前桶是否处于扩容：
  - 若正在扩容，迁移这个桶，并且还另外帮忙多迁移一个桶以及它的溢出桶
```go
为什么需要另外帮忙多迁移一个桶以及它的溢出桶？
（1）加速扩容过程：在扩容过程中，哈希表需要将旧桶中的元素迁移到新的桶中。这个过程需要时间，如果只在插入新元素时迁移当前桶，那么整个扩容过程可能会很慢，特别是在有大量插入操作的时候。通过额外迁移一个桶及其溢出桶，可以加快扩容过程，使哈希表更快地达到新的状态。
（2）减少并发冲突：在高并发环境中，如果多个 goroutine 同时进行插入操作，会导致频繁地触发扩容操作。如果每次插入只迁移当前桶，这会增加并发冲突的概率。通过额外迁移其他桶，可以减少这种冲突，因为每次插入操作都在尽量多地完成迁移工作，减小其他 goroutine 需要处理的工作量。
（3）避免长时间锁定：如果扩容过程中每次只迁移一个桶，可能会导致某些插入操作需要等待较长时间，因为每个插入操作都要等待之前的扩容完成。通过一次性多迁移一些桶，可以减少这种等待时间，提高整体的性能。
```
- 获取目标桶的指针，计算出tophash，开始后面的key查找过程
5. key查找
- 遍历桶和它的溢出桶的每个槽位，按下述方式寻找
- 判断槽位的tophash和目标tophash
  - 不相等
    - 槽位tophash为空，标记这个位置为候选位置（为什么不直接插入？因为后续未遍历到的位置可能已经存在这个key，如果这个key存在则会更新key对应的value，只有当这个key不存在才会插入）
    - 槽位tophash的标志位为“后继空状态”，说明这个key之前没有被插入过，插入key/value
    - tophash标志位不为空，说明存储着其他key，说明当前槽的tophash不符合，继续遍历下一个槽
  - 相等
    - 判断当前槽位的key与目标key是否相等
      - 不相等，继续遍历下一个槽位
      - 相等，找到了目标key的位置，原来已存在键值对，则修改key对应的value，然后执行收尾程序
6. key插入
- 若map中既没有找到key，且根据这个key找到的桶及其这个桶的溢出桶中没有空的槽位了，要申请一个新的溢出桶，在新申请的桶里插入
- 否则在找到的位置插入
7. 收尾程序
- 再次判断map的写状态
- 清除map的写状态

**需要注意：申请一个新的溢出桶的时候并不会直接创建一个，会优先使用map初始化时存储在extra \*mapextra字段中的溢出桶，只有这些预分配的溢出桶使用完了，才会新建**

### map的扩容
两种情况触发扩容：
- map的负载因子已经超过6.5（负载因子 = 哈希表中的元素数量 / 桶的数量）
- 溢出桶的数量过多

**扩容时要判断map是否正处在扩容状态，避免二次扩容。（map扩容不是原子操作，不是一次性完成）**

上述两种情况对应于不同的扩容策略：https://www.bilibili.com/video/BV1hv411x7we/?p=4&vd_source=cc9cc8dfd848ca60c1bf76f733bc2db3
- 负载因子超过6.5：双倍扩容（扩容举例：若B = 2，哈希值与运算（2^B-1 = 0000 0011）确定桶位置，双倍扩容后B = 3，哈希值与运算（2^B - 1 = 0000 0111）散列旧桶中的kv）
- 溢出桶数量过多：等量扩容（一般认为溢出桶的数量接近正常桶数量时）
  - B <= 15，noverflow >= 2^B
  - B > 15, noverflow >= 2^15
  
1. 为什么负载因子6.5触发扩容？
答：源码对负载因子定义为6.5，是经过测试后取出的一个比较合理的值。
负载因子6.5说明桶快用完了，存在溢出的情况下，查找一个key可能要去遍历溢出桶，会造成查找性能下降，有必要扩容

2. 溢出桶数量过多？
答：溢出桶数量过多，查找元素的时候要去遍历溢出桶链表，性能下降，要进行扩容，优化溢出桶数量，提升查找性能。（等量扩容应对的情况：有许多溢出桶，溢出桶内的key-value被删除了一些，也就是溢出桶存在不少空位，导致溢出桶数量很多，但实际存放元素较少的情形，这个时候需要等量重建，将一条溢出桶的链表分散的kv集中起来）

#### 扩容过程
发生在map的赋值操作，在满足上述两个扩容条件时触发。扩容过程用到两个函数，hashGrow()和growWork()，其中hashGrow()只是分配新的buckets，并将老的buckets挂到oldbuckets字段上，并未参与真正的数据迁移，数据迁移由growWork()函数完成。

#### 迁移时机
1. growWork()函数会在mapassign和mapdelete函数中被调用。
2. 迁移过程一般发生在插入或修改、删除key的时候。(为什么读不迁移数据？1.保证读数据性能 2.map不允许并发写，所以增删改迁移数据不存在并发问题，但map允许并发读，如果读也允许迁移数据，会存在并发问题)
3. 扩容完毕后（预分配内存），不会马上进行迁移，而是采取写时复制的方式，当有访问到具体的bucket时，才会逐渐的将oldbucket迁移到新bucket中。

```go
func growWork(t *maptype, h *hmap, bucket uintptr) {
// make sure we evacuate the oldbucket corresponding
// to the bucket we're about to use
    evacuate(t, h, bucket&h.oldbucketmask())

// evacuate one more oldbucket to make progress on growing
    if h.growing() {
        evacuate(t, h, h.nevacuate)
    }
}
```
evacuate函数，大致迁移过程如下：
1. 判断当前bucket是不是已经迁移，没迁移就做迁移操作(该map是否发生迁移？判断oldbuckets是否为空)
```go
b := (*bmap)(add(h.oldbuckets, oldbucket*uintptr(t.BucketSize)))
newbit := h.noldbuckets()
// 判断旧桶是否已经被迁移了
if !evacuated(b) {
	do...   // 做转移操作
}
```
2. evacuated函数直接通过tophash中第一个hash值判断当前bucket是否被转移
```go
func evacuated(b *bmap) bool {
	h := b.tophash[0]
	return h > emptyOne && h < minTopHash
}
```
3. 数据迁移时，根据扩容规则，可能迁移到大小相同的buckets上，也可能迁移到2倍大的buckets上。

（迁移到等量数组，迁移后的目标桶位置还是在原先位置，如果双倍扩容迁移到2倍桶数组，迁移完的目标桶位置有可能在原位置，也有可能是原位置+偏移量（偏移量大小为原桶数组的长度））
x，y标记目标迁移位置，x标识的是迁移到相同的位置，y标识的是迁移到2倍桶数组上的位置
4. 确定完bucket之后，按照bucket内的槽位逐条迁移key/value键值对。
5. 迁移完一个桶后，迁移标记位nevacuate + 1，当nevacuate等于旧桶数组大小时，迁移完成，释放旧的桶数组和旧的溢出桶数组。

### map的删除原理
与map的访问原理类似。

具体原理：
1. 如果找到目标key，则将槽位对应的key和value删除。
2. 将tophash置为emptyOne，如果当前槽位的后面没有元素（tophash[i+1] == emptyRest），则将tophash置为emptyRest，并循环向前检查前一个元素，若前一个元素也为空，槽位状态为emptyOne，则将前一个元素的tophash也设置为emptyRest。
3. 2的目的是尽可能将emptyRest向前推，增加后续查找的效率，因为在查找时发现emptyRest状态就不用继续往后了，后面没有元素了。

删除键值对，内存并不会被释放，对于map的频繁写入和删除可能会造成内存泄露。

### map的遍历
**每次遍历的数据顺序都不同**

描述：每次开始遍历前，会随机选择一个桶下标。（一个桶内遍历的起点槽下标，遍历的时候从这个桶开始，在遍历每个桶的时候，都从这个槽下标开始）

原因：
1. Go扩容不是原子操作，是渐进式，遍历时可能发生扩容，一旦扩容，key位置就会发生变化，下次遍历就不会按原来的顺序了。
2. hash表中数据每次插入的位置是变化的，同一个map变量内，数据删除再添加的位置也有可能变化，因为在同一个桶及溢出链表中数据的位置不分先后。

#### 遍历过程
##### 迭代器
```go
type hiter struct {
    key         unsafe.Pointer  // 键值对的键，key为空指针nil说明遍历结束
    elem        unsafe.Pointer  // 键值对的值
    t           *maptype        // map的类型信息
    h           *hmap           // map的地址
    buckets     unsafe.Pointer  // 桶数组指针，指向迭代器初始化后要遍历的桶数组
    bptr        *bmap           // 指向当前遍历到的桶的指针
    overflow    *[]*bmap        // hmap中正常桶的溢出桶指针
    oldoverflow *[]*bmap        // 发生扩容时，hmap中旧桶的溢出桶指针
    startBucket uintptr         // 开始遍历时，初始化的桶下标
    offset      uint8           // 开始遍历时，初始化的槽位下标
    wrapped     bool            // 标识是否遍历完了，为true则遍历完了
    B           uint8           // 初始化迭代器时，h.B
    i           uint8           // 当前桶已经遍历的键值对数量，i = 0 时，开始遍历当前桶的第一个槽位，i = 8时，当前桶已经遍历完，将it.bptr指向下一个桶
    bucket      uintptr         // 当前遍历桶的偏移量
    checkBucket uintptr         // 桶状态标记位，如果不是noCheck，则表明当前桶还没有迁移
}
```
整体遍历过程：
1. 初始化迭代器
2. 开始一轮遍历

初始化迭代器：
1. 判断map是否为空
2. 随机一个开始遍历的起始桶下标
3. 随机一个槽位下标，后续每个桶内的遍历都从该槽位下标开始
4. 把map置为遍历状态
5. 开始执行一次遍历过程

开始一轮遍历：
1. map并发写检测，判断map是否处于并发写状态，是则fatal error（不可捕获的错误）
2. 判断是否已经遍历完了，遍历完了直接退出
3. 开始遍历
4. 首选确定一个随机开始遍历的起始桶下标作为startBucket，然后确定一个随机的槽位下标作为offset
5. 根据随机起始桶下标和随机槽位下标开始遍历当前桶和当前桶的溢出桶，如果当前桶正在扩容，则进行步骤6，否则进行步骤7
6. 在遍历处于扩容状态的bucket的时候，因为当前bucket正在扩容，我们并不会遍历这个桶，而是会找到这个桶的旧桶old_bucket，遍历旧桶中的一部分key，这些key重新hash计算后能够散列到bucket中，对那些key经过重新hash计算不散列到bucket中的key，则跳过
7. 根据遍历初始化的时候选定的随机槽位开始遍历桶内的各个key/value
8. 继续遍历bucket溢出指针指向的溢出链表中的溢出桶
9. 假如遍历到了起始桶startBucket，则说明遍历完了，结束遍历


针对6的处理举个例子：
1. 双倍扩容2->4（B:1->2）
2. 
- 旧桶oldbuckets[0]中的kv会散列存储在新桶buckets[0]和buckets[2]中 
- 旧桶oldbuckets[1]中的kv会散列存储在新桶buckets[1]和buckets[3]中
3. 若oldbuckets[0]未完成迁移，oldbuckets[1]已完成迁移
4. 遍历新桶，假设从新桶buckets[3]的槽位0开始，因为旧桶oldbuckets[1]已经完成迁移，直接遍历buckets[3]的内容，放入遍历结果集中；
5. 接着遍历新桶buckets[0]，因为旧桶oldbuckets[0]未完成迁移，故需要遍历旧桶oldbuckets[0]，将经过散列后能够放入新桶buckets[0]中的元素放入遍历结果集中，放入不了新桶buckets[0]中的元素跳过；
6. 接着遍历新桶buckets[1]，因为旧桶oldbuckets[1]已经完成迁移，直接遍历buckets[1]的内容，放入遍历结果集中；
7. 接着遍历新桶buckets[2]，因为旧桶oldbuckets[0]未完成迁移，故需要遍历旧桶oldbuckets[0]，将经过散列后能够放入新桶buckets[2]中的元素放入遍历结果集中，放入不了新桶buckets[2]中的元素跳过；
8. 接着遍历新桶buckets[3]，已经遍历过，遍历过程结束。
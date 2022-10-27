package timewheel

import (
	"container/list"
	"math"
	"time"
)

const const_ADD_OPERATION = 1
const const_REMOVE_OPERATION = 2

// Job 延时任务回调函数
type Job func(interface{}, interface{})

// TaskData 回调函数参数类型

// TimeWheel 时间轮
type TimeWheel struct {
	interval time.Duration // 指针每隔多久往前移动一格
	ticker   *time.Ticker
	slots    []*list.List // 时间轮槽
	// key: 定时器唯一标识 value: 定时器所在的槽, 主要用于删除定时器, 不会出现并发读写，不加锁直接访问
	timer                map[interface{}]int
	lastFinishTickerTime time.Time

	currentPos int // 当前指针指向哪一个槽
	slotNum    int // 槽数量
	job        Job // 定时器回调函数
	//addTaskChannel    chan Task        // 新增任务channel
	//removeTaskChannel chan interface{} // 删除任务channel
	changeTaskChannel chan taskOperation //新增删除任务channel
	stopChannel       chan bool          // 停止定时器channel
}

type taskOperation struct {
	opType int
	key    interface{}
	task   Task
}

// Task 延时任务
type Task struct {
	addTime time.Time
	delay   time.Duration // 延迟时间
	circle  int           // 时间轮需要转动几圈
	key     interface{}   // 定时器唯一标识, 用于删除定时器
	data    interface{}   // 回调函数参数
	job     Job
	iScron  bool
}

// New 创建时间轮 //当精度是1ms的时候，准确度较差，建议精度设置最小10ms
func New(interval time.Duration, slotNum int, taskCap int, job Job) *TimeWheel {
	if interval <= 0 || slotNum <= 0 || taskCap < 0 {
		return nil
	}
	tw := &TimeWheel{
		interval:   interval,
		slots:      make([]*list.List, slotNum),
		timer:      make(map[interface{}]int),
		currentPos: 0,
		job:        job,
		slotNum:    slotNum,
		//addTaskChannel:    make(chan Task),
		//removeTaskChannel: make(chan interface{}),
		changeTaskChannel: make(chan taskOperation, taskCap),
		stopChannel:       make(chan bool),
	}

	tw.initSlots()

	return tw
}

// 初始化槽，每个槽指向一个双向链表
func (tw *TimeWheel) initSlots() {
	for i := 0; i < tw.slotNum; i++ {
		tw.slots[i] = list.New()
	}
}

// Start 启动时间轮
func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.start()
}

// Stop 停止时间轮
func (tw *TimeWheel) Stop() {
	tw.stopChannel <- true
}

// Add 添加定时器 key为定时器唯一标识
func (tw *TimeWheel) Add(delay time.Duration, key interface{}, data interface{}) {
	tw.AddFunc(delay, key, data, nil)
}

// Add 添加定时器 key为定时器唯一标识
func (tw *TimeWheel) AddFunc(delay time.Duration, key interface{}, data interface{}, job Job) {
	if delay < 0 {
		return
	}
	//tw.addTaskChannel <- Task{delay: delay, key: key, data: data}
	task := Task{delay: delay, key: key, data: data, job: job, iScron: false}
	task.addTime = time.Now()
	tw.changeTaskChannel <- taskOperation{
		opType: const_ADD_OPERATION,
		key:    key,
		task:   task,
	}
}

// Add 添加定时器 key为定时器唯一标识
func (tw *TimeWheel) AddCron(delay time.Duration, key interface{}, data interface{}) {
	tw.AddCronFunc(delay, key, data, nil)
}

// Add 添加定时器 key为定时器唯一标识
func (tw *TimeWheel) AddCronFunc(delay time.Duration, key interface{}, data interface{}, job Job) {
	if delay <= 0 {
		return
	}
	//tw.addTaskChannel <- Task{delay: delay, key: key, data: data}
	task := Task{delay: delay, key: key, data: data, job: job, iScron: true}
	task.addTime = time.Now()
	tw.changeTaskChannel <- taskOperation{
		opType: const_ADD_OPERATION,
		key:    key,
		task:   task,
	}
}

// Remove 删除定时器 key为添加定时器时传递的定时器唯一标识
func (tw *TimeWheel) Remove(key interface{}) {
	if key == nil {
		return
	}
	tw.changeTaskChannel <- taskOperation{
		opType: const_REMOVE_OPERATION,
		key:    key,
	}
}

// var count uint64 = 0
func (tw *TimeWheel) start() {
	for {
		select {
		case <-tw.ticker.C:
			tw.tickHandler()
			//fmt.Printf("ticker locate %d\n",atomic.LoadUint64(&count))
			//atomic.AddUint64(&count,1)
		case op := <-tw.changeTaskChannel:
			if op.opType == const_ADD_OPERATION {
				//fmt.Printf("addTask locate %d\n",atomic.LoadUint64(&count))
				tw.addTask(&op.task)
			} else if op.opType == const_REMOVE_OPERATION {
				//fmt.Printf("removeTask locate %d\n",atomic.LoadUint64(&count))
				tw.removeTask(op.key)
			}
			//tw.addTask(&task)
		//case key := <-tw.removeTaskChannel:
		//	tw.removeTask(key)
		case <-tw.stopChannel:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) tickHandler() {
	l := tw.slots[tw.currentPos]
	//fmt.Printf("pos %d\n",tw.currentPos)
	tw.scanAndRunTask(l)
	if tw.currentPos == tw.slotNum-1 {
		tw.currentPos = 0
	} else {
		tw.currentPos++
	}
	tw.lastFinishTickerTime = time.Now()
}

// 扫描链表中过期定时器, 并执行回调函数
func (tw *TimeWheel) scanAndRunTask(l *list.List) {
	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.circle > 0 {
			task.circle--
			e = e.Next()
			continue
		}

		//fmt.Printf("task locate %d %v\n",atomic.LoadUint64(&count),task.key)

		if task.job != nil {
			task.job(task.key, task.data)
		} else if tw.job != nil {
			tw.job(task.key, task.data)
		}

		if task.iScron {
			tw.AddCronFunc(task.delay, task.key, task.data, task.job)
		}

		next := e.Next()
		l.Remove(e)
		if task.key != nil {
			delete(tw.timer, task.key)
		}
		e = next
	}
}

// 新增任务到链表中
func (tw *TimeWheel) addTask(task *Task) {
	var pos, circle int
	if task.addTime.Sub(tw.lastFinishTickerTime) <= 0 {
		delay := task.delay - tw.interval
		if delay < 0 {
			delay = 0
		}
		pos, circle = getPositionAndCircle(delay, tw.interval, tw.slotNum, tw.currentPos)
	} else {
		pos, circle = getPositionAndCircle(task.delay, tw.interval, tw.slotNum, tw.currentPos)
	}
	task.circle = circle

	tw.slots[pos].PushBack(task)

	if task.key != nil {
		tw.timer[task.key] = pos
	}
	//fmt.Printf("%s addTask %v %v %v %v %v\n",formatTime(task.addTime), task.delay, task.circle, pos, task.key,task.data)
}

// 获取定时器在槽中的位置, 时间轮需要转动的圈数
func getPositionAndCircle(d time.Duration, inteval time.Duration,
	slotNum int, currentPos int) (pos int, circle int) {
	skip := int(math.Round(float64(d) / float64(inteval)))
	pos = (currentPos + skip) % slotNum
	circle = (currentPos + skip) / slotNum
	if pos < currentPos {
		circle--
	}
	return
}

// 从链表中删除任务
func (tw *TimeWheel) removeTask(key interface{}) {
	// 获取定时器所在的槽
	position, ok := tw.timer[key]
	if !ok {
		return
	}
	// 获取槽指向的链表
	l := tw.slots[position]
	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.key == key {
			delete(tw.timer, task.key)
			l.Remove(e)
			//fmt.Printf("%s removeTask %v\n", formatTime(time.Now()), key)
		}

		e = e.Next()
	}
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.999999")
}

/*
func main()  {
    // 初始化时间轮
    // 第一个参数为tick刻度, 即时间轮多久转动一次
    // 第二个参数为时间轮槽slot数量
    // 第三个参数为回调函数
    tw := timewheel.New(1 * time.Second, 3600, func(data interface{}) {
        // do something
    })

    // 启动时间轮
    tw.Start()

    // 添加定时器
    // 第一个参数为延迟时间
    // 第二个参数为定时器唯一标识, 删除定时器需传递此参数
    // 第三个参数为用户自定义数据, 此参数将会传递给回调函数, 类型为interface{}
    tw.Add(5 * time.Second, conn, map[string]int{"uid" : 105626})

    // 删除定时器, 参数为添加定时器传递的唯一标识
    tw.Remove(conn)

    // 停止时间轮
    tw.Stop()

    select{}
}
*/

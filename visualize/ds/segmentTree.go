package ds

/*
	采用线段树统计区间范围落在[Lowerbound,Upperbound)的次数Record,复杂度O(logN)
	单次修改O(logN)
*/

// 存储结构
type Tree struct {
	lo, ro *Tree // 左/右孩子
	l, r   int64 // 支配的区间
	record int64 // 出现的次数
}

// 创建树
func NewTree(lo, hi int64) *Tree {
	return &Tree{l: lo, r: hi}
}

// 获得对应记录
func (o *Tree) getRecord() int64 {
	if o != nil {
		return o.record
	}
	return 0
}

// 维护值域
func (o *Tree) maintain() {
	o.record = o.lo.getRecord() + o.ro.getRecord()
}

// 进行更新
func (o *Tree) Update(i int64) {
	if o.l == o.r {
		o.record++ // 增加对应记录
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == nil {
			o.lo = &Tree{l: o.l, r: m}
		}
		o.lo.Update(i)
	} else {
		if o.ro == nil {
			o.ro = &Tree{l: m + 1, r: o.r}
		}
		o.ro.Update(i)
	}
	o.maintain()
}

// 进行查询
func (o *Tree) Query(l, r int64) int64 {
	if o == nil || l > o.r || r < o.l {
		return 0
	}
	if l <= o.l && o.r <= r {
		return o.record
	}
	return o.lo.Query(l, r) + o.ro.Query(l, r)
}
